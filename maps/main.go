package main

import "fmt"

func main() {
	user := map[string]string {
		"name": "Claudio",
		"email": "claudinhoandre@gmail.com",
		"role": "Full-stack developer",
	}
	fmt.Println(user)
	fmt.Println(user["name"])

	// Para ver se existe um indice no map preciso passar duas variáveis
	// a primeira para receber o valor, qual pode ser vazio caso não exista o índice
	// a segunda um booleano que confirma se o indice buscado existe de fato

	value, ok := user["salary"]

	if (!ok) {
		fmt.Println("Eu disse que não existia, valor: ", value, ".")
	}

	for key, value := range user {
		fmt.Println("Key:", key, "Value:", value)
	}
}
