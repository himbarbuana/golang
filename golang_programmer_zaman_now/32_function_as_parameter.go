/*
- Function tidak hanya bisa kita simpan di dalam variable sebagai value.
- Namun juga bisa kita gunakan sebagai parameter untuk function lain/
*/

package main

import "fmt"

func main() {
	sayHelloWithFilter("Himbar", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Anjing", filter)
}

func sayHelloWithFilter(name string, filter func(string) string) {
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}
