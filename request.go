package danggo

import "net/http"

type HttpRequest struct {
	*http.Request
	pathVariable map[string]interface{}
	Writer http.ResponseWriter
}

func newRequest(pathVar map[string]interface{},r *http.Request, w http.ResponseWriter) *HttpRequest{
	return &HttpRequest{Request: r,pathVariable: pathVar,Writer: w}
}

func (h *HttpRequest) GetVariablePath(variable string) interface{} {
	return h.pathVariable[variable]
}
