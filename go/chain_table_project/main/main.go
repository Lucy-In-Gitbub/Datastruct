package main

import (
	"fmt"

	"../chain_table"
)

func main() {
	chain1 := chain_table.NodeHead()
	fmt.Printf("Head: %v\n", chain1)
	chain1.AddNode("Node1")
	chain1.AddNode("Node2")
	chain1.AddNode("Node3")
	chain1.AddNode("Node4")
	chain1.Tranversal()
	chain1.Setvalue(2, "New Value")
	chain1.Tranversal()
	chain1.DeleteNode(2)
	chain1.Tranversal()
}
