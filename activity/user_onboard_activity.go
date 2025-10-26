package activity

import (
	"context"
	"time"

	"temporal-orchestration/constants"
	"temporal-orchestration/models"

	"go.temporal.io/sdk/activity"
)

func ProcessUserOnboarding(ctx context.Context, user models.UserData) error {
	activity.GetLogger(ctx).Info("Processing user onboarding completion", "userID", user.ID, "email", user.Email)

	time.Sleep(time.Duration(constants.CompletionEmailDelay) * time.Second)

	activity.GetLogger(ctx).Info("User onboarding processed successfully")
	return nil
}
