package danggo

import (
	"github.com/hsjsjsj009/danggo/response"
	"net/http"
)

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

type handler struct {
	pathSplit []string
	handlerFunc map[string]func(*HttpRequest) response.Response
}

func (h *handler) addMethod(method string,function func(*HttpRequest) response.Response){
	h.handlerFunc[method] = function
}