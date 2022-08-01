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
	gaugeCounter = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "sandbox",
			Name:      "rand_current",
			Help:      "Current random value between 0 and 100",
		},
	)
)

func init() {


	_ = prometheus.Register(gaugeCounter)
}