package bits

import (
	"fmt"
	"testing"
)

func TestParity(t *testing.T) {
	var a uint64 = 0xFF01
	fmt.Printf("%d", parity(a))
	// Output: 1
	t.Log("input: ", a, " output: ", parity(a))
}
