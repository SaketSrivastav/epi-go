package bits

import (
	"fmt"
	"testing"
)

func TestClosestIntSameWeight(t *testing.T) {
	fmt.Printf("%x", ClosestIntSameWeight(0x0006))
	// Output: 0x0005
	defer func() {
		if r := recover(); r != nil {
			t.Log("The code did panic")
		}
	}()
	ClosestIntSameWeight(0x0000)
}
