package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(req *http.Request, parsed_content interface{}) {
	if body, err := ioutil.ReadAll(req.Body); err == nil {
		if err := json.Unmarshal([]byte(body), parsed_content); err != nil {
			return
		}
	}

}
