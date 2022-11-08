package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (d *Data) UpdatingLog(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}
	fmt.Println("Updating...\n")
	res, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Failed read req body : %v -- ", err)
		return
	}
	tmp := make(map[string]any)
	err = json.Unmarshal(res, &tmp)
	service, ok := tmp["service"].(string)
	if !ok {
		fmt.Printf("Failed to convert service field or is does not exist : %v -- ", err)
		return
	}
	if d.Con.CheckConf(res) {
		fmt.Printf("CHECK ERROR : config exist in table")
		return
	}

	err = d.Con.UpdateService(res, service)
	if err != nil {
		fmt.Printf("UPDATE SERVICE DATA ERROR  : %v", err)
		return
	}
	fmt.Println("\n\n...Updated!")

}
