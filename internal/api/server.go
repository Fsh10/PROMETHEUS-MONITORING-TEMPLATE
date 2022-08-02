package api


import (
	pb "github.com/Fsh10/PROMETHEUS-MONITORING-TEMPLATE/pkg/metrics/github.com/Fsh10/PROMETHEUS-MONITORING-TEMPLATE"
	"fmt"

)

type (
	server struct {
		pb.UnimplementedMetricsServiceServer
	}
)

func NewServer() pb.MetricsServiceServer {


	return server{}
}