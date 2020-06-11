package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	"log"
	"osapp/models"
)

var (
	serverAddr         = flag.String("server_addr", "localhost:10000", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "x.test.youtube.com", "The server name used to verify the hostname returned by the TLS handshake")
)

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := models.NewUserClient(conn)

	u := &models.UserInfo{}
	u.AccessKey = "ak_123456"
	u.SecretKey = "sk_123456"

	resp, err := client.AddUser(context.Background(), u)
	if err!=nil{
		log.Fatalf("add user failed\n")
	}else{
		log.Printf("%v\n", resp)
	}
}
