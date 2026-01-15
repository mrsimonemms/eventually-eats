package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/mrsimonemms/eventually-eats/apps/workflow"
	gh "github.com/mrsimonemms/golang-helpers"
	"github.com/mrsimonemms/golang-helpers/temporal"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.temporal.io/sdk/client"
)

func exec() error {
	// The client is a heavyweight object that should be created once per process.
	c, err := temporal.NewConnectionWithEnvvars(
		temporal.WithZerolog(&log.Logger),
	)
	if err != nil {
		return gh.FatalError{
			Cause: err,
			Msg:   "Unable to create client",
		}
	}
	defer c.Close()

	workflowID := "ORDER-" + fmt.Sprintf("%d", time.Now().Unix())

	workflowOptions := client.StartWorkflowOptions{
		ID:        workflowID,
		TaskQueue: workflow.OrderFoodTaskQueue,
	}

	ctx := context.Background()

	order := workflow.OrderState{
		Status:     workflow.OrderStatusDefault,
		Collection: true,
		Products: []workflow.OrderProduct{
			{
				// Chips
				ProductID: 1,
				Quantity:  1,
			},
			{
				// Haddock
				ProductID: 3,
				Quantity:  1,
			},
			{
				// Curry sauce
				ProductID: 4,
				Quantity:  1,
			},
		},
	}
	we, err := c.ExecuteWorkflow(
		ctx,
		workflowOptions,
		workflow.OrderWorkflow,
		order,
	)
	if err != nil {
		return gh.FatalError{
			Cause: err,
			Msg:   "Unable to execute workflow",
		}
	}

	l := log.With().Str("workflowID", we.GetID()).Str("runID", we.GetRunID()).Logger()
	l.Info().Msg("Started workflow")

	time.Sleep(time.Second * 5)

	statuses := []workflow.OrderStatus{
		// Accept order
		workflow.OrderStatusAccepted,
		// Cook it
		workflow.OrderStatusPending,
		// Out for delivery
		workflow.OrderStatusReady,
		// Delivered
		workflow.OrderStatusCompleted,
	}

	for _, status := range statuses {
		l.Info().Any("status", status).Msg("Updating status")

		updateHandle, err := c.UpdateWorkflow(ctx, client.UpdateWorkflowOptions{
			WorkflowID:   we.GetID(),
			UpdateName:   workflow.Updates.UPDATE_STATUS,
			WaitForStage: client.WorkflowUpdateStageAccepted,
			Args: []any{
				status,
			},
		})
		if err != nil {
			return gh.FatalError{
				Cause: err,
				WithParams: func(l *zerolog.Event) *zerolog.Event {
					return l.Any("status", status)
				},
				Msg: "Failed to progress order",
			}
		}

		if err := updateHandle.Get(ctx, nil); err != nil {
			return gh.FatalError{
				Cause: err,
				Msg:   "Error getting update response",
				WithParams: func(l *zerolog.Event) *zerolog.Event {
					return l.Any("status", status)
				},
			}
		}

		printState(ctx, l, c, we)

		time.Sleep(time.Second * 5)
	}

	// Wait for end of workflow
	l.Info().Msg("Waiting for end of workflow")
	if err := we.Get(ctx, nil); err != nil {
		return gh.FatalError{
			Cause: err,
			Msg:   "Failed",
		}
	}

	printState(ctx, l, c, we)

	l.Info().Msg("Order completed")

	return nil
}

func getState(ctx context.Context, c client.Client, we client.WorkflowRun) (*workflow.OrderState, error) {
	resp, err := c.QueryWorkflow(ctx, we.GetID(), "", workflow.Queries.GET_STATUS)
	if err != nil {
		return nil, gh.FatalError{
			Cause: err,
			Msg:   "Failed to query workflow",
		}
	}
	var result *workflow.OrderState
	if err := resp.Get(&result); err != nil {
		return nil, fmt.Errorf("unable to decode state query: %w", err)
	}
	return result, nil
}

func printState(ctx context.Context, l zerolog.Logger, c client.Client, we client.WorkflowRun) {
	if state, err := getState(ctx, c, we); err != nil {
		l.Fatal().Err(err).Msg("Failed to get state")
	} else {
		l.Info().Any("state", *state).Msg("State")
	}
}

func main() {
	if err := exec(); err != nil {
		os.Exit(gh.HandleFatalError(err))
	}
}
