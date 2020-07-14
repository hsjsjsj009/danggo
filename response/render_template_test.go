package response

import (
	"bytes"
	"github.com/stretchr/testify/suite"
	"html/template"
	"net/http/httptest"
	"testing"
)

type renderSuite struct {
	suite.Suite
	template string
	variable interface{}
}

func (r *renderSuite) SetupTest() {
	r.template = "test.html"
	r.variable = struct {
		One int
		Two string
	}{1,"Two"}
}

func (r *renderSuite) TestRenderTemplate() {
	obj := RenderTemplate(r.template,r.variable)

	t, err := template.ParseFiles(r.template)

	r.Equal(err,obj.err)
	r.Equal(t,obj.htmlTemplate)
	r.Equal(r.variable,obj.variable)
}

func (r *renderSuite) TestRenderTemplate_WriteResponse() {
	t, err := template.ParseFiles(r.template)
	obj := &renderTemplate{variable: r.variable,htmlTemplate: t,err: err}
	writer := httptest.NewRecorder()
	buffer := new(bytes.Buffer)
	err1 := t.Execute(buffer, r.variable)
	res, err2 := obj.WriteResponse(writer)

	r.Equal(err1,err2)
	r.Equal(buffer.Bytes(),res)
	r.Equal(HTML,writer.Header().Get("Content-Type"))
}

func (r *renderSuite) TestRenderTemplate_SetHeader() {
	obj := RenderTemplate(r.template,r.variable)
	obj1 := obj.AddHeader("asd","asd")

	r.Equal(obj,obj1)
	r.Equal("asd",obj.header["asd"])

	obj1 = obj.AddHeader("dfg","dfg")
	r.Equal("asd",obj.header["asd"])
	r.Equal("dfg",obj.header["dfg"])
}

func TestRenderTemplate(t *testing.T) {
	suite.Run(t,new(renderSuite))
}
