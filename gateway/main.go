package main

import (
	"flag"
	"fmt"
	"net/http"

	"context"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	hello "github.com/ricardojonathanromero/grpc/gateway/proto/hello"
)

var (
	// the go.micro.srv.greeter address
	endpoint = flag.String("endpoint", "localhost:9090", "go.micro.srv.greeter address")
)

func run() error {
	fmt.Println("trace....")
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := hello.RegisterSayHandlerFromEndpoint(ctx, mux, *endpoint, opts)
	if err != nil {
		return err
	}

	fmt.Println("trace2")

	return http.ListenAndServe(":8080", mux)
}

func main() {
	fmt.Println("test")
	flag.Parse()

	defer glog.Flush()

	err := run()

	fmt.Println("trace3")

	if err != nil {
		glog.Fatal(err)
	}
}
