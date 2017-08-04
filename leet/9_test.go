package leet

import "testing"

func TestIsPalindromeInteger(t *testing.T) {
	if isPalindromeInteger(123) {
		t.Fail()
	}
	// bigger than  2,147,483,647
	if !isPalindromeInteger(1) {
		t.Fail()
	}

	if isPalindromeInteger(-1) {
		t.Fail()
	}

	if !isPalindromeInteger(121) {
		t.Fail()
	}
}
