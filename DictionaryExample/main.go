package main

import (
	"fmt"

	mydict "github.com/yunji/Go-study/DictionaryExample/dict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "First word"}
	definition, err := dictionary.Search("first")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}
}
