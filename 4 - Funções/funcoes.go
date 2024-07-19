package main

import "fmt"

// função é como se fosse uma receita de bolo, uma série de passos que seu programa irá seguir/executar.
// as funções podem ter parametros e retornos. Parametros são dados que você coloca na fun~ção para ela funcionar
// retorno é o que a função te devolver após ser executada

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) { //em go uma função pode dar mais de um retorno
	soma := n1 + n2
	subtracao := n1 + n2
	return soma, subtracao
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(txt string) string { //tipo de função que leva em conta os parametros e o retorno
		fmt.Println(txt)
		return txt
	}

	resultado := f("Texto da função 1")
	fmt.Println(resultado)

	resultadoSoma, resultadoSubtracao := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma, resultadoSubtracao)

	resultadoSoma1, _ := calculosMatematicos(10, 15) // o _ é usado quando você não deseja usar a segunda informação
	fmt.Println(resultadoSoma1)
}
