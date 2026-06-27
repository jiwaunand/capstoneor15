package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "healthy"}`))
	})

	fileServer := http.FileServer(http.Dir("./public"))
	http.Handle("/", fileServer)

	port := "8080"
	fmt.Printf("Server berjalan. Silakan buka http://localhost:%s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("Gagal menjalankan server: ", err)
	}
}
