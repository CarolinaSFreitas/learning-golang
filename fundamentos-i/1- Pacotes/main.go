package main

import (
	"fmt"
	"modulo/auxiliar"
	
	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo no arquivo main")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("carolina@hotmail.com")
	fmt.Println(erro)

}