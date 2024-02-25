/*
For
- Dalam bahasa pemrograman, biasanya ada fitur yang bernama perulangan.
- Salah satu fitur perulangan adalah for loops.

For dengan Statement
- Dalam for, kita bisa menambahkan statement, dimana terdapat 2 statement yang bisa kita tambahkan di for.
- Init statement, yaitu statement sebelum for dieksekusi.
- Post statement, yaitu statement yang akan selalu dieksekusi diakhir tiap perulangan.

For Range
- For bisa digunakan untuk melakukan iterasi terhadap semua data collection.
- Data collection contohnya array, slice dan map.
*/

package main

import "fmt"

func main() {
	counter := 1
	for counter <= 10 {
		fmt.Println("Perulangan ke ", counter)
		counter++
	}
	fmt.Println("Selesai")

	// For dengan statement
	for counter1 := 1; counter1 <= 10; counter1++ {
		fmt.Println("Perulangan yang ke ", counter1)
	}

	// For range
	// Manual:
	names := []string{"Muhammad", "Himbar", "Buana"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// Menggunakan for range:
	names1 := []string{"Almira", "Ulfa", "Danita"}
	for index, name := range names1 {
		fmt.Println("index", index, "=", name)
	}

	nilai := []int{100, 200, 300, 400, 500}
	for i, n := range nilai {
		fmt.Println("Nilai", i, ":", n)
	}

	// Tanpa index:
	jumlah := []int{10, 20, 30, 40, 50}
	for _, j := range jumlah {
		fmt.Println(j)
	}
}
