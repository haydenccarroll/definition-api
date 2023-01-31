package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "strings"

    "github.com/go-sql-driver/mysql"
    "github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func main() {
    // Connect to the MySQL database
    var err error
    db, err = sqlx.Connect("mysql", 
"user:password@tcp(127.0.0.1:3306)/database")
    if err != nil {
        fmt.Println(err)
        return
    }

    // Define the HTTP route for the definition lookup
    http.HandleFunc("/define/", handleDefinitionLookup)

    // Start the server
    fmt.Println("Listening on http://localhost:8080...")
    http.ListenAndServe(":8080", nil)
}

func handleDefinitionLookup(w http.ResponseWriter, r *http.Request) {
    // Get the word from the URL path
    parts := strings.Split(r.URL.Path, "/")
    word := parts[len(parts)-1]

    // Look up the definition via WordsAPI
    client := &http.Client{}
    req, err := http.NewRequest("GET", 
"https://wordsapiv1.p.rapidapi.com/words/"+word, nil)
    if err != nil {
        fmt.Println(err)
        return
    }
    req.Header.Add("X-RapidAPI-Key", "SIGN-UP-FOR-KEY")
    req.Header.Add("Accept", "application/json")

    res, err := client.Do(req)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer res.Body.Close()

    // Decode the JSON response
    var data map[string]interface{}
    json.NewDecoder(res.Body).Decode(&data)

    // Get the definition from the response
    var definition string
    if data["results"] != nil {
        results := data["results"].([]interface{})
        if len(results) > 0 {
            definition = 
results[0].(map[string]interface{})["definition"].(string)
        }
    }

    // Insert the word and definition into the MySQL database
    _, err = db.Exec("INSERT INTO words (word, definition) VALUES (?, ?)", 
word, definition)
    if err != nil {
        fmt.Println(err)
        return
    }

    // Write the definition to the response
    w.Write([]byte(definition))
}

