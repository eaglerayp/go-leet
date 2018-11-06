package ferrum

import "fmt"

const (
	RawC = 702
)

// get string notation of give n index of spreadsheet, which column from A to ZZ
func getSpreadsheetNotation(n int64) string {
	// Write your code here
	var c1 rune = 'A'
	cm := map[int]string{}
	for i := 0; i < 26; i++ {
		s := int(c1) + i
		cm[i] = string(rune(s))
	}
	for j := 0; j < 26; j++ {
		k := j + 1
		for i := 0; i < 26; i++ {
			index := k*26 + i
			cm[index] = cm[j] + cm[i]
			fmt.Println(cm[index])
		}
	}

	r := (n-1)/RawC + 1
	c := (n - 1) % RawC
	return fmt.Sprintf("%d%s", r, cm[int(c)])
}
