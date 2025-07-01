package main

import (
	"log"
	"net/http"
	"social-network/backend/pkg/db"
	"social-network/backend/pkg/route"
)

func main() {
	dbConn := db.InitDB("social_network.db")
	defer dbConn.Close()

	handler := route.InitRoot(dbConn)
	log.Println("Serveur Go lancé sur http://localhost:8081")
	http.ListenAndServe(":8081", withCORS(handler))
}

// Ajoute ce middleware pour autoriser le CORS
func withCORS(h http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }
        h.ServeHTTP(w, r)
    })
}
