package danggo

import "net/http"

type HttpRequest struct {
	*http.Request
	pathVariable map[string]interface{}
}

func newRequest(pathVar map[string]interface{},r *http.Request) *HttpRequest{
	return &HttpRequest{Request: r,pathVariable: pathVar}
}

func (h *HttpRequest) GetVariablePath(variable string) interface{} {
	return h.pathVariable[variable]
}
