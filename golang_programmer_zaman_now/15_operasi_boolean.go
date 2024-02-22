/*
Operasi Boolean
Operator 	Keterangan
&&		Dan
||		Atau
!		Kebalikan

Operasi &&
Nilai 1		Operator	Nilai 2 	Hasil
true		&&		true		true
true		&&		false		false
false		&&		true		false
false		&&		false		false

Notes: dalam operasi && jika nilai dua-duanya true, hasilnya true. Jika salah satunya bernilai false, maka hasilnya false.

Operasi ||
Nilai 1		Operator	Nilai 2 	Hasil
true		||		true		true
true		||		false		true
false		||		true		true
false		||		false		false

Notes: jika salah satu nilainya true, hasilnya akan true.

Operasi !
Operator 	Nilai 		Hasil
!		true		false
!		false		true
*/

package main

import "fmt"

func main() {
	var nilaiAkhir = 90
	var absensi = 80
	var lulusNilaiAkhir bool = nilaiAkhir > 80
	var lulusAbsensi bool = absensi > 80
	var lulus bool = lulusNilaiAkhir && lulusAbsensi
	fmt.Println(lulus)
}
