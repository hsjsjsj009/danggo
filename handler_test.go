package danggo

import (
	"github.com/hsjsjsj009/danggo/response"
	"github.com/stretchr/testify/suite"
	"testing"
)

type handlerSuite struct {
	suite.Suite
	handlerFunc func(*HttpRequest) response.Response
	method string
	handler *handler
}

func (h *handlerSuite) SetupTest() {
	h.handlerFunc = func(request *HttpRequest) response.Response {
		return response.HttpResponse("test")
	}
	h.method = "GET"
	h.handler = &handler{pathSplit: []string{"asd"},handlerFunc: map[string]func(*HttpRequest) response.Response{}}
}

func (h *handlerSuite) TestAddMethod() {
	h.handler.addMethod(h.method,h.handlerFunc)
	_, ok := h.handler.handlerFunc[h.method]
	h.True(ok)
}

func TestHandler(t *testing.T) {
	suite.Run(t,new(handlerSuite))
}


