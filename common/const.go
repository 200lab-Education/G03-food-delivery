package common

import "demo/pubsub"

const (
	DbTypeNote   = 1
	DbTypeUser   = 2
	DbTypeUpload = 3

	CurrentUser = "current_user"
)

const (
	TopicNoteCreated pubsub.Topic = "TopicNoteCreated"
)

type Requester interface {
	GetUserId() int
	GetEmail() string
	GetRole() string
}
