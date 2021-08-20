package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type User struct {
	Email string `json:"email"`
}

var err error

func forgotPasswordAPI(w http.ResponseWriter, r *http.Request) {
	emails := []string{"b@gmail.com", "s@gmail.com", "a@gmail.com"}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	var user User
	json.Unmarshal(reqBody, &user)

	// Validate form input
	if strings.Trim(user.Email, " ") == "" {
		fmt.Println("Parameter's can't be empty")
		http.Redirect(w, r, "/forgot", http.StatusMovedPermanently)
		return
	}

	// Check the array of emails
	if contains(emails, user.Email) {
		fmt.Println("Email exists, check your email!")
		// func sendMail(user.Email) {}
		json.NewEncoder(w).Encode(user)
	} else {
		fmt.Println("Email does not exist!")
		return
	}

}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/forgot", forgotPasswordAPI).Methods("POST")

	fmt.Println("Server starting port 8000")
	// err = http.ListenAndServe(":8000", context.ClearHandler(http.DefaultServeMux))
	err = http.ListenAndServe(":8000", router)
	if err != nil {
		log.Fatal(err)
	}

}
