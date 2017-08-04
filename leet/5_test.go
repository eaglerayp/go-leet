package leet

import (
	"fmt"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	if longestPalindrome("bab") != "bab" {
		t.Fail()
	}
	if longestPalindrome("bbab") != "bab" {
		t.Fail()
	}

	if longestPalindrome("bbbb") != "bbbb" {
		t.Fail()
	}

	if longestPalindrome("bcde") != "b" {
		t.Fail()
	}
	if longestPalindrome("babad") != "bab" {
		t.Fail()
	}
	test := "civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldoftzhatwarWehavecometodedicpateaportionofthatfieldasafinalrestingplaceforthosewhoheregavetheirlivesthatthatnationmightliveItisaltogetherfangandproperthatweshoulddothisButinalargersensewecannotdedicatewecannotconsecratewecannothallowthisgroundThebravelmenlivinganddeadwhostruggledherehaveconsecrateditfaraboveourpoorponwertoaddordetractTgheworldadswfilllittlenotlenorlongrememberwhatwesayherebutitcanneverforgetwhattheydidhereItisforusthelivingrathertobededicatedheretotheulnfinishedworkwhichtheywhofoughtherehavethusfarsonoblyadvancedItisratherforustobeherededicatedtothegreattdafskremainingbeforeusthatfromthesehonoreddeadwetakeincreaseddevotiontothatcauseforwhichtheygavethelastpfullmeasureofdevotionthatweherehighlyresolvethatthesedeadshallnothavediedinvainthatthisnationunsderGodshallhaveanewbirthoffreedomandthatgovernmentofthepeoplebythepeopleforthepeopleshallnotperishfromtheearth"
	fmt.Println(longestPalindrome(test))
}
