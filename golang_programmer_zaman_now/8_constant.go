/*
Constant
- Constant adalah variable yang nilai tidak bisa diubah lagi setelah pertama kali diberi nilai.
- Cara pembuatan Constant mirip dengan variable, yang membedakan hanya kata kunci yang digunakan adalah const, bukan var.
- Saat pembuatan Constant, kita wajib langsung menginisialisasikan datanya.

Constant (const) jika tidak digunakan tidak masalah.
Beda dengan var, jika dideklarasikan harus digunakan, jika tidak akan menyebabkan error pada program.

Deklarasi Multiple Constant
- Sama seperti variable, di Golang juga kita bisa membuat Constant secara sekaligus banyak.
*/

package main

import "fmt"

func main() {
	const firstName string = "Himbar"
	const lastName = "Buana"
	fmt.Println(firstName)
	fmt.Println(lastName)

	/*
		firstName dan lastName tidak bisa dirubah nilainya
		firstName = "Udin"
		lastName = "Bedog"
	*/

	// Multiple Constant
	const (
		nameFirst string = "Heung"
		nameLast         = "Boo"
	)
	fmt.Println(nameFirst)
	fmt.Println(nameLast)
}
