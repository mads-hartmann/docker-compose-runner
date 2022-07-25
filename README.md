
Running integration tests

```sh
go test -tags integration ./...
```

Running tests with code coverage

```sh
go test -v -cover -coverprofile=c.out ./...
go tool cover -html c.out
```
