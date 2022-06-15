### 安装生成工具
go get -u github.com/golang/protobuf/protoc-gen-go
go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
go get github.com/asim/go-micro/cmd/protoc-gen-micro/v3

```
    protoc --proto_path=./proto --go-grpc_out=./proto --micro_out=./proto --go_out=./proto ./proto/*.proto
    
    protoc --proto_path=. --go-grpc_out=. --micro_out=. --go_out=. ./proto/vfile.proto
    protoc --proto_path=. --go-grpc_out=. --micro_out=. --go_out=. ./proto/vcms.proto
```