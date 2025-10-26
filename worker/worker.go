package worker

import (
	"log"

	"temporal-orchestration/activity"
	"temporal-orchestration/constants"
	"temporal-orchestration/workflows"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func StartWorker() {

	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalf("Unable to create Temporal client: %v", err)
	}
	defer c.Close()

	w := worker.New(c, constants.TaskQueue, worker.Options{})

	w.RegisterWorkflow(workflows.UserOnboardingWorkflow)
	w.RegisterActivity(activity.SendWelcomeEmail)
	w.RegisterActivity(activity.ProcessUserOnboarding)

	log.Printf("Starting Worker on Task Queue: %s", constants.TaskQueue)

	if err := w.Run(worker.InterruptCh()); err != nil {
		log.Fatalf("Unable to start worker: %v", err)
	}
}