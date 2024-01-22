package helper

import (
	"encoding/json"
	"net/http"
)

func BodyToRequest(request *http.Request, result interface{}) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(result)
	PanicError(err)
}

func WritteToBody(writter http.ResponseWriter, response interface{}) {
	writter.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writter)
	err := encoder.Encode(response)
	PanicError(err)

}
