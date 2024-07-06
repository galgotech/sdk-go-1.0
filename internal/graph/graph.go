package graph

import (
	"bytes"
	"encoding/json"
	"log"
	"strings"
)

type Node struct {
	value interface{}
	order []string
	edges map[string]*Node
	list  bool
}

func (n *Node) List(list bool) {
	n.list = list
}

func (n *Node) UnmarshalJSON(data []byte) error {
	return unmarshalNode(n, data)
}

func (n *Node) MarshalJSON() ([]byte, error) {
	return marshalNode(n)
}

func (n *Node) Edge(name string) *Node {
	if n.value != nil {
		log.Fatal("value alredy defined, execute clear first")
	}
	if _, ok := n.edges[name]; !ok {
		n.edges[name] = NewNode()
		n.order = append(n.order, name)
	}
	return n.edges[name]
}

func (n *Node) SetString(value string) *Node {
	n.setValue(value)
	return n
}

func (n *Node) SetInt(value int) *Node {
	n.setValue(value)
	return n
}

func (n *Node) SetFloat(value float32) *Node {
	n.setValue(value)
	return n
}

func (n *Node) SetBool(value bool) *Node {
	n.setValue(value)
	return n
}

func (n *Node) setValue(value any) {
	if len(n.edges) > 0 {
		log.Fatal("alredy defined edges, execute clear fist")
	}
	n.value = value
}

func (n *Node) GetString() string {
	return n.value.(string)
}

func (n *Node) GetInt() int {
	return n.value.(int)
}

func (n *Node) GetFloat() float32 {
	return n.value.(float32)
}

func (n *Node) Clear() *Node {
	n.value = nil
	n.edges = map[string]*Node{}
	n.order = []string{}
	return n
}

func (n *Node) Lookup(path string) *Node {
	dotIndex := strings.Index(path, ".")
	var key string
	if dotIndex == -1 {
		key = strings.TrimSpace(path)
	} else {
		key = strings.TrimSpace(path[0:dotIndex])
		path = path[dotIndex+1:]
	}

	var currentNode *Node
	if n.list && key == "*" {
		if dotIndex == -1 {
			return n
		}
		for _, node := range n.edges {
			currentNode = node.Lookup(path)
			if currentNode != nil {
				return currentNode
			}
		}
		return nil
	}

	currentNode = n.Edge(key)
	if currentNode == nil {
		return nil
	}
	if dotIndex == -1 {
		return currentNode
	}

	return currentNode.Lookup(path)
}

func NewNode() *Node {
	return (&Node{}).Clear()
}

func UnmarshalJSON(data []byte) (*Node, error) {
	node := NewNode()
	err := json.Unmarshal(data, &node)
	if err != nil {
		return nil, err
	}

	return node, nil
}

func MarshalJSON(n *Node) ([]byte, error) {
	data, err := json.Marshal(n)
	if err != nil {
		return nil, err
	}

	var out bytes.Buffer
	err = json.Indent(&out, data, "", "  ")
	if err != nil {
		return nil, err
	}

	return out.Bytes(), nil
}

func LoadExternalResource(n *Node) error {
	return nil
}
