package list

import (
	"fmt"
	"testing"
)

func Test_ChainTable(t *testing.T) {
	chain1 := Create("Chain1")
	fmt.Printf("Head: %v\n", chain1)
	chain1.Append("Node1")
	chain1.Append("Node2")
	chain1.Append("Node3")
	chain1.Append("Node4")
	chain1.RemoveAt(2)
	fmt.Println()
	chain1.Tranversal()
}
