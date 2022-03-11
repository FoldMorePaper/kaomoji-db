package config

import (
	"fmt"
	"os"

	"github.com/tkanos/gonfig"
)

type Configuration struct {
	port          int
	db_mongo_port int `env:"APP_PORT"`
	db_mongo_url  string
}

const configuration =  {port: 3000, db_mongo_port: 27017, db_mongo_url: ""}

func get_config() Configuration {

	configuration := Configuration{}

	// mock env variable
	os.Setenv("db_mongo_url", "test")
	os.Setenv("APP_PORT", "8081")

	err := gonfig.GetConf("../config/app.cfg.yaml", &configuration)

	if err != nil {
		fmt.Println(err)
		os.Exit(500)
	}
	fmt.Println(configuration.port)
	fmt.Println(configuration.db_mongo_url)
	fmt.Println(configuration.db_mongo_port)

	return configuration
}
