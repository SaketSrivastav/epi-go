package bits

import "math"

func isNumPalindrome(x uint) bool {
	if x < 0 {
		return false
	}

	if x == 0 {
		return true
	}

	var msdMask uint
	numDigit := math.Floor(math.Log2(float64(x))) + 1
	msdMask = 1 << (uint(numDigit) - 1)
	for x > 0 {
		if (x / msdMask) != (x % 2) {
			return false
		}
		x = x % msdMask
		x = x / 2
		msdMask = msdMask >> 2
	}
	return true
}
