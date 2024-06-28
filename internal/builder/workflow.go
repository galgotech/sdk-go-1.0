package builder

import "github.com/serverlessworkflow/sdk-go/v3/internal/graph"

type WorkflowBuilder struct {
	root     *graph.Node
	document *DocumentBuilder
	do       *DoBuilder
	use      *UseBuilder
}

func (b *WorkflowBuilder) Document() *DocumentBuilder {
	if b.document == nil {
		b.document = NewDocumentBuilder(b.root.Edge("document"))
	}
	return b.document
}

func (b *WorkflowBuilder) Do(dsl string) *DoBuilder {
	if b.do == nil {
		b.do = NewDoBuilder(b.root.Edge("do"))
	}
	return b.do
}

func (b *WorkflowBuilder) Use(dsl string) *UseBuilder {
	if b.use == nil {
		b.use = NewUseBuilder(b.root.Edge("use"))
	}
	return b.use
}
