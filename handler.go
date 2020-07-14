package danggo

import (
	"github.com/hsjsjsj009/danggo/response"
)

type handler struct {
	pathSplit []string
	handlerFunc map[string]func(*HttpRequest) response.Response
}

func (h *handler) addMethod(method string,function func(*HttpRequest) response.Response){
	h.handlerFunc[method] = function
}