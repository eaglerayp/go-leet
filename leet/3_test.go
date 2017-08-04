package leet

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	a1 := lengthOfLongestSubstring("bbbbb")
	a2 := lengthOfLongestSubstring("pwwkew")
	a3 := lengthOfLongestSubstring("abcabcbb")
	a4 := lengthOfLongestSubstring("pwwkewa")
	if a1 != 1 {
		t.Fail()
	}

	if a2 != 3 {
		t.Fail()
	}

	if a3 != 3 {
		t.Fail()
	}

	if a4 != 4 {
		t.Fail()
	}
}
