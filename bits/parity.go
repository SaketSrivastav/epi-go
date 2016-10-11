package bits

// Find the parity of a 64-bit value
func parity(x uint64) int {
	x = x ^ (x >> 32)
	x = x ^ (x >> 16)
	x = x ^ (x >> 8)
	x = x ^ (x >> 4)
	x = x ^ (x >> 2)
	x = x ^ (x >> 1)
	if x&0x1 == 0x1 {
		return 1
	}

	return 0
}
