package main

import (
	. "bookTrade-backend/app"
	"bookTrade-backend/app/router"
	"bookTrade-backend/conf"
	"bookTrade-backend/dao"
	"fmt"
	log "github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("Hello World!")
	config := conf.InitConfig()
	if err := dao.InitDatabase(config); err != nil {
		log.WithError(err).Fatal("Init database failed")
	}
	InitApp(config)
	if err := router.InitRouter(App.Router); err != nil {
		log.WithError(err).Fatal("Init router failed")
	}
	App.Run()
}
