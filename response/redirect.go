package response

import (
	"net/http"
)

type redirect struct {
	responseObject
	url string
	statusCode int
	req *http.Request
}

func (r redirect) WriteResponse(writer http.ResponseWriter) ([]byte, error) {
	statusCode := r.statusCode
	if statusCode == 0 {
		statusCode = http.StatusFound
	}
	http.Redirect(writer,r.req,r.url,statusCode)
	return nil,nil
}

func Redirect(url string,request *http.Request) redirect {
	return redirect{url: url,req: request}
}


