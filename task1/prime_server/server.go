package main

import (
	"endka/task1/protoc"
	"fmt"
"google.golang.org/grpc"
"log"
"math"
"net"
)

type Server struct {
	protoc.UnimplementedNumServiceServer
}

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

	for i := 3; float64(i) <= math.Sqrt(float64(number)); i+=2 {
		for int(number) % i ==0 {
			numbers = append(numbers,int32(i))
			number/=int32(i)
		}
	}

	if number > 2{
		numbers  = append(numbers,number)


	}
	for i:= 0;i<len(numbers);i++{
		result := &protoc.NumberResponse{Res: numbers[i]}
		if err := stream.Send(result); err != nil {
			log.Fatalf("Error: %v", err.Error())
		}
	}

	return nil
}
