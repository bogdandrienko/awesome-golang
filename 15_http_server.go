package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type User6 struct {
	ID   int    `json:"id"`
	Name string `json:"name,omitempty"`
}

var (
	users = []User6{{1, "Vasya"}, {2, "Petya"}, {3, "Alice"}}
)

func handleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getUsers(w, r)
	case http.MethodPost:
		addUser(w, r)
	case http.MethodPut:
	//
	case http.MethodDelete:
		//
	default:
		w.WriteHeader(http.StatusNotImplemented)
		return
	}
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	resp, err := json.Marshal(users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Println(resp)
	_, err = w.Write(resp)
	if err != nil {
		log.Fatal(err)
	}
}

func addUser(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var user User6
	if err = json.Unmarshal(reqBytes, &user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	users = append(users, user)
}

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] %s\n", r.Method, r.URL)
		userID := r.Header.Get("x-id")
		if userID == "" {
			log.Printf("[%s] %s - error: userID is not provided\n", r.Method, r.URL)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, "id", userID)

		r = r.WithContext(ctx)

		next(w, r)
	}
}

func loggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idFromCtx := r.Context().Value("id")
		userID, ok := idFromCtx.(string)
		if !ok {
			log.Printf("[%s] %s - error: userID is invalid\n", r.Method, r.URL)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		log.Printf("[%s] %s by UserID %s\n", r.Method, r.URL, userID)
		next(w, r)
	}
}

func main() {
	http.HandleFunc("/users", authMiddleware(loggerMiddleware(handleUsers)))

	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatal(err)
	}
}

// curl -v localhost:8080/users
// curl -v -X GET -H 'x-id:1' localhost:8080/users
// curl -v -X POST -H 'x-id:1' localhost:8080/users -d '{"id": 5, "name": "Anya"}'
