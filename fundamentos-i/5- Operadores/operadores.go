package main

import "fmt"

func main() {
	// OPERADORES ARIMETICOS
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)


	//não é possível somar/dividir/etc variaveis de tipos diferentes 
	//exemplo: 
	//var numero1 int16 = 10
	// var numero2 int32 = 25 
	//soma := numero1 + numero2	
	var numero1 int16 = 10
	var numero2 int16 = 25
	soma2 := numero1 + numero2

	fmt.Println(soma2)
	//FIM OPERADORES ARITMETICOS


	//OPERADORES DE ATRIBUIÇÃO
	var variavel1 string = "String"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)
	//FIM OPERADORES DE ATRIBUIÇÃO


	//OPERADORES RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 == 2) // == é para comparar, = é para atribuir
	fmt.Println(1 <= 2)
	fmt.Println(1 != 2)
	//FIM OPERADORES RELACIONAIS


	//OPERADORES LOGICOS
	//no Go existem 3: o operador E (&&), o operador OU (||) e o operador NÃO/Negação (!)
	fmt.Println("-----------------------------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso) //E (&&) - só é verdadeiro se as duas expressões forem verdadeiras

	fmt.Println(verdadeiro || falso) //OU (||) - é verdadeiro se pelo menos uma das expressões for verdadeira

	fmt.Println(!verdadeiro) //NEGAÇÃO (!) - inverte o valor da expressão
	fmt.Println(!falso) //NEGAÇÃO (!) - inverte o valor da expressão
	//FIM OPERADORES LOGICOS


	//OPERADORES UNÁRIOS
	//são operadores que trabalham com apenas um operando
	numero := 10
	numero++	//incrementa 1
	numero += 15 //seria o mesmo que numero = numero + 15
	fmt.Println(numero)

	numero-- //decrementa 1
	numero -= 20 //seria o mesmo que numero = numero - 20
	fmt.Println(numero)

	numero *= 3 //seria o mesmo que numero = numero * 3
	fmt.Println(numero)
	numero /= 10 //seria o mesmo que numero = numero / 10
	fmt.Println(numero)
	numero %= 3 //seria o mesmo que numero = numero % 3
	fmt.Println(numero)
	//FIM OPERADORES UNÁRIOS


	//OPERADORES TERNÁRIOS
	//não existe no Go, mas é possível simular. Exemplo:
	//operador ternário em JS: var resultado = expressao ? valor1 : valor2
	//operador ternário em Go: var resultado = valor2; if expressao { resultado = valor1 }
	//exemplo:
	var texto string
	if numero > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5"
	}
	fmt.Println(texto) 			//nesse caso, deu Menor que 5 pq as contas anteriores fizeram o numero ser 1
	//FIM OPERADORES TERNÁRIOS
}