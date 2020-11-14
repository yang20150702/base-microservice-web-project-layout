package config

import "gorm.io/gorm"

var OrmConfig *gorm.Config

func InitOrmConfig() *gorm.Config {
	return OrmConfig
}
