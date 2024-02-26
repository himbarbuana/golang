/*
Pass by Value
- Secara default di golang semua variable itu di-passing by value, bukan by reference.
- Artinya, jika kita mengirim sebuah variable ke dalam function, method atau variable lain, sebenarnya yang dikirim adalah duplikasi value-nya.

Pointer
- Pointer adalah kemampuan membuat reference ke lokasi data di memory yang sama, tanpa menduplikasi data yang sudah ada.
- Sederhananya, dengan kemampuan pointer, kita bisa membuat pass by reference.

Operator &
- Untuk membuat sebuah variable dengan nilai pointer ke variable lain, kita bisa menggunakan operator & diikuti dengan nama variable-nya.

Operator *
- Saat kita mengubah variable pointer, maka yang berubah hanya variable tersebut.
- Semua variable yang mengacu ke data yang sama tidak akan berubah.
- Jika kita ingin mengubah seluruh variable yang mengacu ke data tersebut, kita bisa menggunakan operator *.

Operator New
- Sebelumnya untuk membuat pointer dengan menggunakan operator &.
- Golang juga memiliki function new yang bisa digunakan untuk membuat pointer.
- Namun function new hanya mengembalikan pointer ke data kosong, artinya tidak ada data awal.
*/

package main

import "fmt"

type Address struct {
	City, Province, Country string
}

type Alamat struct {
	Kota, Provinsi, Negara string
}

type AlamatJuga struct {
	Kelurahan, Kecamatan, Kota string
}

func main() {
	address1 := Address{"Bandung", "Jawa Barat", "Indonesia"}
	fmt.Println(address1)

	address2 := address1
	address2.City = "Bogor"
	fmt.Println(address2)

	// var address1 Address = Address{"Bandung", "Jawa Barat", "Indonesia"}
	// var address3 *Address = &address1
	address3 := &address1
	address3.City = "Cianjur"
	fmt.Println(address3)

	alamat1 := Alamat{"Bekasi", "Jawa Barat", "Indonesia"}
	fmt.Println(alamat1)

	alamat2 := &alamat1
	fmt.Println(alamat2)

	alamat2.Kota = "Depok"
	fmt.Println(alamat2)

	var alamatJuga1 AlamatJuga = AlamatJuga{"Lebak Gede", "Coblong", "Bandung"}
	var alamatJuga2 *AlamatJuga = &alamatJuga1
	alamatJuga2.Kelurahan = "Cipaganti"
	fmt.Println(alamatJuga1)
	fmt.Println(alamatJuga2)

	alamatJuga2 = &AlamatJuga{"Sekeloa", "Lebak Gede", "Bandung"}
	fmt.Println(alamatJuga1)
	fmt.Println(alamatJuga2)

	*alamatJuga2 = AlamatJuga{"Pasir Kaliki", "Cicendo", "Bandung"}
	fmt.Println(alamatJuga1)
	fmt.Println(alamatJuga2)

	// var alamatJuga3 *AlamatJuga = &AlamatJuga{}
	// var alamatJuga4 *AlamatJuga = alamatJuga3
	alamatJuga3 := new(AlamatJuga)
	alamatJuga4 := alamatJuga3
	alamatJuga4.Kota = "Samarinda"
	fmt.Println(alamatJuga3)
	fmt.Println(alamatJuga4)
}
