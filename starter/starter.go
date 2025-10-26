package main

import (
	"context"
	"log"

	"temporal-orchestration/constants"
	"temporal-orchestration/mocks"
	"temporal-orchestration/models"
	"temporal-orchestration/workflows"

	"go.temporal.io/sdk/client"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalf("Unable to create Temporal client: %v", err)
	}
	defer c.Close()

	runWorkflow(c)
}

func runWorkflow(c client.Client) {
	const (
		userID    = "user-12345"
		userEmail = "test@example.com"
	)

	userData := models.UserData{
		ID: userID,
		Email:  userEmail,
	}

	we, err := startUserOnboardingWorkflow(c, userID, userData)
	if err != nil {
		log.Fatalf("Unable to start workflow: %v", err)
	}

	log.Println("Started workflow with ID:", we.GetID())

	sendSignals(c, userID)

	result, err := getWorkflowResult(we)
	if err != nil {
		log.Fatalf("Unable to get workflow result: %v", err)
	}

	log.Printf("Workflow completed. Final state: %+v\n", result)
}

func startUserOnboardingWorkflow(c client.Client, workflowID string, input models.UserData) (client.WorkflowRun, error) {
	return c.ExecuteWorkflow(context.Background(), client.StartWorkflowOptions{
		ID:        workflowID,
		TaskQueue: constants.TaskQueue,
	}, workflows.UserOnboardingWorkflow, input)
}

func sendSignals(c client.Client, userID string) {
	mocks.MockEmailSignalForTemporal(c, userID)
	mocks.MockSubscribeSignalForTemporal(c, userID)
}

func getWorkflowResult(we client.WorkflowRun) (*models.UserWorkflowState, error) {
	var result models.UserWorkflowState
	err := we.Get(context.Background(), &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}