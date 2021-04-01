package No741

import (
	"fmt"
	"testing"
)

func TestCherryPickup(t *testing.T) {
	grid := [][]int{
		{0, 1, -1},
		{1, 0, -1},
		{1, 1, 1},
	}
	res := cherryPickup(grid)
	fmt.Println("res: ", res)
}
