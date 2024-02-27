/*
Package Initialization
- Saat kita membuat package, kita bisa membuat sebuah function yang akan bisa diakses ketika package kita diakses.
- Ini sangat cocok ketika contohnya, jika package kita berisi function-function untuk berkomunikasi dengan database, kita membuat function inisialisasi untuk membuka koneksi ke database.
- Untuk membuat function yang bisa diakses secara otomatis ketika package diakses, kita cukup membuat function dengan nama init.

Blank Identifier
- Kadang kita hanya ingin menjalankan init function di package tanpa harus mengeksekusi salah satu function yang ada di package.
- Secara default, golang akan komplen ketika ada package yang di-import namun tidak digunakan.
- Untuk menangani hal tersebut, kita bisa menggunakan blank identifier (_) sebelum nama package ketika melakukan import.
*/

package main

import (
	"fmt"
	"golang_programmer_zaman_now/database"
	_ "golang_programmer_zaman_now/internal"
)

func main() {
	fmt.Println(database.GetDatabase())
}
