package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {/* 
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"id":42}`))

		fmt.Fprintf(w, "path=%s method=%s", r.URL.Path, r.URL)
	}) */

	http.HandleFunc("/created", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated) // 201
	})
	http.HandleFunc("/denied", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusForbidden) // 403
	})
	http.HandleFunc("/boom", func(w http.ResponseWriter, r *http.Request) {
		panic("nil pointer") // Go сам вернёт 500
	})

	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		var in map[string]any
		if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
			w.WriteHeader(http.StatusBadRequest) // 400 из шага 3
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{"you_sent": in})
	})
	http.ListenAndServe(":8080", nil)

	mux := http.NewServeMux()
	 mux.HandleFunc("GET /users/{id}", getUser)      // GET  → 200
	// mux.HandleFunc("POST /users", createUser)       // POST → 201
	// mux.HandleFunc("DELETE /users/{id}", deleteUser) // DELETE → 204
	http.ListenAndServe(":8080", mux)


	


}


	func getUser(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")   // "42"
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"id":%s,"name":"Alice"}`, id)
	}