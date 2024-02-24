/*
Tipe Data Slice
- Tipe data slice adalah potongan dari data array.
- Slice mirip dengan array, yang membedakan adalah ukuran slice bisa berubah.
- Slice dan array selalu terkoneksi, dimana slice adalah data yang mengakses sebagian atau seluruh data di array.

Detail Tipe Data Slice
- Tipe data slice memiliki 3 data, yaitu pointer, length dan capacity.
- Pointer adalah penunjuk data pertama di array pada slice.
- Length adalah panjang dari slice, dan
- Capacity adalah kapasitas dari slice, dimana length tidak boleh lebih dari capacity.

Membuat Slice dari Array
Membuat Slice		Keterangan
array[low:high]		Membuat slice dari array dimulai index low sampai index sebelum high
array[low:]			Membuat slice dari array dimulai index low sampai index akhir di array
array[:high]		Membuat slice dari array dimulai dengan index 0 sampai index sebelum high
array[:]			Membuat slice dari array dimulai index 0 sampai index akhir di array

Function Slice
Operasi									Keterangan
len(slice)								Untuk mendapatkan panjang
cap(slice)								Untuk mendapatkan kapasitas
append(slice, data)						Membuat slice baru dengan menambah data ke posisi terakhir slice
										Jika kapasitas sudah penuh, maka akan membuat array baru
make([]TypeData, length, capacity)		Membuat slice baru
copy(destination, source)				Menyalin slice dari source ke destination
*/

package main

import (
	"fmt"
	"reflect"
)

func main() {
	names := [...]string{"Jasmine", "Aurora", "Seraphina", "Isabella", "Penelope", "Genevieve", "Valentina", "Arabella", "Anastasia", "Celeste", "Vivienne", "Camille"}

	slice1 := names[4:6]
	fmt.Println(names)
	fmt.Println(slice1)

	slice2 := names[:3]
	fmt.Println(slice2)

	slice3 := names[3:]
	fmt.Println(slice3)

	slice4 := names[:]
	fmt.Println(slice4)

	// Manual:
	var slice5 []string = names[:]
	fmt.Println(slice5)

	days := [...]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	fmt.Println(days)

	daySlice1 := days[5:]
	fmt.Println(daySlice1)

	daySlice1[0] = "New Saturday"
	fmt.Println(daySlice1)

	daySlice1[1] = "New Sunday"
	fmt.Println(daySlice1)

	daySlice2 := append(daySlice1, "New Holiday")
	fmt.Println(daySlice2)

	daySlice2[0] = "Ups"
	fmt.Println(daySlice2)
	fmt.Println(days)

	// Make slice
	newSlice := make([]string, 2, 5)
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice[0] = "Himbar"
	fmt.Println(newSlice)

	newSlice[1] = "Himbar"
	fmt.Println(newSlice)

	newSlice1 := append(newSlice, "Himbar")
	fmt.Println(newSlice1)
	fmt.Println(len(newSlice1))
	fmt.Println(cap(newSlice1))

	newSlice1[0] = "Buana"
	fmt.Println(newSlice1)
	fmt.Println(newSlice)

	// Copy slice
	fromSlice := days[:]
	fmt.Println(fromSlice)

	toSlice := make([]string, len(fromSlice), cap(fromSlice))
	copy(toSlice, fromSlice)
	fmt.Println(toSlice)

	// Notes: Hati-hati saat membuat array atau slice.
	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(iniArray)
	fmt.Println(iniSlice)

	fmt.Println(reflect.TypeOf(iniArray))
	fmt.Println(reflect.TypeOf(iniSlice))
}
