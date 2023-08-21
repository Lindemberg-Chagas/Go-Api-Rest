package controle

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Lindemberg-Chagas/Go-Api-Rest/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "HOME PAGE")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalidades) //função para encodar o formato json e aprensentar no endpoint
}
func RetornaId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //criando variavel para pegar o endpoint usando o mux
	id := vars["id"]    // pegando apenas 1 endpoint para comparar

	for _, persona := range models.Personalidades { // range dentro da variavel persona usando a struct models para fazer a validação
		if strconv.Itoa(persona.Id) == id {
			json.NewEncoder(w).Encode(persona) // fazendo o teste de validação com if,
			//se o id da pagina que foi pegada com mux for igual ao Id da Persona com
			// o que foi usado no Range da Struct ele mostra apenas o valor de um ou outro
		}
	}
}
