package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/santhosh-tekuri/jsonschema/v6"
)

func Example_fromStrings() {
	catSchema, err := jsonschema.UnmarshalJSON(strings.NewReader(`{
        "type": "object",
        "properties": {
            "speak": { "const": "meow" }
        },
        "required": ["speak"]
    }`))
	if err != nil {
		log.Fatal(err)
	}
	// note that dog.json is loaded from file ./testdata/examples/dog.json
	petSchema, err := jsonschema.UnmarshalJSON(strings.NewReader(`{
        "oneOf": [
            { "$ref": "cat.json" }
        ]
    }`))
	if err != nil {
		log.Fatal(err)
	}
	inst, err := jsonschema.UnmarshalJSON(strings.NewReader(`{"speak": "bow"}`))
	if err != nil {
		log.Fatal(err)
	}

	c := jsonschema.NewCompiler()
	if err := c.AddResource("./testdata/examples/cat.json", catSchema); err != nil {
		log.Fatal(err)
	}
	if err := c.AddResource("./testdata/examples/pet.json", petSchema); err != nil {
		log.Fatal(err)
	}
	sch, err := c.Compile("./testdata/examples/pet.json")
	if err != nil {
		log.Fatal(err)
	}
	err = sch.Validate(inst)
	fmt.Println("valid:", err == nil)
	// Output:
	// valid: true
}

func main() {
	Example_fromStrings()
}
