package main

import (
	"fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("src/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/js/login-script.js", func(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "src/assets/js/login-script.js") // Adjust the path as needed
})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "src/assets/index.html")
	})

	// Use the GetAccountIdsHandler from the API package
	//http.HandleFunc("/api/getAccountIds", api.GetAccountIdsHandler)
	//http.HandleFunc("/api/getAccountData", api.GetAccountDataHandler)

	fmt.Println("Server started on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}


