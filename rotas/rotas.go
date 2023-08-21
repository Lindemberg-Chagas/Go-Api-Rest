package rotas

import (
	"log"
	"net/http"

	"github.com/Lindemberg-Chagas/Go-Api-Rest/controle"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controle.Home)
	r.HandleFunc("/api/personalidades", controle.TodasPersonalidades).Methods("GET")
	r.HandleFunc("/api/personalidades/{id}", controle.RetornaId).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))
}
