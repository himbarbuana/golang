/*
Type Assertions
- Type assertions merupakan kemampuan merubah tipe data menjadi tipe data yang diinginkan.
- Fitur ini sering sekali digunakan ketika kita bertemu dengan data interface kosong.

Type Assertions Menggunakan Switch
- Saat salah menggunakan type assertions, maka bisa berakibat terjadi panic di aplikasi kita.
- Jika panic dan tidak ter-recover, maka otomatis program kita akan mati.
- Agar lebih aman, sebaiknya kita menggunakan switch expression untuk melakukan type assertions.
*/

package main

import "fmt"

func main() {
	var result any = random()
	var resultString string = result.(string)
	fmt.Println(resultString)

	result1 := random()
	switch value := result1.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown", value)

	}
}

func random() any {
	return "Ok"
}
