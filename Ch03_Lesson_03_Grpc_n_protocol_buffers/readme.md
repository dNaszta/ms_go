1. create .proto for your functionality1. Download Protocol Buffers - https://developers.google.com/protocol-buffers/docs/downloads
2. Create `.proto` file to describe services and messages
3. Install gen go tool for generate go source - 
```go get google.golang.org/protobuf/cmd/protoc-gen-go google.golang.org/grpc/cmd/protoc-gen-go-grpc```
4. Use protoc to generate go source - 
```protoc --go_out=. --go_opt=paths=source_relative  --go-grpc_out=. --go-grpc_opt=paths=source_relative  fibonacci/fibonacci.proto```