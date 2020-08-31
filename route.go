package danggo

import (
	"fmt"
	"github.com/hsjsjsj009/danggo/response"
	"strings"
)

type route struct {
	app *app
	route string
}

func checkPathVariable(in []string){
	for _,variable := range in {
		if string(variable[0]) == "<" && string(variable[len(variable)-1]) == ">"{
			if correct,dataType := checkVariableType(variable); !correct {
				panic(fmt.Sprintf("Unsupported type %s in %s",dataType,variable))
			}
		}
	}
}

func (r *route) Handle(function func(*HttpRequest) response.Response, method ...string) {
	app := r.app
	routeHandler,ok := app.routeHandler[r.route]
	if !ok {
		pathSplit := PathToSlice(r.route)
		checkPathVariable(pathSplit)
		app.routeHandler[r.route] = &handler{handlerFunc: map[string]func(*HttpRequest) response.Response{},pathSplit: pathSplit}
		routeHandler = app.routeHandler[r.route]
	}

	for _, reqMethod := range method {
		reqMethod = strings.ToUpper(reqMethod)
		routeHandler.addMethod(reqMethod,function)
	}
}

func (r *route) Get(function func(*HttpRequest) response.Response){
	r.Handle(function,"GET")
}

func (r *route) Post(function func(*HttpRequest) response.Response){
	r.Handle(function,"POST")
}

func (r *route) Put(function func(*HttpRequest) response.Response){
	r.Handle(function,"PUT")
}

func (r *route) Delete(function func(*HttpRequest) response.Response){
	r.Handle(function,"DELETE")
}

func (r *route) SubRoute(path string) *route {
	newRoute := r.route+path
	if r.route == "/" {
		newRoute = path
	}
	return &route{app: r.app,route: newRoute}
}