package main

import "fmt"

//No Go, os parâmetros de uma função são tipados, ou seja, é necessário informar o tipo de dado que a função irá receber
//E o tipo de dado que a função irá retornar
func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	//função abaixo é uma função anônima, ou seja, não possui nome e é atribuída a uma variável
	var f = func(){
		fmt.Println("Função f")
	}
	f()

	//função anônima com parâmetro, é necessário informar o tipo de dado que a função irá receber
	//e o tipo de dado que a função irá retornar
	var f1 = func(txt string){
		fmt.Println(txt)
	}
	f1("Texto da Função 1")

	//função anônima com parâmetro e retorno
	var f2 = func(txt string) string{
		fmt.Println(txt)
		return txt
	}

	resultado := f2("Texto da Função 2")
	fmt.Println(resultado)

	//função com múltiplos retornos
	resultadosSoma, resultadosSubtracao := calculosMatematicos(10, 15)
	fmt.Println(resultadosSoma, resultadosSubtracao)

	//caso queira ignorar um dos retornos, basta utilizar o underline. No exemplo abaixo, o retorno da subtração será ignorado
	resultadosSoma, _ = calculosMatematicos(10, 15)
	fmt.Println(resultadosSoma)

	//exemplo igrano o retorno da soma
	_ , resultadosSubtracao = calculosMatematicos(10, 15)
	fmt.Println(resultadosSubtracao)
}

//funções no Go podem ter múltiplos retornos, como no exemplo abaixo que retorna a soma e a subtração de dois números
//é necessário informar o tipo de dado que a função irá retornar
//no caso abaixo, a função irá retornar dois inteiros de 8 bits

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}
