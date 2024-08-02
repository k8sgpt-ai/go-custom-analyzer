package analyzer

import (
	"context"
	"fmt"

	rpc "buf.build/gen/go/k8sgpt-ai/k8sgpt/grpc/go/schema/v1/schemav1grpc"
	v1 "buf.build/gen/go/k8sgpt-ai/k8sgpt/protocolbuffers/go/schema/v1"
	"github.com/ricochet2200/go-disk-usage/du"
)

type Handler struct {
	rpc.AnalyzerServiceServer
}
type Analyzer struct {
	Handler *Handler
}

func (a *Handler) Run(context.Context, *v1.AnalyzerRunRequest) (*v1.AnalyzerRunResponse, error) {
	println("Running analyzer")
	usage := du.NewDiskUsage("/")
	diskUsage := int((usage.Size() - usage.Free()) * 100 / usage.Size())
	return &v1.AnalyzerRunResponse{
		Result: &v1.Result{
			/*
				a lowercase RFC 1123 subdomain must consist of lower case alphanumeric characters, '-' or '.', and must start and end with an alphanumeric character (e.g. 'example.com',
				regex used for validation is '[a-z0-9]([-a-z0-9]*[a-z0-9])?(\\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*')
			*/
			Name:    "diskuse",
			Details: fmt.Sprintf("Disk usage is %d", diskUsage),
			Error: []*v1.ErrorDetail{
				{
					Text: fmt.Sprintf("Disk usage is %d", diskUsage),
				},
			},
		},
	}, nil
}
