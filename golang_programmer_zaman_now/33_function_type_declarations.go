/*
- Kadang jika function terlalu panjang, agak ribet untuk menuliskannya di dalam parameter.
- Type declarations juga bisa digunakan untuk membuat alias function, sehingga akan mempermudah kita menggunakan function sebagai parameter.
*/

package main

import "fmt"

type Filter func(string) string

func main() {
	sayHelloWithFilter("Himbar", spamFilters)

	filter := spamFilters
	sayHelloWithFilters("Anjing", filter)
}

func sayHelloWithFilters(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}

func spamFilters(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}
