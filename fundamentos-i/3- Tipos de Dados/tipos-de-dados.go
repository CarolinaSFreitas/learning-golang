package main

import (
	"errors"
	"fmt"
)

func main(){
	//Números inteiros e reais = int e float. Porém no Go se subdividem em mais categorias.
	//Existem 4 tipos de números inteiros
	//int8, int16, int32, int64: int8 = 8 bits, int16 = 16 bits, int32 = 32 bits, int64 = 64 bits. 
	//Cada um desses tipos de inteiros tem um limite de valores que podem ser armazenados
	var numero0 int8 = 100
	fmt.Println(numero0)

	var numero1 int64 = 100000000000
	fmt.Println(numero1)
	
	//o 'int' é um tipo de dado que varia de acordo com a arquitetura do computador.
	var numero2 int = 100000000000
	fmt.Println(numero2)

	//Se não declarar o tipo de dado, o Go irá inferir o tipo de dado baseado no valor atribuido a variável.
	numero3 := 100000000000
	fmt.Println(numero3)

	//unsigned int = uint. Não aceita valores negativos.
	var numero4 uint32 = 1000
	fmt.Println(numero4)

	//alias
	//byte = uint8
	var numero5 byte = 100
	fmt.Println(numero5)

	//alias
	//rune = int32
	var numero6 rune = 1000
	fmt.Println(numero6)

	
	//Números reais
	//float32 e float64
	var numeroReal1 float32 = 123.45 //é sempre '.' e não ','
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 1234500000.45
	fmt.Println(numeroReal2)

	//o tipo 'float' só aparece quando não se declara o tipo de dado, o Go irá inferir o tipo de dado baseado no valor atribuido a variável.
	numeroReal3 := 12345.67
	fmt.Println(numeroReal3)
	//FIM NÚMEROS INTEIROS E REAIS


	//STRINGS
	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto 2"
	fmt.Println(str2)

	//char não existe no Go!! O Go trata o char como um int32!!
	//o retorno será o equivalente em número da tabela ASCII
	//não aceita mais de um caractere
	char := 'A'
	fmt.Println(char)
	//FIM STRINGS

	var texto string
	fmt.Println(texto) //valor zero de uma string é uma string vazia

	var texto0 int16
	fmt.Println(texto0) //valor zero de um int16 é 0

	//quando a declaração é implicita, é obrigatório atribuir um valor a variável
	text := 5
	fmt.Println(text) //valor zero de um int é 0


	//BOOLEANOS
	var booleano1 bool = true
	fmt.Println(booleano1)
	
	var booleano2 bool = false
	fmt.Println(booleano2)

	//se deixar sem declaração, o valor zero de um booleano é false, então o retorno será false
	var booleano3 bool
	fmt.Println(booleano3)


	//ERROS
	var erro error
	fmt.Println(erro) //valor zero de um erro é nil
	//um tipo de dado que serve como valor 0 pra muitas coisas: interfaces, ponteiros, canais, funções, slices, maps e strings

	//Para criar um erro, é necessário importar a biblioteca 'errors', que é uma biblioteca padrão do Go
	//errors.New() é uma função que cria um erro, e o valor passado como argumento é a mensagem de erro
	//não é uma string, é um dado tipo error
	var erro1 error = errors.New("Erro interno")
	fmt.Println(erro1)
}