package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(r *http.Request, result interface{}){
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(result)
	PanicError(err)
}

func WriteToResponseBody(w http.ResponseWriter, response interface{}){
	w.Header().Add("Content-type","application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(response)
	PanicError(err)
}