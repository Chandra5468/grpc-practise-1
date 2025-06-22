gen:
  @protoc \
    --proto_path=protobuf "protobuf/orders.proto" \
	--go_out=services/common/genproto/orders --go_opt=paths=source_relative \
	--go-grpc_out=services/common/genproto/orders
	--go-grpc_opt=paths=source_relative

# What is happening in gen:
# telling protoc : proto compiler
# --proto_path=protobuf "" Here we are mentioning proto file path
# --go_out this is folder path for outputting the generated proto code.
# 