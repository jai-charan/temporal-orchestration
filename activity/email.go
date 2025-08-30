package activity

import (
	"context"
	"time"

	"temporal-orchestration/constants"
	"temporal-orchestration/types"

	"go.temporal.io/sdk/activity"
)

func SendWelcomeEmail(ctx context.Context, user types.UserData) error {
	activity.GetLogger(ctx).Info("Sending welcome email", "userID", user.UserID, "email", user.Email)

	time.Sleep(time.Duration(constants.WelcomeEmailDelay) * time.Second)

	activity.GetLogger(ctx).Info("Welcome email sent successfully")
	return nil
}

func ProcessUserOnboarding(ctx context.Context, user types.UserData) error {
	activity.GetLogger(ctx).Info("Processing user onboarding completion", "userID", user.UserID, "email", user.Email)

	time.Sleep(time.Duration(constants.CompletionEmailDelay) * time.Second)

	activity.GetLogger(ctx).Info("User onboarding processed successfully")
	return nil
}
