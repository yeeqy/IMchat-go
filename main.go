package main

import (
	"IM-chat/api"
	"IM-chat/config"
	"IM-chat/router"
	"log"
)

func main() {
	log.SetFlags(log.Llongfile)

	config.Init()

	go api.Manage.Run()

	r := router.NewRouter()

	_ = r.Run("127.0.0.1:8888")
}
