package main

import (
	"Assigment/test"
	"fmt"
	"os"
	"strconv"
)

type mahasiswa struct {
	id        int
	name      string
	alamat    string
	pekerjaan string
	alasan    string
}

var allMahasiswa = []mahasiswa{
	{id: 1, name: "Yuu", alamat: "Binus", pekerjaan: "deadwood", alasan: "Non"},
	{id: 2, name: "Dery", alamat: "Sumut", pekerjaan: "IT", alasan: "kakkakaa"},
	{id: 3, name: "Aldary", alamat: "Anggrek", pekerjaan: "Prog", alasan: "lapopasd"},
	{id: 4, name: "Van", alamat: "Reccc", pekerjaan: "Tidak Ada", alasan: "D DD"},
}

func getAbsen(str string) string {
	absen, _ := strconv.Atoi(str)

	var a int = len(allMahasiswa)

	for i, mhs := range allMahasiswa {

		if mhs.id == absen {

			fmt.Println(mhs.id, " Nama : ", mhs.name, " Alamat :", mhs.alamat, " Pekerjaan : ", mhs.pekerjaan+" Alasan :"+mhs.alasan)
			fmt.Println(i)

		}

	}
	if absen > a || absen == 0 {

		fmt.Println("No Absen Tidak Ditemukan")

	}

	return str
}

func main() {

	str := os.Args[1]
	getAbsen(str)

	names := []string{"Yudistiro", "Wahyu", "Aldary"}
	test.Student(names...)

}
