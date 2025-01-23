package main

import (
	"fmt"
)

// Structs são tipos de dados compostos, ou seja, são formados por outros tipos de dados
//É uma coleção de campos, cada um tem um nome e um tipo

//Struct de usuário
//o struct é um tipo
type usuario struct {
	nome string
	idade uint8
	endereco endereco //um struct pode ter outro struct como campo. Nesse caso, o campo endereco é do tipo endereco, que é outro struct
}

type endereco struct{
	logradouro string
	numero uint8
}

func main() {
	fmt.Println("Arquivo Structs")

	var u usuario //criando uma variável do tipo usuario, que é um struct
	//para popular os campos do struct, usamos o ponto
	u.nome = "Carol"
	u.idade = 24
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua dos Bobos", 0}

	//Outra forma de declarar um struct
	usuario2 := usuario{"Fernanda", 27, enderecoExemplo}
	fmt.Println(usuario2)

	//Terceira forma de declarar um struct, quando não sabe ainda o valor de algum campo por ex, de forma instanciada
	usuario3 := usuario{nome: "Carolina"} //caso não seja passado o valor de idade, ele será 0. Se fosse um campo string, seria vazio.
	//não é possível passar algo como usuario{"Carolina"} pois o struct tem dois campos
	fmt.Println(usuario3)	
}