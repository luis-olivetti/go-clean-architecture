# go-clean-architecture

Agora é a hora de botar a mão na massa. Pra este desafio, você precisará criar o usecase de listagem das orders.

Esta listagem precisa ser feita com:

- Endpoint REST (GET /order)
- Service ListOrders com GRPC
- Query ListOrders GraphQL

Não esqueça de criar as migrações necessárias e o arquivo api.http (para executar com a extensão Rest Client VSCode) com a request para criar e listar as orders.

---

## Como executar a aplicação?

Pré-requisito:
Possuir um banco de dados chamado **orders** com a entidade **orders**

```sql
CREATE TABLE orders (id varchar(255) NOT NULL, price float NOT NULL, tax float NOT NULL, final_price float NOT NULL, PRIMARY KEY (id))
```

Execute os seguintes comandos para iniciar a aplicação REST, GraphQL e GRPC
```
cd cmd/ordersystem
go run main.go wire_gen.go
```

## REST

No diretório **api** foi disponibilizado os arquivos **http** para serem executados com a extensão REST Client (VSCode).

## GraphQL Playground

localhost:8080

```
mutation orders {
  createOrder(input: {id: "ab", Price: 50.90, Tax: 0.10}) {
    id
    Price
    Tax
    FinalPrice
  }
}

query list_orders {
  ListOrders{
    id
    Price
    Tax
    FinalPrice
  }
}
```

## GRPC

Para executar os comandos via GRPC você pode usar o client que lhe agrade, como por exemplo o Evans através do comando `evans -r rpl` ou o Postman e Insomnia.

Exemplo com Insomnia:

![Exemplo de uso de GRPC com Insomnia](insomnia_grpc.png)

## Anotações para desenvolvedor

### Para gerar/atualizar os protobuffers

No diretório raíz do projeto, execute:

```
protoc --proto_path=internal/infra/grpc/protofiles --go_out=internal/infra/grpc/pb --go_opt=paths=source_relative \
  --go-grpc_out=internal/infra/grpc/pb --go-grpc_opt=paths=source_relative \
  internal/infra/grpc/protofiles/order.proto
```

### Para gerar/atualizar os arquivos do GraphQL, execute:

```
go run github.com/99designs/gqlgen generate
```
