package main

//Run Server with "go run main.go" in the client dir
//Run Client with "go run main.go" in the client dir

//To use client either open up a browser and type http://localhost:8080/GetCourses/1
//or type curl http://localhost:8080/GetCourses/1 in terminal
//GetCourse being the endpoint
//1 being the course we wanna get (also supports 2 and 3)

import (
	"context"
	"grpc/proto"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func WhichCourse(id int64) string {
	if id == 1 {
		return "ID: \"01\", Name: \"BDSA\", Teacher: \"Alice\", Rating: 7, Workload: \"5 exercises\""
	} else if id == 2 {
		return "ID: \"02\", Name: \"DISYS\", Teacher: \"Bob\", Rating: 8, Workload: \"4 projects\""
	} else if id == 3 {
		return "ID: \"03\", Name: \"IDBS\", Teacher: \"Carl\", Rating: 6, Workload: \"3 assignments\""

	}
	return "no course with this id"
}

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

func (s *server) GetCourses(ctx context.Context, req *proto.Request) (*proto.Response, error) {
	a := req.GetA()

	result := WhichCourse(a)
	return &proto.Response{Result: string(result)}, nil
}
