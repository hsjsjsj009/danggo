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
	pathVar,found ,handlerObj, err:= ParsePath(path,a.routeHandler)
	if err != nil {
		writeError(
			err.Error(),
			method,
			path,
			http.StatusNotFound,
			w)
		return
	}
	if !found {
		writeError(
			"404 Not Found",
			method,
			path,
			http.StatusNotFound,
			w)
		return
	}

	requestObj := newRequest(pathVar,r,w)
	handler := handlerObj.handlerFunc[method]

	if handler == nil {
		writeError(
			fmt.Sprintf("Path %s doesn't accept method %s",path,r.Method),
			method,
			path,
			http.StatusNotFound,
			w)
		return
	}

	responseObj = handler(requestObj)

	responseObj.WriteHeader(w)

	data,err := responseObj.WriteResponse(w)

	if data == nil && err == nil { // Redirect
		if statusCode = responseObj.GetStatusCode(); statusCode == 0 {
			statusCode = http.StatusFound
		}
		printLog(method,path,statusCode)
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

	printLog(method,path,statusCode)
}

func printLog(method string, path string, code int){
	log.Println(fmt.Sprintf("[%s] %s status code %d",method,path,code))
}

func writeError(message string,method string, path string, code int, w http.ResponseWriter) {
	w.WriteHeader(code)
	_, _ = w.Write([]byte(message))
	printLog(method,path,code)
}

func (a *app) Start(url string, plugins ...func(http.Handler) http.Handler) {
	log.Println(fmt.Sprintf("Server run in localhost port %s",url))
	var handler http.Handler = a
	if len(plugins) != 0 {
		handler = plugins[0](a)
		for _,plugin := range plugins[1:] {
			handler = plugin(handler)
		}
	}
	log.Fatal(http.ListenAndServe(url, handler))
}
