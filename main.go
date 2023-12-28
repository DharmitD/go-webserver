package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

// Define a struct to represent user data.
type User struct {
	Username string
	Password string
	Name     string
	Email    string
}

// Define a struct to represent form submissions.
type FormSubmission struct {
	ID      int
	User    string
	Message string
}

// Define a struct to represent data for the template.
type PageData struct {
	Title       string
	Message     string
	User        *User
	Submissions []FormSubmission
}

// Global session store.
var store = sessions.NewCookieStore([]byte("your-secret-key"))

func main() {
	// Create a new router using gorilla/mux.
	router := mux.NewRouter()

	// Define routes.
	router.HandleFunc("/", HomeHandler).Methods("GET")
	router.HandleFunc("/form", FormPageHandler).Methods("GET")
	router.HandleFunc("/submit", FormSubmitHandler).Methods("POST")
	router.HandleFunc("/login", LoginHandler).Methods("GET", "POST")
	router.HandleFunc("/logout", LogoutHandler).Methods("GET")
	router.HandleFunc("/dashboard", DashboardHandler).Methods("GET")

	// Serve static files.
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Start the HTTP server on port 8080.
	fmt.Println("Server is listening on port 8080...")
	http.Handle("/", router)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

// Define user data.
var users = map[string]*User{
	"user1": {
		Username: "user1",
		Password: "password1",
		Name:     "User One",
		Email:    "user1@example.com",
	},
	"user2": {
		Username: "user2",
		Password: "password2",
		Name:     "User Two",
		Email:    "user2@example.com",
	},
}

// In-memory data storage for form submissions.
var submissions []FormSubmission

// Define session names.
const sessionName = "your-session-name"

// HomeHandler handles the homepage route.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title:   "Home Page",
		Message: "Welcome to the Home Page!",
	}

	session, _ := store.Get(r, sessionName)
	username, ok := session.Values["username"].(string)
	if ok {
		data.User = users[username]
	}

	tmpl, err := template.ParseFiles("templates/home.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, data)
}

// FormPageHandler handles the form page route.
func FormPageHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/form.html")
}

// FormSubmitHandler handles form submissions.
func FormSubmitHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the user is authenticated.
	session, _ := store.Get(r, sessionName)
	username, ok := session.Values["username"].(string)
	if !ok {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	message := r.FormValue("message")

	// Simulate a database insert.
	submission := FormSubmission{
		ID:      len(submissions) + 1,
		User:    username,
		Message: message,
	}
	submissions = append(submissions, submission)

	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

// LoginHandler handles user login.
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")

		user, ok := users[username]
		if ok && user.Password == password {
			session, _ := store.Get(r, sessionName)
			session.Values["username"] = username
			session.Save(r, w)

			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	}

	tmpl, err := template.ParseFiles("templates/login.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

// LogoutHandler handles user logout.
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, sessionName)
	delete(session.Values, "username")
	session.Save(r, w)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// DashboardHandler handles the user dashboard.
func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, sessionName)
	username, ok := session.Values["username"].(string)
	if !ok {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Filter form submissions for the logged-in user.
	var userSubmissions []FormSubmission
	for _, submission := range submissions {
		if submission.User == username {
			userSubmissions = append(userSubmissions, submission)
		}
	}

	data := PageData{
		Title:       "Dashboard",
		Message:     "Welcome to your Dashboard!",
		User:        users[username],
		Submissions: userSubmissions,
	}

	tmpl, err := template.ParseFiles("templates/dashboard.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, data)
}
