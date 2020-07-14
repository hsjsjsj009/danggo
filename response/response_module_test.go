package response

import (
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type responseSuite struct {
	suite.Suite
	statusCode int
	response *responseObject
	header map[string]string
}

func (r *responseSuite) SetupTest() {
	r.statusCode = http.StatusOK
	r.header = map[string]string{
		"asd":"123",
	}
	r.response = &responseObject{
		statusCode: r.statusCode,
		child: &httpResponse{},
		header: r.header,
	}
}

func (r *responseSuite) TestResponse_WriteResponse() {
	writer := httptest.NewRecorder()
	t, err := r.response.WriteResponse(writer)

	r.Equal(nil,err)
	r.Equal([]byte{},t)
}

func (r *responseSuite) TestResponseObject_GetStatusCode() {
	r.Equal(http.StatusOK,r.response.statusCode)
}

func (r *responseSuite) TestResponseObject_AddHeader() {
	obj := r.response.AddHeader("asd","asd")
	r.Equal(obj,r.response.child)
	r.Equal("asd",r.response.header["asd"])
}

func (r *responseSuite) TestResponseObject_SetStatusCode() {
	obj := r.response.SetStatusCode(http.StatusNotFound)
	r.Equal(obj,r.response.child)
	r.Equal(http.StatusNotFound,r.response.statusCode)
}

func (r *responseSuite) TestResponseObject_SetHeader() {
	header := map[string]string{
		"asd":"asd",
	}
	obj := r.response.SetHeader(header)
	r.Equal(obj,r.response.child)
	r.Equal(header,r.response.header)
}

func (r *responseSuite) TestResponseObject_WriteHeader() {
	writer := httptest.NewRecorder()
	r.response.WriteHeader(writer)
	r.Equal(r.header["asd"],writer.Header().Get("asd"))
}

func TestResponseObject(t *testing.T) {
	suite.Run(t,new(responseSuite))
}