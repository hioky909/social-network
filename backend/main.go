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

	log.Println("Serveur Go lancé sur http://localhost:8080")
	http.ListenAndServe(":8080", handler)
}
