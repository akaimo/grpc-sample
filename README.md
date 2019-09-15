# grpc-sample

## setup
```
$ go get google.golang.org/grpc
$ brew install protobuf
$ go get github.com/golang/protobuf/protoc-gen-go
```

## cli
```
$ brew install grpcurl
```

## Run
```
$ make server
$ grpcurl -plaintext -d '{"latitude": 10, "longitude": 20}' localhost:10000 RouteGuide.GetFeature
```
