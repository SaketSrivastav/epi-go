package bits

/*
BitwiseLSBCount return the number of 1's in the given integer.
Description:
	It AND's the LSB of the given number with 0x01 and increment the counter if
	it is 1.
*/
func BitwiseLSBCount(x uint32) uint32 {
	if x == 0 {
		return 0
	}

	var cnt uint32
	cnt = 0
	for x != 0 {
		if x&1 == 1 {
			cnt++
		}
		x = x >> 1
	}
	return cnt
}

/*
KerninghamBitCount returns the number of set bits in uin32 number
Description:
	x = x & x-1 always flips the rightmost set bit
*/
func KerninghamBitCount(x uint32) uint32 {
	if x == 0 {
		return 0
	}
	var cnt uint32
	cnt = 0
	for {
		x &= x - 1
		cnt++
		if x == 0 {
			break
		}
	}

	return cnt
}

/*
LookupTblBitCount returns the number of set bits using precomputed lookup table
*/
func LookupTblBitCount(x uint32) uint32 {
	if x == 0 {
		return 0
	}

	var cnt uint32
	var lookUpTbl [256]uint8
	cnt = 0
	for i := 0; i < 256; i++ {
		lookUpTbl[i] = (uint8(i) & 1) + lookUpTbl[uint8(i)/2]
	}

	for i := 0; i < 32 && x != 0; i++ {
		cnt += uint32(lookUpTbl[uint8(x)&255])
		x = x >> 8
	}
	return cnt
}

