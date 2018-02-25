package intowow

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func CountCountry(A [][]int) int {
	// write your code in Go 1.4
	country := 1
	row := len(A)
	col := len(A[0])
	cMap := [][]int{}
	for i := 0; i < row; i++ {
		r := []int{}
		for j := 0; j < col; j++ {
			r = append(r, 0)
		}
		cMap = append(cMap, r)
	}
	cMap[0][0] = country
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if cMap[i][j] == 0 {
				left, top := false, false
				// as left
				if j > 0 && A[i][j] == A[i][j-1] {
					cMap[i][j] = cMap[i][j-1]
					left = true
				}
				// as up
				if i > 0 && A[i][j] == A[i-1][j] {
					cMap[i][j] = cMap[i-1][j]
					top = true
				}
				if left && top {
					if cMap[i][j-1] != cMap[i-1][j] {
						// make left = up
						cMap[i][j-1] = cMap[i-1][j]
						country--
					}
					continue
				}
				if left || top {
					continue
				}
				// new country
				country++
				cMap[i][j] = country
			}
		}
	}
	// for i := 0; i < row; i++ {
	// 	log.Println(cMap[i])
	// }
	return country
}
