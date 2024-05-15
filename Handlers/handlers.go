package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	st "socialNetwork/Structs"
	"strings"
)

func LoginHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	fmt.Println("entered LoginHandler")

	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == "POST" {

		fmt.Println("Method is POST")

		// Setting response header
		w.Header().Set("Content-Type", "application/json")

		// Get login creds from front
		var loginData st.LoginData
		err := json.NewDecoder(r.Body).Decode(&loginData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Println("creds = ", loginData)

		// Get user
		var user st.User
		err = db.QueryRow("SELECT * FROM users WHERE email = ?", loginData.Email).Scan(
			&user.Uuid,
			&user.Username,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.Password,
			&user.DateOfBirth,
			&user.AboutMe,
			&user.HaveImage,
		)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "Invalid email or password", http.StatusUnauthorized)
				return
			}
			http.Error(w, "Database error", http.StatusInternalServerError)
			fmt.Println(err)
			return
		}

		fmt.Println("user found in db")

		fmt.Println(user)

		// Check password
		if user.Password == loginData.Password {
			// login success
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode((map[string]string{"message": "Login succesfull"}))
			fmt.Println("Provided password is good")
		} else {
			// login fail
			http.Error(w, "Invalid email or password", http.StatusUnauthorized)
			return
		}

	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		fmt.Println("Looks like method is not POST")
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Add logout logic here

		fmt.Fprintln(w, "Logout Endpoint")
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func SignupHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == "POST" {

		w.Header().Set("Content-Type", "application/json")

		var newUser st.User
		err := json.NewDecoder(r.Body).Decode(&newUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Println(newUser)

		uuid := 1
		haveImage := false

		username := newUser.Username
		firstName := newUser.FirstName
		lastName := newUser.LastName
		email := newUser.Email
		password := newUser.Password
		dateOfBirth := newUser.DateOfBirth
		aboutMe := newUser.AboutMe
		haveImage = newUser.HaveImage

		query, err := db.Prepare("INSERT INTO users(uuid, username, first_name, last_name, email, password, date_of_birth, about_me, have_image) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)")
		if err != nil {
			fmt.Println("Error in query to add user")
			return
		}

		_, err = query.Exec(uuid, username, firstName, lastName, email, password, dateOfBirth, aboutMe, haveImage)
		if err != nil {
			fmt.Printf("Error executing query to add user with username: %s", username)
			return
		}

		fmt.Fprintln(w, "New user registered")
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func ProfileHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	fmt.Println("Entered Profile Handler")

	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	path := strings.TrimPrefix(r.URL.Path, "/api/profile/")
	if path == "" {
		http.Error(w, "Profile UUID is missing", http.StatusBadRequest)
		fmt.Println("profile UUID is missing")
		return
	}

	uuid := path

	fmt.Printf("Profile requested for UUID: %s", uuid)
}

func GetPostsHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	fmt.Println("entered GetPostHandler")

	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Example of handling multiple types of requests in one function
	if r.Method != "GET" {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	} else {

		// Setting response header
		w.Header().Set("Content-Type", "application/json")

		var posts []st.Post

		rows, err := db.Query("SELECT * FROM posts")
		if err != nil {
			fmt.Println("Error reading post table")
			return
		}

		defer rows.Close()

		for rows.Next() {
			var post st.Post
			err := rows.Scan(&post.Uuid, &post.UserUuid, &post.Content, &post.Date, &post.Type, &post.GroupUuid, &post.HaveImage)
			if err != nil {
				http.Error(w, "Error getting posts in db", http.StatusInternalServerError)
				fmt.Println(err)
				return
			}
			posts = append(posts, post)
		}
		fmt.Println(posts)
	}
}

func GroupsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprintln(w, "Get Groups or Group Details")
	case "POST":
		fmt.Fprintln(w, "Manage Group Members Endpoint")
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func SessionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprintln(w, "Session Check Endpoint")
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
