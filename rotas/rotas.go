package rotas

import (
	"log"
	"net/http"

	"github.com/Lindemberg-Chagas/Go-Api-Rest/controle"
)

func HandleRequest() {
	http.HandleFunc("/", controle.Home)
	http.HandleFunc("/api/personalidades", controle.TodasPersonalidades)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
