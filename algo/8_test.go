package algo

import (
	"fmt"
	"testing"
)

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
