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
      "test2": "val"
    }
  ]
}`)

	root, err := UnmarshalJSON(source)
	if !assert.NoError(t, err) {
		return
	}

	t.Run("marshal", func(t *testing.T) {
		data, err := MarshalJSON(root)
		assert.NoError(t, err)
		assert.Equal(t, string(source), string(data))
	})

	t.Run("lookup key", func(t *testing.T) {
		nodeTest := root.Lookup("test")
		assert.NoError(t, err)
		assert.Equal(t, "val", nodeTest.value)
	})

	t.Run("lookup not found", func(t *testing.T) {
		nodeTest := root.Lookup("list2")
		assert.Nil(t, nodeTest)
	})

	t.Run("lookup list", func(t *testing.T) {
		nodeTest := root.Lookup("list")
		assert.Nil(t, nodeTest.value)
		assert.Equal(t, 1, len(nodeTest.edges))

		nodeTest = root.Lookup("listObject.*")
		assert.Nil(t, nodeTest.value)
		assert.Equal(t, 2, len(nodeTest.edges))
	})

	t.Run("lookup list index", func(t *testing.T) {
		nodeTest := root.Lookup("list.0")
		assert.Equal(t, "test1", nodeTest.value)
	})

	t.Run("lookup search in a list", func(t *testing.T) {
		nodeTest := root.Lookup("listObject.*.test")
		assert.Equal(t, "val", nodeTest.value)
	})
}
