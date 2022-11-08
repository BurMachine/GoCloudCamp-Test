package handlers

import (
	"fmt"
	"net/http"
)

func (d *Data) GetLog(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		return
	}
	service := r.URL.Query().Get("service")
	if service == "" {
		fmt.Printf("SERVICE QUERY PARAMETR DOES NOT EXIST")
		return
	}
	json, err := d.Con.GetConfFromService(service)
	if err != nil {
		fmt.Printf("SELECT SERVICE CON ERROR : %v", err)
		return
	}
	_, err = w.Write(json)
	if err != nil {
		fmt.Printf("WRITE ERROR : %v", err)
		return
	}
}
