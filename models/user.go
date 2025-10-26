package models

type UserData struct {
	ID string
	Email  string
}

type UserWorkflowState struct {
	IsEmailSent  bool
	IsSubscribed bool
}
