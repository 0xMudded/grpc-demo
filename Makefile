generate_grpc_code:
	protoc \
	--go_out=number-generator \
	--go_opt=paths=source_relative \
	--go-grpc_out=number-generator \
	--go-grpc_opt=paths=source_relative \
	NumberGenerator.proto