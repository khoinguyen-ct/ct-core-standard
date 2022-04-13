package main

import (
	"context"
	"fmt"
	svr "github.com/ct-core-standard/pkg/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, _ := grpc.Dial("127.0.0.1:8989", grpc.WithTransportCredentials(insecure.NewCredentials()))
	client := svr.NewAdListingServiceClient(conn)
	req := &svr.GetAdInfoRequest{ListId: 102993}
	res, err := client.GetAdInfo(context.Background(), req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}
