package response

import (
	"fmt"
	"net/http"
)


type httpResponse struct {
	responseObject
	data interface{}
}

func HttpResponse(data interface{}) *httpResponse {
	obj := &httpResponse{data: data}
	obj.setChild(obj)
	return obj
}

func (h *httpResponse) WriteResponse(w http.ResponseWriter) ([]byte, error) {
	if w.Header().Get("Content-Type") == "" {
		w.Header().Set("Content-Type",HTML)
	}
	return []byte(fmt.Sprint(h.data)),nil
}






