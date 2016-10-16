package bits

import "testing"

func TestBitwiseLSBCount(t *testing.T) {
	var result uint32
	result = BitwiseLSBCount(0x1011)
	if result != 3 {
		t.Errorf("expected 3, but got %d", result)
	}
	t.Log("BitwiseLSBCount(0x1011) = 3")

	result = BitwiseLSBCount(0x11111111)
	if result != 8 {
		t.Errorf("expected 8, but got %d", result)
	}
	t.Log("BitwiseLSBCount(0x11111111) = 8")
}

func TestKerninghamBitCount(t *testing.T) {
	var result uint32
	result = KerninghamBitCount(0x1011)
	if result != 3 {
		t.Errorf("expected 3, but got %d", result)
	}
	t.Log("KerninghamBitCount(0x1011) = 3")

	result = KerninghamBitCount(0x11111111)
	if result != 8 {
		t.Errorf("expected 8, but got %d", result)
	}
	t.Log("KerninghamBitCount(0x11111111) = 8")
}

func TestLookupTblBitCount(t *testing.T) {
	var result uint32
	result = LookupTblBitCount(0x1011)
	if result != 3 {
		t.Errorf("expected 3, but got %d", result)
	}
	t.Log("LookupTblBitCount(0x1011) = 3")

	result = LookupTblBitCount(0xff)
	if result != 8 {
		t.Errorf("expected 8, but got %d", result)
	}
	t.Log("LookupTblBitCount(0x11111111) = 8")
}

func TestPopulationCount(t *testing.T) {
	var result uint32
	result = PopulationCount(0x1011)
	if result != 3 {
		t.Errorf("expected 3, but got %d", result)
	}
	t.Log("numBitsSet(0x1011) = 3")

	result = PopulationCount(0x11111111)
	if result != 8 {
		t.Errorf("expected 8, but got %d", result)
	}
	t.Log("numBitsSet(0x11111111) = 8")
}
