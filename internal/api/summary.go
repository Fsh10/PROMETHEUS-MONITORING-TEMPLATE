package api


import (
	"context"

	pb "github.com/Fsh10/PROMETHEUS-MONITORING-TEMPLATE/pkg/metrics/github.com/Fsh10/PROMETHEUS-MONITORING-TEMPLATE"

	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/protobuf/types/known/emptypb"
	"fmt"
	"math/rand"

)

var (
	requestSizeSummary = prometheus.NewSummary(
		prometheus.SummaryOpts{
			Namespace:  "sandbox",
			Name:       "request_size_bytes",
			Help:       "Summary of request sizes in bytes",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		},
	)
)

func init() {


	_ = prometheus.Register(requestSizeSummary)
}