/*
- Closure adalah kemampuan sebuah function berinteraksi dengan data-data disekitarnya dalam scope yang sama.
- Harap gunakan fitur closure ini dengan bijak saat kita membuat aplikasi.
*/

package main

import "fmt"

func main() {
	// Membuat variable dengan nama counter
	counter := 0

	// Membuat anonymous function yang disimpan dalam variable increment
	increment := func() {
		fmt.Println("Increment")
		counter++
	}

	increment()
	increment()
	fmt.Println(counter)
}
