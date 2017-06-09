package algo

import (
	"fmt"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	if longestCommonPrefix([]string{"abc", "ab", "abcd"}) != "ab" {
		fmt.Println("T1")
		t.Fail()
	}
	if longestCommonPrefix([]string{"abc"}) != "abc" {
		fmt.Println("T2")
		t.Fail()
	}

	if longestCommonPrefix([]string{"abcdef", "bc"}) != "" {
		fmt.Println("T3")
		t.Fail()
	}
}
