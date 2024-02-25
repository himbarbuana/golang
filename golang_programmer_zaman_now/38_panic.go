/*
Panic
- Panic function adalah function yang bisa kita gunakan untuk menghentikan program.
- Panic function biasanya dipanggil ketika terjadi panic pada saat program kita berjalan.
- Saat panic function dipanggil, program akan terhenti, namun defer function tetap akan dieksekusi.
*/

package main

import "fmt"

func main() {
	runApp(false)
}

func endApp() {
	fmt.Println("End app")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("error")
	}
}
