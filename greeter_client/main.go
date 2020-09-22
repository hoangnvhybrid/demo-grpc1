package main

import (
	"context"
	"fmt"
	pb "github.com/hoangnvhybrid/demo-grpc1/pb"
	"google.golang.org/grpc"
	"log"
	"os"
)

const (
	defaultName = "world"
)

func main() {
	fmt.Println("Hoang ơi, cố lên, vươn lên để phụng sự nhân sinh và tu tập để về thế giới tây phương cực lạc")
	//Thiết lập kết nối đến server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}
