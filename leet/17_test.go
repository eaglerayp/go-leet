package leet

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	if !reflect.DeepEqual(letterCombinations("23"), []string{"ad", "bd", "cd", "ae", "be", "ce", "af", "bf", "cf"}) {
		fmt.Println("t1 fail", letterCombinations("23"))
		t.Fail()
	}
	if !reflect.DeepEqual(letterCombinations("2"), []string{"a", "b", "c"}) {
		fmt.Println("t2 fail", letterCombinations("2"))
		t.Fail()
	}
}
