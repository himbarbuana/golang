/*
Struct
- Struct adalah sebuah template data yang digunakan untuk menggabungkan nol atau lebih tipe data lainnya dalam satu kesatuan.
- Struct biasanya representasi data dalam program aplikasi yang kita buat.
- Data di struct disimpan dalam field.
- Sederhananya struct adalah kumpulan field.

Membuat Data Struct
- Struct adalah template data atau prototype data.
- Struct tidak bisa langsung digunakan.
- Namun kita bisa membuat data/object dari struct yang telah kita buat.

Struct Literals
- Sebelumnya kita telah membuat data dari struct, namun sebenarnya ada banyak cara yang bisa kita gunakan untuk membuat struct.

Struct Method
- Struct adalah tipe data seperti tipe data lainnya, dia bisa digunakan sebagai parameter untuk function.
- Namun jika kita ingin menambahkan method ke dalam struct, sehingga seakan-akan sebuah struct memiliki function.
- Method adalah function.
*/

package main

import "fmt"

func main() {
	var cust Customer
	cust.Name = "Himbar"
	cust.Address = "Bandung"
	cust.Age = 17
	fmt.Println(cust)

	var cust1 Customer1
	cust1.Name1 = "Buana"
	cust1.Address1 = "Jakarta"
	cust1.Age1 = 18
	fmt.Println(cust1.Name1)
	fmt.Println(cust1.Address1)
	fmt.Println(cust1.Age1)

	// Struct literals
	rimba := Customer{
		Name:    "Rimba",
		Address: "Jakarta",
		Age:     21,
	}
	fmt.Println(rimba)

	almaira := Customer{"Almaira", "Malaysia", 19}
	fmt.Println(almaira)

	//Struct method
	almaira.sayHello("Faiz")
}

type Customer struct {
	Name    string
	Address string
	Age     int
}

type Customer1 struct {
	Name1, Address1 string
	Age1            int
}

// Struct method
func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "my name is", customer.Name)
}
