package main


import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Fsh10/PROMETHEUS-MONITORING-TEMPLATE/internal/api"
	metrics "github.com/Fsh10/PROMETHEUS-MONITORING-TEMPLATE/pkg/metrics/github.com/Fsh10/PROMETHEUS-MONITORING-TEMPLATE"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	defaultHTTPPort    = "80"
	defaultGRPCPort    = "82"
	defaultMetricsPort = "84"
	shutdownTimeout    = 10 * time.Second















)
