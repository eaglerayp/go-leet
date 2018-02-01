package main

func MaxMStr(s string, m int) int {
	Max := 0
	Cur := 0
	charSet := map[rune]bool{}
	for i, c := range s {
		if charSet[c] {
			// same case
			Cur = Cur + 1
			if Cur > Max {
				Max = cur
			}
			continue
		}
		if len(charSet) < m {
			//initial case
			charSet[c] = true
			Cur = cur + 1
			if Cur > Max {
				Max = cur
				continue
			}

		}
		// change char case, re init Cur, charSet
		Cur = 0
		clearMap(charSet)
		charSet[c] = true
		for j := i; j > 0; j-- {
			prevC := s[j]
			if charSet[prevC] {
				// same case
				Cur = Cur + 1
				if Cur > Max {
					Max = cur
				}
				continue
			}
			if len(charSet) < m {
				//initial case
				charSet[prevC] = true
				Cur = cur + 1
				if Cur > Max {
					Max = cur
				}
				continue
			}
			//max m char
			break
		}
	}
	return max
}
