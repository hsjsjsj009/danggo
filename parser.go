package danggo

import (
	"strings"
)

func ParsePath(path string,listHandler map[string]*handler) (map[string]string,bool,*handler) {
	var (
		param map[string]string
		handler *handler = nil
	)
	pathSplit := PathToSlice(path)

	found := false

	for _,obj := range listHandler {
		pathSlice := obj.pathSplit
		if param,found = checkPathConfig(pathSplit,pathSlice) ; found {
			handler = obj
			break
		}
	}

	return param,found,handler
}

func checkPathConfig(path []string,pattern []string) (map[string]string,bool){
	if len(path) != len(pattern){
		return nil,false
	}

	data := map[string]string{}

	same := true

	for idx,pathString := range path {
		patternIdx := pattern[idx]

		if patternIdx == pathString {
			continue
		}

		if string(patternIdx[0]) == ":" && len(pathString) != 0{
			patternIdx = strings.Replace(patternIdx,":","",1)
			data[patternIdx] = pathString
			continue
		}

		same = false
		break

	}

	return data,same
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