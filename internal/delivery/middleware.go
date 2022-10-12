package delivery

import (
	"encoding/json"
	"net/http"
	types "play3/internal/model"
)

type MiddlewareInferface interface {
	GET(path string, targetFunc http.HandlerFunc)
	POST(path string, targetFunc http.HandlerFunc)
	HandlerRESP(w http.ResponseWriter, status int, data interface{})
}

type Middleware struct {
}

func (m *Middleware) GET(path string, targetFunc http.HandlerFunc) {
	http.HandleFunc(path, m.Handler("GET", targetFunc))
}
func (m *Middleware) POST(path string, targetFunc http.HandlerFunc) {
	http.HandleFunc(path, m.Handler("POST", targetFunc))
}

func (m *Middleware) Handler(method string, next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			m.HandlerRESP(w, http.StatusForbidden, nil)
		} else {
			//here authentication
			//if not, next
			next.ServeHTTP(w, r)
		}
	})
}

func (m *Middleware) HandlerRESP(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	resp := types.JSONResponse{
		Header: types.JSONResponseHeader{
			Code: status,
		},
		Data: data,
	}
	json.NewEncoder(w).Encode(resp)
}
