package storage

import (
	"fmt"

	"github.com/yang20150702/base-microservice-web-project-layout/internal/configs"
	ormConfig "github.com/yang20150702/base-microservice-web-project-layout/internal/services/mailbox/storage/configs"
	"github.com/yang20150702/base-microservice-web-project-layout/internal/services/mailbox/storage/message"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type OrmBackend struct {
	conf *configs.ServerConf
	db   *gorm.DB
	*message.BackendMessageOrm
}

func NewOrmBackend(conf *configs.ServerConf) (Backend, error) {
	var db *gorm.DB
	var err error

	ormBackend := &OrmBackend{conf: conf}

	switch conf.Orm.Sqlite.String() {
	case `sqlite`:
		db, err = gorm.Open(sqlite.Open("test.db"), ormConfig.InitOrmConfig())
	}
	if err != nil {
		return nil, err
	}

	ormBackend.db = db
	ormBackend.registerImpl()
	ormBackend.init()
	fmt.Println("数据库初始化完成")

	return ormBackend, nil
}

func (s *OrmBackend) registerImpl() {
	s.BackendMessageOrm = message.NewBackendMessageOrm(s.conf, s.db).(*message.BackendMessageOrm)
}

func (s *OrmBackend) init() {
	// 进行初始化操作：建表、建索引
	s.BackendMessageOrm.InitTable()
}
