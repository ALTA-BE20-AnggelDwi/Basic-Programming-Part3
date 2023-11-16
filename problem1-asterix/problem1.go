package main

import "fmt"

func PlayWithAsterix(n int) string {
	// your code here
	var result string

	for i := 1; i <= n; i++ {
		// Menambahkan spasi sebelum bintang agar membentuk segitiga
		for j := 1; j <= n-i; j++ {
			result += " "
		}

		// Menambahkan bintang pada setiap baris sesuai dengan nomor baris
		for k := 1; k <= i; k++ {
			result += "* "
		}

		result += "\n"
	}

	return result
}

func main() {
	fmt.Println(PlayWithAsterix(10))
}
