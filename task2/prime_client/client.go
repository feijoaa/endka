package main

import (
	"context"
	"endka/task2/protoc"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func main() {
	fmt.Println("Hello I'm a client")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	c := protoc.NewNumServiceClient(conn)
	findAvg(c)
}

func findAvg(c protoc.NumServiceClient) {
	ctx:= context.Background()
	arr := []int{1, 2, 3, 4}
	req := &protoc.NumberRequest{Num: &protoc.Number{
		Number: arr,
	},
	}

}