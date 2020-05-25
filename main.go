package main

import (
	"fmt"
	"io/ioutil"

	"jsonschemaTest/schema"

	"github.com/xeipuuv/gojsonschema"
)

func main() {
	jsonContent, err := ioutil.ReadFile("document.json")
	if err != nil {
		panic(err.Error())
	}

	documentLoader := gojsonschema.NewStringLoader(string(jsonContent))
	result, err := schema.ProductSchema.Validate(documentLoader)
	if err != nil {
		panic(err.Error())
	}

	if result.Valid() {
		fmt.Println("The document is valid")
	} else {
		fmt.Println("The document is not valid, see errors:")
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
	}
}
