package handlers

import (
	mongoDb "burmachine/goCloudCamp/mongo"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (d *HandlerData) CreateConfigHander(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fmt.Printf("Method error")
		return
	}
	fmt.Printf("[CREATING] ...\n")

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
	err = mongoDb.CreatingInsert(service, res, d.MongoCollection)
	fmt.Printf("[CREATED] ...\n")
}
