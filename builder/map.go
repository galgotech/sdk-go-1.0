package builder

import "github.com/serverlessworkflow/sdk-go/v3/internal/graph"

type MapBuilder struct {
	root *graph.Node
}

func (b *MapBuilder) Set(name string, value string) {
	b.root.Edge(name).SetString(value)
}

func (b *MapBuilder) Get(name string) string {
	return b.root.Edge(name).GetString()
}

func NewMapBuilder(root *graph.Node) *MapBuilder {
	return &MapBuilder{
		root: root,
	}
}
