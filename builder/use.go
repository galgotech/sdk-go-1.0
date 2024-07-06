package builder

import "github.com/serverlessworkflow/sdk-go/v3/internal/graph"

type UseBuilder struct {
	root *graph.Node
}

func NewUseBuilder(root *graph.Node) *UseBuilder {
	return &UseBuilder{
		root: root,
	}
}
