package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type t_metrics struct {
	myapp_start       prometheus.Gauge
	myapp_running_sec prometheus.Counter
}

type t_vars struct {
	startDt time.Time
}

var (
	g_mt  t_metrics
	g_var t_vars
)

func main() {
	{
		g_var.startDt = time.Now()

		fmt.Printf("init completed. \n")
		fmt.Printf("\t- startDt: %s \n", g_var.startDt.Format(time.RFC3339))
	}

	{
		g_mt.myapp_start = promauto.NewGauge(prometheus.GaugeOpts{
			Name: "myapp_start",
			Help: "프로세스 시작 시간(unix초)",
			ConstLabels: map[string]string{
				"rfc3339": g_var.startDt.Format(time.RFC3339),
				"sec":     strconv.FormatInt(g_var.startDt.Unix(), 10),
				"msec":    strconv.FormatInt(g_var.startDt.UnixMilli(), 10),
				"usec":    strconv.FormatInt(g_var.startDt.UnixMicro(), 10),
			},
		})
		g_mt.myapp_start.Set(1)

		g_mt.myapp_running_sec = promauto.NewCounter(prometheus.CounterOpts{
			Name: "myapp_running_sec",
			Help: "프로세스 시작 시간(unix초)",
		})
	}

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":9101", nil)
}

func Loop1() {
	tickerSec := time.NewTicker(time.Second)

	for {
		select {
		case <-tickerSec.C:
			g_mt.myapp_running_sec.Add(1)
		}
	}
}
