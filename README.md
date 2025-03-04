# CEP Resolver Go

Neste projeto constitui o assunto sobre Multithreading e APIs para buscar o resultado mais rápido entre duas APIs distintas.

## Estrutura do Projeto

```
cep-resolver-go/
├── cmd/
│   └── main.go
├── internal/
│   ├── handlers/
│   │   ├── address_handler.go
│   │   └── address_handler_test.go
│   ├── services/
│   │   ├── brasil_api.go
│   │   ├── brasil_api_test.go
│   │   ├── viacep_api.go
│   │   └── viacep_api_test.go
│   └── models/
│       └── address.go
├── pkg/
│   └── utils/
│       └── http_client.go
├── go.mod
└── go.sum
```

## Descrição

As duas requisições serão feitas simultaneamente para as seguintes APIs:

- [BrasilAPI](https://brasilapi.com.br/api/cep/v1/01153000)
- [ViaCEP](http://viacep.com.br/ws/01153000/json/)

## Requisitos

- Acatar a API que entregar a resposta mais rápida e descartar a resposta mais lenta.
- O resultado da request deverá ser exibido no command line com os dados do endereço, bem como qual API a enviou.
- Limitar o tempo de resposta em 1 segundo. Caso contrário, o erro de timeout deve ser exibido.

## Como Executar

1. Clone o repositório:
    ```sh
    git clone https://github.com/AugustoSandim/cep-resolver-go.git
    cd cep-resolver-go
    ```

2. Baixar as dependências:
    ```sh
    go mod tidy
    ```

3. Execute o comando abaixo para rodar o programa.
    ```sh
    go run cmd/main.go
    ```

4. Insira o CEP que deseja buscar o endereço.
    ```sh
    curl -X GET "http://localhost:8080/address?cep=01153000"
    ```

