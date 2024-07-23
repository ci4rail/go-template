package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"

	// This controls the maxprocs environment variable in container runtimes.
	// see https://martin.baillie.id/wrote/gotchas-in-the-go-network-packages-defaults/#bonus-gomaxprocs-containers-and-the-cfs
	_ "go.uber.org/automaxprocs"
)

var (
	sensorValue = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "sensor_value",
			Help: "Current value of the sensor",
		},
		[]string{"sensor"},
	)
)

func init() {
	prometheus.MustRegister(sensorValue)
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "an error occurred: %s\n", err)
		os.Exit(1)
	}
}

func run() error {
	logger, err := newJSONLogger(os.Getenv("LOG_LEVEL"))
	if err != nil {
		return err
	}

	defer func() {
		err = logger.Sync()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error syncing logger: %s\n", err)
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	go func() {
		for {
			// Fetch and update the sensor values here
			value := getSensorValue("sensor1")
			sensorValue.WithLabelValues("sensor1").Set(value)
			logger.Info("Sensor value updated", zap.String("sensor", "sensor1"), zap.Float64("value", value))

			// Add logic for other sensors as needed

			// Sleep for a certain interval before fetching the values again
			time.Sleep(10 * time.Second)
		}
	}()

	logger.Info("Starting server on :8080")
	return http.ListenAndServe(":8080", nil)
}

func newJSONLogger(level string) (*zap.Logger, error) {
	config := zap.NewProductionConfig()
	if level != "" {
		var zapLevel zap.AtomicLevel
		err := zapLevel.UnmarshalText([]byte(level))
		if err != nil {
			return nil, err
		}
		config.Level = zapLevel
	}
	return config.Build()
}

func getSensorValue(sensor string) float64 {
	// Implement the logic to fetch the sensor values here
	// Example: returning a dummy value
	return 42.0 // Example value
}
