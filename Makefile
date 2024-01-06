gen-api: gen-curppb gen-xlinepb gen-mockpb

gen-curppb:
	rm -rf api/gen/curppb && \
	mkdir -p api/gen/curppb && \
	cd api/proto/curp-proto/src && \
	protoc \
		--experimental_allow_proto3_optional \
		--go_out=../../../gen/curppb \
		--go_opt=paths=source_relative \
		--go-grpc_out=../../../gen/curppb \
		--go-grpc_opt=paths=source_relative \
	*.proto

gen-xlinepb:
	rm -rf api/gen/xlinepb && \
	mkdir -p api/gen/xlinepb && \
	cd api/proto/xline-proto/src && \
	protoc \
		--experimental_allow_proto3_optional \
		--go_out=../../../gen/xlinepb \
		--go_opt=paths=source_relative \
		--go-grpc_out=../../../gen/xlinepb \
		--go-grpc_opt=paths=source_relative \
	*.proto
	
gen-mockpb:
	rm -rf api/gen/mockpb && \
	mkdir -p api/gen/mockpb && \
	cd api/proto/mock-proto && \
	protoc \
		--experimental_allow_proto3_optional \
		--go_out=../../gen/mockpb \
		--go_opt=paths=source_relative \
		--go-grpc_out=../../gen/mockpb \
		--go-grpc_opt=paths=source_relative \
	*.proto
