package constants

const (
	// Signal names
	EmailSentSignal  = "email_sent"
	SubscribedSignal = "subscribed"

	// Task queue
	TaskQueue = "User-Onboarding-Task-Queue"

	// Workflow ID prefix
	WorkflowIDPrefix = "user-onboarding-"

	// Activity timeouts
	ActivityTimeout = 15 // seconds

	// Signal delays
	EmailSentDelay  = 15 // seconds
	SubscribedDelay = 25  // seconds

	// Activity delays
	WelcomeEmailDelay    = 2 // seconds
	CompletionEmailDelay = 1 // second
)
