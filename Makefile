gen-api: gen-xlinepb gen-curppb

gen-xlinepb:
	rm -rf gen/xlinepb && \
	mkdir -p gen/xlinepb && \
	cd api/xline-proto/src && \
	protoc \
		--experimental_allow_proto3_optional \
		--go_out=../../../gen/xlinepb \
		--go_opt=paths=source_relative \
		--go-grpc_out=../../../gen/xlinepb \
		--go-grpc_opt=paths=source_relative \
	*.proto

gen-curppb:
	rm -rf gen/curppb && \
	mkdir -p gen/curppb && \
	cd api/curp-proto/src && \
	protoc \
		--experimental_allow_proto3_optional \
		--go_out=../../../gen/curppb \
		--go_opt=paths=source_relative \
		--go-grpc_out=../../../gen/curppb \
		--go-grpc_opt=paths=source_relative \
	*.proto
