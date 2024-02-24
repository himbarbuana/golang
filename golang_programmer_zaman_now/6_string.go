/*
Tipe Data String
- String adalah tipe data kumpulan karakter.
- Jumlah karakter di dalam String bisa nol sampai tak terhingga.
- Tipe data String di Golang direpresentasikan dengan kata kunci string.
- Nilai data String di Golang selalu diawali dengan karakter " (petik dua) dan diakhiri dengan karakter " (petik dua).

Function untuk String
Function		Keterangan
len("string")		Menghitung jumlah karakter di String
"string"[number]	Mengambil karakter pada posisi yang ditentukan
*/

package main

import "fmt"

func main() {
	fmt.Println("Muhammad")
	fmt.Println("Muhammad Himbar")
	fmt.Println("Muhammad Himbar Buana")
	fmt.Println(len("Himbar"))
	fmt.Println("Himbar"[0]) // Output dalam bentuk byte
	fmt.Println("Himbar"[1]) // Output dalam bentuk byte
}
