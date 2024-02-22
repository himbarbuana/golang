/*
Operasi Matematika
Operator 	Keterangan
+			Penjumlahan
-			Pengurangan
*			Perkalian
/			Pembagian
%			Sisa Pembagian
*/

package main

import "fmt"

func main() {
	var a = 20
	var b = 3
	var c = a + b
	var d = a - b
	var e = a * b
	var f = a / b
	var g = a % b

	fmt.Println(c, d, e, f, g)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

	var w = 4
	var x = 15
	var y = 2
	var z = 3
	var hitung = w + x - y*z
	fmt.Println("Hasilnya =", hitung)

}
