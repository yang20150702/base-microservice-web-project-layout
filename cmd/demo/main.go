package main

import (
	"io/ioutil"

	"github.com/yang20150702/base-microservice-web-project-layout/framework"
	"github.com/yang20150702/base-microservice-web-project-layout/internal/configs"
	_ "github.com/yang20150702/base-microservice-web-project-layout/internal/services/mailbox"
	"github.com/yang20150702/base-microservice-web-project-layout/internal/services/mailbox/storage"
	"gopkg.in/yaml.v2"
)

var log = framework.GetLogger()

func main() {
	log.Infoln("主程序启动")
	conf := ParseConf("../../configs/server.yaml")
	if err := storage.Init(conf); err != nil {
		log.Fatal(err)
		return
	}
	router := framework.GlobalEngine

	router.Run(":8080")
}

func ParseConf(name string) *configs.ServerConf {
	f, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}
	conf := new(configs.ServerConf)
	err = yaml.Unmarshal(f, conf)
	if err != nil {
		panic(err)
	}
	log.Infof("conf: %+v", conf)
	return conf
}
