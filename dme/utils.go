package dme

import (
	"log"
	"strings"
)

func StripQuotes(word string) string {
	if strings.HasPrefix(word, "\"") && strings.HasSuffix(word, "\"") {
		return strings.TrimPrefix(strings.TrimSuffix(word, "\""), "\"")
	} else if word == "{}" {
		word = ""
		return word
	}
	return word
}

func toListOfString(configured interface{}) []string {
	vs := make([]string, 0, 1)
	log.Println(configured.([]interface{}))
	for _, value := range configured.([]interface{}) {
		vs = append(vs, value.(string))
	}
	return vs
}

func toListOfInterface(name interface{}) []interface{} {
	nameList := make([]interface{}, 0)
	nameList = append(nameList, name)
	return nameList
}
