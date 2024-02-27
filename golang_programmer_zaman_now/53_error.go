/*
Error Interface
- Golang memiliki interface yang digunakan sebagai kontrak untuk membuat error, nama interface-nya adalah error.

Membuat Error
- Untuk membuat error, kita tidak perlu manual.
- Golang sudah menyediakan library untuk membuat helper secara mudah, yang terdapat di package errors (package akan kita bahas secara detail di materi tersendiri).
*/

package main

import (
	"errors"
	"fmt"
)

func main() {
	hasil, err := Pembagian(100, 0)
	if err == nil {
		fmt.Println("Hasil", hasil)
	} else {
		fmt.Println("Error", err.Error())
	}
}

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("pembagian dengan nol")
	} else {
		return nilai / pembagi, nil
	}
}
