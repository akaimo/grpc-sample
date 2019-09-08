gen:
	protoc -I api/ api/route_guide.proto --go_out=plugins=grpc:api

.PHONY: server
server:
	go run server/main.go
