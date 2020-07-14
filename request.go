package danggo

import "net/http"

type HttpRequest struct {
	Request *http.Request
	pathVariable map[string]string
}

func newRequest(pathVar map[string]string,r *http.Request) *HttpRequest{
	return &HttpRequest{Request: r,pathVariable: pathVar}
}

func (h *HttpRequest) GetVariablePath(variable string) string {
	return h.pathVariable[variable]
}
