package serv

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jmoiron/sqlx"
)

func Service(db *sqlx.DB) http.Handler {
	r := chi.NewRouter()

	// request real ip and log
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)

	// request timeout
	r.Use(middleware.Timeout(60 * time.Second))

	// api path
	r.Get("/odds", Oddshandler(db))

	return r
}
