# About

Simple example usage tg.

# Update transport

```bash
go generate ./...
```

# Run servcie

```bash
go run ./cmd/example/main.go
```

# API Example

Request: 
```bash
curl -X POST "http://localhost:9000/" -H "Content-Type: application/json" -d '[{"id":"123", "jsonrpc":"2.0", "method":"adder.add", "params":{"a":1,"b":2}}]'
```

Response: 
```bash
[{"id":"123","jsonrpc":"2.0","result":{"c":3}}]
```

Log on server:
```
{"time":"2022-01-24T10:12:16+03:00","message":"hello world"}
{"level":"info","bind":":9000","time":"2022-01-24T10:12:16+03:00","message":"listen on"}
{"level":"info","method":"add","request":"{A:1 B:2}","response":"{C:3}","service":"Adder","took":"21.375µs","time":"2022-01-24T10:13:31+03:00","message":"call add"}
```
