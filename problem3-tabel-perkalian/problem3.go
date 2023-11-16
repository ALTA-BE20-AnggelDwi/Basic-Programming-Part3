package main

import "fmt"

func CetakTabelPerkalian(number int) string {
	// your code here
	result := ""

	// Validasi nilai number
	if number < 1 || number > 30 {
		return "Nilai number harus antara 1 dan 30."
	}

	// Membangun tabel perkalian
	for i := 1; i <= number; i++ {
		for j := 1; j <= number; j++ {
			result += fmt.Sprintf("%v\t", i*j)
		}
		result += "\n"
	}

	return result
}

func main() {
	fmt.Println(CetakTabelPerkalian(9))
}
