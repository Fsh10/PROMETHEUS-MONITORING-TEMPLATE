# Prometheus Metrics Template

<div align="center">

![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go)
![License](https://img.shields.io/badge/License-MIT-green.svg?style=for-the-badge)
![Docker](https://img.shields.io/badge/Docker-Ready-2496ED?style=for-the-badge&logo=docker)
![Prometheus](https://img.shields.io/badge/Prometheus-Ready-E6522C?style=for-the-badge&logo=prometheus)

**Production-ready Go application with gRPC API for demonstrating and working with Prometheus metrics**

[Architecture](#-architecture) • [Quick Start](#-quick-start) • [Documentation](#-documentation) • [Development](#-development)

</div>

---

## 📋 Table of Contents

- [About the Project](#-about-the-project)
- [Features](#-features)
- [Architecture](#-architecture)
- [Requirements](#-requirements)
- [Quick Start](#-quick-start)
- [Configuration](#-configuration)
- [API Documentation](#-api-documentation)
- [Monitoring](#-monitoring)
- [Development](#-development)
- [Deployment](#-deployment)
- [Troubleshooting](#-troubleshooting)
- [License](#-license)

---

## 🎯 About the Project

**Prometheus Metrics Template** is a production-ready Go application template that demonstrates best practices for working with Prometheus metrics through gRPC API. The project includes a complete monitoring infrastructure with Prometheus and Grafana, ready for use in production environments.

### Core Principles

- ✅ **Production-ready** — ready for production use
- ✅ **Best practices** — follows Go development best practices
- ✅ **Graceful shutdown** — proper shutdown of all services
- ✅ **Observability** — full observability through metrics
- ✅ **Containerization** — optimized Docker images
- ✅ **Documentation** — comprehensive documentation and examples

---

## ✨ Features

### Metric Types

| Type | Description | Usage |
|------|------------|-------|
| **Counter** | Monotonically increasing counter | Counting total number of requests |
| **Gauge** | Value that can increase and decrease | Current value of a random metric |
| **Histogram** | Distribution of values into buckets | Analyzing request durations |
| **Summary** | Statistical summary with quantiles | Request sizes with percentiles |

### Technical Features

- 🚀 **Three independent servers**: HTTP, gRPC, Metrics
- 🔄 **Graceful shutdown** with timeouts
- ⚙️ **Configuration via environment variables**
- 🐳 **Multi-stage Docker build** for minimal image size
- 🏥 **Healthchecks** for all services
- 📊 **Ready-to-use monitoring infrastructure** (Prometheus + Grafana)
- 🔒 **Production-ready** error handling and logging

---

## 🏗 Architecture

### System Components

```
┌─────────────────────────────────────────────────────────────┐
│                    Prometheus Metrics App                    │
├─────────────────────────────────────────────────────────────┤
│                                                               │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐      │
│  │ HTTP Server  │  │ gRPC Server  │  │Metrics Server│      │
│  │   Port 80    │  │   Port 82    │  │  Port 84     │      │
│  │              │  │              │  │              │      │
│  │ /ping        │  │ Metrics API  │  │ /metrics     │      │
│  └──────────────┘  └──────────────┘  └──────────────┘      │
│         │                 │                  │             │
│         └─────────────────┼──────────────────┘             │
│                           │                                │
│                  ┌─────────▼─────────┐                      │
│                  │  Prometheus       │                      │
│                  │  Client Library   │                      │
│                  └───────────────────┘                      │
│                                                               │
└─────────────────────────────────────────────────────────────┘
                            │
                            ▼
┌─────────────────────────────────────────────────────────────┐
│              Monitoring Infrastructure                      │
├─────────────────────────────────────────────────────────────┤
│                                                               │
│  ┌──────────────┐              ┌──────────────┐            │
│  │  Prometheus  │◄─────────────┤   Grafana    │            │
│  │  Port 9090   │   Scrapes    │   Port 3000  │            │
│  └──────────────┘              └──────────────┘            │
│                                                               │
└─────────────────────────────────────────────────────────────┘
```

### Project Structure

```
prometheus-template/
├── api/                    # Protocol Buffers definitions
│   └── metrics.proto
├── cmd/                    # Application entry point
│   └── main.go
├── internal/               # Internal packages
│   └── api/                # gRPC handlers and metrics
│       ├── counter.go
│       ├── gauge.go
│       ├── histogram.go
│       ├── summary.go
│       ├── server.go
│       └── rand.go
├── pkg/                    # Generated protobuf files
│   └── metrics/
├── prometheus/             # Prometheus configuration
│   ├── prometheus.yml
│   └── alert.yml
├── docker-compose.yml      # Docker Compose configuration
├── Dockerfile             # Multi-stage Docker build
├── Makefile               # Task automation
└── README.md              # Documentation
```

---

## 📦 Requirements

### Required

- **Go** 1.21 or higher
- **Docker** 20.10+ and **Docker Compose** 2.0+
- **protoc** (Protocol Buffers compiler) 3.0+

### Optional

- **golangci-lint** for code linting
- **grpcurl** for testing gRPC API

### Installing Dependencies

<details>
<summary><b>macOS</b></summary>

```bash
# Homebrew
brew install go protobuf docker docker-compose

# Install protoc-gen-go plugins
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

</details>

<details>
<summary><b>Linux (Ubuntu/Debian)</b></summary>

```bash
# Install Go
wget https://go.dev/dl/go1.21.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.21.linux-amd64.tar.gz

# Install protoc
sudo apt-get update
sudo apt-get install -y protobuf-compiler

# Install Docker
curl -fsSL https://get.docker.com -o get-docker.sh
sudo sh get-docker.sh

# Install protoc-gen-go plugins
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

</details>

<details>
<summary><b>Windows</b></summary>

```powershell
# Chocolatey
choco install golang protoc docker-desktop

# Install protoc-gen-go plugins
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

</details>

---

## 🚀 Quick Start

### Option 1: Docker Compose (Recommended)

The fastest way to start the entire stack:

```bash
# Clone the repository
git clone https://github.com/Fsh10/PROMETHEUS-MONITORING-TEMPLATE.git
cd prometheus_metrics_example

# Start all services
make docker-up

# Check status
docker-compose ps

# View logs
make docker-logs
```

After startup, the following are available:
- 🌐 **HTTP API**: http://localhost:8080/ping
- 📊 **Metrics**: http://localhost:8084/metrics
- 🔍 **Prometheus**: http://localhost:9000
- 📈 **Grafana**: http://localhost:3000 (admin/admin)

### Option 2: Local Development

For development and debugging:

```bash
# Install dependencies
make deps

# Generate protobuf files
make generate

# Build application
make build-local

# Run
make run
```

The application will start on ports:
- HTTP: `8080` (if HTTP_PORT=8080 is set)
- gRPC: `8082` (if GRPC_PORT=8082 is set)
- Metrics: `8084` (if METRICS_PORT=8084 is set)

---

## ⚙️ Configuration

### Environment Variables

The application is configured via environment variables:

| Variable | Description | Default |
|----------|-------------|---------|
| `HTTP_PORT` | HTTP server port | `80` |
| `GRPC_PORT` | gRPC server port | `82` |
| `METRICS_PORT` | Metrics server port | `84` |

### Configuration Example

```bash
# Set environment variables
export HTTP_PORT=8080
export GRPC_PORT=8082
export METRICS_PORT=8084

# Run application
./bin/app
```

Or via Docker Compose:

```yaml
services:
  app:
    environment:
      - HTTP_PORT=8080
      - GRPC_PORT=8082
      - METRICS_PORT=8084
```

---

## 📚 API Documentation

### gRPC API

The application provides a gRPC API for working with metrics. All methods accept `google.protobuf.Empty` and return `BaseResponse`.

#### Service: MetricsService

```protobuf
service MetricsService {
  rpc Counter(google.protobuf.Empty) returns (BaseResponse);
  rpc Gauge(google.protobuf.Empty) returns (BaseResponse);
  rpc Histogram(google.protobuf.Empty) returns (BaseResponse);
  rpc Summary(google.protobuf.Empty) returns (BaseResponse);
}
```

#### Methods

##### Counter

Increments the request counter by 1.

**Request:**
```json
{}
```

**Response:**
```json
{
  "status": "ok"
}
```

**Metric:** `sandbox_requests_total` (Counter)

**Usage Example:**
```bash
grpcurl -plaintext localhost:8082 \
  prometheus.example.api.MetricsService/Counter
```

##### Gauge

Sets a random value from 0 to 100.

**Request:**
```json
{}
```

**Response:**
```json
{
  "status": "ok"
}
```

**Metric:** `sandbox_rand_current` (Gauge)

**Usage Example:**
```bash
grpcurl -plaintext localhost:8082 \
  prometheus.example.api.MetricsService/Gauge
```

##### Histogram

Records request duration in a histogram.

**Request:**
```json
{}
```

**Response:**
```json
{
  "status": "ok"
}
```

**Metric:** `sandbox_request_duration_seconds` (Histogram)

**Buckets:** Standard Prometheus buckets (0.005, 0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1, 2.5, 5, 10)

**Usage Example:**
```bash
grpcurl -plaintext localhost:8082 \
  prometheus.example.api.MetricsService/Histogram
```

##### Summary

Records request size in a summary.

**Request:**
```json
{}
```

**Response:**
```json
{
  "status": "ok"
}
```

**Metric:** `sandbox_request_size_bytes` (Summary)

**Objectives:** 
- 50th percentile (p50)
- 90th percentile (p90)
- 99th percentile (p99)

**Usage Example:**
```bash
grpcurl -plaintext localhost:8082 \
  prometheus.example.api.MetricsService/Summary
```

### HTTP API

#### GET /ping

Healthcheck endpoint to verify application health.

**Request:**
```bash
curl http://localhost:8080/ping
```

**Response:**
```
pong!
```

**Status Code:** `200 OK`

### Metrics API

#### GET /metrics

Prometheus metrics in Prometheus exposition format.

**Request:**
```bash
curl http://localhost:8084/metrics
```

**Response:**
```
# HELP sandbox_requests_total Total number of counter requests
# TYPE sandbox_requests_total counter
sandbox_requests_total 42

# HELP sandbox_rand_current Current random value between 0 and 100
# TYPE sandbox_rand_current gauge
sandbox_rand_current 73.5

# HELP sandbox_request_duration_seconds Histogram of request durations in seconds
# TYPE sandbox_request_duration_seconds histogram
sandbox_request_duration_seconds_bucket{le="0.005"} 10
sandbox_request_duration_seconds_bucket{le="0.01"} 25
...

# HELP sandbox_request_size_bytes Summary of request sizes in bytes
# TYPE sandbox_request_size_bytes summary
sandbox_request_size_bytes{quantile="0.5"} 5000
sandbox_request_size_bytes{quantile="0.9"} 9000
...
```

---

## 📊 Monitoring

### Prometheus

Prometheus automatically collects metrics from the application every 5 seconds.

**Access:** http://localhost:9000

**Main Queries:**

```promql
# Total number of Counter requests
sandbox_requests_total

# Current Gauge value
sandbox_rand_current

# Average request duration
rate(sandbox_request_duration_seconds_sum[5m]) / rate(sandbox_request_duration_seconds_count[5m])

# 95th percentile of request duration
histogram_quantile(0.95, rate(sandbox_request_duration_seconds_bucket[5m]))

# 99th percentile of request size
sandbox_request_size_bytes{quantile="0.99"}
```

### Grafana

Grafana is pre-installed and ready to use.

**Access:** http://localhost:3000

**Credentials:**
- Username: `admin`
- Password: `admin` (change on first login)

**Setting up Prometheus as a data source:**

1. Go to **Configuration → Data Sources**
2. Click **Add data source**
3. Select **Prometheus**
4. URL: `http://prometheus:9090`
5. Click **Save & Test**

### Alerts

Prometheus is configured to send alerts when services are unavailable.

**Rule:** `InstanceDown` — triggers if a service is unavailable for more than 5 minutes.

**Alert configuration:** `prometheus/alert.yml`

---

## 🔧 Development

### Code Structure

```
internal/api/
├── server.go      # gRPC server implementation
├── counter.go     # Counter metric handler
├── gauge.go       # Gauge metric handler
├── histogram.go   # Histogram metric handler
├── summary.go     # Summary metric handler
└── rand.go        # Shared random number generator
```

### Adding a New Metric

1. **Define the metric in the corresponding file:**

```go
var (
    newMetric = prometheus.NewCounter(
        prometheus.CounterOpts{
            Namespace: "sandbox",
            Name:      "new_metric_total",
            Help:      "Description of the new metric",
        },
    )
)

func init() {
    _ = prometheus.Register(newMetric)
}
```

2. **Add the method to server.go:**

```go
func (s server) NewMethod(ctx context.Context, _ *emptypb.Empty) (*pb.BaseResponse, error) {
    newMetric.Inc()
    return &pb.BaseResponse{Status: "ok"}, nil
}
```

3. **Update the proto file:**

```protobuf
service MetricsService {
  ...
  rpc NewMethod(google.protobuf.Empty) returns (BaseResponse);
}
```

4. **Regenerate protobuf files:**

```bash
make generate
```

### Make Commands

| Command | Description |
|---------|-------------|
| `make deps` | Install Go dependencies |
| `make generate` | Generate protobuf files |
| `make build` | Build for Linux (production) |
| `make build-local` | Build for local platform |
| `make run` | Run application locally |
| `make test` | Run tests |
| `make fmt` | Format code |
| `make lint` | Run linter (requires golangci-lint) |
| `make tidy` | Clean and update dependencies |
| `make clean` | Remove build artifacts |
| `make docker-build` | Build Docker image |
| `make docker-up` | Start Docker Compose |
| `make docker-down` | Stop Docker Compose |
| `make docker-logs` | View application logs |

### Testing

```bash
# Run all tests
make test

# Run tests with coverage
go test -v -cover ./...

# Run specific test
go test -v ./internal/api -run TestCounter
```

### Linting

```bash
# Install golangci-lint
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Run linter
make lint
```

### Formatting

```bash
# Automatic formatting
make fmt

# Check formatting
gofmt -d .
```

---

## 🚢 Deployment

### Docker

#### Building the Image

```bash
# Build image
make docker-build

# Or directly
docker build -t prometheus-metrics-app:latest .
```

#### Running the Container

```bash
docker run -d \
  --name metrics-app \
  -p 8080:80 \
  -p 8082:82 \
  -p 8084:84 \
  -e HTTP_PORT=80 \
  -e GRPC_PORT=82 \
  -e METRICS_PORT=84 \
  prometheus-metrics-app:latest
```

### Kubernetes

Example Kubernetes manifest:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: prometheus-metrics-app
spec:
  replicas: 3
  selector:
    matchLabels:
      app: prometheus-metrics-app
  template:
    metadata:
      labels:
        app: prometheus-metrics-app
    spec:
      containers:
      - name: app
        image: prometheus-metrics-app:latest
        ports:
        - containerPort: 80
          name: http
        - containerPort: 82
          name: grpc
        - containerPort: 84
          name: metrics
        env:
        - name: HTTP_PORT
          value: "80"
        - name: GRPC_PORT
          value: "82"
        - name: METRICS_PORT
          value: "84"
        livenessProbe:
          httpGet:
            path: /ping
            port: 80
          initialDelaySeconds: 10
          periodSeconds: 10
        readinessProbe:
          httpGet:
            path: /ping
            port: 80
          initialDelaySeconds: 5
          periodSeconds: 5
---
apiVersion: v1
kind: Service
metadata:
  name: prometheus-metrics-app
spec:
  selector:
    app: prometheus-metrics-app
  ports:
  - name: http
    port: 80
    targetPort: 80
  - name: grpc
    port: 82
    targetPort: 82
  - name: metrics
    port: 84
    targetPort: 84
```

### Production Recommendations

1. **Use environment variables** for configuration
2. **Configure healthchecks** for all services
3. **Use reverse proxy** (nginx/traefik) for HTTP
4. **Set up monitoring** and alerts
5. **Use secrets management** for sensitive data
6. **Configure logging** to a centralized system
7. **Use resource limits** in Kubernetes/Docker

---

## 🐛 Troubleshooting

### Issue: Application Won't Start

**Symptoms:** Error on startup, port is busy

**Solution:**
```bash
# Check occupied ports
lsof -i :80 -i :82 -i :84

# Change ports via environment variables
export HTTP_PORT=8080
export GRPC_PORT=8082
export METRICS_PORT=8084
```

### Issue: Protobuf Files Not Generating

**Symptoms:** Error when running `make generate`

**Solution:**
```bash
# Check protoc installation
protoc --version

# Install protoc-gen-go plugins
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Check PATH
echo $PATH | grep -q "$(go env GOPATH)/bin" || export PATH=$PATH:$(go env GOPATH)/bin
```

### Issue: Prometheus Not Collecting Metrics

**Symptoms:** Metrics not showing in Prometheus

**Solution:**
```bash
# Check metrics availability
curl http://localhost:8084/metrics

# Check Prometheus configuration
docker-compose exec prometheus cat /etc/prometheus/prometheus.yml

# Check Prometheus logs
docker-compose logs prometheus
```

### Issue: Docker Image Too Large

**Symptoms:** Large Docker image size

**Solution:**
- Multi-stage build is used to minimize size
- Final image is based on Alpine Linux (~10MB)
- If image is still large, check Docker cache:
```bash
docker system prune -a
```

### Issue: Graceful Shutdown Not Working

**Symptoms:** Hanging connections remain when stopping the application

**Solution:**
- Check timeouts in `cmd/main.go`
- Ensure correct context is used for shutdown
- Check logs on shutdown:
```bash
docker-compose logs app | grep -i shutdown
```

---

## 📝 Best Practices

### Metrics

1. **Use meaningful names** for metrics with prefixes
2. **Add Help descriptions** for all metrics
3. **Group metrics** by namespace
4. **Use correct types** of metrics for tasks
5. **Don't create too many metrics** — this affects performance

### Code

1. **Handle all errors** — don't ignore them
2. **Use contexts** for operation cancellation
3. **Implement graceful shutdown** for all services
4. **Use structured logging**
5. **Write tests** for critical functionality

### Deployment

1. **Use healthchecks** to verify readiness
2. **Set up monitoring** before deployment
3. **Use versioning** for images
4. **Test in staging** before production
5. **Document changes** in CHANGELOG

---

## 📄 License

This project is licensed under the MIT License — see the [LICENSE](LICENSE) file for details.

---

## 🙏 Acknowledgments

- [Prometheus](https://prometheus.io/) for the excellent monitoring system
- [gRPC](https://grpc.io/) for the high-performance RPC framework
- [Grafana](https://grafana.com/) for the powerful visualization platform

---

<div align="center">

**Made with ❤️ for the Go developer community**

[⬆ Back to Top](#prometheus-metrics-template)

</div>
