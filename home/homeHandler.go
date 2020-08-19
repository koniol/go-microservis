package home

import (
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"time"
)

type Handlers struct {
	logger *log.Logger
	db     *sqlx.DB
}


func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("<h1>Ready to works</h1>"))
}

func (h *Handlers) Login(w http.ResponseWriter, r *http.Request) {
	type D struct {
		Name int `db:"name"`
		Row  int` db:"row"`
	}

	var f = D{}
	if err := h.db.Select(&f, "SELECT 123 as name, 1 as row"); err != nil {
		h.logger.Println("err", err)
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("<h1>Login works</h1>"))
}

func (h *Handlers) LoggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		defer h.logger.Printf("Request time %s\n", time.Since(startTime))
		next(w, r)
	}

}

func (h *Handlers) SetupRouters(mux *http.ServeMux) {
	mux.HandleFunc("/", h.LoggerMiddleware(h.Home))
	mux.HandleFunc("/login", h.LoggerMiddleware(h.Login))
}

func NewHandlers(logger *log.Logger, db *sqlx.DB) *Handlers {
	return &Handlers{
		logger: logger,
		db: db,
	}
}
