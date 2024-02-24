/*
If Expression
- If adalah salah satu kata kunci yang digunakan untuk percabangan.
- Percabangan artinya kita bisa mengeksekusi kode program tertentu ketika suatu kondisi terpenuhi.
- Hampir di semua bahasa pemrograman mendukung if expression.

Else Expression
- Blok if akan dieksekusi ketika kondisi if bernilai true.
- Kadang kita ingin melakukan eksekusi program tertentu jika kondisi if bernilai false.
- Hal ini bisa dilakukan dengan menggunakan else expression.

Else If Expression
- Kadang dalam if, kita butuh membuat beberapa kondisi.
- Kasus seperti ini, kita bisa menggunakan else if expression.

If dengan Short Statement
- If mendukung short statement sebelum kondisi.
- Hal ini sangat cocok untuk membuat statement yang sederhana sebelum melakukan pengecekan terhadap kondisi.
*/

package main

import "fmt"

func main() {
	// If
	name := "Himbar"
	if name == "Himbar" {
		fmt.Println("Hello Himbar")
	}

	// Else
	// Notes: Kalau benar maka blok if akan dieksekusi, kalau salah blok else yang akan dieksekusi.
	name1 := "Buana"
	if name1 == "Himbar" {
		fmt.Println("Hello Himbar")
	} else {
		fmt.Println("Hi, boleh kenalan")
	}

	// Else If
	name2 := "Buana"
	if name2 == "Himbar" {
		fmt.Println("Hello Himbar")
	} else if name2 == "Buana" {
		fmt.Println("Hello Buana")
	} else {
		fmt.Println("Hi, boleh kenalan")
	}

	// Short statement
	name3 := "Muhammad"
	length := len(name3)
	if length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}

	// Ini menggunakan short statement:
	name4 := "Ahmad"
	if length1 := len(name4); length1 > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
