package No93

import (
	"fmt"
	"testing"
)

func TestRestoreIpAddresses(t *testing.T) {
	s := "25525511135"
	res := restoreIpAddresses(s)
	fmt.Println("res: ", res)
}
