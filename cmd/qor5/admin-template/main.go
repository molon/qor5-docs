package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/qor5/docs/v3/cmd/qor5/admin-template/admin"
)

func main() {
	// Setup project
	mux := admin.Initialize()

	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}

	fmt.Println("Served at http://localhost:" + port + "/admin")

	http.Handle("/", mux)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}
