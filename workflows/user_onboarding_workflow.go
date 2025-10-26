package workflows

import (
	"time"

	"go.temporal.io/sdk/log"
	"temporal-orchestration/activity"
	"temporal-orchestration/constants"
	"temporal-orchestration/models"

	"go.temporal.io/sdk/workflow"
)

func UserOnboardingWorkflow(ctx workflow.Context, user models.UserData) (*models.UserWorkflowState, error) {
	logger := workflow.GetLogger(ctx)
	state := &models.UserWorkflowState{}

	ao := workflow.ActivityOptions{
		StartToCloseTimeout: time.Duration(constants.ActivityTimeout) * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	logger.Info("Starting UserOnboardingWorkflow", "userID", user.ID)

	err := executeSendWelcomeEmail(ctx, user, logger)
	if err != nil {
		return nil, err
	}

	waitForSignals(ctx, state, logger)

	err = executeProcessUserOnboarding(ctx, user, logger)
	if err != nil {
		return nil, err
	}

	logger.Info("User onboarding process completed successfully", "userID", user.ID)

	return state, nil
}

func executeSendWelcomeEmail(ctx workflow.Context, user models.UserData, logger log.Logger) error {
	logger.Info("Executing SendWelcomeEmail activity")
	err := workflow.ExecuteActivity(ctx, activity.SendWelcomeEmail, user).Get(ctx, nil)
	if err != nil {
		logger.Error("Failed to send welcome email", "error", err)
		return err
	}
	logger.Info("Welcome email activity completed successfully")
	return nil
}

func executeProcessUserOnboarding(ctx workflow.Context, user models.UserData, logger log.Logger) error {
	logger.Info("Executing final activity: ProcessUserOnboarding")
	err := workflow.ExecuteActivity(ctx, activity.ProcessUserOnboarding, user).Get(ctx, nil)
	if err != nil {
		logger.Error("Error executing final activity ProcessUserOnboarding", "error", err)
		return err
	}
	return nil
}

func waitForSignals(ctx workflow.Context, state *models.UserWorkflowState, logger log.Logger) {
	emailSentCh := workflow.GetSignalChannel(ctx, constants.EmailSentSignal)
	subscribedCh := workflow.GetSignalChannel(ctx, constants.SubscribedSignal)

	logger.Info("Signal channels created and ready to receive",
		"emailSignal", constants.EmailSentSignal,
		"subscribedSignal", constants.SubscribedSignal)

	selector := workflow.NewSelector(ctx)

	selector.AddReceive(emailSentCh, handleEmailSentSignal(ctx, state, logger))
	selector.AddReceive(subscribedCh, handleSubscribedSignal(ctx, state, logger))

	logger.Info("Waiting for signals...")

	for !state.IsEmailSent || !state.IsSubscribed {
		logger.Debug("Waiting in signal loop...", "emailSent", state.IsEmailSent, "subscribed", state.IsSubscribed)
		selector.Select(ctx)
		logger.Debug("Signal received, checking state...", "emailSent", state.IsEmailSent, "subscribed", state.IsSubscribed)
	}

	logger.Info("All required signals received. Exiting signal waiting loop.")
}

func handleEmailSentSignal(ctx workflow.Context, state *models.UserWorkflowState, logger log.Logger) func(workflow.ReceiveChannel, bool) {
	return func(c workflow.ReceiveChannel, more bool) {
		var signalData string
		c.Receive(ctx, &signalData)
		if !state.IsEmailSent {
			state.IsEmailSent = true
			logger.Info("Received 'email_sent' signal",
				"signalData", signalData,
				"state", state)
		}
	}
}

func handleSubscribedSignal(ctx workflow.Context, state *models.UserWorkflowState, logger log.Logger) func(workflow.ReceiveChannel, bool) {
	return func(c workflow.ReceiveChannel, more bool) {
		var signalData string
		c.Receive(ctx, &signalData)
		if !state.IsSubscribed {
			state.IsSubscribed = true
			logger.Info("Received 'subscribed' signal",
				"signalData", signalData,
				"state", state)
		}
	}
}
