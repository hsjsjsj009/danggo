package response

import (
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type redirectSuite struct {
	suite.Suite
	req      *http.Request
}

func (r *redirectSuite) SetupTest() {
	r.req = httptest.NewRequest("","/asdt",nil)
}

func (r *redirectSuite) TestRedirect() {
	obj := Redirect("/asd",r.req)

	r.Equal(r.req,obj.req)
	r.Equal("/asd",obj.url)
}

func (r *redirectSuite) TestRedirect_WriteResponse() {
	obj := redirect{url: "/asd",req: r.req}
	writer := httptest.NewRecorder()
	res, err := obj.WriteResponse(writer)

	var temp []byte
	r.Equal(nil,err)
	r.Equal(temp,res)
	r.Equal("/asd",writer.Header().Get("Location"))
	r.Equal(302,writer.Code)
}

func (r *redirectSuite) TestRedirect_SetHeader() {
	obj := Redirect("/asd",r.req)
	obj1 := obj.AddHeader("asd","asd")

	r.Equal(obj,obj1)
	r.Equal("asd",obj.header["asd"])

	obj1 = obj.AddHeader("dfg","dfg")
	r.Equal("asd",obj.header["asd"])
	r.Equal("dfg",obj.header["dfg"])
}

func TestRedirectSuite(t *testing.T) {
	suite.Run(t,new(redirectSuite))
}
