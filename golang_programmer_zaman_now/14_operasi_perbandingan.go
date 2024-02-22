/*
Operasi Perbandingan
- Operasi perbandingan adalah operasi untuk membandingkan dua buah data.
- Operasi perbandingan adalah operasi yang menghasilkan nilai boolean (benar atau salah).
- Jika hasil operasinya adalah benar, maka nilainya adalah true.
- Jika hasil operasinya salah, maka nilainya false.

Operator Perbandingan
Operator 		Keterangan
>				Lebih dari
<				Kurang dari
>=				Lebih dari sama dengan
<=				Kurang dari sama dengan
==				Sama dengan
!=				Tidak sama dengan
*/

package main

import "fmt"

func main() {
	var name1 = "Himbar"
	var name2 = "Himbar"
	var nameResult bool = name1 == name2
	fmt.Println(nameResult)
}
