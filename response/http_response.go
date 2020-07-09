package response

const JSON = "application/json"
const HTML = "text/html"


type httpResponse struct {
	responseObject
	data interface{}
	contentType string
}





