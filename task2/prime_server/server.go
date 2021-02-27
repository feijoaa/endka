package main

import (
	"endka/task2/protoc"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"math"
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

	numbers := []int32{}

	for i := 0; i < len(numbers); i++ {
		result := &protoc.NumberResponse{Res: numbers[i]}
		if err := stream.Send(result); err != nil {
			log.Fatalf("Error: %v", err.Error())
		}
		sum := 0

		for i := 0; i < len(a); i++ {
			sum += (arr[i])
		}

		// declaring a variable
		// avg to find the average
		avg := (int32(sum)) / (int32(n))
	}

	return nil
}
