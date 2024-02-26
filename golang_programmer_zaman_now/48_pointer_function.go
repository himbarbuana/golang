/*
Pointer di Function
- Saat kita membuat parameter di function, secara default adalah pass by value, artinya data akan di-copy lalu dikirim ke function tersebut.
- Oleh karena itu, jika kita mengubah data di dalam function, data yang aslinya tidak akan pernah berubah.
- Hal ini membuat variable menjadi aman, karena tidak akan bisa diubah.
- Namun kadang kita ingin membuat function yang bisa mengubah data asli parameter tersebut.
- Untuk melakukan ini, kita juga bisa menggunakan pointer di function.
- Untuk menjadikan sebuah parameter sebagai pointer, kita bisa menggunakan operator * di paramaternya.
*/

package main

import "fmt"

type MyAddress struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(myaddress *MyAddress) {
	myaddress.Country = "Indonesia"
}

func main() {
	//var myaddress *MyAddress = &MyAddress{}
	myaddress := MyAddress{}
	ChangeCountryToIndonesia(&myaddress)
	fmt.Println(myaddress)
}
