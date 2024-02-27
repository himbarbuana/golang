/*
Access Modifier
- Dalam bahasa pemrograman lain, biasanya ada kata kunci yang bisa digunakan untuk menentukan access modifier terhadap suatu function atau variable.
- Di golang, untuk menentukan access modifier, cukup dengan nama function atau variable.
- Jika namanya diawali dengan huruf besar, maka artinya bisa diakses dari package lain, jika dimulai dengan huruf kecil, artinya tidak bisa diakses dari package lain.
*/

package main

var version = "1.0.0"      // Tidak bisa diakses dari luar
var Application = "Golang" // Bisa diakses dari luar

func main() {

}

// sayYouLoveMe tidak bisa diakses dari luar package
func sayYouLoveMe(name string) string {
	return "Love you " + name
}

// WillYouMarryMe bisa diakses dari luar package
func WillYouMarryMe(name string) string {
	return "Will you marry me " + name
}
