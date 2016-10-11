package bits

import "fmt"

/*
Add adds 2 uint64 numbers using bitwise operations
*/
func Add(x uint64, y uint64) uint64 {
	var result uint64
	var tempX = x
	var tempY = y
	var bit uint64 = 1
	var xBit uint64
	var yBit uint64
	var carryin uint64
	var carryout uint64

	for (tempX > 0) || (tempY > 0) {
		xBit = x & bit
		yBit = y & bit
		result = result | (xBit ^ yBit ^ carryin)
		carryout = (xBit & yBit) | (xBit & carryin) | (yBit & carryin)
		carryin = carryout << 1
		tempX = tempX >> 1
		tempY = tempY >> 1
		bit = bit << 1
	}
	return result | carryin
}

/*
Mul multiplies two given number without using arithmetic multiply operation
*/
func Mul(x uint64, y uint64) uint64 {
	var sum uint64
	for x > 0 {
		if (x & 1) > 0 {
			sum = Add(sum, y)
		}
		x = x >> 1
		y = y << 1
	}
	return sum
}

/*
Divide divide x by y and returns the quotient of the division
*/
func Divide(x uint64, y uint64) uint64 {
	var result uint64
	var power uint64
	power = 32
	if y == 0 {
		panic(fmt.Sprintf("Math Error: Division by 0 error"))
	}

	var yPower = y << power
	for x >= y {
		for yPower > x {
			yPower = yPower >> 1
			power--
		}
		result = result + 0x1<<power
		x = x - yPower
	}
	return result
}

/*
Power returns x power y
Description:
	x ^ y can be broken down to x (y0 + y1) = x^y0 * x^y1. Generalizing the
	rule in binary, if LSB of y is 1 then x^y = x * (x^y/2)2 and when LSB of
	y = 0 then x^y = (x^y/2)2
*/
func Power(x uint, y uint) uint {
	var result uint
	result = 1
	for y > 0 {
		if (y & 1) > 0 {
			result *= x
		}
		x *= x
		y = y >> 1
	}
	return result
}

/*
ReverseDigit reverses the digits of decimal number
*/
func ReverseDigit(x uint) uint {
	var xRem uint
	var result uint
	xRem = x
	result = 0
	for xRem > 0 {
		result = result*10 + xRem%10
		xRem = xRem / 10
	}
	return uint(result)
}
