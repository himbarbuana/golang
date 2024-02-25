/*
- Saat membuat function, kadang-kadang kita membutuhkan data dari luar, atau kita sebut parameter.
- Kita bisa menambahkan parameter di function, bisa lebih dari satu.
- Parameter tidaklah wajib, jadi kita bisa membuat function tanpa parameter seperti sebelumnya yang sudah kita buat.
- Namun jika kita menambahkan parameter di function, maka ketika memanggil function tersebut, kita wajib memasukkan data ke parameternya.
*/

package main

import "fmt"

func main() {
	// Memanggil fungsi sayHelloTo dengan memasukkan parameter pertama, "Himbar", dan parameter kedua "Buana", tipe datanya string.
	sayHelloTo("Himbar", "Buana")
}

func sayHelloTo(firstName string, lastName string) {
	// firstName: parameter pertama, lastName: parameter kedua dengan tipe data string.
	fmt.Println("Hello :", firstName, lastName)
}
