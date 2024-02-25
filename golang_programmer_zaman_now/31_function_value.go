/*
- Function adalah first class citizen
- Function juga merupakan tipe data, dan bisa disimpan di dalam variable.
*/

package main

import "fmt"

func main() {
	goodBye := getGoodBye
	fmt.Println(goodBye("Honey"))
	fmt.Println(goodBye("My Lovely"))

	contoh := getGoodBye
	misal := getGoodBye
	fmt.Println(contoh("Contoh"))
	fmt.Println(misal("Misal"))

}

func getGoodBye(name string) string {
	return "Good Bye " + name
}
