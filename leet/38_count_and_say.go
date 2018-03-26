package leet

import (
	"bytes"
	"fmt"
)

func countAndSay(n int) string {
	if n < 1 {
		return ""
	}

	s := "1"
	for i := 1; i < n; i++ {
		s = say(s)
	}
	return s
}

func say(s string) string {
	n := len(s)
	lastC := s[0]
	count := 1
	str := bytes.Buffer{}
	for i := 1; i < n; i++ {
		if s[i] == lastC {
			count++
		} else {
			str.WriteString(fmt.Sprintf("%d%c", count, lastC))
			count = 1
			lastC = s[i]
		}
	}
	str.WriteString(fmt.Sprintf("%d%c", count, lastC))
	return str.String()
}
