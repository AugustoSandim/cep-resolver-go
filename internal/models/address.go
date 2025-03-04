package models

type Address struct {
	CEP          string `json:"zip_code"`
	Street       string `json:"street"`
	Neighborhood string `json:"neighborhood"`
	City         string `json:"city"`
	State        string `json:"state"`
	Service      string `json:"service"`
}
