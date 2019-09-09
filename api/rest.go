package api

import (
	"metrics"
	"metrics/manager"

	"github.com/gin-gonic/gin"
)

type apiController struct {
	mm manager.MetricManager
}

//NewAPIController  инициализация обработчика
func NewAPIController(mm manager.MetricManager, router *gin.Engine) {
	api := &apiController{mm}
	router.Handle("POST", "/collect", api.AddMetric)
}

//AddMetric обработчик
func (api *apiController) AddMetric(g *gin.Context) {
	metric := metrics.Metric{ //разбираем POST форму
		ID:   g.PostForm("tid"),
		Type: g.PostForm("t"),
		URL:  g.PostForm("dp"),
	}
	api.mm.AddMetric(metric)
}
