package handlers

import (
	"fmt"
	"io"
	"net/http"
)

func (d *Data) DeleteLog(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fmt.Printf("METHOD ERROR")
		return
	}
	res, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Failed read req body : %v", err)
		return
	}
	//if id == "" {
	//	fmt.Printf("SERVICE QUERY PARAMETR DOES NOT EXIST")
	//	return
	//}
	err = d.Con.Delete(res)
	if err != nil {
		fmt.Printf("ATOI ERROR%v", err)
		return
	}

}
