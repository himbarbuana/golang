/*
Package
- Package adalah tempat yang bisa digunakan untuk mengorganisir kode program yang kita buat di golang.
- Dengan menggunakan package, kita bisa merapikan kode program yang kita buat.
- Package sendiri sebenarnya hanya direktori folder di sistem operasi kita.

Import
- Secara standar, file golang hanya bisa mengakses file golang lainnya yang berada dalam package yang sama.
- Jika kita ingin mengakses file golang yang berada di luar package, maka kita bisa menggunakan import.
*/

package main

import (
	"fmt"
	"golang_programmer_zaman_now/helper"
)

func main() {
	result := helper.SayHi("Udin")
	fmt.Println(result)
}
