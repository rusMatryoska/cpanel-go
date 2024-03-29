package incidents

import (
	"context"
	"fmt"

	"github.com/jackc/pgconn"
	"github.com/rusMatryoska/cpanel-go/internal/storage/postgresql"
)

func GetIncidentByID(ctx context.Context, db *postgresql.Database, id string) (pgconn.CommandTag, error) {
	query := fmt.Sprintf("select incident_id, incident_type_cd, camera_rk, start_dttm, end_dttm, creator from %s.incidents where incident_id = %s", db.Schema, id)
	res, err := db.Exec(ctx, query)
	return res, err
}
