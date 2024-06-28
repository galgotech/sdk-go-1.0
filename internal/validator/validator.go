// Copyright 2020 The Serverless Workflow Specification Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package validator

import (
	"bytes"
	"log"

	"github.com/santhosh-tekuri/jsonschema/v6"
	"sigs.k8s.io/yaml"

	"github.com/serverlessworkflow/sdk-go/v3/internal/workflow"
)

var schema *jsonschema.Schema

func Valid(source []byte) (bool, error) {
	inst, err := jsonschema.UnmarshalJSON(bytes.NewReader(source))
	if err != nil {
		log.Fatal(err)
	}

	err = schema.Validate(inst)
	if err != nil {
		return false, err
	}

	return true, nil
}

func init() {
	var err error

	jsonBytes, err := yaml.YAMLToJSON([]byte(workflow.WorkflowSpec))
	if err != nil {
		log.Fatal(err)
	}
	readerJsonSchema, err := jsonschema.UnmarshalJSON(bytes.NewReader(jsonBytes))
	if err != nil {
		log.Fatal(err)
	}

	c := jsonschema.NewCompiler()
	err = c.AddResource("workflow.json", readerJsonSchema)
	if err != nil {
		log.Fatal(err)
	}

	schema, err = c.Compile("workflow.json")
	if err != nil {
		log.Fatal(err)
	}
}
