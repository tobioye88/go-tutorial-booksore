package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(request *http.Request, x interface{}) {
	// json.NewDecoder(request.Body).Decode(&x)
	if body, err := ioutil.ReadAll(request.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err == nil {
			return
		}
	}
}
