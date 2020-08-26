package main

import (
	"strconv"
	"strings"
	"time"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	//HTTPReqDuration metric:http_request_duration_seconds
	HTTPReqDuration *prometheus.HistogramVec
	//HTTPReqTotal metric:http_request_total
	HTTPReqTotal *prometheus.CounterVec
)

func init() {
	// 监控接口请求耗时
	HTTPReqDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "http_request_duration_seconds",
		Help:    "The HTTP request latencies in seconds",
		Buckets: nil,
	}, []string{"method", "path"})

	// 监控接口请求次数
	HTTPReqTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total number of HTTP requests made.",
	}, []string{"method", "path", "status"})

	prometheus.MustRegister(HTTPReqDuration, HTTPReqTotal)
}

// /api/epgInfo/1371648200  -> /api/epgInfo
func parsePath(path string) string {
	itemList := strings.Split(path, "/")
	if len(path) >= 4 {
		return strings.Join(itemList[0:3], "/")
	}
	return path
}

func Metric() gin.HandlerFunc {
	return func(c *gin.Context) {
		tBegin := time.Now()
		c.Next()
		path := parsePath(c.Request.URL.Path)

		HTTPReqTotal.With(prometheus.Labels{
			"method": c.Request.Method,
			"path":   path,
			"status": strconv.Itoa(c.Writer.Status()),
		}).Inc()

		HTTPReqDuration.With(prometheus.Labels{
			"method": c.Request.Method,
			"path":   path,
		}).Observe(time.Since(tBegin).Seconds())
	}
}

func DealAPI1(c *gin.Context) {
	time.Sleep(time.Microsecond * 10)
	c.Writer.Write([]byte("/api/api1"))
}

func DealAPI2(c *gin.Context) {
	time.Sleep(time.Microsecond * 20)
	c.Writer.Write([]byte("/api/api2"))
}

func main() {
	router := gin.Default()
	pprof.Register(router, "/debug/pprof")
	g := router.Group("/api")

	g.Use(Metric())

	g.GET("api1", DealAPI1)
	g.GET("api2", DealAPI2)

	// 暴露给Prometheus
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	router.Run(":28181")
}
