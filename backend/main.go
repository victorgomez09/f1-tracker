package main

import (
	"log"
	"net/http"

	"github.com/victorgomez09/f1-tracker.git/internal"
)

func main() {
	internal.InitStream()
	
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal("Error running up server", err)
	}
}