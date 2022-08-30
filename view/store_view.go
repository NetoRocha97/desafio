package view

import "modulo/models"

type ServiceStore struct {
	Nome            string
	RazaoSocial     string
	Endereco        string
	Estado          string
	Cidade          string
	Cep             string
	Numero          int
	EstablishmentID uint
}

func StoreView(v models.Store) ServiceStore {
	return ServiceStore{Nome: v.Nome,
		RazaoSocial:     v.RazaoSocial,
		Endereco:        v.Endereco,
		Cep:             v.Cep,
		Estado:          v.Estado,
		Cidade:          v.Cidade,
		Numero:          v.Numero,
		EstablishmentID: v.EstablishmentID}
}
