.PHONY: gen_coordinator

gen_coordinator:
	protoc \
		--go_out=. --go_opt=module=github.com/Gurveer1510/dfps \
		--go-grpc_out=. --go-grpc_opt=module=github.com/Gurveer1510/dfps \
		proto/dfps.proto
