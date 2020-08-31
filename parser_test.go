package danggo

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckVariableType(t *testing.T) {
	assertions := assert.New(t)

	testCase := []struct{
		variable string
		correct bool
		errorDataType string
	}{
		{
			variable: "<test:int>",
			correct: true,
			errorDataType: "",
		},
		{
			variable: "<test:str>",
			correct: true,
			errorDataType: "",
		},
		{
			variable: "<test:bool>",
			correct: true,
			errorDataType: "",
		},
		{
			variable: "<test:asdsad>",
			correct: false,
			errorDataType: "asdsad",
		},
	}

	for _,test := range testCase {
		correct, dataType := checkVariableType(test.variable)
		assertions.Equal(test.correct,correct)
		assertions.Equal(test.errorDataType,dataType)
	}

}

func TestParsePath(t *testing.T){
	assertions := assert.New(t)

	//list := [][]string{
	//	{"asd",":asd"}, // Path => /asd/:asd
	//	{"dfg"}, // Path => /dfg or /dfg/
	//	{"dfg","asd",":dfg"}, // Path => /dfg/asd/:dfg
	//	{"asd",":asd","dfg",":dfg"}, // Path => /asd/:asd/dfg/:dfg
	//}

	list := map[string]*handler{
		"/asd/<asd>":{pathSplit: []string{"asd","<asd>"}},
		"/dfg":{pathSplit: []string{"dfg"}},
		"/dfg/asd/<dfg>":{pathSplit: []string{"dfg","asd","<dfg>"}},
		"/asd/<asd>/dfg/<dfg>":{pathSplit: []string{"asd","<asd>","dfg","<dfg>"}},
		"/dasd/<number:int>":{pathSplit: []string{"dasd","<number:int>"}},
		"/sdfsdfsf/<text:str>":{pathSplit: []string{"sdfsdfsf","<text:str>"}},
	}

	testCase := []struct{
		path string
		url string
		err error
		founded bool
		param map[string]interface{}
	}{
		{
			path: "/asd/<asd>",
			url: "/asd/asd",
			err: nil,
			founded: true,
			param: map[string]interface{}{
				"asd":"asd",
			},
		},
		{
			path: "/dfg",
			url: "/dfg",
			err: nil,
			founded: true,
			param: map[string]interface{}{},
		},
		{
			path: "/dfg",
			url: "/dfg/",
			err: nil,
			founded: true,
			param: map[string]interface{}{},
		},
		{
			url: "/dfg/asd/dfg",
			path: "/dfg/asd/<dfg>",
			err: nil,
			founded: true,
			param: map[string]interface{}{
				"dfg":"dfg",
			},
		},
		{
			url: "/asd/asd/dfg/dfg",
			path: "/asd/<asd>/dfg/<dfg>",
			param: map[string]interface{}{
				"asd":"asd",
				"dfg":"dfg",
			},
			err: nil,
			founded: true,
		},
		{
			url: "/dfg/asd/",
			path: "",
			err: nil,
			founded: false,
		},
		{
			url: "/dasd/12/",
			path: "/dasd/<number:int>",
			err: nil,
			founded: true,
			param: map[string]interface{}{
				"number":12,
			},
		},
	}

	var nilHandler *handler

	for _, test := range testCase {
		param,found,handler, err := ParsePath(test.url,list)

		assertions.Equal(err, test.err)
		if test.path != "" {
			assertions.Equal(list[test.path],handler)
			assertions.Equal(fmt.Sprint(test.param),fmt.Sprint(param))
		}else{
			assertions.Equal(nilHandler,handler)
		}
		assertions.Equal(test.founded,found)
	}
}