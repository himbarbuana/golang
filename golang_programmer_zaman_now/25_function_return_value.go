/*
- Function bisa mengembalikan data.
- Untuk memberitahu bahwa function mengembalikan data, kita harus menuliskan tipe data kembalian dari function tersebut.
- Jika function tersebut kita deklarasikan dengan tipe data pengembalian, maka wajib di dalam function-nya kita harus mengembalikan data.
- Untuk mengembalikan data dari function, kita bisa menggunakan kata kunci return, diikuti dengan datanya.
*/

package main

import "fmt"

func main() {
	result := getHello("Himbar")
	fmt.Println(result)
	fmt.Println(getHello("Buana"))
}

func getHello(name string) string {
	hello := "Hello " + name
	return hello
}
