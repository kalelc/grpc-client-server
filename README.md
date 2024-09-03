# gRPC Server/Client

1. Define Proto Service

```
user.proto
```

2. Generate code using proto File

```
protoc --go_out=./ --go-grpc_out=./ user.proto
```

3. Execute gRPC Server

```
go run server/server.go
```

4. Test using client to consume gRPC Server

```
go run client/client.go
```
