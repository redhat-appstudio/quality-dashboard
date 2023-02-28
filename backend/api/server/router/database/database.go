package database

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redhat-appstudio/quality-studio/api/types"
	"github.com/redhat-appstudio/quality-studio/pkg/storage/ent/client"
	"github.com/redhat-appstudio/quality-studio/pkg/utils/httputils"
)

func (d *databaseRouter) getDbConnection(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	cfg := client.GetPostgresConnectionDetails()
	_, err := cfg.Open()
	fmt.Println(err)

	if err != nil {
		return httputils.WriteJSON(w, http.StatusBadRequest, types.ErrorResponse{
			Message:    "database connection is down",
			StatusCode: 400,
		})
	}

	return httputils.WriteJSON(w, http.StatusOK, "database connection is up")
}
