package storage

import "github.com/yang20150702/base-microservice-web-project-layout/internal/services/mailbox/storage/message"

type Backend interface {
	message.BackendMessage
}
