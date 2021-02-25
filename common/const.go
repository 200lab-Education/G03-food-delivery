package common

const (
	DbTypeNote = 1

	CurrentUser = "current_user"
)

type Requester interface {
	GetUserId() int
	GetEmail() string
	GetRole() string
}
