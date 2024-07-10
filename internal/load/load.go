package load

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/serverlessworkflow/sdk-go/v3/internal/graph"
	"sigs.k8s.io/yaml"
)

const (
	extJSON = ".json"
	extYAML = ".yaml"
	extYML  = ".yml"
)

var supportedExt = []string{extYAML, extYML, extJSON}

func FromFile(path string) (*graph.Node, []byte, error) {
	if err := checkFilePath(path); err != nil {
		return nil, nil, err
	}

	fileBytes, err := os.ReadFile(filepath.Clean(path))
	if err != nil {
		return nil, nil, err
	}

	if strings.HasSuffix(path, extYAML) || strings.HasSuffix(path, extYML) {
		return FromYAMLSource(fileBytes)
	}

	return FromJSONSource(fileBytes)
}

func FromYAMLSource(source []byte) (*graph.Node, []byte, error) {
	jsonBytes, err := yaml.YAMLToJSON(source)
	if err != nil {
		return nil, nil, err
	}
	return FromJSONSource(jsonBytes)
}

func FromJSONSource(fileBytes []byte) (*graph.Node, []byte, error) {
	root, err := graph.UnmarshalJSON(fileBytes)
	if err != nil {
		return nil, nil, err
	}

	err = graph.LoadExternalResource(root)
	if err != nil {
		return nil, nil, err
	}

	return root, fileBytes, nil
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
