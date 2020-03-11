## How to run app
1. ```go build```
2. ```go run server.go```

## How to run test
1. ```go test -v ./...```

## How to run with docker-compose
1. ```docker-compose build```
2. ```docker-compose up```
3. ```curl localhost:8080/load -d '["foobar", "aabb", "baba", "boofar", "test"]'```
4. ```curl 'localhost:8080/get?word=foobar'```

