package graph

import (
	"encoding/json"
)

func marshalNode(n *Node) ([]byte, error) {
	if n.value != nil {
		return json.Marshal(n.value)
	}

	var out []byte
	if n.list {
		out = append(out, '[')
	} else {
		out = append(out, '{')
	}

	nEdge := len(n.order) - 1
	for i, edge := range n.order {
		node := n.edges[edge]
		val, err := json.Marshal(node)
		if err != nil {
			return nil, err
		}

		if n.list {
			out = append(out, val...)
		} else {
			out = append(out, []byte("\""+edge+"\":")...)
			out = append(out, val...)
		}

		if nEdge != i {
			out = append(out, byte(','))
		}
	}

	if n.list {
		out = append(out, ']')
	} else {
		out = append(out, '}')
	}

	return out, nil
}
