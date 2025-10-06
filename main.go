package main

import (
	"fmt"
	"jwt-tool/handlers"
	"log"
	"net/http"
	"os"
)

func Routes(mux *http.ServeMux) {
	
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	
	mux.Handle("POST /tool", http.HandlerFunc(handlers.CreateJWT))
}

func main() {
	mux := http.NewServeMux()
	Routes(mux)

	
	handler := handlers.Cors(mux)

	
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" 
	}

	serverAddr := ":" + port
	fmt.Println("Server running on", serverAddr)

	
	err := http.ListenAndServe(serverAddr, handler)
	if err != nil {
		log.Fatal("Error Starting The Server: ", err)
	}
}
