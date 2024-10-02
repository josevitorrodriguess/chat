package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("./.env")
	port := os.Getenv("PORT")
	fmt.Print(port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("Server is running on port %s", port)))
	})

	log.Fatal(http.ListenAndServe(port, nil))
}
