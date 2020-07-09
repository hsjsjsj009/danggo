package response

import (
	"encoding/json"
	"net/http"
)

type jsonResponse struct {
	responseObject
	Data map[string]interface{} `json:"data"`
}

func JsonResponse(data map[string]interface{}) jsonResponse{
	return jsonResponse{Data: data}
}

func (j jsonResponse) AddHeader(key string,value string) Response {
	if j.header == nil {
		j.header = map[string]string{}
	}
	j.header[key] = value
	return j
}

func (j jsonResponse) WriteResponse(w http.ResponseWriter) ([]byte,error) {
	js,err := json.Marshal(j)

	if err != nil {
		return nil,err
	}

	w.Header().Set("Content-Type", "application/json")

	return js,nil
}
