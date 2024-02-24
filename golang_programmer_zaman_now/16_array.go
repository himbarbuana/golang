/*
Tipe Data Array
- Array adalah tipe data yang berisikan kumpulan data dengan tipe yang sama.
- Saat membuat array, kita perlu menentukan jumlah data yang bisa ditampung oleh array tersebut.
- Daya tampung array tidak bisa bertambah setelah array dibuat.

Index di Array
Index		Data
0		Muhammad
1		Himbar
2		Buana

Notes: Index dalam array dimulai dari 0.

Membuat Array Langsung
- Di Golang kita juga bisa membuat array secara langsung saat deklarasi variable.

Function Array
Operasi				Keterangan
len(array)			Untuk mendapatkan panjang array
array[index]			Mendapat data di posisi index
array[index] = value		Mengubah data di posisi index
*/

package main

import "fmt"

func main() {
	var names [3]string   // Membuat variable names, dengan tipe data array yang hanya bisa menampung 3 data bertipe string
	names[0] = "Muhammad" // Menambahkan data ke index pertama
	names[1] = "Himbar"   // Menambahkan data ke index kedua
	names[2] = "Buana"    // Menambahkan data ke index ketiga

	fmt.Println(names)
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// Membuat data array secara langsung
	// Bisa ditulis seperti ini:
	var value = [5]int{1, 2, 3, 4, 5}

	// Atau seperti ini:
	// Notes: Jika menggunakan enter, wajib menambahkan tanda koma (,) dibelakang nilai index terakhir
	// Default value pada array dengan tipe data int adalah 0
	var values = [3]int{
		90,
		80,
	}

	var anotherValues = [4]int{
		100,
		200,
		300,
		400,
	}

	fmt.Println(value)
	fmt.Println(values)
	fmt.Println(anotherValues)

	var nilai = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(nilai)

	// Membuat array tanpa menentukan jumlah isi arraynya bisa menggunakan ...
	var nilai1 = [...]int{
		50,
		60,
		70,
	}

	fmt.Println(nilai1)
	fmt.Println(len(nilai1))
	fmt.Println(cap(nilai1))

	// Mengubah data
	nilai1[2] = 80
	fmt.Println(nilai1)

}
