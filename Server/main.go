package main
//Client can be run with "go run main.go" in the client dir

//To use client either open up a browser and type http://localhost:8080/add/5/10
//or type curl http://localhost:8080/add/5/10 in terminal
//add being the thing we wanna do (also supports multiply)
//5 and 10 being the numbers you wanna add. (You ar free to change these)


import (
	"context"
	"grpc/proto"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {}

func main() {
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	proto.RegisterAddServiceServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(err)
	}
}

func (s *server) Add(ctx context.Context, req *proto.Request) (*proto.Response, error) {
	a, b := req.GetA(), req.GetB()

	result := a + b

	return &proto.Response{Result: result}, nil
}

func (s *server) Multiply(ctx context.Context, req *proto.Request) (*proto.Response, error) {
	a, b := req.GetA(), req.GetB()

	result := a * b

	return &proto.Response{Result: result}, nil
}