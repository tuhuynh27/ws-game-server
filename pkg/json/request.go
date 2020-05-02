package json

import (
	"encoding/json"
	"log"
	"net/http"
)

type Request struct {
	Request *http.Request
}

// Decode request JSON to struct
func (r *Request) GetJSONBody(model interface{}) {
	decoder := json.NewDecoder(r.Request.Body)
	err := decoder.Decode(&model)
	if err != nil {
		log.Println("json decode error: ", err.Error())
	}
}
