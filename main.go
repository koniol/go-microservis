package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"microservices/m/home"
	"microservices/m/server"
	"net/http"
	"os"
)

var (
	SecureServerAddress = ":443" // CHANGE THIS TO PROPER ADDRESS
	ServerCert         = "./certs/MyCertificate.crt"
	ServerKey          = "./certs/MyKey.key"
)

func main() {

	logger := log.New(os.Stdout, "microservices ", log.LstdFlags|log.Lshortfile)
	// Redirect all http request to secure https protocol
	go server.RedirectToTls()

	mux := http.NewServeMux()

	db, err := sqlx.Connect("postgres", "user=root dbname=workflow password=root sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	if err != nil {
		log.Fatalf("DB error %v", err)
	}

	if err = db.Ping(); err !=nil {
		log.Fatalf("Ping error %v", err)
	}

	h := home.NewHandlers(logger, db)
	h.SetupRouters(mux)

	srv := server.New(mux, SecureServerAddress)
	if err := srv.ListenAndServeTLS(ServerCert, ServerKey); err != nil {
		logger.Fatalf("Server error %v", err)
	}

}
