package configs

type ServerConf struct {
	Storage string  `yaml:"storage"`
	Orm     OrmConf `yaml:"orm"`
}

type OrmConf struct {
	Sqlite SqliteConf `yaml:"sqlite"`
}

func (o OrmConf) String() string {
	return `orm`
}

type SqliteConf struct {
	User           string `yaml:"user"`
	Password       string `yaml:"password"`
	DbName         string `yaml:"db_name"`
	Host           string `yaml:"host"`
	ConnectTimeout int    `yaml:"connect_timeout"`
}

func (s SqliteConf) String() string {
	return `sqlite`
}

func (s *ServerConf) GetUri() string {
	return ""
}

func (s *ServerConf) GetDB() string {
	return s.Orm.Sqlite.DbName
}
