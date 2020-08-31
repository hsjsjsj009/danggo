package danggo

import (
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type requestSuite struct {
	suite.Suite
	variable map[string]interface{}
	req *http.Request
}

func (r *requestSuite) SetupTest() {
	r.variable = map[string]interface{}{
		"asd":"asd",
	}

	r.req = httptest.NewRequest("GET","/asd",nil)
}

func (r *requestSuite) TestNewRequest() {
	obj := newRequest(r.variable,r.req)

	r.Equal(r.variable,obj.pathVariable)
	r.Equal(r.req,obj.Request)
}

func (r *requestSuite) TestGetVariablePath() {
	obj := newRequest(r.variable,r.req)

	r.Equal("asd",obj.pathVariable["asd"])
}

func TestRequest(t *testing.T) {
	suite.Run(t,new(requestSuite))
}