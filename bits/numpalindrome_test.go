package bits

import "testing"

func TestIsNumPalindrome(t *testing.T) {
	result := isNumPalindrome(5)
	if result != true {
		t.Error("expected true but got ", result)
	}
	t.Log("isNumPalindrome(5) = ", result)
}
