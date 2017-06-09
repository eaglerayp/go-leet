package algo

import (
	"fmt"
	"testing"
)

func TestIntToRoman(t *testing.T) {
	if intToRoman(3999) != "MMMCMXCIX" {
		fmt.Println("T1")
		t.Fail()
	}
	if intToRoman(100) != "C" {
		fmt.Println("T2")
		t.Fail()
	}
	if intToRoman(499) != "CDXCIX" {
		fmt.Println("T3")
		t.Fail()
	}
}
