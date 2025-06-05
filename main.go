package main

import (
	"errors"
	"fmt"
	"net"
	"net/http"

	rpc "buf.build/gen/go/k8sgpt-ai/k8sgpt/grpc/go/schema/v1/schemav1grpc"
	"github.com/k8sgpt-ai/go-custom-analyzer/pkg/analyzer"
	"github.com/k8sgpt-ai/go-custom-analyzer/pkg/query"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	var err error
	address := fmt.Sprintf(":%s", "8085")
	lis, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	aa := analyzer.Analyzer{}
	qh := query.Query{}

	// Some stub examples of other capabilities that can be registered
	rpc.RegisterServerQueryServiceServer(grpcServer, qh)
	//
	rpc.RegisterCustomAnalyzerServiceServer(grpcServer, aa.Handler)
	fmt.Println("Starting!")
	if err := grpcServer.Serve(
		lis,
	); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return
	}
}
