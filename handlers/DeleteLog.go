package handlers

import (
	"fmt"
	"net/http"
	"strconv"
)

func (d *Data) DeleteLog(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		fmt.Printf("METHOD ERROR")
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		fmt.Printf("ATOI ERROR : %v", err)
		return
	}
	//if id == "" {
	//	fmt.Printf("SERVICE QUERY PARAMETR DOES NOT EXIST")
	//	return
	//}
	err = d.Con.Delete(id)
	if err != nil {
		fmt.Printf("ATOI ERROR%v", err)
		return
	}

}
