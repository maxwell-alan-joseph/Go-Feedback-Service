package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
)

type Feedback struct {
    Name    string `json:"name"`
    Email   string `json:"email"`
    Message string `json:"message"`
}

func feedbackHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
        return
    }

    var fb Feedback
    err := json.NewDecoder(r.Body).Decode(&fb)
    if err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    log.Printf("Received feedback from %s (%s): %s", fb.Name, fb.Email, fb.Message)
    w.WriteHeader(http.StatusCreated)
    fmt.Fprintln(w, "Feedback received successfully!")
}

func main() {
    http.HandleFunc("/feedback", feedbackHandler)
    fmt.Println("Server is running on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
