package workflow

import (
	"time"

	"temporal-orchestration/activity"
	"temporal-orchestration/constants"
	"temporal-orchestration/types"

	"go.temporal.io/sdk/workflow"
)

func UserOnboardingWorkflow(ctx workflow.Context, user types.UserData) (*types.UserWorkflowState, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: time.Duration(constants.ActivityTimeout) * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	state := &types.UserWorkflowState{}

	if err := workflow.ExecuteActivity(ctx, activity.SendWelcomeEmail, user).Get(ctx, nil); err != nil {
		return nil, err
	}

	emailSentCh := workflow.GetSignalChannel(ctx, constants.EmailSentSignal)
	subscribedCh := workflow.GetSignalChannel(ctx, constants.SubscribedSignal)

	selector := workflow.NewSelector(ctx)

	selector.AddReceive(emailSentCh, func(c workflow.ReceiveChannel, more bool) {
		var signalData string
		c.Receive(ctx, &signalData)
		state.IsEmailSent = true
	})

	selector.AddReceive(subscribedCh, func(c workflow.ReceiveChannel, more bool) {
		var signalData string
		c.Receive(ctx, &signalData)
		state.IsSubscribed = true
	})

	for !state.IsEmailSent || !state.IsSubscribed {
		selector.Select(ctx)
	}

	if err := workflow.ExecuteActivity(ctx, activity.ProcessUserOnboarding, user).Get(ctx, nil); err != nil {
		return nil, err
	}

	return state, nil
}
