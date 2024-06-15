# REST API Project

Este é um projeto de API REST construído em Go usando o `net/http` para o servidor e cliente. A API integra-se com a seguinte API externa: `https://api.restful-api.dev/objects`.

## Estrutura do Projeto

- `cmd/server/main.go`: Ponto de entrada da aplicação.
- `internal/handlers/handler.go`: Manipuladores de requisições HTTP.
- `internal/services/service.go`: Lógica de negócios e integração com a API externa.
- `internal/models/object.go`: Definição dos modelos de dados.
- `internal/handlers/handler_test.go`: Testes unitários para os manipuladores.
- `internal/services/service_test.go`: Testes unitários para os serviços.
- `go.mod`: Arquivo de gerenciamento de dependências.

## Pré-requisitos

- Go 1.22.4 ou superior

## Instalação de Dependências

Para instalar as dependências do projeto, execute o comando abaixo no diretório raiz do projeto:

```bash
go mod tidy
```

## Executando o Servidor

Para iniciar o servidor, execute o seguinte comando:

```bash
go run cmd/server/main.go
```

O servidor estará disponível em `http://localhost:8080`.

## Executando os Testes

Para executar os testes unitários, utilize o comando abaixo:

```bash
go test ./...
```

Este comando executará todos os testes presentes nos pacotes do projeto.

## Estrutura de Diretórios

```plaintext
.
├── cmd
│   └── server
│       └── main.go
├── internal
│   ├── handlers
│   │   ├── handler.go
│   │   └── handler_test.go
│   ├── models
│   │   └── object.go
│   └── services
│       ├── service.go
│       └── service_test.go
├── go.mod
└── README.md
```

## Contato

Para mais informações, entre em contato com o desenvolvedor do projeto.
