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
	if myAtoi("999999999999999999999999999") != 2147483647 {
		fmt.Println(myAtoi("999999999999999999999999999"))
		t.Fail()
	}

	if myAtoi("+-2") != 0 {
		t.Fail()
	}

	if myAtoi("  - 2") != 0 {
		t.Fail()
	}

	if myAtoi("-2147483648") != -2147483648 {
		fmt.Println("-2147483648:", myAtoi("2147483648"))
		t.Fail()
	}
	if myAtoi("-1234a568") != -1234 {
		fmt.Println(myAtoi("-1234a568"))
		t.Fail()
	}
	if myAtoi("2147483648") != 2147483647 {
		fmt.Println("2147483648:", myAtoi("2147483648"))
		t.Fail()
	}
	if myAtoi("-2147483649") != -2147483648 {
		fmt.Println("-2147483649:", myAtoi("-2147483649"))
		t.Fail()
	}

}

func TestMaxArea(t *testing.T) {
	if maxArea([]int{1, 2, 3, 4, 5}) != 6 {
		fmt.Println("T1")
		t.Fail()
	}
	if maxArea([]int{1}) != 0 {
		fmt.Println("T2")
		t.Fail()
	}
	if maxArea([]int{3, 3}) != 3 {
		fmt.Println("T3")
		t.Fail()
	}
}
