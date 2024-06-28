package builder

import "github.com/serverlessworkflow/sdk-go/v3/internal/graph"

type DocumentBuilder struct {
	root *graph.Node
}

func (b *DocumentBuilder) SetDSL(dsl string) *DocumentBuilder {
	node := b.root.Edge("dsl")
	node.SetString(dsl)
	return b
}

func (b *DocumentBuilder) GetDSL() string {
	node := b.root.Edge("dsl")
	return node.GetString()
}

func (b *DocumentBuilder) SetNamespace(dsl string) *DocumentBuilder {
	node := b.root.Edge("namespace")
	node.SetString(dsl)
	return b
}

func (b *DocumentBuilder) GetNamespace() string {
	node := b.root.Edge("namespace")
	return node.GetString()
}

func NewDocumentBuilder(root *graph.Node) *DocumentBuilder {
	return &DocumentBuilder{
		root: root,
	}
}
