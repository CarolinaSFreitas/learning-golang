package main

import "fmt"

func main() {

	//Primeira forma de declarar uma variável (string) de forma explicita
	var variavel1 string = "Variável 1"

	//Segunda forma de declarar uma variável (:=) de forma implicita. Inferencia de tipo.
	variavel2 := "Variável 2"

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	//Declarando várias variáveis de uma vez, de forma explicita.
	var(
		variaviel3 string = "lalala"
		variavel4 string = "lalalalaaa"
	)

	fmt.Println(variaviel3, variavel4)

	//Declarando várias variáveis de uma vez, de forma implicita. Inferencia de tipo.
	variavel5, variavel6 := "Variável 5", "Variável 6"
	fmt.Println(variavel5, variavel6)

	//Declarando uma constante. Forma explicita.
	const constante1 string = "Constante 1"
	fmt.Println(constante1)

	//Inverter valores de variáveis
	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)
}
