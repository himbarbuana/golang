/*
Recursive Function
- Recursive function adalah function yang memanggil function dirinya sendiri.
- Kadang dalam pekerjaan, kita sering menemui kasus dimana menggunakan recursive function lebih mudah dibandingkan tidak menggunakan recursive function.
- Contoh kasus yang lebih mudah diselesaikan menggunakan recursive function adalah factorial.
*/

package main

import "fmt"

func main() {
	fmt.Println(factorialLoop(10))

	resultManual := 10 * 9 * 8 * 7 * 6 * 5 * 4 * 3 * 2 * 1
	fmt.Println(resultManual)

	fmt.Println(factorialRecursive(10))
}

func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}
