package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/dgrijalva/jwt-go"

	"github.com/Mark-Otte/EMS/graph"
)

// Handler for the /login endpoint
func loginHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve the username and password from the query parameters
	username := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")

	// For this example, check if the username and password are valid
	if username == "admin" && password == "password" {
		// Create a new JWT token
		token := jwt.New(jwt.SigningMethodHS256)

		// Set the claims for the token
		claims := token.Claims.(jwt.MapClaims)
		claims["username"] = username

		// Generate the token string
		tokenString, err := token.SignedString([]byte("your-secret-key"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Return the JWT token
		response := map[string]interface{}{
			"token": tokenString,
		}
		json.NewEncoder(w).Encode(response)
	} else {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
	}
}

func main() {

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	// Define the API endpoints and their corresponding handlers
	http.HandleFunc("/login", loginHandler)
	http.Handle("/employee", srv)

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", nil))
}
