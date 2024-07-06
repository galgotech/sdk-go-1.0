package validate

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"sigs.k8s.io/yaml"

	"github.com/serverlessworkflow/sdk-go/v3/internal/validator"
)

const (
	extJSON = ".json"
	extYAML = ".yaml"
	extYML  = ".yml"
)

var supportedExt = []string{extYAML, extYML, extJSON}

// FromFile parses the given Serverless Workflow file into the Workflow type.
func FromFile(path string) error {
	if err := checkFilePath(path); err != nil {
		return err
	}
	fileBytes, err := os.ReadFile(filepath.Clean(path))
	if err != nil {
		return err
	}
	if strings.HasSuffix(path, extYAML) || strings.HasSuffix(path, extYML) {
		return FromYAMLSource(fileBytes)
	}
	return FromJSONSource(fileBytes)
}

// FromYAMLSource parses the given Serverless Workflow YAML source into the Workflow type.
func FromYAMLSource(source []byte) error {
	jsonBytes, err := yaml.YAMLToJSON(source)
	if err != nil {
		return err
	}
	return FromJSONSource(jsonBytes)
}

// FromJSONSource parses the given Serverless Workflow JSON source into the Workflow type.
func FromJSONSource(source []byte) error {
	return validator.Valid(source)
}

// checkFilePath verifies if the file exists in the given path and if it's supported by the parser package
func checkFilePath(path string) error {
	info, err := os.Stat(path)
	if err != nil {
		return err
	}
	if info.IsDir() {
		return fmt.Errorf("file path '%s' must stand to a file", path)
	}
	for _, ext := range supportedExt {
		if strings.HasSuffix(path, ext) {
			return nil
		}
	}
	return fmt.Errorf("file extension not supported for '%s'. supported formats are %s", path, supportedExt)
}
