package algo

import (
	"fmt"
	"testing"
)

func TestRomanToInt(t *testing.T) {
	if romanToInt("MMMCMXCIX") != 3999 {
		fmt.Println("T1")
		t.Fail()
	}

	if romanToInt("I") != 1 {
		fmt.Println("T2")
		t.Fail()
	}

	if romanToInt("CDXCIX") != 499 {
		fmt.Println("T3")
		t.Fail()
	}
	if romanToInt("") != 0 {
		fmt.Println("T4")
		t.Fail()
	}
}
