package message

type Message struct {
	ID      int64 `gorm:"primaryKey"`
	Content string
}

func (m *Message) TableName() string {
	return `message`
}

func (m *Message) GetContent() string {
	return `content`
}
