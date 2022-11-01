package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("========== 1 Dimension =============")

	var arrSize int
	fmt.Print("Enter the size of array: ")
	fmt.Scan(&arrSize)

	array := make([]string, arrSize)
	var hasil string
	for i := 0; i < arrSize; i++ {
		fmt.Print("Enter the array elements: ")
		fmt.Scan(&array[i])
	}

	fmt.Println("")
	fmt.Println("===============Hasil array===============")
	for i := 0; i < arrSize; i++ {
		hasil = string(array[i])
		fmt.Printf("| %s |\t", hasil)
	}
	fmt.Println("")
	fmt.Println("============= Masukkan Keyword =============")

	var Keyword string
	fmt.Print("Enter keyword: ")
	fmt.Scan(&Keyword)

	fmt.Println("============= Result Keyword =============")
	for i := 0; i < arrSize; i++ {
		hasil = string(array[i])

		if strings.Count(hasil, Keyword) != 0 {
			fmt.Printf("| %s |\t", hasil)
		}
	}

}
