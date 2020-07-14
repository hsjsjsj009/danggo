package response

import "net/http"

type responseObject struct {
	header map[string]string
	statusCode int
	child Response
}

func (r *responseObject) setChild(child Response) {
	r.child = child
}

func (r *responseObject) WriteResponse(writer http.ResponseWriter) ([]byte, error) {
	return []byte{},nil
}

func (r *responseObject) WriteHeader(writer http.ResponseWriter) {
	if r.header != nil {
		for key,value := range r.header{
			writer.Header().Set(key,value)
		}
	}
}

func (r *responseObject) SetHeader(m map[string]string) Response {
	r.header = m
	return r.child
}

func (r *responseObject) AddHeader(key string, value string) Response {
	if r.header == nil {
		r.header = map[string]string{}
	}
	r.header[key] = value
	return r.child
}

func (r *responseObject) GetStatusCode() int {
	return r.statusCode
}

func (r *responseObject) SetStatusCode(i int) Response {
	r.statusCode = i
	return r.child
}

type Response interface {
	WriteResponse(http.ResponseWriter) ([]byte,error)
	WriteHeader(http.ResponseWriter)
	SetHeader(map[string]string) Response
	AddHeader(string,string) Response
	GetStatusCode() int
	SetStatusCode(int) Response
}