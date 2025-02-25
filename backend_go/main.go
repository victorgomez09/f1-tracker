package main

import (
	"log"
	"net/http"

	socket "github.com/victorgomez09/f1-tracker.git/internal/socket"
)

func main() {
	socket.InitStream()

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal("Error running up server", err)
	}
}
