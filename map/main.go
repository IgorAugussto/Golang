package main

import "fmt"

func main() {
	//x := make(map[string]int)
	//x["chave"] = 10
	//fmt.Println(x["chave"])
	//x := make(map[int]int)
	//x[1] = 20
	//x[2] = 30
	//fmt.Println(x[1], x[2])

	element := make(map[string]string)
	element["H"] = "Hidrogênio"
	element["He"] = "Hélio"
	element["Li"] = "Lítio"
	fmt.Println(element["Li"])
}
