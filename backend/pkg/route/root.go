package route

import (
	"database/sql"
	"net/http"

	"social-network/backend/pkg/handler"
)

func InitRoot(dbConn *sql.DB) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/register", handler.RegisterHandler(dbConn))
	return mux
}
