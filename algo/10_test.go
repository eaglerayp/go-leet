package algo

import (
	"fmt"
	"testing"
)

func TestIsMatch(t *testing.T) {
	if isMatch("aa", "a") {
		fmt.Println("T1")
		t.Fail()
	}
	if !isMatch("aabbc", "a*abb*cc*") {
		fmt.Println("T2")
		t.Fail()
	}
	if isMatch("a", "") {
		fmt.Println("T3")
		t.Fail()
	}
	if isMatch("ab", ".*c") {
		fmt.Println("T4")
		t.Fail()
	}

	if !isMatch("aaa", "a*a") {
		fmt.Println("T5")
		t.Fail()
	}

	if !isMatch("", ".*") {
		fmt.Println("T6")
		t.Fail()
	}
}
