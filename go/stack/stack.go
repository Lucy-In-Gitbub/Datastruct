package stack

import "fmt"

type Node struct {
	Index int
	Value any
	blow  *Node
}

type Stack struct {
	Top    *Node
	Name   string
	Length int
}

func Create(init any, name string) *Stack {
	top := &Node{1, init, nil}
	stack := &Stack{top, name, 1}
	return stack
}

func (s *Stack) Push(value any) {
	node := &Node{s.Length + 1, value, s.Top}
	s.Top = node
	s.Length++
}

func (N *Node) Show() {
	fmt.Printf("Index: %d, Value: %v\n", N.Index, N.Value)
}

func (s *Stack) Show() {
	node := s.Top
	if node == nil {
		fmt.Println("Stack is empty")
		return
	}
	for node != nil {
		node.Show()
		node = node.blow
	}
}

func (s *Stack) Pop() any {
	node := s.Top
	if node == nil {
		fmt.Println("Stack is empty")
		return nil
	}
	s.Top = node.blow
	s.Length--
	return node.Value
}
