1. create .proto for your functionality1. Download Protocol Buffers - https://developers.google.com/protocol-buffers/docs/downloads
2. Create `.proto` file to describe services and messages
3. Install gen go tool for generate go source - `go install google.golang.org/protobuf/cmd/protoc-gen-go`
4. Use protoc to generate go source - `protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/fibonacci.proto`