package builder

import "github.com/serverlessworkflow/sdk-go/v3/internal/graph"

type WaitBuilder struct {
	root     *graph.Node
	duration *DurationBuilder
}

func (b *WaitBuilder) SetWait(wait string) {
	b.root.Edge("wait").Clear().SetString(string(wait))
}

func (b *WaitBuilder) GetWait() string {
	return b.root.Edge("wait").GetString()
}

func (b *WaitBuilder) Duration() *DurationBuilder {
	if b.duration == nil {
		node := b.root.Edge("wait").Clear()
		b.duration = NewDurationBuilder(node)
	}
	return b.duration
}

func NewWaitBuilder(root *graph.Node) *WaitBuilder {
	return &WaitBuilder{
		root: root,
	}
}