/*
PopulationCount returns the number of bits set in a given unsigned integer

Description:

	i = i - ((i >> 1) & 0x55555555);
	First of all, the significance of the constant 0x55555555 is that, written
    using the Java / GCC style binary literal notation),

	0x55555555 = 0b01010101010101010101010101010101
	That is, all its odd-numbered bits (counting the lowest bit as bit 1 = odd)
    are 1, and all the even-numbered bits are 0.

	The expression ((i >> 1) & 0x55555555) thus shifts the bits of i right by
    one, and then sets all the even-numbered bits to zero. (Equivalently, we
    could've first set all the odd-numbered bits of i to zero with & 0xAAAAAAAA
    and then shifted the result right by one bit.) For convenience, let's call
    this intermediate value j.

	What happens when we subtract this j from the original i? Well, let's see
    what would happen if i had only two bits:

		i           j         i - j
	----------------------------------
	0 = 0b00    0 = 0b00    0 = 0b00
	1 = 0b01    0 = 0b00    1 = 0b01
	2 = 0b10    1 = 0b01    1 = 0b01
	3 = 0b11    1 = 0b01    2 = 0b10
	Hey! We've managed to count the bits of our two-bit number!

	OK, but what if i has more than two bits set? In fact, it's pretty easy to
    check that the lowest two bits of i - j will still be given by the table
    above, and so will the third and fourth bits, and the fifth and sixth bits,
    and so and. In particular:

	despite the >> 1, the lowest two bits of i - j are not affected by the
    third or higher bits of i, since they'll be masked out of
    j by the & 0x55555555; and since the lowest two bits of j can never have a
    greater numerical value than those of i, the subtraction will never borrow
    from the third bit of i: thus, the lowest two bits of i also cannot affect
    the third or higher bits of i - j.
	In fact, by repeating the same argument, we can see that the calculation on
    this line, in effect, applies the table above to each of the 16 two-bit
    blocks in i in parallel. That is, after executing this line, the lowest two
    bits of the new value of i will now contain the number of bits set among
    the corresponding bits in the original value of i, and so will the next two
    bits, and so on.

	Line 2:

	i = (i & 0x33333333) + ((i >> 2) & 0x33333333);
	Compared to the first line, this one's quite simple. First, note that

	0x33333333 = 0b00110011001100110011001100110011
	Thus, i & 0x33333333 takes the two-bit counts calculated above and throws
    away every second one of them, while (i >> 2) & 0x33333333 does the same
    after shifting i right by two bits. Then we add the results together.

	Thus, in effect, what this line does is take the bitcounts of the lowest
    two and the second-lowest two bits of the original input, computed on the
    previous line, and add them together to give the bitcount of the lowest
    four bits of the input. And, again, it does this in parallel for all the 8
    four-bit blocks (= hex digits) of the input.

	Line 3:

	return (((i + (i >> 4)) & 0x0F0F0F0F) * 0x01010101) >> 24;
	OK, what's going on here?

	Well, first of all, (i + (i >> 4)) & 0x0F0F0F0F does exactly the same as
    the previous line, except it adds the adjacent four-bit bitcounts together
    to give the bitcounts of each eight-bit block (i.e. byte) of the input.
    (Here, unlike on the previous line, we can get away with moving the
    & outside the addition, since we know that the eight-bit bitcount can never
    exceed 8, and therefore will fit inside four bits without overflowing.)

	Now we have a 32-bit number consisting of four 8-bit bytes, each byte
    holding the number of 1-bit in that byte of the original input. (Let's call
    these bytes A, B, C and D.) So what happens when we multiply this value
    (let's call it k) by 0x01010101?

	Well, since 0x01010101 = (1 << 24) + (1 << 16) + (1 << 8) + 1, we have:

	k * 0x01010101 = (k << 24) + (k << 16) + (k << 8) + k
	Thus, the highest byte of the result ends up being the sum of:

	its original value, due to the k term, plus
	the value of the next lower byte, due to the k << 8 term, plus
	the value of the second lower byte, due to the k << 16 term, plus
	the value of the fourth and lowest byte, due to the k << 24 term.
	(In general, there could also be carries from lower bytes, but since we
    know the value of each byte is at most 8, we know the addition will never
    overflow and create a carry.)

	That is, the highest byte of k * 0x01010101 ends up being the sum of the
    bitcounts of all the bytes of the input, i.e. the total bitcount of the
    32-bit input number. The final >> 24 then simply shifts this value down
    from the highest byte to the lowest.

	Ps. This code could easily be extended to 64-bit integers, simply by
    changing the 0x01010101 to 0x0101010101010101 and the >> 24 to >> 56.
    Indeed, the same method would even work for 128-bit integers; 256 bits
    would require adding one extra shift / add / mask step, however, since the
    number 256 no longer quite fits into an 8-bit byte.

Example:
    Consider only 8 bits in which case this algorithm would need only
    3 bit masks.
    Let x = 11111111
    mask1 = 0x01010101 (0x55) pair of one 0 and 1
    mask2 = 0x00110011 (0x33) pair of two 0 and 1
    mask3 = 0x00001111 (0x0f) pair of four 0 and 1

Step 1:
	x = (x & mask1) + ((x >> 1) & mask1)
    11111111 & 01010101 = 01010101
                        +
    01111111 & 01010101 = 01010101
                      x = 10101010
Step 2:
	x = (x & mask2) + ((x >> 2) & mask2)
    10101010 & 00110011 = 00100010
    00101010 & 00110011 = 00100010
                      x = 01000100

Step 3:
	x = (x & mask3) + ((x >> 4) & mask3)
    01000100 & 00001111 = 00000100
    00000100 & 00001111 = 00000100
                      x = 00001000 = 8
    To calculate 8 bits we took 3 masks therefore worst case complexity is
    O(logn)
*/
func PopulationCount(x uint32) uint32 {
	var mask1 uint32
	mask1 = 0x55555555
	var mask2 uint32
	mask2 = 0x33333333
	var mask3 uint32
	mask3 = 0x0f0f0f0f
	var mask4 uint32
	mask4 = 0x00ff00ff
	var mask5 uint32
	mask5 = 0x0000ffff

	x = (x & mask1) + ((x >> 1) & mask1)
	x = (x & mask2) + ((x >> 2) & mask2)
	x = (x & mask3) + ((x >> 4) & mask3)
	x = (x & mask4) + ((x >> 8) & mask4)
	x = (x & mask5) + ((x >> 16) & mask5)
	return x
}
