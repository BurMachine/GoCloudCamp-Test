package handlers

import (
	"burmachine/configService/postgres"
	"encoding/json"
	"fmt"
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
	err = d.Con.Insert(res, service)
	if err != nil {
		fmt.Printf("INSERT ERROR  : %v", err)
		return
	}
	fmt.Println("\n\nCreated!")
}

func (d *Data) ComposeHandlers(router *http.ServeMux) *http.ServeMux {
	router.HandleFunc("/config", d.CreateLog)

	return router
}
