/*
- Parameter yang berada di posisi terakhir, memiliki kemampuan dijadikan varargs/variable arguments.
- Varargs artinya datanya bisa menerima lebih dari satu input, atau anggap saja semacam Array.
- Apa bedanya dengan parameter biasa dengan tipe data Array?
  * Jika parameter tipe Array, kita wajib membuat array-nya terlebih dahulu sebelum mengirimkan ke function.
  * Jika parameter menggunakan varargs, kita bisa langsung mengirim datanya, jika lebih dari satu, cukup gunakan tanda koma.
*/

package main

import "fmt"

func main() {
	total := sumAll(100, 200, 300, 400, 500)
	fmt.Println(total)
}

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
