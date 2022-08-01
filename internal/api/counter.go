package api


import (
	"context"

	pb "github.com/Fsh10/PROMETHEUS-MONITORING-TEMPLATE/pkg/metrics/github.com/Fsh10/PROMETHEUS-MONITORING-TEMPLATE"

	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/protobuf/types/known/emptypb"
	"fmt"

)

var (
	requestsCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: "sandbox",
			Name:      "requests_total",
			Help:      "Total number of counter requests",
		},
	)
)
