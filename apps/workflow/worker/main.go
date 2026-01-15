package main

import (
	"os"

	"github.com/mrsimonemms/eventually-eats/apps/workflow"
	gh "github.com/mrsimonemms/golang-helpers"
	"github.com/mrsimonemms/golang-helpers/temporal"
	"github.com/rs/zerolog/log"
	"go.temporal.io/sdk/worker"
)

func exec() error {
	// The client and worker are heavyweight objects that should be created once per process.
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

	w := worker.New(c, workflow.OrderFoodTaskQueue, worker.Options{})

	w.RegisterWorkflow(workflow.OrderWorkflow)

	activities, err := workflow.NewActivities()
	if err != nil {
		return gh.FatalError{
			Cause: err,
			Msg:   "Unable to create activities",
		}
	}
	w.RegisterActivity(activities)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		return gh.FatalError{
			Cause: err,
			Msg:   "Unable to start worker",
		}
	}

	return nil
}

func main() {
	if err := exec(); err != nil {
		os.Exit(gh.HandleFatalError(err))
	}
}
