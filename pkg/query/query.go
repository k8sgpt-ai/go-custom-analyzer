package query

import (
	"context"

	rpc "buf.build/gen/go/k8sgpt-ai/k8sgpt/grpc/go/schema/v1/schemav1grpc"
	v1 "buf.build/gen/go/k8sgpt-ai/k8sgpt/protocolbuffers/go/schema/v1"
)

type Handler struct {
	rpc.ServerQueryServiceServer
}
type Query struct {
	Handler *Handler
}

func (Query) Query(context.Context, *v1.QueryRequest) (*v1.QueryResponse, error) {

	return &v1.QueryResponse{
		// Not implemented
	}, nil
}
