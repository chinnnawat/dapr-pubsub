package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func logRequest(r *http.Request) {
	log.Printf("Received %s request for %s", r.Method, r.URL.Path)
}

func getAPIData(w http.ResponseWriter, r *http.Request) {
	logRequest(r) // Log the request details
	if r.Method == http.MethodGet {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("Failed to read request body:", err)
			http.Error(w, "Failed to read request body", http.StatusInternalServerError)
			return
		}
		log.Printf("Request body: %s", string(body))
		datetimeNow := time.Now()
		log.Printf("Current time: %s", datetimeNow.Format("2006-01-02 15:04:05"))
		fmt.Fprintln(w, "message from service-c")
	} else {
		log.Println("Method not allowed")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func apiData(w http.ResponseWriter, r *http.Request) {
	logRequest(r) // Log the request details
	if r.Method == http.MethodPost {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("Failed to read request body:", err)
			http.Error(w, "Failed to read request body", http.StatusInternalServerError)
			return
		}

		log.Printf("Request body: %s", string(body))

		datetimeNow := time.Now()
		log.Printf("Current time: %s", datetimeNow.Format("2006-01-02 15:04:05"))

		log.Println("Responding with: ok")
		fmt.Fprintln(w, "ok")
	} else {
		log.Println("Method not allowed")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/api", getAPIData)
	http.HandleFunc("/api/data", apiData)

	log.Println("Server is starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
