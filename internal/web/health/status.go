// Package health - 服务健康状态监控
package health

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// SessionStatus 会话状态结构
type SessionStatus struct {
	IsRunning bool `json:"is_running"` // 运行状态
}

// Status 服务状态监控
func Status(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	data, err := json.Marshal(&SessionStatus{IsRunning: true})
	if err != nil {
		log.Fatal(err)
	}
	w.Write(data)
}
