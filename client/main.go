package main

import (
	"log"
	"context"
	"time"

	"google.golang.org/grpc"

	pb "akaimo.com/grpc-sample/api"
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial("127.0.0.1:10000", opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewRouteGuideClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	feature, err := client.GetFeature(ctx, &pb.Point{Latitude:400000000, Longitude: -750000000})
	if err != nil {
		log.Fatalf("%v.GetFeature(_) = _, %v", client, err)
	}
	log.Println(feature)
}
