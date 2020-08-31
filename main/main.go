package main

import (
	"github.com/hsjsjsj009/danggo"
	"github.com/hsjsjsj009/danggo/response"
)

func main() {
	//var test = "asdasd"
	//val, err := strconv.ParseInt(test,10,64)
	//fmt.Println(err)
	app := danggo.NewApp()
	app.MainRoute("/<asdasdasd:asdasd>").Get(func(request *danggo.HttpRequest) response.Response {
		return nil
	})
}
