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

package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGraph(t *testing.T) {
	source := []byte(`{
  "test": "val",
  "test2": "val2",
  "list": [
    "test1"
  ],
  "listObject": [
    {
      "test": "val"
    },
    {
      "test": "val"
    }
  ]
}`)

	root, err := UnmarshalJSON(source)
	if !assert.NoError(t, err) {
		return
	}

	t.Run("lookup", func(t *testing.T) {
		nodeTest, err := root.Lookup("test")
		assert.NoError(t, err)
		assert.Equal(t, "val", nodeTest.value)

		nodeTest, err = root.Lookup("list")
		assert.NoError(t, err)
		assert.Nil(t, nodeTest.value)
		assert.Equal(t, 1, len(nodeTest.edges))
	})

	t.Run("marshal", func(t *testing.T) {
		data, err := MarshalJSON(root)
		assert.NoError(t, err)
		assert.Equal(t, string(source), string(data))
	})
}
