/*
- Function tidak hanya dapat mengembalikan satu value, tapi bisa juga multiple value.
- Untuk memberitahu jika function mengembalikan multiple value, kita harus menulis semua tipe data return value-nya di function.
*/

package main

import "fmt"

func main() {
	firstName, lastName := getMyFullName()
	fmt.Println(firstName, lastName)
}

func getMyFullName() (string, string) {
	return "Himbar", "Buana"
}
