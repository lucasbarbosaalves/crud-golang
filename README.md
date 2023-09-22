# CRUD com Golang (Aprendendo)

Este documento fornece informações sobre uma API REST simples em Golang, que utiliza os pacotes nativos da linguagem, como net/http, e a biblioteca externa mux para roteamento. A API oferece operações básicas para gerenciar recursos de usuário.

## Pré-requisitos
- [Golang](https://golang.org/doc/install)
- [Git](https://git-scm.com/downloads)

## Endpoints


- POST http://localhost:5000/usuarios
```
HTTP/1.1 201 CREATED
Content-Length: 129
Content-Type: application/json

{
    "nome": "Lucas,
    "email": "lucas@email.com"
}
```

- GET /usuarios/{id}
```
HTTP/1.1 200 OK
Content-Length: 129
Content-Type: application/json
 
```

- GET /usuarios
```
HTTP/1.1 200 OK
Content-Type: application/json

```

- PUT /usuarios/{id}
```
HTTP/1.1 204 no Content
Content-Length: 142
Content-Type: application/json

{
    "nome": "Lucas,
    "email": "lucas@email.com"
}
```

- DELETE /usuarios/{id}
```
HTTP/1.1 204 no Content
Content-Length: 142
Content-Type: application/json

```

## Executando o Projeto

Para iniciar a API, siga as etapas abaixo:

1. Execute o seguinte comando no terminal dentro do diretório do projeto:

   ```bash
   go run main.go
A API estará disponível em http://localhost:5000.

## Autor

- Lucas Alves

