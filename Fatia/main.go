package main

import "fmt"

func main() {
	//----- EXAMPLE 1 -----
	/*arr := [7]float64{1, 2, 3, 4, 5, 6, 7}
	slice := make([]float64, 5)
	slice := arr[0:7]
	slice := arr[2:5]
	fmt.Println(slice)*/

	//----- EXAMPLE 2 -----
	slice1 := []int{1, 2, 3}
	//slice2 := append(slice1, 4, 5)
	slice2 := make([]int, 2 /* <- 2 indica a quantidade de número a ser copiada da fatia 1*/)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)
}
