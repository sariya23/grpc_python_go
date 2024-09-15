## Generate protoc

### Python

```shell
nikita:py python3 -m grpc_tools.protoc --proto_path=../api/proto --python_out=. --grpc_python_out=.  ../api/proto/calc.proto
```

### Go

```shell
nikita:solver protoc --go_out=pkg/api --go_opt=paths=source_relative --go-grpc_out=pkg/api --go-grpc_opt=paths=source_relative api/proto/calc.proto
```