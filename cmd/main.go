package main

import (
	"burmachine/configService/config"
	"flag"
	"fmt"
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
	mux := http.NewServeMux()
	fmt.Println(mux)
}
