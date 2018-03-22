package util

import "encoding/json"
import "net/http"

func GetJsonBody(r *http.Request, out interface{}) error {
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&out)
	if err != nil {
		return err
	}
	return nil
}

func GetRequestForm(r *http.Request, valueName, defaultValue string) string {
	r.ParseForm()
	if _, ok := r.Form[valueName];ok {
		return r.Form.Get(valueName)
	} else {
		return defaultValue
	}
}