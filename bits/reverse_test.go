package bits

import (
	"fmt"
	"testing"
)

func TestSwap(t *testing.T) {
	var a uint = 0x0001
	var result = Swap(a, 0, 1)
	fmt.Printf("%x", result)
	// Output: 0x0002
	t.Log("input: ", a, " output: ", Swap(a, 0, 1))
}

func TestReverse(t *testing.T) {
	var a uint64 = 0x8421
	fmt.Printf("%x", Reverse(a))
	// Output: 0x1248
	t.Log("input: ", a, " output: ", Reverse(a))
}
