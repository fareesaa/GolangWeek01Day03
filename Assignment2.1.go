package main

import (
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
		fmt.Println("========== 1 Dimension =============")

		var arrSize int
		fmt.Print("Enter the size of array: ")
		fmt.Scan(&arrSize)
		fmt.Print("Enter the array elements: ")

		array := make([]float32, arrSize)
		var hasil int
		for i := 0; i < arrSize; i++ {
			fmt.Scan(&array[i])
		}
		fmt.Println("Hasil array")
		for i := 0; i < arrSize; i++ {
			hasil = int(array[i])
			fmt.Printf("%d \t", hasil)
		}
		break
	case 2:
		fmt.Println("========== 2 Dimension =============")
		fmt.Println()

		var dim [100][100]int

		var row, col int

		fmt.Println("Enter the rows of array: ")
		fmt.Scan(&row)
		fmt.Println("Enter the cols of array: ")
		fmt.Scan(&col)

		fmt.Println("Enter the array elements: ")

		for i := 0; i < row; i++ {
			for j := 0; j < col; j++ {
				fmt.Printf("Enter the element for dimension %d %d :", i+1, j+1)
				fmt.Scan(&dim[i][j])
			}
		}

		for i := 0; i < row; i++ {
			for j := 0; j < col; j++ {
				fmt.Printf(" %d ", dim[i][j])
				if j == col-1 {
					fmt.Println("")
				}
			}
		}

		break
	default:
		fmt.Println("====Exit Program====")

	}

}
