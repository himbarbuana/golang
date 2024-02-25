/*
- Multiple return value wajib ditangkap semua value-nya.
- Jika kita ingin menghiraukan value tersebut, kita bisa menggunakan tanda _ (underscore/garis bawah).
*/

package main

import "fmt"

func main() {
	firstName, _ := getFullName()
	fmt.Println(firstName)
}

func getFullName() (string, string) {
	return "Himbar", "Buana"
}
