package main

import (
	"errors"
	"fmt"
)

// Existem tipos diferentes de int8, int10, int32, int64 (a diferença é a quantidade de bits que cada um comporta)

func main() {

	// NUMEROS INTEIROS
	var numero int8 = 100
	println(numero)

	var numero1 int = 1 // o int sozinho infere de forma implicita a quantidade bits do seu computador)
	println(numero1)

	numeroNegativo := -100000 // também suportam números negativos
	fmt.Println(numeroNegativo)

	var numero2 uint //unsygned int, int sem sinal, não caeita número negativo
	fmt.Println(numero2)

	// ALIAS = APELIDO

	var numero3 rune //int32 = rune, normalmente usado com numeros que representam caracteres, principalmente da tabela ASCII
	fmt.Print(numero3)

	var numero4 byte //int8 = byte, byte possui 8 bits cada
	fmt.Println(numero4)

	//NUMERO REAL
	var numeroReal float32 = 123.45
	fmt.Println(numeroReal)

	var numeroReal1 float64 = 12345.67
	fmt.Println(numeroReal1)

	numeroReal2 := 12345.67 // inferencia de tipo, também se baeara nos bits do seu pc, não é permitido declarar apenas float puro
	fmt.Println(numeroReal2)

	//STRINGS

	var str string = "texto"
	fmt.Println(str)

	str2 := "texto2" //inferencia de tipo
	fmt.Println(str2)

	char := "b"
	fmt.Println(char) // não existe char, ea acaba devolvendo um int/posição na tabela ascii

	// valor zero = valor que será atribuido a uma variavel quando vc noa a inicializa com valor, no caso da

	// variavel abaixo a string será retornada em branco/vazia
	var texto string
	fmt.Println(texto)

	// no caso abaixo, por se tratar de uma variavel int, o go irá atribui o numero zero em vez de estar vazia
	var texto2 int16
	fmt.Println(texto2)

	//BOOLEANO
	var booleano bool = true // or false, se deixar vazia, o valor zero do booleano é o false
	fmt.Println(booleano)

	//ERROR
	var erro error
	fmt.Println(erro) // o valor zero do error é nil

	var erro1 error = errors.New("Erro interno") // erro é o nome da variavel, error é tipo da variavel, e errors é o nome do pacote
	fmt.Println(erro1)

}
