/*
Interface Kosong
- Golang bukanlah bahasa pemrograman yang berorientasi objek.
- Biasanya dalam pemrograman berorientasi objek, ada satu data parent di puncak yang bisa dianggap sebagai semua implementasi data yang ada di bahasa pemrograman tersebut.
- Contoh di Java ada java.lang.Object.
- Untuk menangani kasus seperti ini, di Golang kita bisa menggunakan interface kosong.
- Interface kosong adalah interface yang tidak memiliki deklarasi method satupun, hal ini membuat secara otomatis semua tipe data yang akan menjadi implementasinya.
- Interface kosong, juga memiliki type alias bernama any.

Penggunaan Interface Kosong di Golang
- fmt.Println(a ...interface{})
- panic(v interface{})
- recover() interface{}
- dan lain-lain
*/

package main

import "fmt"

func main() {
	kosong := Ups()
	fmt.Println(kosong)
}

func Ups() interface{} {
	return "Ups"
}
