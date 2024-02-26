/*
Interface
- Interface adalah tipe data abstract, dia tidak memiliki implementasi langsung.
- Sebuah interface berisikan definisi-definisi method.
- Biasanya interface digunakan sebagai kontrak.

Implementasi Interface
- Setiap data yang sesuai dengan kontrak interface, secara otomatis dianggap sebagai interface tersebut.
- Sehingga kita tidak perlu mengimplementasikan interface secara manual.
- Hal ini agak berbeda dengan bahasa pemrograman lain yang ketika membuat interface, kita harus menyebutkan secara eksplisit akan menggunakan interface mana.
*/

package main

import "fmt"

func main() {
	person := Person{Name: "Himbar"}
	SayHello(person)

	animal := Animal{Name: "Kucing"}
	SayHello(animal)
}

type HasName interface {
	GetName() string
}

type Person struct {
	Name string
}

type Animal struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

func SayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

func (animal Animal) GetName() string {
	return animal.Name
}
