package danggo

import (
	"fmt"
	"github.com/hsjsjsj009/danggo/response"
	"log"
	"net/http"
	"strings"
)

type app struct {
	routeHandler map[string]*handler
}

func NewApp() *app{
	return &app{routeHandler: map[string]*handler{}}
}

func (a *app) MainRoute(path string) *route{
	return &route{a,path}
}

func (a *app) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	var (
		responseObj response.Response
		statusCode int
	)
	path := r.URL.Path
	method := strings.ToUpper(r.Method)
	pathVar,found ,handlerObj:= ParsePath(path,a.routeHandler)
	if !found {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte("404 Not Found"))
		PrintLog(method,path,http.StatusNotFound)
		return
	}

	requestObj := newRequest(pathVar,r)
	handler := handlerObj.handlerFunc[method]

	if handler == nil {
		PrintLog(method,path,http.StatusNotFound)
		w.WriteHeader(404)
		_, _ = w.Write([]byte(fmt.Sprintf("Path %s doesn't accept method %s",path,r.Method)))
		return
	}

	responseObj = handler(requestObj)

	responseObj.WriteHeader(w)

	data,err := responseObj.WriteResponse(w)

	if data == nil && err == nil { // Redirect
		if statusCode = responseObj.GetStatusCode(); statusCode == 0 {
			statusCode = http.StatusFound
		}
		PrintLog(method,path,statusCode)
		return
	}

	if err != nil {
		statusCode = http.StatusInternalServerError
		http.Error(w,err.Error(), statusCode)
		log.Fatal(fmt.Sprintf("Error %s\n",err.Error()))
		return
	}

	if statusCode = responseObj.GetStatusCode(); statusCode != 0 {
		w.WriteHeader(statusCode)
	}else{
		statusCode = 200
	}

	_, _ = w.Write(data)

	PrintLog(method,path,statusCode)
}

func PrintLog(method string, path string, code int){
	log.Println(fmt.Sprintf("[%s] %s status code %d",method,path,code))
}

func (a *app) Start(url string) {
	log.Println(fmt.Sprintf("Server run in localhost port %s",url))
	log.Fatal(http.ListenAndServe(url, a))
}
