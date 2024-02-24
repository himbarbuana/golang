/*
Tipe Data Map
- Pada array atau slice, untuk mengakses data, kita menggunakan index number dimulai dari 0.
- Map adalah tipe data lain yang berisikan kumpulan data yang sama, namun kita bisa menentukan jenis tipe data index yang akan kita gunakan.
- Sederhananya, map adalah tipe data kumpulan key-value (kata kunci - nilai), dimana kata kuncinya bersifat unik, tidak boleh sama.
- Berbeda dengan array dan slice, jumlah data yang kita masukkan ke dalam map boleh sebanyak-banyaknya, asalkan kata kuncinya berbeda, jika kita gunakan kata kunci yang sama, maka secara otomatis data sebelumnya akan diganti dengan data baru.

Function Map
Operasi 						Keterangan
len(map)						Untuk mendapatkan jumlah data di map
map[key]						Mengambil data di map dengan key
map[key] = value				Mengubah data di map dengan key
make(map[TypeKey]TypeValue)		Membuat map baru
delete(map, key)				Menghapus data di map dengan key
*/

package main

import "fmt"

func main() {
	// Map kosong, akan mengembalikan nilai kosong sesuai dengan tipe data pada key-value pairnya.
	var person map[string]string = map[string]string{}
	fmt.Println(person)

	var nilai map[int]int = map[int]int{}
	fmt.Println(nilai)

	person["name"] = "Himbar"
	person["address"] = "Bandung"
	fmt.Println(person)

	// Menulis map lebih ringkas
	person1 := map[string]string{
		"name":    "Buana",
		"address": "Jakarta",
	}
	fmt.Println(person1["name"])
	fmt.Println(person1["address"])
	fmt.Println(person1)

	book := make(map[string]string)
	book["title"] = "Buku Golang"
	book["author"] = "Himbar"
	book["wrong"] = "Ups"
	fmt.Println(book)

	delete(book, "wrong")
	fmt.Println(book)
}
