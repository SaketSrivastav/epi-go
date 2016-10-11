package bits

/*
Swap swap the bits mentioned in ith and jth location
*/
func Swap(x uint, i uint, j uint) uint {
	if ((x >> i) & 0x1) == ((x >> j) & 0x1) {
		return x
	}

	var bitmask uint = ((0x1 << i) | (0x1 << j))
	return (x ^ bitmask)
}

/*
Reverse reverse a 64-bit word
Input: 1000 0100 0010 0001
Output: 0001 0010 0100 1000
Description:
	Take a cache of precomputed reverse of 2 words. Pass a 64-bit value
	which is nothing but a 4 word tuple. Pass 2 words each time to reverse
*/
func Reverse(x uint64) uint64 {
	precomputedReverse := [4]uint16{0x0, 0x2, 0x1, 0x3}
	var wordLength uint64 = 16
	var bitmask uint64 = 0xFFFF

	return ((uint64(precomputedReverse[uint16(x>>(3*wordLength))]) & bitmask) |
		(uint64(precomputedReverse[uint16(x>>(2*wordLength))]) & bitmask << wordLength) |
		(uint64(precomputedReverse[uint16(x>>wordLength)]) & bitmask << (2 * wordLength)) |
		(uint64(precomputedReverse[uint16(x<<(3*wordLength))]) & bitmask << (3 * wordLength)))
}
