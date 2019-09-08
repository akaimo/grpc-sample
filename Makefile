gen:
	protoc -I api/ api/route_guide.proto --go_out=plugins=grpc:api
