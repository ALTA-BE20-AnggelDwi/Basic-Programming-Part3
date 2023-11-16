package main

import "fmt"

func UbahHuruf(sentence string) string {
	// your code here
	var result string

	for _, char := range sentence {
		if char >= 'A' && char <= 'Z' {
			// Enkripsi huruf kapital
			var encrypted int32 = 'A' + (char-'A'+int32(10))%26
			result += string(encrypted)
		} else if char >= 'a' && char <= 'z' {
			// Enkripsi huruf kecil
			var encrypted int32 = 'a' + (char-'a'+int32(10))%26
			result += string(encrypted)
		} else {
			// Karakter selain huruf tidak dienkripsi
			result += string(char)
		}
	}

	return result
}

func main() {
	fmt.Println(UbahHuruf("SEPULSA OKE"))     // COZEVCK YUO
	fmt.Println(UbahHuruf("ALTERRA ACADEMY")) // KVDOBBK KMKNOWI
	fmt.Println(UbahHuruf("INDONESIA"))       // SXNYXOCSK
	fmt.Println(UbahHuruf("GOLANG"))          // QYVKXQ
	fmt.Println(UbahHuruf("PROGRAMMER"))      // ZBYQBKWWOB
	fmt.Println(UbahHuruf("programmer"))      // ZBYQBKWWOB
}
