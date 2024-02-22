/*
- Di Golang, function dalam module/project adalah unik, artinya kita tidak boleh membuat nama function yang sama.
- Oleh karena itu, jika kita membuat file baru, misal sample.go, lalu membuat nama function yang sama yaitu main.
- Maka kita tidak bisa melakukan build module, karena main function tersebut duplikat dengan yang ada di main function helloworld.go.
*/

package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}
