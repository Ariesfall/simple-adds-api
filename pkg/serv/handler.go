package serv

import (
	"encoding/json"
	"net/http"

	"github.com/Ariesfall/simple-odds-api/pkg/data"
	"github.com/jmoiron/sqlx"
)

func Oddshandler(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sports := r.URL.Query().Get("sport")

		in := &data.Match{SportKey: sports}
		res, err := data.ListMatch(db, in)
		if err != nil {
			writeErrorStatus(w, err)
			return
		}

		writeJsonStatus(w, res)
		return
	}
}

func writeJsonStatus(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func writeErrorStatus(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusBadRequest)
}
