package main

import (
	"assignmentplus/shape"
	"fmt"
)

func main() {

	var pil int

	fmt.Println("Menu")
	fmt.Println("1.Menghitung Luas Segitiga")
	fmt.Println("2.Menghitung Luas Lingkaran")
	fmt.Println("3.Menghitung Luas Persegi")
	fmt.Println("4.EXIT")
	fmt.Println("====================")
	fmt.Print("Silahkan pilih menu: ")
	fmt.Scanf("%d\n", &pil)
	fmt.Println("====================")

	switch pil {
	case 1:
		fmt.Println("Menghitung Luas Segitiga")
		fmt.Println("________________________")
		var a int
		var t int

		fmt.Print("Masukkan panjang alas segitiga: ")
		fmt.Scanf("%d\n", &a)
		fmt.Print("Masukkan panjang tinggi segitiga: ")
		fmt.Scanf("%d\n", &t)

		triangle := shape.Triangle{A: a, T: t}

		fmt.Printf("Triangle area: %d cm^2 \n", triangle.AreaTri())

		break
	case 2:
		fmt.Println("Menghitung Luas Lingkaran")
		fmt.Println("________________________")

		var r float64

		fmt.Print("Masukkan panjang jari-jari: ")
		fmt.Scanf("%f\n", &r)

		circle := shape.Circle{Radius: r}
		fmt.Printf("Circle area: %2f cm^2 \n", circle.AreaCir())
		break
	case 3:
		fmt.Println("Menghitung Luas Persegi")
		fmt.Println("________________________")

		var s int

		fmt.Print("Masukkan panjang sisi: ")
		fmt.Scanf("%d\n", &s)

		rectangle := shape.Rectangle{S: s}

		fmt.Printf("Triangle area: %d cm^2 \n", rectangle.AreaRect())
		break
	default:
		fmt.Println("====Exit Program====")

	}
}
