# Books CRUD 📚

## Api de cadastro de livros.

**Stack estudada** 🛠️

 > * GoLang 
 > * Postgres 
 > * Docker
 > * GRPC 

 --- 

 ## Run (Server Http or Grpc) 🏃
> $ docker-compose up -d

> $ go run main.go 
or
> $ make run

 ## Run Client Grpc
 > $ make run-client

 Obs: Para rodar o client grpc precisamos do servidor grpc rodando junto com o banco de dados. Caso queira trocar a função de criar para visualizar, vá em client/cmd/main.go e altere o valor da variável "routehttp". Deixei dessa forma estática e não tão usual para demonstrar apenas a utilização de um client Grpc. 