package api

import (
	"github.com/gin-gonic/gin"
	"metrics"
	"metrics/manager"
)

type apiController struct {
	mm manager.MetricManager
}

func NewApiController(mm manager.MetricManager, router *gin.Engine) {
	api := &apiController{mm}
	router.Handle("POST", "/collect", api.AddMetric)
}

func (api *apiController) AddMetric(g *gin.Context) {
	metric := metrics.Metric{
		ID:   g.PostForm("tid"),
		Type: g.PostForm("t"),
		URL:  g.PostForm("dp"),
	}
	api.mm.AddMetric(metric)
}
