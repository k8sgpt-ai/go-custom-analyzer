package analyzer

import (
	rpc "buf.build/gen/go/k8sgpt-ai/k8sgpt/grpc/go/schema/v1/schemav1grpc"
	v1 "buf.build/gen/go/k8sgpt-ai/k8sgpt/protocolbuffers/go/schema/v1"
	"context"
	"github.com/ricochet2200/go-disk-usage/du"
)

type Handler struct {
	rpc.AnalyzerServiceServer
}
type Analyzer struct {
	Handler *Handler
}

func (a *Handler) Run(context.Context, *v1.AnalyzerRunRequest) (*v1.AnalyzerRunResponse, error) {

	usage := du.NewDiskUsage("/")
	diskUsage := int((usage.Size() - usage.Free()) * 100 / usage.Size())
	var response = &v1.AnalyzerRunResponse{}
	if diskUsage > 90 {
		response = &v1.AnalyzerRunResponse{
			Result: &v1.Result{
				Name:    "Disk Usage",
				Details: "Disk usage is above 90%",
				Error: []*v1.ErrorDetail{
					&v1.ErrorDetail{
						Text: "Disk usage is above 90%",
					},
				},
			},
		}
	}
	return response, nil
}
