package stack

type Node struct {
	Value any
	blow  *Node
}

type Stack struct {
	Top  *Node
	size int
}

func Create(init any) *Stack {
	top := &Node{init, nil}
	stack := &Stack{top, 1}
	return stack
}

func (s *Stack) IsEmpty() bool {
	return s.Top == nil
}

func (s *Stack) Push(value any) {
	node := &Node{value, s.Top}
	s.Top = node
	s.size++
}

func (s *Stack) Pop() any {
	if s.IsEmpty() {
		return nil
	}
	value := s.Top.Value
	s.Top = s.Top.blow
	s.size--
	return value
}

func (s *Stack) Size() int {
	return s.size
}
func (s *Stack) Peek() any {
	if s.IsEmpty() {
		return nil
	}
	return s.Top.Value
}
