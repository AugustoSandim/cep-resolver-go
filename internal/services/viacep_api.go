package services

import (
	"cep-resolver-go/internal/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ViaCEPResponse struct {
	CEP          string `json:"cep"`
	Street       string `json:"logradouro"`
	Neighborhood string `json:"bairro"`
	City         string `json:"localidade"`
	State        string `json:"uf"`
}

func FetchViaCEP(cep string, ch chan<- models.Address) {
	url := fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep)
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var result ViaCEPResponse
	if err := json.Unmarshal(body, &result); err == nil {
		ch <- models.Address{
			CEP:          result.CEP,
			Street:       result.Street,
			Neighborhood: result.Neighborhood,
			City:         result.City,
			State:        result.State,
			Service:      "ViaCEP",
		}
	}
}
