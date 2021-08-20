package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
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
		fmt.Println("Email exists in our record, check your email!")
		sendMail(user.Email)

		json.NewEncoder(w).Encode(user)
	} else {
		fmt.Println("Email does not exist!")
		return
	}

}

func sendMail(email string) {

	err = godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	from := "blessedmadukoma@gmail.com"
	password := os.Getenv("PASSWD")

	// toList is list of email address that email is to be sent.
	toList := []string{email}

	host := "smtp.gmail.com"

	port := "587"

	// Test run OTP
	OTP := 1234

	msg := "Hello " + email + ". This is you OTP: " + strconv.Itoa(OTP)

	body := []byte(msg)

	auth := smtp.PlainAuth("", from, password, host)

	err := smtp.SendMail(host+":"+port, auth, from, toList, body)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Successfully sent mail")
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
