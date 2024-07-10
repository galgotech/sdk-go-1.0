package validator

import (
	"bytes"
	"log"

	"github.com/santhosh-tekuri/jsonschema/v6"
	"sigs.k8s.io/yaml"

	"github.com/serverlessworkflow/sdk-go/v3/internal/dsl"
	"github.com/serverlessworkflow/sdk-go/v3/internal/graph"
)

var schema *jsonschema.Schema

func Valid(root *graph.Node, source []byte) error {
	inst, err := jsonschema.UnmarshalJSON(bytes.NewReader(source))
	if err != nil {
		return err
	}

	err = schema.Validate(inst)
	if err != nil {
		return err
	}

	err = integrityValidate(root)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	var err error

	jsonBytes, err := yaml.YAMLToJSON([]byte(dsl.DSLSpec))
	if err != nil {
		log.Fatal(err)
	}
	readerJsonSchema, err := jsonschema.UnmarshalJSON(bytes.NewReader(jsonBytes))
	if err != nil {
		log.Fatal(err)
	}

	c := jsonschema.NewCompiler()
	err = c.AddResource("dslspec.json", readerJsonSchema)
	if err != nil {
		log.Fatal(err)
	}

	schema, err = c.Compile("dslspec.json")
	if err != nil {
		log.Fatal(err)
	}
}
