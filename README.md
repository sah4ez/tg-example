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
curl -X POST "http://localhost:9000/" -H "Content-Type: application/json" -d '[{"id":"123", "jsonrpc":"2.0", "method":"user.getUserNameByID", "params":{"id":1}}]'
```

Response for error: 
```bash
[{"id":"123","jsonrpc":"2.0","error":{"code":-32603,"message":"user not found: record not found","data":{}}}]
```

Prepare user data:
```bash
psql 'host=localhost port=5432 user=postgres dbname=postgres password=secretpass sslmode=disable client_encoding=UTF8' -c "insert into users(name) values('user2');"

# INSERT 0 1
```

Response with username: 
```bash
[{"id":"123","jsonrpc":"2.0","result":{"name":"user"}}]
```

Log on server:
```
./bin/example
2022-03-25T00:52:57+03:00 ??? hello world
2022-03-25T00:52:58+03:00 INF listen on bind=:9000

2022/03/25 00:53:02 /home/sah4ez/go/src/github.com/sah4ez/tg-example/pkg/storage/user.go:11 record not found
[2.101ms] [rows:0] SELECT * FROM "users" WHERE id = 1 ORDER BY "users"."id" LIMIT 1
2022-03-25T00:53:02+03:00 ERR call getUserNameByID error="user not found: record not found" method=getUserNameByID request={Id:1} response={Name:} service=User took=2.545998ms
2022-03-25T00:54:36+03:00 INF call getUserNameByID method=getUserNameByID request={Id:1} response={Name:user} service=User took=2.123242ms
```

# Serve binary data

Prepare embeded files:

```
cd pkg/files/
rice embed-go 
```
