BINARY = ${USER_SERVICE_BINARY}
.PHONY: protos

run:
	cd cmd/user-service-server && make run

protos:
	protoc -I protos/ protos/user-service.proto --go_out=plugins=grpc:protos/