package main

import (
	"log"
	"time"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/service/grpc"
	hello "github.com/ricardojonathanromero/grpc/greeter/srv/proto/hello"

	"context"
)

// Say struct
type Say struct{}

// Hello func
func (s *Say) Hello(ctx context.Context, req *hello.Request, rsp *hello.Response) error {
	log.Print("Received Say.Hello request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

func main() {
	service := grpc.NewService(
		micro.Name("go.micro.srv.greeter"),
		micro.Address(":9090"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	// optionally setup command line usage
	service.Init()

	// Register Handlers
	hello.RegisterSayHandler(service.Server(), new(Say))

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
