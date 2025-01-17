package builder

import "github.com/serverlessworkflow/sdk-go/v3/internal/graph"

type CallKind string

const (
	CallKindHttp CallKind = "http"
	CallKindGrpc CallKind = "grpc"
)

type CallBuilder struct {
	root *graph.Node
	with *MapBuilder
}

func (b *CallBuilder) SetCall(call CallKind) *CallBuilder {
	b.root.Edge("call").SetString(string(call))
	return b
}

func (b *CallBuilder) GetCall() string {
	return b.root.Edge("call").GetString()
}

func (b *CallBuilder) With() *MapBuilder {
	if b.with == nil {
		b.with = NewMapBuilder(b.root.Edge("with"))
	}
	return b.with
}

func NewCallBuilder(root *graph.Node) *CallBuilder {
	return &CallBuilder{
		root: root,
	}
}
