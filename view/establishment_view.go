package view

import "modulo/models"

type ServiceEstablishment struct {
	Nome        string
	RazaoSocial string
	Endereco    string
	Estado      string
	Cidade      string
	Cep         string
	Numero      int
	Stores      []models.Store
}

func EstablishmentView(v models.Establishment) ServiceEstablishment{
	return ServiceEstablishment{Nome: v.Nome,
	RazaoSocial: v.RazaoSocial,
	Endereco: v.Endereco,
	Cep: v.Cep,
	Estado: v.Estado,
	Cidade: v.Cidade,
	Numero: v.Numero,
	Stores: v.Stores}
}