package message

import (
	"github.com/yang20150702/base-microservice-web-project-layout/framework"
	"github.com/yang20150702/base-microservice-web-project-layout/internal/configs"
	"gorm.io/gorm"
)

var log = framework.GetLogger()

type BackendMessageOrm struct {
	conf *configs.ServerConf
	db   *gorm.DB
	meta *Message
}

func NewBackendMessageOrm(conf *configs.ServerConf, db *gorm.DB) BackendMessage {
	b := &BackendMessageOrm{conf: conf, db: db, meta: &Message{}}
	return b
}

func (b *BackendMessageOrm) InitTable() {
	log.Debug("初始化操作")
	b.db.AutoMigrate(&Message{})
}

func (b *BackendMessageOrm) CreateMessage(content string) (Message, error) {
	m := Message{Content: content}
	result := b.db.Create(&m)
	if result.Error != nil {
		log.Error(result.Error)
		return Message{}, result.Error
	}
	return m, nil
}

func (b *BackendMessageOrm) GetMessage(id int64) (Message, error) {
	var m Message
	result := b.db.First(&m, id)
	if result.Error != nil {
		return Message{}, result.Error
	}
	return m, nil
}

func (b *BackendMessageOrm) CountMessages() (int64, error) {
	var count int64
	result := b.db.Table(b.meta.TableName()).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}

	return count, nil
}
