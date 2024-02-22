/*
Variable
- Variable adalah tempat untuk menyimpan data.
- Variable digunakan agar kita bisa mengakses data yang sama dimanapun kita mau.
- Di Golang Variable hanya bisa menyimpan tipe data yang sama, jika kita ingin menyimpan data yang berbeda-beda jenis, kita harus membuat beberapa Variable.
- Untuk membuat Variable, kita bisa menggunakan kata kunci var, lalu diikuti dengan nama Variable dan tipe datanya.

Tipe Data Variable
- Saat kita membuat Variable, maka kita wajib menyebutkan tipe data Variable tersebut.
- Namun jika kita langsung menginisialisasikan data pada Variablenya, maka kita tidak wajib menyebutkan tipe data Variablenya.

Kata Kunci Var
- Di Golang, kata kunci var saat membuat Variable tidaklah wajib.
- Asalkan saat membuat Variable kita langsung menginisialisasikan datanya.
- Agar tidak perlu menggunakan kata kunci var, kita perlu menggunakan kata kunci := saat menginisialisasikan data pada Variable tersebut.

Deklarasi Multiple Variable
- Di Golang kita bisa membuat Variable secara sekaligus banyak.
- Code yang dibuat akan lebih bagus dan mudah dibaca.
*/

package main

import "fmt"

func main() {
	// Deklarasi variable dengan tipe data string.
	var name string

	name = "Himbar Buana"
	fmt.Println(name)

	name = "Muhammad Himbar Buana"
	fmt.Println(name)

	// Deklarasi Variable langsung tanpa menuliskan tipe datanya.
	var nama = "M. Himbar Buana"
	fmt.Println(nama)

	nama = "Himbar"
	fmt.Println(nama)

	// Membuat Variable dengan ":=", hanya digunakan untuk deklarasi pertama saja.
	kota := "Bandung"
	fmt.Println(kota)

	kota = "Jakarta"
	fmt.Println(kota)

	// Deklarasi Multiple Variable.
	var (
		firstName = "Himbar"
		lastName  = "Buana"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
}
