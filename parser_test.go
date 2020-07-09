package danggo

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParsePathCorrect(t *testing.T){
	assertions := assert.New(t)

	//list := [][]string{
	//	{"asd",":asd"}, // Path => /asd/:asd
	//	{"dfg"}, // Path => /dfg or /dfg/
	//	{"dfg","asd",":dfg"}, // Path => /dfg/asd/:dfg
	//	{"asd",":asd","dfg",":dfg"}, // Path => /asd/:asd/dfg/:dfg
	//}

	list := map[string]*handler{
		"/asd/:asd":{pathSplit: []string{"asd",":asd"}},
		"/dfg":{pathSplit: []string{"dfg"}},
		"/dfg/asd/:dfg":{pathSplit: []string{"dfg","asd",":dfg"}},
		"/asd/:asd/dfg/:dfg":{pathSplit: []string{"asd",":asd","dfg",":dfg"}},
	}

	url := "/asd/asd"

	param,found,handler := ParsePath(url,list)

	output := map[string]string{
		"asd":"asd",
	}

	assertions.Equal(list["/asd/:asd"],handler)
	assertions.Equal(true,found)
	assertions.Equal(fmt.Sprint(output),fmt.Sprint(param))

	url = "/dfg"

	param,found,handler = ParsePath(url,list)

	output = map[string]string{}

	assertions.Equal(list["/dfg"],handler)
	assertions.Equal(true,found)
	assertions.Equal(fmt.Sprint(output),fmt.Sprint(param))

	url = "/dfg/"

	param,found,handler = ParsePath(url,list)

	output = map[string]string{}

	assertions.Equal(list["/dfg"],handler)
	assertions.Equal(true,found)
	assertions.Equal(fmt.Sprint(output),fmt.Sprint(param))

	url = "/dfg/asd/dfg"

	param,found,handler = ParsePath(url,list)

	output = map[string]string{
		"dfg":"dfg",
	}

	assertions.Equal(list["/dfg/asd/:dfg"],handler)
	assertions.Equal(true,found)
	assertions.Equal(fmt.Sprint(output),fmt.Sprint(param))

	url = "/asd/asd/dfg/dfg"

	param,found,handler = ParsePath(url,list)

	output = map[string]string{
		"asd":"asd",
		"dfg":"dfg",
	}

	assertions.Equal(list["/asd/:asd/dfg/:dfg"],handler)
	assertions.Equal(true,found)
	assertions.Equal(fmt.Sprint(output),fmt.Sprint(param))

}

func TestParsePathNotFound(t *testing.T){
	assertions := assert.New(t)

	//list := [][]string{{"asd",":asd"},{"dfg"},{"asd"}}

	list := map[string]*handler{
		"/asd/:asd":{pathSplit: []string{"asd","asd"}},
		"/dfg":{pathSplit: []string{"dfg"}},
		"/asd":{pathSplit: []string{"asd"}},
	}

	var nilHandler *handler

	url := "/dfg/asd/"

	_,found,handler := ParsePath(url,list)

	assertions.Equal(nilHandler,handler)
	assertions.Equal(false,found)

}