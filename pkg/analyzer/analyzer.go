package analyzer

import (
	rpc "buf.build/gen/go/k8sgpt-ai/k8sgpt/grpc/go/schema/v1/schemav1grpc"
	v1 "buf.build/gen/go/k8sgpt-ai/k8sgpt/protocolbuffers/go/schema/v1"
	"context"
)

type Handler struct {
	rpc.AnalyzerServiceServer
}
type Analyzer struct {
	Handler *Handler
}

func (a *Handler) Run(context.Context, *v1.AnalyzerRunRequest) (*v1.AnalyzerRunResponse, error) {

	response := &v1.AnalyzerRunResponse{
		Result: &v1.Result{
			Name:    "example",
			Details: "example",
			Error: []*v1.ErrorDetail{
				&v1.ErrorDetail{
					Text: "This is an example error message!",
				},
			},
		},
	}

	return response, nil
}
