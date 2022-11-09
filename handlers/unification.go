package handlers

import (
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type HandlerData struct {
	MongoCollection *mongo.Collection
}

func (d *HandlerData) ComposeHandlers(router *mux.Router) *mux.Router {
	router.HandleFunc("/config", d.CreateConfigHander).Methods("POST")

	return router
}
