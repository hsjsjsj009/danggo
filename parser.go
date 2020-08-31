package danggo

import (
	"strconv"
	"strings"
)

func ParsePath(path string,listHandler map[string]*handler) (map[string]interface{},bool,*handler, error) {
	var (
		param map[string]interface{}
		handler *handler
		err error
	)
	pathSplit := PathToSlice(path)

	found := false

	for _,obj := range listHandler {
		pathSlice := obj.pathSplit
		param,found,err = checkPathConfig(pathSplit,pathSlice)
		if err != nil {
			return nil, found, nil, err
		}
		if found {
			handler = obj
			break
		}
	}

	return param,found,handler,nil
}

func checkPathConfig(path []string,pattern []string) (map[string]interface{},bool,error){
	if len(path) != len(pattern){
		return nil,false,nil
	}

	data := map[string]interface{}{}

	same := true

	for idx,pathString := range path {
		patternIdx := pattern[idx]

		if patternIdx == pathString {
			continue
		}

		if string(patternIdx[0]) == "<" && string(patternIdx[len(patternIdx)-1]) == ">" && len(pathString) != 0 {
			variable,processedData,err := processVariable(clearBrackets(patternIdx),pathString)
			if err != nil {
				return nil, false, err
			}
			data[variable] = processedData
			continue
		}

		same = false
		break

	}

	return data,same,nil
}

func PathToSlice(path string) []string{
	path = strings.Replace(path,"/","",1)
	path = strings.ReplaceAll(path,"/"," ")
	pathSplit := strings.Split(path," ")
	lengthPath := len(pathSplit)
	if pathSplit[lengthPath - 1] == "" {
		pathSplit = pathSplit[:lengthPath - 1]
	}

	return pathSplit
}

func checkVariableType(variable string) (bool,string) {
	processedVariable := clearBrackets(variable)
	if strings.Contains(processedVariable,":") {
		splitVariable := strings.Split(processedVariable,":")
		switch dataType := splitVariable[1]; dataType {
		case "int",
			"bool",
			"str":
				return true,""
		default:
			return false,dataType
		}
	}
	return true,""
}

func clearBrackets(in string) string {
	length := len(in)
	return in[1:length-1]
}

func processVariable(variable string, path string) (string,interface{},error) {
	if strings.Contains(variable,":") {
		var (
			err error
			convertedData interface{}
		)
		variableArr := strings.Split(variable,":")
		dataType := variableArr[1]
		pathVar := variableArr[0]
		switch dataType {
		case "int" :
			convertedData, err = strconv.ParseInt(path,10,64)
		case "str" :
			convertedData = path
		case "bool":
			convertedData, err = strconv.ParseBool(path)
		}
		if err != nil {
			return "", nil, WrongTypeError(dataType)
		}
		return pathVar, convertedData,nil
	}

	return variable,path,nil
}