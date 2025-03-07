package main

import (
	"log"
	"net/http"

	"k8s-pod-deployer/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/pods", handlers.CreatePod).Methods("POST")
	r.HandleFunc("/pods/{name}", handlers.GetPod).Methods("GET")
	r.HandleFunc("/pods/{name}", handlers.DeletePod).Methods("DELETE")

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("could not start server: %v\n", err)
	}
}
