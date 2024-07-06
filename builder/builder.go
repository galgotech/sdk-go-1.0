package builder

import (
	"encoding/json"

	"github.com/serverlessworkflow/sdk-go/v3/validate"
	"sigs.k8s.io/yaml"
)

func Validate(builder *WorkflowBuilder) error {
	data, err := Json(builder)
	if err != nil {
		return err
	}

	err = validate.FromJSONSource(data)
	if err != nil {
		return err
	}

	return nil
}

func Json(builder *WorkflowBuilder) ([]byte, error) {
	data, err := json.MarshalIndent(builder.node(), "", "  ")
	if err != nil {
		return nil, err
	}

	return data, nil
}

func Yaml(builder *WorkflowBuilder) ([]byte, error) {
	data, err := Json(builder)
	if err != nil {
		return nil, err
	}
	return yaml.JSONToYAML(data)
}
