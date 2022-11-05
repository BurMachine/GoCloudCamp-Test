package main

import (
	"burmachine/configService/config"
	"burmachine/configService/handlers"
	"burmachine/configService/postgres"
	"flag"
	"log"
	"net/http"
)

func main() {
	cfgPath := flag.String("config", "./config.yaml", "Path to yaml configuration file")

	flag.Parse()

	if *cfgPath == "" {
		log.Fatalln("Path to configuration file was not provided")
	}
	conf := config.NewConfigStruct()
	err := conf.LoadConfig(*cfgPath)
	if err != nil {
		log.Fatalln("Config loading error")
	}

	con := postgres.NewConnStruct(conf.DbUrl)
	var data handlers.Data
	data.Con = con

	err = data.Con.InitDbTables()
	if err != nil {
		log.Println(err)
	}

	mux := http.NewServeMux()
	mux = data.ComposeHandlers(mux)
	logMux := handlers.MiddlewareLog(mux)

	http.ListenAndServe(conf.Addr, logMux)
}
