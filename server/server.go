package main

import (
	"log"
	"net/http"

	"github.com/dayfine/typescript-react-bazel/server/handlers"
)

// var globalSessions *session.Manager

// func init() {
// 	globalSessions, _ = NewManager("memory", "gosessionid", 3600)
// }

func main() {
	http.HandleFunc("/", handlers.SayhelloName)
	http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("/upload", handlers.Upload)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
