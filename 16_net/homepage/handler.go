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
	h.logger.Println("Request processed")
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Homepage loaded"))
}

func New(logger *log.Logger) *Handler {
	return &Handler{logger: logger}
}

// middleware
func (h *Handler) Logger(next http.HandlerFunc) http.HandlerFunc  {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer h.logger.Printf("Request processed in: %v", time.Now().Sub(start))
		next(w, r)
	}
}

// Defining routes locally, makes it easier to tests
// and keeps context in one place
func (h *Handler) SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", h.Logger(h.Home))
}