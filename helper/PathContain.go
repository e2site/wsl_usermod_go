package helper

import "os"


func PathContain(pathSearch string,pathPart string) bool {
	var result bool = true
	var listPathSearch = Explode(pathSearch,string(os.PathSeparator))
	var listPathPart = Explode(pathPart,string(os.PathSeparator))

	for indSearch,valSearch := range listPathSearch {
		if(!result) {
			break
		}
		if(indSearch>=len(listPathPart)) {
			break
		}
		if(len(valSearch)!=len(listPathPart[indSearch])) {
			result = false
			break
		}
		for i:=0; i < len(valSearch);i++{
			if(valSearch[i] != listPathPart[indSearch][i]) {
				result = false
				break
			} 
		}
	}

	return result
}