package main

import(
	"Nomer_Nol/sum"
	"fmt"
)

func main(){
	hidup1, mati2 := sum.Sum("omama")
	fmt.Printf("test omama have %v vocal and %v consonant\n", hidup1, mati2)
	fmt.Println("")
	fmt.Println("===================||==================||============")
	hidup, mati := sum.Sum("osama")
	fmt.Printf("test omama have %v vocal and %v consonant\n", hidup, mati)
}
