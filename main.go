package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/jgrancell/logger"
	"github.com/jgrancell/logger/formats"
	"github.com/julienschmidt/httprouter"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	l    *logger.Logger
	port int = 8080

	totalRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Count of GET requests.",
		},
		[]string{"path"},
	)
)

func main() {
	os.Exit(Run())
}

func Run() int {
	l = &logger.Logger{
		Format: &formats.KubernetesFormat{},
	}
	if err := l.Init(); err != nil {
		fmt.Println("unable to initialize logger")
		return 1
	}

	prometheus.Register(totalRequests)

	// Healthcheck endpoint for Kubernetes
	router := httprouter.New()
	router.GET("/", IndexController)
	l.Debug("endpoint registered: /")

	router.GET("/health", HealthController)
	l.Debug("endpoint registered: /health")

	router.GET("/healthz", HealthController)
	l.Debug("endpoint registered: /healthz")

	// The DRY configuration would be to use `/api/:currency`
	//   but that isn't what was asked for in the doc
	router.GET("/api/:currency", CoinbaseController)
	l.Debug("endpoint registered: /api/:currency")

	// This impements /:currency in a non-DRY way without the /api namespace
	router.GET("/USD", CoinbaseController)
	router.GET("/EUR", CoinbaseController)
	router.GET("/GBP", CoinbaseController)
	router.GET("/JPY", CoinbaseController)
	l.Debug("endpoint registered: /{USD|EUR|GBP|JPY}")

	// Metrics endpoint for Prometheus
	router.Handler("GET", "/metrics", promhttp.Handler())
	l.Debug("endpoint registered: /metrics")

	l.Info(fmt.Sprintf("listening on port %d", port))
	return l.LogAndExit(http.ListenAndServe(":8080", router))
}

func LogRequest(status int, r *http.Request, l *logger.Logger) {
	message := fmt.Sprintf(
		"%s %s %d - %s",
		r.Method,
		r.RequestURI,
		status,
		r.RemoteAddr,
	)
	if status == 200 {
		l.Info(message)
	} else {
		l.Warning(message)
	}
}
