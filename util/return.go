package util

import (
	"net/http"
	"encoding/json"
)

func JsonReturnOK(j interface{}, w http.ResponseWriter) {
	JsonReturnErr(j, http.StatusOK, w)
	return
}

func JsonReturnErr (j interface{}, status int, w http.ResponseWriter) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	if j != nil {
		bj, err := json.Marshal(j)
		if err != nil {
			panic(err)
		}
		w.Write(bj)
	}
	return
}
