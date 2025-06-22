# grpc-practise-1
Learning grpc (intercom with nodejs and golang)
<!-- https://grpc.io/docs/languages/go/basics/ -->

protoc --go_out=. --go-grpc_out=. input.proto

## protoc :- protoc is compiler
## --go_out=. :- Generates Protobuf Go code in current directory  ex : orders.pb.go
## --go-grpc_out=. :- Generates grpc Go code in current directory ex : orders_grpc.pb.go
## input.proto :- proto file name which is given as input


You can specify the Generated files output path : Where you want to generate
protoc --go_out=../orderservice --go-grpc_out=../orderservice orders.proto

<!-- INSTALL THESE dependenciees aprat from go install grpc and protobuff -->
go get google.golang.org/protobuf
go get google.golang.org/grpc

<!-- COMMAND TO GENERATE grpc and protobuf files  -->
protoc \
    --proto_path=protobuf "protobuf/orders.proto" \
    --go_out=services/common/genproto/orders --go_opt=paths=source_relative \
    --go-grpc_out=services/common/genproto/orders \
    --go-grpc_opt=paths=source_relative


# protoc: The Protocol Buffers compiler.
# --proto_path=protobuf: Specifies the directory (protobuf) where protoc looks for .proto files
# protobuf/orders.proto: The input proto file
# --go_out=services/common/genproto/orders : Specifies the output directory for generated go code
# --go_opt=paths=source_relative : Ensures generated go code uses relative path for imports
