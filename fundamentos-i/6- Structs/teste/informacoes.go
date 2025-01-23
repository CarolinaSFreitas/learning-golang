package main

import "fmt"

type eu struct {
	nome string
	idade uint8
	signo string
	bandaFavorita bandaFavorita
}

type bandaFavorita struct{
	nome string
}

func main() {
	fmt.Println("Informações")

	euMesma := eu{"Carolina", 23, "Aquário", bandaFavorita{"Atualmente, Stone Temple Pilots"}}
	fmt.Println(euMesma)
}
