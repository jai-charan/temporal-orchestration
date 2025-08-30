package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"go.temporal.io/sdk/client"

	"temporal-orchestration/constants"
	"temporal-orchestration/types"
	"temporal-orchestration/workflow"
)

func generateRandomUser() types.UserData {
	rand.Seed(time.Now().UnixNano())

	userID := fmt.Sprintf("user-%d", rand.Intn(100000))
	email := fmt.Sprintf("user%d@example.com", rand.Intn(100000))

	return types.UserData{
		UserID: userID,
		Email:  email,
	}
}

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create Temporal client:", err)
	}
	defer c.Close()

	user := generateRandomUser()
	workflowID := fmt.Sprintf("%s%s", constants.WorkflowIDPrefix, user.UserID)

	// Start the user onboarding workflow
	we, err := c.ExecuteWorkflow(context.Background(), client.StartWorkflowOptions{
		ID:        workflowID,
		TaskQueue: constants.TaskQueue,
	}, workflow.UserOnboardingWorkflow, user)
	if err != nil {
		log.Fatalln("Unable to start workflow:", err)
	}
	log.Println("Started workflow:", workflowID)

	// --- Mock: Simulate signal from the Notification Service ---
	go func() {
		time.Sleep(time.Duration(constants.EmailSentDelay) * time.Second)
		err := c.SignalWorkflow(context.Background(), workflowID, "", constants.EmailSentSignal, nil)
		if err != nil {
			log.Println("Error sending 'email_sent' signal:", err)
		}
	}()

	// --- Mock: Simulate signal from the Subscription Service ---
	go func() {
		time.Sleep(time.Duration(constants.SubscribedDelay) * time.Second)
		err := c.SignalWorkflow(context.Background(), workflowID, "", constants.SubscribedSignal, nil)
		if err != nil {
			log.Println("Error sending 'subscribed' signal:", err)
		}
	}()

	// Wait for the workflow to complete and retrieve its final state
	var result *types.UserWorkflowState
	if err := we.Get(context.Background(), &result); err != nil {
		log.Fatalln("Unable to get workflow result:", err)
	}

	log.Printf("Workflow completed. Final state: %+v\n", result)
}
