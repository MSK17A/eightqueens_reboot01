package piscine

import "fmt"

var x = [8]int{0,0,0,0,0,0,0,0} // x[k] = i k queen in column i (k might be the row that have the queen)
func Abs(n int) int {
	if n < 0 {
		return -1*n
	}
	return n
}
func isSafe(k, i int) bool {
	for j := 1; j < k; j++ {
		if x[j] == i || Abs(x[j]-i) == Abs(j-k) {
			return false
		}
	}
	return true
}

func NQueens(k, n int) {
	for i:=1; i<= n; i++ {
		if isSafe(k, i) {
			x[k-1] = i
			if k==n {
				for j:=0;j<len(x);j++ {
					fmt.Print(x[j])
				}
				fmt.Println("")
			} else {
				NQueens(k+1, n)
			}
		}
	}
}

func EightQueens() {
	NQueens(1, 4)
}
