package metrics

import (
	"regexp"
	"testing"

	"github.com/candiddev/shared/go/assert"
	"github.com/prometheus/client_golang/prometheus"
)

func TestNew(t *testing.T) {
	r := regexp.MustCompile("^homechart_")

	Setup("homechart")
	CacheRequestTotal.WithLabelValues("AuthAccount", "hit").Add(1)
	ControllerEventTotal.WithLabelValues("signup", "web", "latest").Add(1)
	ControllerRequestTotal.WithLabelValues("/test", "POST", "500").Add(1)
	ControllerRequestDuration.WithLabelValues("/", "POST").Observe(1)
	ControllerRequestSize.WithLabelValues("/", "POST").Observe(1)
	ControllerResponseSize.WithLabelValues("/", "POST").Observe(1)
	ControllerSSEClients.WithLabelValues().Set(1)
	ControllerUserAgentTotal.WithLabelValues("Go").Add(1)
	Notifications.WithLabelValues("fcm", "cancelled").Add(1)
	PostgreSQLQueryDuration.WithLabelValues("auth_account", "insert").Observe(1)
	TaskRunner.WithLabelValues().Set(1)
	TelemetryErrorTotal.WithLabelValues("/home", "1.1.1").Add(1)

	metrics, _ := prometheus.DefaultGatherer.Gather()
	for _, metric := range metrics {
		n := metric.GetName()
		v := metric.GetMetric()

		if r.MatchString(n) {
			assert.Equal(t, len(v), 1)
		}
	}
}
