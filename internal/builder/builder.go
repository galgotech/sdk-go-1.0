package builder

import "github.com/serverlessworkflow/sdk-go/v3/internal/graph"

func NewWorkflow() *WorkflowBuilder {
	node := graph.NewNode()
	return &WorkflowBuilder{
		root: node,
	}
}
