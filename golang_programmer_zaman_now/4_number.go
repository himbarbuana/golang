/*
Tipe Data Number
- Ada dua jenis tipe data Number, yaitu:
  - Integer (Bilangan bulat).
  - Floating Point (Bilangan desimal).

Tipe Data Integer (1):
Tipe Data	Nilai Minimum			Nilai Maksimum
int8		-128				127
int16		-32768				32767
int32		-2147483648			2147483647
int64		-9223372036854775808		9223372036854775807

Tipe Data Integer (2):
Tipe Data	Nilai Minimum		Nilai Maksimum
uint8		0			255
uint16		0			65535
uint32		0			4294967295
uint64		0			18446744073709551615

Tipe Data Floating Point:
Tipe Data		Nilai Minimum			Nilai Maksimum
float32			1.18x10 pangkat -38		3.4x10 pangkat 38
float 64		2.23x10 pangkat -308		1.80x10 pangkat 308
complex64		complex number with float32 real and imaginary parts
complex128		complex number with float64 real and imaginary parts

Tipe Data Alias:
Tipe Data	Alias untuk
byte		uint8
rune		int32
int		Minimal int32
uint		Minimal uint32
*/

package main

import "fmt"

func main() {
	fmt.Println("Satu =", 1)
	fmt.Println("Dua ", 2)
	fmt.Println("Tiga Koma Lima =", 3.5)
}
