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

func main() {


	httpPort := getEnv("HTTP_PORT", defaultHTTPPort)
	grpcPort := getEnv("GRPC_PORT", defaultGRPCPort)
	metricsPort := getEnv("METRICS_PORT", defaultMetricsPort)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	errCh := make(chan error, 3)

	go func() {
		if err := runHTTP(ctx, httpPort); err != nil {
			errCh <- fmt.Errorf("HTTP server error: %w", err)
		}
	}()

	go func() {
		if err := runGRPC(ctx, grpcPort); err != nil {
			errCh <- fmt.Errorf("gRPC server error: %w", err)
		}
	}()

	go func() {
		if err := runMetrics(ctx, metricsPort); err != nil {
			errCh <- fmt.Errorf("metrics server error: %w", err)
		}
	}()

	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	select {
	case sig := <-stopCh:
		log.Printf("Received signal: %v, shutting down gracefully...", sig)
		cancel()

		shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), shutdownTimeout)
		defer shutdownCancel()

		// Wait for all servers to shutdown
		done := make(chan struct{})
		go func() {
			time.Sleep(shutdownTimeout)
			close(done)
		}()

		select {
		case <-done:
			log.Println("Shutdown completed")
		case <-shutdownCtx.Done():
			log.Println("Shutdown timeout exceeded")
		}
	case err := <-errCh:
		log.Fatalf("Server error: %v", err)
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func runHTTP(ctx context.Context, port string) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("pong!\n"))
	})

	httpServer := &http.Server{
		Addr:         ":" + port,
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		<-ctx.Done()
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			log.Printf("HTTP server shutdown error: %v", err)
		}
	}()

	log.Printf("HTTP server starting on port %s", port)
	if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}
