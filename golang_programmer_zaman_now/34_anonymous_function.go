/*
- Sebelumnya setiap membuat function, kita akan selalu memberikan sebuah nama pada function tersebut.
- Namun kadang ada kalanya lebih mudah membuat function secara langsung di variable atau parameter tanpa harus membuat function terlebih dahulu.
- Hal tersebut dinamakan anonymous function, atau function tanpa nama.
*/

package main

import "fmt"

type Blacklist func(string) bool

func main() {
	blacklist := func(name string) bool {
		return name == "Babi"
	}

	registerUser("Buana", blacklist)

	registerUser("Himbar", func(name string) bool {
		return name == "Anjing"
	})
}

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}
