package incidents

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rusMatryoska/cpanel-go/internal/storage/postgresql"
	"golang.org/x/exp/slog"

	"github.com/rusMatryoska/cpanel-go/internal/storage/postgresql/incidents"
)

func GetIncidentByID(ctx context.Context, log *slog.Logger, db *postgresql.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.Incidents.GetIncidentByID"

		log := log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		id := chi.URLParam(r, "id")
		_, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "ID parameter must be int type", http.StatusBadRequest)
			return
		}
		res, err := incidents.GetIncidentByID(ctx, db, id)
		if err != nil {

			log.Error(fmt.Sprint(err))
			http.Error(w, fmt.Sprint(err), http.StatusBadRequest)
			return
		}
		fmt.Println(res)
		w.Write([]byte(id))
	}
}
