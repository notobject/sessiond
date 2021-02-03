package debug

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Debug 调试接口
func Debug(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	w.Write([]byte("You want debug modlue -> " + p.ByName("modlue")))
}
