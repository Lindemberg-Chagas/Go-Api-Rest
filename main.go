package main

import (
	"fmt"

	"github.com/Lindemberg-Chagas/Go-Api-Rest/models"
	"github.com/Lindemberg-Chagas/Go-Api-Rest/rotas"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Nome: "Nome 1", Historia: "Historia 1"},
		{Nome: "Nome 2", Historia: "Historia 2"},
	}

	fmt.Println("Iniciando o Servido Rest em GO")
	rotas.HandleRequest()
}
