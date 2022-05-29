user-service-proto:
	protoc -I proto-contracts/ proto-contracts/user-service/user-service.proto --go_out=plugins=grpc:proto-contracts/user-service/

status-proto:
	protoc -I proto-contracts/ proto-contracts/status/status.proto --go_out=plugins=grpc:proto-contracts/status/