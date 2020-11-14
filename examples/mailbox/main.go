package main

import (
	"log"

	"github.com/yang20150702/base-microservice-web-project-layout/framework"
	"github.com/yang20150702/base-microservice-web-project-layout/internal/configs"
	"github.com/yang20150702/base-microservice-web-project-layout/internal/services/mailbox/storage"

	_ "github.com/yang20150702/base-microservice-web-project-layout/internal/services/mailbox"
)

func main() {
	if err := storage.Init(&configs.ServerConf{}); err != nil {
		log.Fatal(err)
		return
	}
	router := framework.GlobalEngine

	router.Run(":8080")
}
