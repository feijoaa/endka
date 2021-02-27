package main

import (
	"endka/task2/protoc"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen:%v", err)
	}
	s := grpc.NewServer()
	protoc.RegisterNumServiceServer(s, &Server{})
	log.Println("Server is running on port:50051")
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve:%v", err)
	}
}
func (s *Server) Number(req *protoc.NumberRequest, stream protoc.NumService_NumberServer) error {
	fmt.Printf("GreetManyTimes function was invoked with %v \n", req)
	number := req.GetNum().GetNumber()

	numbers :=  []int32 {}

	for number%2 == 0 {
		numbers = append(numbers, 2)
		number/=2
	}