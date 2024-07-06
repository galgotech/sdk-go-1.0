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

func (b *DocumentBuilder) SetName(dsl string) *DocumentBuilder {
	node := b.root.Edge("name")
	node.SetString(dsl)
	return b
}

func (b *DocumentBuilder) GetName() string {
	node := b.root.Edge("name")
	return node.GetString()
}

func (b *DocumentBuilder) SetVersion(dsl string) *DocumentBuilder {
	node := b.root.Edge("version")
	node.SetString(dsl)
	return b
}

func (b *DocumentBuilder) GetVersion() string {
	node := b.root.Edge("version")
	return node.GetString()
}

func NewDocumentBuilder(root *graph.Node) *DocumentBuilder {
	documentBuilder := &DocumentBuilder{
		root: root,
	}
	documentBuilder.SetDSL("1.0.0-alpha1")
	return documentBuilder
}
