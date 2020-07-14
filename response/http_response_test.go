package response

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"net/http/httptest"
	"testing"
)

type httpSuite struct {
	suite.Suite
	data interface{}
	response *httptest.ResponseRecorder
}

func (h *httpSuite) SetupTest(){
	h.data = []int{1,2,3}
	h.response = httptest.NewRecorder()
}

func (h *httpSuite) TestHttpResponse_WriteResponse() {
	obj := &httpResponse{data: h.data}

	h.response.Header().Set("Content-Type",JSON)

	res, err := obj.WriteResponse(h.response)

	h.Equal(nil,err)
	h.Equal(JSON,h.response.Header().Get("Content-Type"))
	h.Equal([]byte(fmt.Sprint(h.data)),res)
}

func (h *httpSuite) TestHttpResponse_WriteResponseNoContentType() {
	obj := httpResponse{data: h.data}
	res, err := obj.WriteResponse(h.response)

	h.Equal(nil,err)
	h.Equal(HTML,h.response.Header().Get("Content-Type"))
	h.Equal([]byte(fmt.Sprint(h.data)),res)
}

func (h *httpSuite) TestHttpResponse_SetHeader() {
	obj := HttpResponse(h.data)
	obj1 := obj.SetHeader(map[string]string{
		"asd":"asd",
	})

	h.Equal(obj,obj1)
	h.Equal("asd",obj.header["asd"])
}

func TestHttpResponse(t *testing.T) {
	suite.Run(t,new(httpSuite))
}