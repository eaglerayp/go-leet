package algo

import (
	"fmt"
	"reflect"
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

func TestThreeSum(t *testing.T) {
	if reflect.DeepEqual(threeSum([]int{-1, 0, 1, 2, -1, -4}), [][]int{[]int{-1, 0, 1}, []int{-1, -1, 2}}) {
		fmt.Println("T1")
		t.Fail()
	}
	if reflect.DeepEqual(threeSum([]int{-1, 0, 2}), []int{}) {
		fmt.Println("T2")
		t.Fail()
	}
	if reflect.DeepEqual(threeSum([]int{-1, 0}), []int{}) {
		fmt.Println("T3")
		t.Fail()
	}
	// if reflect.DeepEqual(threeSum([]int{-2, 0, 0, 2, 2}), [][]int{[]int{-2, 0, 2}}) {
	// 	fmt.Println("T4")
	// 	t.Fail()
	// }
	if reflect.DeepEqual(threeSum([]int{-2, -3, 0, 0, -2}), []int{}) {
		fmt.Println("T5")
		t.Fail()
	}
	// fmt.Println(threeSum([]int{13, 14, 1, 2, -11, -11, -1, 5, -1, -11, -9, -12, 5, -3, -7, -4, -12, -9, 8, -13, -8, 2, -6, 8, 11, 7, 7, -6, 8, -9, 0, 6, 13, -14, -15, 9, 12, -9, -9, -4, -4, -3, -9, -14, 9, -8, -11, 13, -10, 13, -15, -11, 0, -14, 5, -4, 0, -3, -3, -7, -4, 12, 14, -14, 5, 7, 10, -5, 13, -14, -2, -6, -9, 5, -12, 7, 4, -8, 5, 1, -10, -3, 5, 6, -9, -5, 9, 6, 0, 14, -15, 11, 11, 6, 4, -6, -10, -1, 4, -11, -8, -13, -10, -2, -1, -7, -9, 10, -7, 3, -4, -2, 8, -13}))
}
