package common

const (
	DbTypeNote   = 1
	DbTypeUser   = 2
	DbTypeUpload = 3

	CurrentUser = "current_user"
)

type Requester interface {
	GetUserId() int
	GetEmail() string
	GetRole() string
}
