package algo

import "testing"

func TestZigconvert(t *testing.T) {
	if convert("PAYPALISHIRING", 3) != "PAHNAPLSIIGYIR" {
		t.Fail()
	}
	if convert("PAYPALISHIRING", 1) != "PAYPALISHIRING" {
		t.Fail()
	}
}
