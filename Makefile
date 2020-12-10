compile_grpc:
	protoc -I ./proto \
		--go_out . --go_opt paths=source_relative \
		--go-grpc_out . --go-grpc_opt paths=source_relative \
		./proto/user_service/user_service.proto

compile_reverse_proxy:
	protoc -I ./proto --grpc-gateway_out . \
     --grpc-gateway_opt logtostderr=true \
     --grpc-gateway_opt paths=source_relative \
     --grpc-gateway_opt generate_unbound_methods=true \
     ./proto/user_service/user_service.proto