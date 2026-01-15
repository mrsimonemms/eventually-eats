package workflow

import (
	"context"
	"time"

	"go.temporal.io/sdk/activity"
)

type activities struct{}

func (a *activities) RefundPayment(ctx context.Context) error {
	logger := activity.GetLogger(ctx)
	logger.Info("Activity started")

	time.Sleep(time.Second * 5)

	logger.Info("Activity finished")

	return nil
}

func (a *activities) SendTextMessage(ctx context.Context, status OrderState) error {
	logger := activity.GetLogger(ctx)
	logger.Info("Activity started")

	time.Sleep(time.Second)

	logger.Info("Activity finished")

	return nil
}

func (a *activities) TakePayment(ctx context.Context) error {
	logger := activity.GetLogger(ctx)
	logger.Info("Activity started")

	time.Sleep(time.Second * 5)

	logger.Info("Activity finished")

	return nil
}

func NewActivities() (*activities, error) {
	return &activities{}, nil
}
