package helpers

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(request *http.Request, result any) any {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(result)
	PanicIfError(err)
	return result
}
