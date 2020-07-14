package danggo

import (
	"github.com/hsjsjsj009/danggo/response"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type appSuite struct {
	suite.Suite
	req *http.Request
	redirect *http.Request
	notFound *http.Request
	statusCode *http.Request
	wrongMethod *http.Request
	writer 	*httptest.ResponseRecorder
	handler map[string]*handler
	app *app
}

func (a *appSuite) SetupTest() {
	a.req = httptest.NewRequest("GET","/asd",nil)
	a.wrongMethod = httptest.NewRequest("POST","/asd",nil)
	a.redirect = httptest.NewRequest("GET","/redirect",nil)
	a.notFound = httptest.NewRequest("GET","/notfound",nil)
	a.statusCode = httptest.NewRequest("GET","/statuscode",nil)
	a.writer = httptest.NewRecorder()
	a.handler = map[string]*handler{
		"/asd":{pathSplit: []string{"asd"},handlerFunc: map[string]func(*HttpRequest) response.Response{
				"GET": func(request *HttpRequest) response.Response {
					return response.HttpResponse("asd").AddHeader("asd","asd")
				},
			},
		},
		"/redirect":{pathSplit: []string{"redirect"},handlerFunc: map[string]func(*HttpRequest) response.Response{
				"GET": func(request *HttpRequest) response.Response {
					return response.Redirect("/asd",request.Request)
				},
			},
		},
		"/statuscode":{pathSplit: []string{"statuscode"},handlerFunc: map[string]func(*HttpRequest) response.Response{
				"GET": func(request *HttpRequest) response.Response {
					return response.HttpResponse("test").SetStatusCode(http.StatusNotFound)
				},
			},
		},
	}
	a.app = &app{routeHandler: a.handler}
}

func (a *appSuite) TestNewApp() {
	obj := NewApp()
	a.Equal(map[string]*handler{},obj.routeHandler)
}

func (a *appSuite) TestMainRoute() {
	route := a.app.MainRoute("/")

	a.Equal(a.app,route.app)
	a.Equal("/",route.route)
}

func (a *appSuite) TestServeHttp() {
	a.app.ServeHTTP(a.writer,a.req)
	a.Equal(200,a.writer.Code)
	a.Equal("asd",a.writer.Body.String())
}

func (a *appSuite) TestServeHttpRedirect() {
	a.app.ServeHTTP(a.writer, a.redirect)
	a.Equal(http.StatusFound, a.writer.Code)
	a.Equal("/asd", a.writer.Header().Get("Location"))
}

func (a *appSuite) TestServeHttpStatusCode() {
	a.app.ServeHTTP(a.writer, a.statusCode)
	a.Equal(http.StatusNotFound,a.writer.Code)
	a.Equal("test",a.writer.Body.String())
}

func (a *appSuite) TestServeHttpNotFound() {
	a.app.ServeHTTP(a.writer,a.notFound)
	a.Equal(http.StatusNotFound,a.writer.Code)
	a.Equal("404 Not Found",a.writer.Body.String())
}

func (a *appSuite) TestServeHttpWrongMethod() {
	a.app.ServeHTTP(a.writer,a.wrongMethod)
	a.Equal(http.StatusNotFound,a.writer.Code)
	a.Equal("Path /asd doesn't accept method POST",a.writer.Body.String())
}


func TestApp(t *testing.T) {
	suite.Run(t,new(appSuite))

}