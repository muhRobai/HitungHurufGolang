package main

import(
	"Nomer_Nol/sum"
	"fmt"
)

func main(){
	hidup1, mati2 := sum.Sum("omama")
	fmt.Println(hidup1, mati2)
	fmt.Println("")
	fmt.Println("===================||==================||============")
	hidup, mati := sum.Sum("osama")
	fmt.Println(hidup, mati)
}
