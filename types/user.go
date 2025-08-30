package types

type UserData struct {
	UserID string
	Email  string
}

type UserWorkflowState struct {
	IsEmailSent  bool
	IsSubscribed bool
}
