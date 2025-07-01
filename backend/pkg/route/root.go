package route

import (
	"database/sql"
	"net/http"

	"social-network/backend/pkg/handler"
)

func InitRoot(dbConn *sql.DB) http.Handler {
    mux := http.NewServeMux()
    mux.HandleFunc("/api/register", handler.RegisterHandler(dbConn))
    mux.HandleFunc("/api/login", handler.LoginHandler(dbConn))

    // Exemple de route protégée
    mux.Handle("/api/protected", handler.JWTMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Accès autorisé avec un token valide !"))
    })))

    return mux
}