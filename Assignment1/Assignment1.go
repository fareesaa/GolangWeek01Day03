package main

import (
	"assignment1/geometry"
	"fmt"
)

func main() {

	var pil int

	fmt.Println("Menu")
	fmt.Println("1.Menghitung Volume Kubus")
	fmt.Println("2.Menghitung Volume Bola")
	fmt.Println("3.Menghitung Umur")
	fmt.Println("4.EXIT")
	fmt.Println("====================")
	fmt.Print("Silahkan pilih menu: ")
	fmt.Scanf("%d\n", &pil)
	fmt.Println("====================")

	switch pil {
	case 1:
		fmt.Println("Menghitung Volume Kubus")
		fmt.Println("________________________")
		var s int

		fmt.Print("Masukkan panjang sisi: ")
		fmt.Scanf("%d\n", &s)
		cube := geometry.Cube{S: s}

		fmt.Printf("Volume kubus tersebut adalah : %d cm^3 \n", cube.VolCube())

		break
	case 2:
		fmt.Println("Menghitung Volume Bola")
		fmt.Println("________________________")

		var r float64

		fmt.Print("Masukkan panjang jari-jari: ")
		fmt.Scanf("%f\n", &r)

		ball := geometry.Ball{Radius: r}

		fmt.Printf("Volume Bola tersebut adalah : %2f cm^3 \n", ball.VolBall())

		break
	case 3:
		fmt.Println("Menghitung Umur")
		fmt.Println("________________________")

		var b int
		var l int

		fmt.Print("Masukkan Tahun Lahir: ")
		fmt.Scanf("%d\n", &b)
		fmt.Print("Masukkan Tahun Saat Ini: ")
		fmt.Scanf("%d\n", &l)

		countAge := geometry.CountAge{B: b, L: l}

		fmt.Printf("Umur anda saat ini adalah: %d tahun\n", countAge.CountingAge())
		break
	default:
		fmt.Println("====Exit Program====")

	}
}
