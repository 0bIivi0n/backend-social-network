package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	sndb "socialNetwork/Database"
	h "socialNetwork/Handlers"
)

func main() {

	os.Remove("socialNetwork.db")

	db, err := sndb.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/api/login", func(w http.ResponseWriter, r *http.Request) {
		h.LoginHandler(w, r, db)
	})

	http.HandleFunc("/api/logout", h.LogoutHandler)

	http.HandleFunc("/api/signup", func(w http.ResponseWriter, r *http.Request) {
		h.SignupHandler(w, r, db)
	})
	http.HandleFunc("/api/profile", h.ProfileHandler)
	http.HandleFunc("/api/posts", h.PostsHandler)
	http.HandleFunc("/api/groups", h.GroupsHandler)
	http.HandleFunc("/api/session", h.SessionHandler)

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
