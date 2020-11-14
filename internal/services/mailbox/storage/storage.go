package storage

import (
	"errors"

	"github.com/yang20150702/base-microservice-web-project-layout/internal/configs"
)

var defaultBackend Backend

func Init(conf *configs.ServerConf) (err error) {
	if conf == nil {
		return errors.New("StorageConf is nil")
	}

	switch conf.Storage {
	case `orm`:
		defaultBackend, err = NewOrmBackend(conf)
	default:
		err = errors.New("do not support this backend")
	}

	return err
}

func GetBackendImpl() Backend {
	if defaultBackend == nil {
		panic("you should init storage first")
	}

	return defaultBackend
}
