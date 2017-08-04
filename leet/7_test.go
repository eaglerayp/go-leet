package leet

import "testing"

func TestReverse(t *testing.T) {
	if reverse(123) != 321 {
		t.Fail()
	}
	// bigger than  2,147,483,647
	if reverse(1111111119) != 0 {
		t.Fail()
	}

	if reverse(0) != 0 {
		t.Fail()
	}

	if reverse(-123456789) != -987654321 {
		t.Fail()
	}
}
