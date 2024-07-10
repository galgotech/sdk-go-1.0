package validate

import (
	"github.com/serverlessworkflow/sdk-go/v3/internal/load"
	"github.com/serverlessworkflow/sdk-go/v3/internal/validator"
)

// FromFile parses the given Serverless Workflow file into the Workflow type.
func FromFile(path string) error {
	root, fileBytes, err := load.FromFile(path)
	if err != nil {
		return err
	}

	return validator.Valid(root, fileBytes)
}

// FromYAMLSource parses the given Serverless Workflow YAML source into the Workflow type.
func FromYAMLSource(source []byte) error {
	root, jsonBytes, err := load.FromYAMLSource(source)
	if err != nil {
		return err
	}

	return validator.Valid(root, jsonBytes)
}

// FromJSONSource parses the given Serverless Workflow JSON source into the Workflow type.
func FromJSONSource(source []byte) error {
	root, jsonBytes, err := load.FromJSONSource(source)
	if err != nil {
		return err
	}

	return validator.Valid(root, jsonBytes)
}
