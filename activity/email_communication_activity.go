package activity

import (
	"context"
	"time"

	"temporal-orchestration/constants"
	"temporal-orchestration/models"

	"go.temporal.io/sdk/activity"
)

func SendWelcomeEmail(ctx context.Context, user models.UserData) error {
	activity.GetLogger(ctx).Info("Sending welcome email", "userID", user.ID, "email", user.Email)

	time.Sleep(time.Duration(constants.WelcomeEmailDelay) * time.Second)

	activity.GetLogger(ctx).Info("Welcome email sent successfully")
	return nil
}


