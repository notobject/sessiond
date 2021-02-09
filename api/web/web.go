package web

// Web服务, 提供pull模式的监控接口

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/notobject/sessiond/internal/web/debug"
	"github.com/notobject/sessiond/internal/web/health"
	"github.com/notobject/sessiond/pkg/config"
)

// Start 开启web服务
func Start(cfg *config.Config) {

	router := httprouter.New()

	router.GET("/debug/:modlue", debug.Debug)
	router.GET("/health/quantity", health.Quantity)
	router.GET("/health/status", health.Status)

	go http.ListenAndServe(cfg.Addr.Web, router)
	log.Println("start web service at ", cfg.Addr.Web)
}
