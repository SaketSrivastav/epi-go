package bits

import "fmt"

/*
ClosestIntSameWeight return the closest integer of same number of 1's
Input: uint64 110
Output: uint64 101
Description:
	Take 2 bits k1 and k2, to have the next closest integer k1 > k2 and
	|2^k1 - 2^k2| should be minimum. To satisfy this k1 should be minimum value
	and k2 should be less than k1 and close to k1. K1 and k2 should be
	consecutive. Swap the rightmost 2 bits that differ.
*/
func ClosestIntSameWeight(x uint64) uint64 {
	if (x == 0xFFFF) || (x == 0x0000) {
		panic(fmt.Sprintf(string(x), " is invalid number"))
	}

	for i := 0; i < 64-1; i++ {
		if ((x >> uint64(i)) & 1) != (x >> uint64(i+1) & 1) {
			x ^= (uint64(1) << uint64(i)) | (uint64(1) << uint64((i + 1)))
		}
	}
	return x
}
