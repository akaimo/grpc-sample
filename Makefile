gen:
	protoc -I api/ api/route_guide.proto --go_out=plugins=grpc:api

.PHONY: server
server:
	GODEBUG=http2debug=2 go run server/main.go

grpcurl:
	grpcurl -plaintext -d '{"latitude": 10, "longitude": 20}' localhost:10000 RouteGuide.GetFeature

list:
	grpcurl -plaintext localhost:10000 list RouteGuide
