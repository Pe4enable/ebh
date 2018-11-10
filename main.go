package main

import (
	"flag"
	"github.com/BankEx/ebh/config"
	"github.com/BankEx/ebh/handlers"
	"github.com/BankEx/ebh/router"
	"log"
	"net/http"
	"github.com/BankEx/ebh/repositories"
)

func main() {
	var configPathPtr string

	flag.StringVar(&configPathPtr, "config-path", "./config/config.yml", "A path to config file")
	flag.Parse()

	// Config
	conf, err := config.LoadConfig(configPathPtr)
	if err != nil {
		panic(err)
	}

	//nodeService, err := services.NewNodeReader(conf)
	//if err != nil {
	//	panic(err)
	//}
	//log.Printf("ETHService is initialised")

	//cache := make(chan ratestates.RateState, conf.Cache)
	//err = btcService.ConnectToDB()
	//if err != nil {
	//	panic(err)
	//}
	//defer btcService.CloseDBConnection()
	//
	//err = btcService.StartListenBTC()
	//if err != nil {
	//	panic(err)
	//}
	//defer btcService.StopListenBTC()

	mongo, err := repositories.New(conf.DBConfig)
	if err != nil {
		panic(err)
	}
	log.Printf("MongoWriter is initialised")
	//go mongo.Start(cache)

	handlers := handlers.New(nil, mongo)
	r := router.New(handlers)

	log.Printf("Server is listening on %s port", conf.Port)
	log.Panic(http.ListenAndServe(conf.Port, r))
}
