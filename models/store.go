package models

type Store struct {
	ID              uint   `json:"id" gorm:"primaryKey"`
	Nome            string `json:"nome"`
	RazaoSocial     string `json:"razao_social"`
	Endereco        string `json:"endereco"`
	Estado          string `json:"estado"`
	Cidade          string `json:"cidade"`
	Cep             string `json:"cep"`
	Numero          int    `json:"numero"`
	EstablishmentID uint   `json:"establishment_id"`
}
