/*
Augmented Assignments
Operasi Matematika 	Augmented Assignments
a = a + 10		a += 10
a = a - 10		a -= 10
a = a * 10		a *= 10
a = a / 10		a /= 10
a = a % 10		a %= 10
*/
package main

import "fmt"

func main() {
	var i = 10
	fmt.Println(i)

	i += 10 // i = i + 10
	fmt.Println(i)

	i += 5 // i = i + 5
	fmt.Println(i)
}
