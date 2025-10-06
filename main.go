package main

import (
	"fmt"
	"jwt-tool/handlers"
	"net/http"
)

func Routes(mux *http.ServeMux) {
    mux.HandleFunc("/tool", handlers.CreateJWT) // use HandleFunc
}


func main() {

	mux := http.NewServeMux()

	

	Routes(mux)

	fmt.Println("Server running on http://localhost:8080")
	err := http.ListenAndServe(":8080", handlers.Cors(mux))
	if err != nil {
		fmt.Println("Error Starting The Server")
	}
}
