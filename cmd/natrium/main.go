package main

import (
	"log"

	"github.com/natrongmbh/natrium/pkg/config"
	"github.com/natrongmbh/natrium/pkg/controllers"
	"github.com/natrongmbh/natrium/pkg/env"
	"github.com/natrongmbh/natrium/pkg/server"
	"github.com/natrongmbh/natrium/pkg/util"
)

var (
	err error
)

func init() {
	err := env.Init()
	if err != nil {
		log.Fatal(err)
	}
	util.InitLogger()
	err = config.Init()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	err = server.Setup()
	if err != nil {
		util.Logger.Error(err)
	}

	bindingController := controllers.BindingController{}
	err = bindingController.Setup()
	if err != nil {
		util.Logger.Error(err)
	}

	// start the pocketbase server
	if err := server.App.Start(); err != nil {
		log.Fatal(err)
	}
}
