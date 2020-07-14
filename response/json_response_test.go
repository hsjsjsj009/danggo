package response

import (
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"net/http/httptest"
	"testing"
)

type jsonSuite struct {
	suite.Suite
	writer *httptest.ResponseRecorder
	data map[string]interface{}
}

func (j *jsonSuite) SetupTest(){
	j.writer = httptest.NewRecorder()
	j.data = map[string]interface{}{
		"asd":"asdasd",
		"asd123":[]int{1,2,3},
	}
}

func (j *jsonSuite) TestJsonResponse_WriteResponse() {
	response := jsonResponse{Data: j.data}
	res, err := response.WriteResponse(j.writer)
	js,err1 := json.Marshal(map[string]interface{}{
		"data":j.data,
	})

	j.Equal(err1,err)
	j.Equal(js,res)
	j.Equal(JSON,j.writer.Header().Get("Content-Type"))
}

func (j *jsonSuite) TestJsonResponse_SetHeader() {
	obj := JsonResponse(j.data)
	obj1 := obj.AddHeader("asd","asd")

	j.Equal(obj,obj1)
	j.Equal("asd",obj.header["asd"])

	obj1 = obj.AddHeader("dfg","dfg")
	j.Equal("asd",obj.header["asd"])
	j.Equal("dfg",obj.header["dfg"])
}

func TestJsonResponse(t *testing.T) {
	suite.Run(t,new(jsonSuite))
}
