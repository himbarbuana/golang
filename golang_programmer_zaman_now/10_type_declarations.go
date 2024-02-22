/*
Type Declarations
- Type Declarations adalah kemampuan membuat ulang tipe data baru dari tipe data yang sudah ada.
- Type Declarations biasanya digunakan untuk membuat alias terhadap tipe data yang sudah ada, dengan tujuan agar lebih mudah dimengerti.
*/

package main

import "fmt"

func main() {
	type NoKTP string

	var ktpHimbar NoKTP = "1111111111111111"
	fmt.Println(ktpHimbar)
	fmt.Println(NoKTP("2222222222222222"))

	var contoh string = "3333333333333333"
	var contohKtp NoKTP = NoKTP(contoh)

	fmt.Println(contohKtp)
}
