package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type bugData struct {
	Text string `json:"text"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "TheraPigeon Bug Report Microservice - Now hosted on EB")
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong - The service is online.")
	fmt.Println(r.URL.Query())
}

func soupHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "Sorry, no GET method is available for this endpoint!")
	case "POST":
		q := r.URL.Query()
		value, ok := q["bug"]
		if ok {
			bug := value[0]
			bugReport := bugData{Text: bug}
			jsonReport, err := json.Marshal(bugReport)

			fmt.Println(r.URL.Query())
			resp, err := http.Post(os.Getenv("SOUP_HOOK"), "application/json", bytes.NewReader(jsonReport))

			if err != nil {
				fmt.Fprintf(w, fmt.Sprintln(err))
			} else {
				fmt.Println(resp.Status)
				fmt.Fprintf(w, "OK")
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Please provide a bug url param containing a string.")
		}

	default:
		fmt.Fprintf(w, "Sorry! Only GET and POST methods are supported on this endpoint.")
	}
}

func main() {
	// Loads environmental variables
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Can't load a .env file...")
		if os.Getenv("SOUP_HOOK") != "" {
			fmt.Println("That's ok though! I found the variable I was looking for.")
		} else {
			panic(err)
		}
	}

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/ping/", pingHandler)
	http.HandleFunc("/soup/", soupHandler)

	fmt.Printf("Starting server for TheraPigeon bug reporter...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
