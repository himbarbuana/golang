/*
Switch Expression
- Selain if expression, untuk melakukan percabangan, kita juga bisa menggunakan switch expression.
- Switch expression sangat sederhana dibandingkan if.
- Biasanya switch expression digunakan untuk melakukan pengecekan ke kondisi dalam satu variable.

Switch dengan Short Statement
- Sama dengan if, switch juga mendukung short statement sebelum variable yang akan dicek kondisinya.

Switch Tanpa Kondisi
- Kondisi di switch expression tidak wajib.
- Jika kita tidak menggunakan kondisi di switch expression, kita bisa menambahkan kondisi tersebut di setiap casenya.
*/

package main

import "fmt"

func main() {
	name := "Himbar"
	switch name {
	case "Himbar":
		fmt.Println("Hello Himbar")
	case "Buana":
		fmt.Println("Hello Buana")
	default:
		fmt.Println("Hallo, boleh kenalan?")
	}

	// Short statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	// Switch tanpa kondisi
	name1 := "Buana"
	length1 := len(name1)
	switch {
	case length1 > 10:
		fmt.Println("Nama terlalu panjang")
	case length1 > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
