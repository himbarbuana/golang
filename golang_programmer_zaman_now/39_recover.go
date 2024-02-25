/*
Recover
- Recover adalah function yang bisa kita gunakan untuk menangkap data panic.
- Dengan recover proses panic akan terhenti, sehingga program akan tetap berjalan.
*/

package main

import "fmt"

func main() {
	runMyApp(true)
	runYourApp(true)
}

// Recover yang salah
func endMyApp() {
	fmt.Println("End app")
}

func runMyApp(error bool) {
	defer endMyApp()
	if error {
		panic("Error")
	}
	message := recover()
	fmt.Println("Terjadi error", message)
}

// Recover yang benar
func endYourApp() {
	fmt.Println("End your app")
	messages := recover()
	fmt.Println("Terjadi error", messages)
}

func runYourApp(error bool) {
	defer endYourApp()
	if error {
		panic("Error")
	}
}
