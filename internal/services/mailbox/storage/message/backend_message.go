package message

type BackendMessage interface {
	CreateMessage(content string) (Message, error)
	GetMessage(id int64) (Message, error)
	CountMessages() (int64, error)
	// UpdateMessage(ctx context.Context, id int64) error
	// DeleteMessage(ctx context.Context, id int64) error
}
