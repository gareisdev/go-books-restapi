package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func ParseBody(request *http.Request, result_interface interface{} ) {
	body, err := ioutil.ReadAll(request.Body);

	if err != nil {
		log.Fatal("[ERROR] ParseBody failed")
		return
	}

	err = json.Unmarshal([]byte(body), result_interface)
	
	if err != nil {
		log.Fatal("[ERROR] JSON Unmarshal failed")
		return
	}

	
}