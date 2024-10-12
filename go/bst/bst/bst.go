package bst

import "../../stack/stack"

type Tnode[T comparable] struct {
	Vale  T
	Left  *Tnode[T]
	Right *Tnode[T]
}

type Bst[T comparable] struct {
	Root *Tnode[T]
}

func Create[T comparable](Vale T) *Bst[T] {
	Node := &Tnode[T]{Vale, nil, nil}

	return &Bst[T]{Node}
}
func (t *Bst[T]) Find(Vale T) {
	stack := stack.Create[T]()
	stack.Push()
}
func Add[T comparable](Vale T) {

}
