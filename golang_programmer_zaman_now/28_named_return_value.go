/*
- Biasanya saat kita memberi tahu bahwa sebuah funtion mengembalikan value, maka kita hanya mendeklarasikan tipe data return value di funtion.
- Namun kita juga bisa membuat variable secara langsung di tipe data return function-nya.
*/

package main

import "fmt"

func main() {
	a, b, c := getCompleteName()
	fmt.Println(a, b, c)

	d, e, f := getCompleteName1()
	fmt.Println(d, e, f)

	g, h, i := getCompleteName2()
	fmt.Println(g, h, i)
}

func getCompleteName() (string, string, string) {
	firstName := "Muhammad"
	middleName := "Himbar"
	lastName := "Buana"
	return firstName, middleName, lastName
}

// Bisa ditulis seperti di bawah ini:
func getCompleteName1() (firstName1 string, middleName1 string, lastName1 string) {
	firstName1 = "Shama"
	middleName1 = "Inara"
	lastName1 = "Yasofa"
	return firstName1, middleName1, lastName1
}

// Karena tipe datanya sama string, penulisannya bisa diringkas sebagai berikut:
func getCompleteName2() (firstName2, middleName2, lastName2 string) {
	firstName2 = "Salman"
	middleName2 = "Fariz"
	lastName2 = "Hafizi"
	return firstName2, middleName2, lastName2

}
