package main

import "fmt"

// Pense em ponteiros como uma variavel que irá salvar dentro dela , não necessáriamente um valor,
// mas o endereço de memória que hospeda esse valor.

func main() {
	fmt.Println("Ponteiros")

	// QUando você adiciona uma variavel como valor de outra variavel, você está fazendo somente uma cópia

	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	// ao incrementar +1 na variavel1, o valor é somente alterado na variavel1,
	// o valor acrescentado não é transferido para a variavel2

	variavel1++
	fmt.Println(variavel1, variavel2)

	// PONTEIRO É UMA REFERÊNCIA DE MEMÓRIA

	var variavel3 int
	var ponteiro *int

	variavel3 = 100
	ponteiro = &variavel3

	fmt.Println(variavel3, ponteiro)
	fmt.Println(variavel3, *ponteiro) // desreferenciação, está tirando sua referencia, e pedindo para trazer seu conteudo.

	variavel3 = 150 // apesar de o valor da variavel mudar dentro do endereço de memoria, o endereço em si nao sofre alteração.
	fmt.Println(variavel3, ponteiro)
	fmt.Println(variavel3, *ponteiro) // agpra com a desreferenciação, ele irá apresentar o valor, que no caso é o novo valor de 150.
}
