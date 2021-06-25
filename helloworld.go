package main

import "fmt"

func main() {
	var nama = "dimas"
	fmt.Println(nama)
	fmt.Println("hello, world")

	const kota = "lamongan"
	fmt.Println(kota)

	type KTP int
	type alumni string

	var nomorKtp KTP = 12345
	var alum alumni = "smk"
	fmt.Println(nomorKtp)
	fmt.Println(alum)

	var nilai32 int32 = 100000
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	var jeneng = "dwi"
	var a = jeneng[0]
	var aString = string(a)

	fmt.Println(jeneng)
	fmt.Println(aString)
}
