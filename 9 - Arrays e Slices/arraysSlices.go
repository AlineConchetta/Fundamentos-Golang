package main

import (
	"fmt"
)

func main() {
	fmt.Println("Arrays e Slices")

	//ARRAYS

	// É uma variavel que guarda uma série de valores que precisam ser do mesmo tipo.
	// Existem duas formas de criar arrays.
	// O array é fixo,inflexivel, ou seja, não permite mais itens do que seu tamanho

	var array1 [5]string
	array1[0] = "Posição 1"
	fmt.Println(array1)

	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}
	fmt.Println(array2)

	// é a forma menos inflexivel de usar o array, aqui ele define o tamanho do array a partir da quantidade de dados
	// que você passar, não faz com que ele seja flexivel.
	// Por sua inflexibilidade, o Array não costuma ser muito utilizado no dia a dia.

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	// SLICES

	// Slice é uma estrutura baseada no Array , mas possui agumas flexibilidades a mais.
	// Criar um Slice é bem parecido com Array.

	// Não é colocado o tamanho, para não fixa-lo. Ainda é necessário os itens do slice sejam do mesmo tipo
	slice := []int{10, 12, 13, 14, 15, 16, 17, 18, 19}
	fmt.Println(slice)

	// é como se um Slice referenciasse um Array.A tradução da palavra Slice significa fatia, então é como se
	// o Slice fosse uma fatia do Array

	// fmt.Println(reflect.TypeOf(slice))
	// fmt.Println(reflect.TypeOf(array3)) Type off diz qual o tipo da sua variavel

	slice = append(slice, 20) // O append adiciona o novo item no Slice, e retorna o mesmo Slice.
	fmt.Println(slice)

	// buscando diretamente em um Arrau a fatia/slice
	// Lembrando que de acordo com o exemplo abaixo, o indice 1 é inclusive, vai pegar a partir do indice 1
	// O indice 3 é exclusivo, ou seja, a partir do indice 3 será exclusivo (o indice 3 nao é incluido na fatia)

	slice2 := array2[1:3] // buscando diretamente em um Arrau a fatia/slice
	fmt.Println(slice2)

	array2[1] = "Posição alerada" // alterando conteúdo do indice 1
	fmt.Println(slice2)
}
