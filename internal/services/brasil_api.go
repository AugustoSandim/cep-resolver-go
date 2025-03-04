package services

import (
	"cep-resolver-go/internal/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func FetchBrasilAPI(cep string, ch chan<- models.Address) {
	url := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var result models.Address
	if err := json.Unmarshal(body, &result); err == nil {
		result.Service = "BrasilAPI"
		ch <- result
	}
}
