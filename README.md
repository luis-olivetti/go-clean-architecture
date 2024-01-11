# go-clean-architecture

Agora é a hora de botar a mão na massa. Pra este desafio, você precisará criar o usecase de listagem das orders.

Esta listagem precisa ser feita com:

- Endpoint REST (GET /order)

- Service ListOrders com GRPC

- Query ListOrders GraphQL

Não esqueça de criar as migrações necessárias e o arquivo api.http (para executar com a extensão Rest Client VSCode) com a request para criar e listar as orders.

---

localhost:8080 - GraphQL

```
mutation orders {
  createOrder(input: {id: "ab", Price: 50.90, Tax: 0.10}) {
    id
    Price
    Tax
    FinalPrice
  }
}
```

Como executar:

```
go run main.go wire_gen.go
```


Anotações para desenvolvimento

Para gerar o proto buffers

protoc --proto_path=internal/infra/grpc/protofiles --go_out=internal/infra/grpc/pb --go_opt=paths=source_relative \
  --go-grpc_out=internal/infra/grpc/pb --go-grpc_opt=paths=source_relative \
  internal/infra/grpc/protofiles/order.proto