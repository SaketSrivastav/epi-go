package bits

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	if result != 5 {
		t.Error("expected 5 but got ", result)
	}
	t.Log("Add(2, 3) = ", result)

	result = Add(15, 1)
	if result != 16 {
		t.Error("expected 16 but got ", result)
	}
	t.Log("Add(15, 1) = ", result)
}

func TestMul(t *testing.T) {
	result := Mul(2, 3)
	if result != 6 {
		t.Error("expected 6 but got ", result)
	}
	t.Log("Mul(2, 3) = ", result)

	result = Mul(15, 1)
	if result != 15 {
		t.Error("expected 15 but got ", result)
	}
	t.Log("Mul(15, 1) = ", result)
}

func TestDivide(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from Divide: ", r)
		}
	}()
	result := Divide(4, 2)
	if result != 2 {
		t.Error("expected 2 but got ", result)
	}
	t.Log("Divide(4, 2) = ", result)

	result = Divide(4, 0)
	t.Errorf("Divide did not panic")
}

func TestPower(t *testing.T) {
	result := Power(4, 2)
	if result != 16 {
		t.Error("expected 16 but got ", result)
	}
	t.Log("Power(4, 2) = ", result)
}

func TestReverseDigit(t *testing.T) {
	result := ReverseDigit(3343)
	if result != 3433 {
		t.Error("expected 3433 but got ", result)
	}
	t.Log("ReverseDigit(3343) = ", result)
}
