package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("port")
	if port == "" {
		port = "8080"
	}
	startWeb(port)
}

func startWeb(port string) {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		switch request.URL.Path {
		default:
			writer.WriteHeader(http.StatusNotFound)
		case "/":
			_, _ = fmt.Fprint(writer, getContent())
		case "/health":
			writer.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(writer, "healthy")
		}
	})
	_ = http.ListenAndServe(":"+port, nil)
}

func getContent() string {
	return "Hello World"
}
