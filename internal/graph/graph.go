package graph

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type Node struct {
	value interface{}
	order []string
	edges map[string]*Node
	list  bool
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

func (n *Node) Lookup(path string) (*Node, error) {
	pathSplit := strings.Split(path, ".")
	edge := n

	walked := []string{}
	for _, key := range pathSplit {
		walked = append(walked, key)
		if val, ok := edge.edges[key]; ok {
			edge = val
		} else {
			return nil, fmt.Errorf("path not found: %s", strings.Join(walked, "."))
		}
	}
	return edge, nil
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
