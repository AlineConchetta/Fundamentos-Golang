package main

import "fmt"

// Struct dentro de go é o que tem de mais próximo de um objeto que há numa linguagem orientada a objetos.
// Struct é uma coleção de campos.Cada campo tem um nome e um tipo. Struct é um tipo.

type usuario struct {
	nome     string
	idade    int8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint
}

func main() {
	fmt.Println("Arquivo Struct")

	var u usuario
	u.nome = "Davi"
	u.idade = 12

	fmt.Println(u)

	enderecoExemplo := endereco{"Rua dos bobos", 0}

	usuario2 := usuario{"Davi", 12, enderecoExemplo}
	fmt.Println(usuario2)

	usuario3 := usuario{idade: 21} //especificar o campo para retornar ele de volta unitária
	fmt.Println(usuario3)
}
