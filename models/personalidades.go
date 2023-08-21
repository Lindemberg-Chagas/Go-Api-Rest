package models

type Personalidade struct {
	Id       int    `json:"id"`
	Nome     string `json:"nome"`     //serializar
	Historia string `json:"histotia"` //serializar
}

var Personalidades []Personalidade //instanciando objeto em Slice
