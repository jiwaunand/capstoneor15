package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Menentukan direktori tempat file HTML/statis berada
	fileServer := http.FileServer(http.Dir("./public"))

	// Menangani semua request (routing ke root "/")
	http.Handle("/", fileServer)

	port := "8080"
	fmt.Printf("Server berjalan. Silakan buka http://localhost:%s\n", port)

	// Menjalankan server pada port 8080
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("Gagal menjalankan server: ", err)
	}
}
