package list

import "fmt"

type Chainfunc interface {
	Tranversal()
	Append(initValue any)
	RemoveAt(index int)
	Change(index int, value any)
	Find(value any)
	Check()
}
type Nodefunc interface {
	Show()
}

type HashTableFunc interface {
	CheckHash()
}
type HashTable struct {
	capacity int
	bucket   [][]Node
}

type Node struct {
	Index int
	Value any
	Next  *Node
	Prev  *Node
}
type Chain struct {
	Head   *Node
	Tail   *Node
	Hash   HashTable
	Length int
}

func (h HashTable) Hashfunc(Key int) int {
	return Key % h.capacity
}

func (n Node) Show() {
	fmt.Println(n.Index, n.Value)
}

func (c *Chain) Tranversal() {
	p := c.Head
	for i := 0; i <= c.Length; i++ {
		fmt.Println(p.Index, p.Value)
		p = p.Next
	}
}

func (c *Chain) Find(index int) *Node {
	key := c.Hash.Hashfunc(index)
	bucket := c.Hash.bucket[key]
	for i, v := range bucket {
		if v.Index == index {
			fmt.Println(v.Index, v.Value)
			x := &c.Hash.bucket[key][i]
			return x
		}
	}
	fmt.Println("Not Found")
	return nil
}

func (c *Chain) Check() {
	p := c.Head
	for i := 0; i <= c.Length; i++ {
		fmt.Println(p)
		p = p.Next
	}
	fmt.Println()
	n := c.Tail
	for i := 0; i < c.Length; i++ {
		fmt.Println(n)
		n = n.Prev
	}
}
func (c *Chain) RemoveAt(index int) {
	find := c.Find(index)
	p := find.Prev
	n := p.Next.Next
	n.Prev = p
	p.Next = n
	c.Length--
}

func (c *Chain) Append(initValue any) {
	p := &Node{0, initValue, nil, nil}
	p.Value = initValue
	c.Tail.Next = p
	p.Prev = c.Tail
	c.Tail = p
	c.Tail.Next = c.Head
	c.Length++
	p.Index = c.Length
	key := c.Hash.Hashfunc(p.Index)
	c.Hash.bucket[key] = append(c.Hash.bucket[key], *p)
}

func (c HashTable) CheckHash() {
	for i, v := range c.bucket {
		for j, n := range v {
			fmt.Println(i, j, n)
		}
	}
}

func Create(initValue any) Chain {
	head := &Node{0, initValue, nil, nil}
	tail := head
	tail.Next = head
	tail.Prev = nil
	hashlist := HashTable{10, make([][]Node, 10)}
	hashlist.bucket[0] = append(hashlist.bucket[0], *head)
	Chain := Chain{head, tail, hashlist, 0}
	return Chain
}
