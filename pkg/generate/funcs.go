package generate

import (
	"html/template"
	"strings"
)

func trimScheme(url string) (trimmed string) {
	trimmed = strings.TrimPrefix(url, "https://")
	trimmed = strings.TrimPrefix(trimmed, "http://")
	trimmed = strings.TrimSuffix(trimmed, "/")
	return
}

func getRandomColor(i int) (color string) {
	colors := []string{"#79addc", "#ffc09f", "#ffee93", "#fcf5c7", "#adf7b6"}
	return colors[i%len(colors)]
}

func checkStringNil(val *string) bool {
	if val == nil {
		return true
	}
	return false
}

// FuncMap is the default func map that comes w/ generator
var FuncMap = template.FuncMap{
	"Title":      strings.Title,
	"TrimScheme": trimScheme,
	"GetColor":   getRandomColor,
	"CheckSkip":  checkStringNil,
}
