package danggo

import (
	"fmt"
	"github.com/hsjsjsj009/danggo/response"
	"github.com/stretchr/testify/suite"
	"testing"
)

type routeSuite struct {
	suite.Suite
	app *app
	route *route
	routeRoot *route
	method func(*HttpRequest) response.Response
	method1 func(*HttpRequest) response.Response
}

func (r *routeSuite) SetupTest() {
	r.app = NewApp()
	r.route = r.app.MainRoute("/asd")
	r.routeRoot = r.app.MainRoute("/")
	r.method = func(request *HttpRequest) response.Response {
		return response.HttpResponse("asd")
	}
	r.method1 = func(request *HttpRequest) response.Response {
		return response.HttpResponse("123")
	}
}

func (r *routeSuite) TestHandle() {
	r.route.Handle(r.method,"GET","POST")

	r.Equal(fmt.Sprintf("%p",r.method),fmt.Sprintf("%p",r.app.routeHandler["/asd"].handlerFunc["GET"]))
	r.Equal(fmt.Sprintf("%p",r.method),fmt.Sprintf("%p",r.app.routeHandler["/asd"].handlerFunc["POST"]))

	r.route.Handle(r.method1,"GET")
	r.Equal(fmt.Sprintf("%p",r.method1),fmt.Sprintf("%p",r.app.routeHandler["/asd"].handlerFunc["GET"]))
	r.Equal(fmt.Sprintf("%p",r.method),fmt.Sprintf("%p",r.app.routeHandler["/asd"].handlerFunc["POST"]))
}

func (r *routeSuite) TestGet() {
	r.route.Get(r.method)
	r.Equal(fmt.Sprintf("%p",r.method),fmt.Sprintf("%p",r.app.routeHandler["/asd"].handlerFunc["GET"]))

	r.route.Get(r.method1)
	r.Equal(fmt.Sprintf("%p",r.method1),fmt.Sprintf("%p",r.app.routeHandler["/asd"].handlerFunc["GET"]))
}

func (r *routeSuite) TestPost() {
	r.route.Post(r.method)
	r.Equal(fmt.Sprintf("%p",r.method),fmt.Sprintf("%p",r.app.routeHandler["/asd"].handlerFunc["POST"]))

	r.route.Post(r.method1)
	r.Equal(fmt.Sprintf("%p",r.method1),fmt.Sprintf("%p",r.app.routeHandler["/asd"].handlerFunc["POST"]))
}

func (r *routeSuite) TestPut() {
	r.route.Put(r.method)
	r.Equal(fmt.Sprintf("%p",r.method),fmt.Sprintf("%p",r.app.routeHandler["/asd"].handlerFunc["PUT"]))

	r.route.Put(r.method1)
	r.Equal(fmt.Sprintf("%p",r.method1),fmt.Sprintf("%p",r.app.routeHandler["/asd"].handlerFunc["PUT"]))
}

func (r *routeSuite) TestDelete() {
	r.route.Delete(r.method)
	r.Equal(fmt.Sprintf("%p",r.method),fmt.Sprintf("%p",r.app.routeHandler["/asd"].handlerFunc["DELETE"]))

	r.route.Delete(r.method1)
	r.Equal(fmt.Sprintf("%p",r.method1),fmt.Sprintf("%p",r.app.routeHandler["/asd"].handlerFunc["DELETE"]))
}

func (r *routeSuite) TestSubRouteRoot() {
	asd := r.routeRoot.SubRoute("/123")

	r.Equal("/123",asd.route)
}

func (r *routeSuite) TestSubRoute() {
	asd := r.route.SubRoute("/123")

	r.Equal("/asd/123",asd.route)
}

func (r *routeSuite) TestCheckPathVariable() {
	r.Panics(func() {
		r.route.SubRoute("/<asdasdasd:asdasdasd>").Get(r.method)
	},"Unsupported type asdasdasd in <asdasdasd:asdasdasd>")
}

func TestRoute(t *testing.T) {
	suite.Run(t,new(routeSuite))
}
