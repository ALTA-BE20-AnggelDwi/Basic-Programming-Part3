package main

import "fmt"

func DrawXYZ(N int) string {
	// your code here
	result := ""
	for i := 1; i <= N*N; i++ {
		switch {
		case i%3 == 0:
			result += "X"
		case i%2 == 1:
			result += "Y"
		default:
			result += "Z"
		}

		if i%N == 0 {
			result += " \n"
		} else {
			result += " "
		}
	}

	return result
}

func main() {
	fmt.Println(DrawXYZ(3))
	fmt.Println(DrawXYZ(5))
}
