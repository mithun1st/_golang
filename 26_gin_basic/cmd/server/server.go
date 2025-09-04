package server

import (
	"gin-basic/internal/config"
	"gin-basic/internal/database"
	"gin-basic/internal/router"
	"gin-basic/pkg/enum"
	"log"
)

func init() {

}

func Run(enviroment enum.EnviromentE) {

	//* Setup Enviroment
	config.SetEnviroment(enviroment)

	//* Load Config
	appConfig, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	//* Initialize Database
	db, err := database.NewPostgresDb(appConfig.Db)
	if err != nil {
		log.Fatal(err)
	}

	//* DB Checking
	err = database.Check(db)
	if err != nil {
		log.Fatal(err)
	} else {
		defer database.Close(db)
	}

	//* Setup Route
	router := router.SetupRoute(db)

	//* Server Running
	err = router.Run(appConfig.Server.IP + ":" + appConfig.Server.Port)
	if err != nil {
		log.Fatal(err)
	}

}
