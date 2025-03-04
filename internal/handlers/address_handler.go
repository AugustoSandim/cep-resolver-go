package handlers

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"cep-resolver-go/internal/models"
	"cep-resolver-go/internal/services"
)

func GetAddress(w http.ResponseWriter, r *http.Request) {
	cep := r.URL.Query().Get("cep")
	if cep == "" {
		http.Error(w, "Missing cep parameter", http.StatusBadRequest)
		return
	}

	ch := make(chan models.Address, 2)
	var wg sync.WaitGroup
	wg.Add(2)

	go services.FetchBrasilAPI(cep, ch)
	go services.FetchViaCEP(cep, ch)

	timeout := time.After(1 * time.Second)

	select {
	case result := <-ch:
		json.NewEncoder(w).Encode(result)
	case <-timeout:
		http.Error(w, "Request timed out", http.StatusGatewayTimeout)
	}
}
