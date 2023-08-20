package models

type Personalidade struct {
	Nome     string `json:"nome"`     //serializar
	Historia string `json:"histotia"` //serializar
}

var Personalidades []Personalidade //instanciando objeto em Slice
