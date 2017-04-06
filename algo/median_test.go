package algo

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	if isPalindrome("bbab") {
		t.Fail()
	}

	if !isPalindrome("bbbb") {
		t.Fail()
	}

	if isPalindrome("baba") {
		t.Fail()
	}

	if !isPalindrome("bb") {
		t.Fail()
	}

	if !isPalindrome("b") {
		t.Fail()
	}
	if isPalindrome("ba") {
		t.Fail()
	}
}

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

func TestZigconvert(t *testing.T) {
	if convert("PAYPALISHIRING", 3) != "PAHNAPLSIIGYIR" {
		t.Fail()
	}
	if convert("PAYPALISHIRING", 1) != "PAYPALISHIRING" {
		t.Fail()
	}
}

func TestMyAtoi(t *testing.T) {
	if myAtoi("123") != 123 {
		t.Fail()
	}
	if myAtoi("0") != 0 {
		t.Fail()
	}
	if myAtoi("") != 0 {
		fmt.Println(myAtoi(""))
		t.Fail()
	}
	if myAtoi("-123") != -123 {
		t.Fail()
	}
	if myAtoi("999999999999999999999999999") != 0 {
		fmt.Println(myAtoi("999999999999999999999999999"))
		t.Fail()
	}

	if myAtoi("+-2") != 0 {
		t.Fail()
	}

	if myAtoi("  - 2") != 0 {
		t.Fail()
	}
}
