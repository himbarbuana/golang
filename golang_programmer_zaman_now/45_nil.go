/*
Nil
- Biasanya di dalam bahasa pemrograman lain, object yang belum diinisialisasi maka secara otomatis nilainya null atau nil.
- Berbeda dengan golang, di golang saat kita buat variable dengan tipe data tertentu, maka secara otomatis akan dibuatkan default valuenya.
- Namun di golang ada data nil, yaitu data kosong.
- Nil sendiri hanya bisa digunakan di beberapa tipe data, seperti interface, map, slice, pointer dan channel.
*/

package main

import "fmt"

func main() {
	data := NewMap("")
	if data == nil {
		fmt.Println("Data kosong")
	} else {
		fmt.Println(data)
	}
}

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}
