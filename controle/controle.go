package controle

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Lindemberg-Chagas/Go-Api-Rest/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "HOME PAGE")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalidades) //função para encodar o formato json e aprensentar no endpoint
}
