package chain_table

import "fmt"

type node struct {
	Index int
	Value string
	Next  *node
	Prev  *node
}
type Chain struct {
	Head  *node
	Tail  *node
	count int
}

func (c *Chain) Tranversal() {
	p := c.Head
	p = p.Prev
	for i := 1; i <= c.count; i++ {
		fmt.Println(p.Index, p.Value)
		p = p.Prev
	}
}
func (c *Chain) DeleteNode(index int) {
	p := c.Head
	for i := 0; i < index; i++ {
		p = p.Prev
	}
	cp := p.Prev
	for i := p.Index; i <= c.count-1; i++ {
		cp.Index--
		cp = cp.Prev
	}
	p.Prev.Next = p.Next
	p.Next.Prev = p.Prev
	c.count--
}
func (c *Chain) Setvalue(index int, value string) {
	p := c.Head
	for i := 0; i < index; i++ {
		p = p.Prev
	}
	p.Value = value
}

func (c *Chain) AddNode(value string) {
	p := new(node)
	p.Value = value
	c.Tail.Prev = p
	p.Next = c.Tail
	c.Tail = p
	c.Tail.Prev = c.Head
	c.count++
	p.Index = c.count
}

func NodeHead() Chain {
	head := new(node)
	tail := head
	tail.Index = 0
	tail.Prev = head
	tail.Next = nil
	Chain := Chain{head, tail, 0}
	return Chain
}
