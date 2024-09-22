package list

import "fmt"

type node[T int | string | float64 | float32] struct {
	Index int
	Value T
	Next  *node[T]
	Prev  *node[T]
}
type Chain[T int | string | float64 | float32] struct {
	Head  *node[T]
	Tail  *node[T]
	count int
}

func (c *Chain[T]) Tranversal() {
	p := c.Head
	p = p.Next
	for i := 0; i <= c.count; i++ {
		fmt.Println(p.Index, p.Value)
		p = p.Next
	}
}

// func (c *Chain) DeleteNode(index int) {
// 	p := c.Head
// 	for i := 0; i < index; i++ {
// 		p = p.Prev
// 	}
// 	cp := p.Prev
// 	for i := p.Index; i <= c.count-1; i++ {
// 		cp.Index--
// 		cp = cp.Prev
// 	}
// 	p.Prev.Next = p.Next
// 	p.Next.Prev = p.Prev
// 	c.count--
// }
// func (c *Chain) Setvalue(index int, value string) {
// 	p := c.Head
// 	for i := 0; i < index; i++ {
// 		p = p.Next
// 	}
// 	p.Value = value
// }

func (c *Chain[T]) Append(initValue T) {
	p := &node[T]{0, initValue, nil, nil}
	p.Value = initValue
	c.Tail.Next = p
	p.Prev = c.Tail
	c.Tail = p
	c.Tail.Next = c.Head
	c.count++
	p.Index = c.count
}

func Create[T int | string | float64 | float32](initValue T) Chain[T] {
	head := &node[T]{0, initValue, nil, nil}
	tail := head
	tail.Next = head
	tail.Prev = nil
	Chain := Chain[T]{head, tail, 0}
	return Chain
}
