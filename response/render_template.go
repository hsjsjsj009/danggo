package response

import (
	"bytes"
	"html/template"
	"net/http"
)

type renderTemplate struct {
	responseObject
	htmlTemplate *template.Template
	variable interface{}
	err error
}

func RenderTemplate(path string,variable interface{}) *renderTemplate {
	t,err := template.ParseFiles(path)
	obj := &renderTemplate{htmlTemplate: t,variable: variable,err: err}
	obj.setChild(obj)
	return obj
}

func (r *renderTemplate) WriteResponse(writer http.ResponseWriter) ([]byte, error) {
	if r.err != nil {
		return nil, r.err
	}
	out := new(bytes.Buffer)
	err := r.htmlTemplate.Execute(out, r.variable)
	if err != nil {
		return nil, err
	}
	writer.Header().Set("Content-Type", "text/html")
	return out.Bytes(),nil
}
