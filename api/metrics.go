package api

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	// "github.com/prometheus/client_golang/prometheus"
)


func (server *Server) CounterRequest(ctx *gin.Context) {
	// prometheus.
	metric := promhttp.Handler()
	metric.ServeHTTP(ctx.Writer,ctx.Request)
}