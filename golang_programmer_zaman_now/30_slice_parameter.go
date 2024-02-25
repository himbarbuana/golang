/*
- Kadang ada kasus dimana kita menggunakan variadic function, namun memiliki variable berupa slice.
- Kita bisa menjadikan slice sebagai varargs parameter.
*/

package main

import "fmt"

func main() {
	total := sumAlls(10, 10, 10, 10, 10, 10)
	fmt.Println(total)

	numbers := []int{10, 10, 10, 10, 10}
	total = sumAlls(numbers...)
	fmt.Println(total)
}

func sumAlls(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
