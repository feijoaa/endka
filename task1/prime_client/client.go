package main
import (
	"context"
	"endka/task1/protoc"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main() {
	fmt.Println("Hello I'm a client")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No connection: %v", err)
	}
	defer conn.Close()

	c := protoc.NewNumServiceClient(conn)
	findPrimes(c)
}


func findPrimes(c protoc.NumServiceClient) {

	ctx := context.Background()
	//req := &protoc.NumberRequest{Number : &protoc.Number{
	//	Number: 100,
	//}}
	req := &protoc.NumberRequest{Num: &protoc.Number{
		Number: 100,
	},
	}

	stream, err := c.Number(ctx, req)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	defer stream.CloseSend()

LOOP:
	for {
		res, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				// we've reached the end of the stream
				break LOOP
			}
			log.Fatalf("Error  %v", err)
		}
		log.Printf("Responce : %v \n", res.GetRes())
	}
}
