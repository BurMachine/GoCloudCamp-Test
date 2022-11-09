package main

import (
	"burmachine/goCloudCamp/config"
	"burmachine/goCloudCamp/handlers"
	"context"
	"flag"
	"fmt"
	mux2 "github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
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

	cOpts := options.Client().ApplyURI(conf.MongoUrlConnection)
	if err != nil {
		fmt.Printf("[MAIN]: %v", err)
		return
	}
	mongoClient, err := mongo.Connect(context.Background(), cOpts)
	if err != nil {
		log.Fatal(err)
	}

	defer mongoClient.Disconnect(context.Background())
	if err = mongoClient.Ping(context.TODO(), readpref.Primary()); err != nil {
		// Can't connect to Mongo server
		log.Fatal(err)
	}
	mCollection := mongoClient.Database("configurations").Collection("services_configs")

	handlerData := new(handlers.HandlerData)
	handlerData.MongoCollection = mCollection

	mux := mux2.NewRouter()
	mux = handlerData.ComposeHandlers(mux)

	http.ListenAndServe(conf.Addr, mux)
}
