package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "akaimo.com/grpc-sample/api"
)

type routeGuideServer struct{}

func (s *routeGuideServer) GetFeature(ctx context.Context, point *pb.Point) (*pb.Feature, error) {
	return &pb.Feature{Location: point}, nil
}

func newServer() *routeGuideServer {
	s := &routeGuideServer{}
	return s
}

func main() {
	lis, err := net.Listen("tcp", "localhost:10000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	pb.RegisterRouteGuideServer(grpcServer, newServer())
	_ = grpcServer.Serve(lis)
}
