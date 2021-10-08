package main

import "fmt"

func main() {
	var name string = "Jocel" // declarar tipagem depois do nome da variavel
	//var idade int -- nao deixa declarar uma variavel e nao usar
	var version float32 = 1.01
	// go aceita o var com inferencia de tipo
	// outra opção é :
	idade := 27

	// não usar ; ao final da instrução
	fmt.Println("Hello World,", name)  // concatenar com a virgula
	fmt.Println("Sua idade é,", idade) // concatenar com a virgula
	fmt.Println("Versão do programa", version)
}
