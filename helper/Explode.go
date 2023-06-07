package helper

import (
	"strings"
)
package helper

import (
	"strings"
)

func Explode(text string,delimetr string) []string {
	if(len(delimetr)> len(text)) {
		return strings.Split(delimetr,text)
	} else {
		return strings.Split(text,delimetr)
	}
}
func Explode(text string,delimetr string) []string {
	if(len(delimetr)> len(text)) {
		return strings.Split(delimetr,text)
	} else {
		return strings.Split(text,delimetr)
	}
}