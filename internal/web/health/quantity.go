// Package health 在线用户数监控
package health

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// SessionQuntity 会话数量结构
type SessionQuntity struct {
	TotalSession int64 `json:"total_session"` // 总会话数量
}

// Quantity 数量监控
func Quantity(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	data, err := json.Marshal(&SessionQuntity{TotalSession: 999})
	if err != nil {
		log.Fatal(err)
	}
	w.Write(data)
}
