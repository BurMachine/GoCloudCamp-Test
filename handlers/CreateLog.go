package handlers

import (
	"burmachine/configService/postgres"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

type Data struct {
	Con *postgres.ConnectData
}

func (d *Data) CreateLog(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}
	fmt.Println("Creating...\n")
	res, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Failed read req body : %v", err)
		return
	}
	tmp := make(map[string]any)
	err = json.Unmarshal(res, &tmp)
	service, ok := tmp["service"].(string)
	if !ok {
		fmt.Printf("Failed to convert service field or is does not exist : %v", err)
		return
	}
	if d.Con.CheckConf(res) {
		fmt.Printf("CHECK ERROR : config exist in table")
		return
	}
	err = d.Con.InsertConf(res, service)
	if err != nil {
		fmt.Printf("INSERT ERROR  : %v", err)
		return
	}
	err = d.Con.UpdateService(res, service)
	if err != nil {
		fmt.Printf("UPDATE SERVICE DATA ERROR  : %v", err)
		return
	}
	fmt.Println("\n\nCreated!")
}

func (d *Data) ComposeHandlers(router *mux.Router) *mux.Router {
	router.HandleFunc("/config", d.CreateLog).Methods("POST")
	router.HandleFunc("/update", d.UpdatingLog).Methods("POST")

	router.HandleFunc("/config", d.GetLog).Methods("GET")
	router.HandleFunc("/delete", d.DeleteLog).Methods("POST")
	return router
}
