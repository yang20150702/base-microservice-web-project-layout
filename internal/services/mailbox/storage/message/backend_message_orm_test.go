package message

import (
	"testing"

	"github.com/yang20150702/base-microservice-web-project-layout/internal/configs"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func init() {
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Message{})
}

func TestPrint(t *testing.T) {
	t.Log("test")
}

func TestCreateMessage(t *testing.T) {
	s := NewBackendMessageOrm(&configs.ServerConf{}, db)

	rslt, err := s.CreateMessage("test content")
	if err != nil {
		t.Error(err)
	}
	t.Log(rslt)

}

func TestGetMessage(t *testing.T) {
	s := NewBackendMessageOrm(&configs.ServerConf{}, db)

	result, err := s.GetMessage(int64(1))
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}

func TestCountMessages(t *testing.T) {
	s := NewBackendMessageOrm(&configs.ServerConf{}, db)

	result, err := s.CountMessages()
	if err != nil {
		t.Error(err)
	}

	t.Log(result)
}
