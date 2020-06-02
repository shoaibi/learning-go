package homepage

import (
	"log"
	"net/http"
	"time"
)

type Handler struct {
	logger *log.Logger
}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	h.logger.Println("[Handler] Loaded Homepage")
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Homepage loaded"))
}

func New(logger *log.Logger) *Handler {
	return &Handler{logger: logger}
}

// middlewares
func (h *Handler) Timed(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer h.logger.Printf("Request processed in: %v", time.Now().Sub(start))
		// could also do next(w,r)
		next.ServeHTTP(w, r)
	}
}

func (h *Handler) Log(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.logger.Println("Request Method: ", r.Method)
		h.logger.Println("Request URL: ", r.URL)
		h.logger.Println("Request UserAgent: ", r.UserAgent())
		h.logger.Println("Request Host: ", r.Host)
		h.logger.Println("Request Protocol: ", r.Proto)
		h.logger.Println("Requester IP: ", r.RemoteAddr)
		// could also do next(w,r)
		next.ServeHTTP(w, r)
	}
}

// Defining routes locally, makes it easier to test
// and keeps context in one place
func (h *Handler) SetupRoutes(mux *http.ServeMux) {
	commonMiddleware := []Middleware{h.Timed, h.Log}
	mux.HandleFunc("/", with(h.Home, commonMiddleware...))
}
