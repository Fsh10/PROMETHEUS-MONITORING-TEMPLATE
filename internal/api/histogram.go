package api


import (
	"context"
	"time"

	pb "github.com/Fsh10/PROMETHEUS-MONITORING-TEMPLATE/pkg/metrics/github.com/Fsh10/PROMETHEUS-MONITORING-TEMPLATE"

	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/protobuf/types/known/emptypb"
	"fmt"
	"math/rand"

)

var (
	requestDurationHistogram = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Namespace: "sandbox",
			Name:       "request_duration_seconds",
			Help:       "Histogram of request durations in seconds",
			Buckets:    prometheus.DefBuckets,
		},
	)
)

func init() {


	_ = prometheus.Register(requestDurationHistogram)
}