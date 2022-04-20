package api

import (
	// "golang/internal/model/domain"
	"golang/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
	// metric "golang/pkg/metric"
)

type Server struct {
	// store  domain.Store
	router *gin.Engine
	service service.Service
}

func NewServer(service service.Service) *Server {

	gin.ForceConsoleColor()
	gin.SetMode(gin.ReleaseMode)
	server := &Server{service: service}
	router := gin.New()
	m := ginmetrics.GetMonitor()

	// +optional set metric path, default /debug/metrics
	m.SetMetricPath("/metrics")
	// +optional set slow time, default 5s
	m.SetSlowTime(10)
	// +optional set request duration, default {0.1, 0.3, 1.2, 5, 10}
	// used to p95, p99
	m.SetDuration([]float64{0.1, 0.3, 1.2, 5, 10})

	// set middleware for gin
	m.Use(router)


	// http.Handle("/metrics", promhttp.Handler())
	// router.POST("/users", server.CreateUser)
	// router.GET("/users", server.ListUser)
	router.GET("/healthy",func(ctx *gin.Context) {
		var str = "application healthy"
		ctx.JSON(200,str)
	})
	router.GET("/role/:id", server.GetRoleById)
	router.GET("/product/:id", server.GetProductById)
	router.GET("/order/:id",server.GetOrderByid)
	// router.GET("/metrics",server.CounterRequest)
	server.router = router
	return server
}
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
