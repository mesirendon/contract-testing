package main

import (
	"log"
	"net"
	"net/http"

	"github.com/mesirendon/contract-testing/provider/internal/middleware"
	"github.com/mesirendon/contract-testing/provider/internal/model"
)

func main() {
	mux := middleware.GetHTTPHandler(&map[int]model.User{
		10: {
			FirstName: "John",
			LastName:  "Doe",
			Username:  "drwho",
			Type:      "user",
			ID:        10,
		},
	})

	ln, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	log.Printf("API starting: port %d (%s)", 8082, ln.Addr())
	log.Printf("API terminating: %v", http.Serve(ln, mux))
}
