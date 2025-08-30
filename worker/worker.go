package main

import (
	"log"
	"temporal-orchestration/activity"
	"temporal-orchestration/constants"
	"temporal-orchestration/workflow"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {

	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create Temporal client", err)
	}

	w := worker.New(c, constants.TaskQueue, worker.Options{})
	w.RegisterWorkflow(workflow.UserOnboardingWorkflow)
	w.RegisterActivity(activity.SendWelcomeEmail)
	w.RegisterActivity(activity.ProcessUserOnboarding)

	if err := w.Run(worker.InterruptCh()); err != nil {
		log.Fatalln("Unable to start worker", err)
	}

}
