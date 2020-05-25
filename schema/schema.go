package schema

import (
	"io/ioutil"

	"github.com/xeipuuv/gojsonschema"
)

var (
	// ProductSchema schema for product
	ProductSchema *gojsonschema.Schema
)

func initProductSchema() {
	schemaContent, err := ioutil.ReadFile("schema.json")
	if err != nil {
		panic(err.Error())
	}

	loader := gojsonschema.NewStringLoader(string(schemaContent))
	ProductSchema, err = gojsonschema.NewSchema(loader)
	if err != nil {
		panic(err.Error())
	}
}

func init() {
	initProductSchema()
}
