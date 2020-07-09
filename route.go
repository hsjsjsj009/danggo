package danggo

import (
	"github.com/hsjsjsj009/danggo/response"
	"strings"
)

type route struct {
	app *app
	route string
}

func (r *route) Handle(function func(*HttpRequest) response.Response, method ...string) {
	app := r.app
	routeHandler := app.routeHandler[r.route]
	if routeHandler == nil {
		pathSplit := PathToSlice(r.route)
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
	r.Handle(function,"Delete")
}

func (r *route) SubRoute(path string) *route {
	return &route{app: r.app,route: r.route+path}
}