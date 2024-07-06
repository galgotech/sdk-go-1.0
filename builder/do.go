package builder

import (
	"fmt"

	"github.com/serverlessworkflow/sdk-go/v3/internal/graph"
)

type DoBuilder struct {
	root  *graph.Node
	tasks []any
}

func (b *DoBuilder) AddCall(name string) (*CallBuilder, int) {
	index := len(b.tasks)
	nodeIndex := b.root.Edge(fmt.Sprintf("%d", index))
	nodeName := nodeIndex.Edge(name)

	callBuilder := NewCallBuilder(nodeName)
	b.tasks = append(b.tasks, callBuilder)
	return callBuilder, index
}

func (b *DoBuilder) AddWait(name string) (*WaitBuilder, int) {
	index := len(b.tasks)
	nodeIndex := b.root.Edge(fmt.Sprintf("%d", index))
	nodeName := nodeIndex.Edge(name)

	waitBuilder := NewWaitBuilder(nodeName)
	b.tasks = append(b.tasks, waitBuilder)
	return waitBuilder, index
}

func (b *DoBuilder) RemoveTask(index int) *DoBuilder {
	b.tasks = append(b.tasks[:index], b.tasks[index+1:]...)
	return b
}

func NewDoBuilder(root *graph.Node) *DoBuilder {
	root.List(true)
	return &DoBuilder{
		root:  root,
		tasks: []any{},
	}
}
