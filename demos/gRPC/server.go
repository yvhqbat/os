package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"osapp/models"
	"sync"
)

var (
	port = flag.Int("port", 10000, "The server port")
)



type adminServer struct {
	models.UnimplementedUserServer
	mu         sync.Mutex // protects routeNotes
}

func (server *adminServer) AddUser(ctx context.Context, req *models.UserInfo) (*models.Response, error) {
	log.Printf("input: %v\n", req)

	resp := &models.Response{}
	resp.Code = 200
	resp.Msg = "add user success"

	return resp, nil
}

func newServer() *adminServer {
	s := &adminServer{}
	return s
}

func main(){
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	models.RegisterUserServer(grpcServer, newServer())
	err = grpcServer.Serve(lis)
	if err!=nil{
		log.Fatalf("%v\n", err)
	}
}
