package No86

import (
	"fmt"
	"testing"
)

func TestPartition(t *testing.T) {
	head := partition(CreateList(), 3)
	//head := CreateList()
	cur := head
	for ; cur != nil; cur = cur.Next {
		fmt.Print(cur.Val, "->")
	}
	fmt.Print("\n")
}
