package bits

import "testing"

func TestNumBitsSet(t *testing.T) {
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
