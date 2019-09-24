package gopherconbr2019

import (
    "fmt"
    "net/http"
)

// This function returns for the Browsers some informations about the API
// This is mandary for some Browsers allows the access to the API
func enableCors(w *http.ResponseWriter) {
    (*w).Header().Set("Access-Control-Allow-Origin", "*")
}

// This function returns a "Welcome message" to the API
func Index_handler(w http.ResponseWriter, r *http.Request) {
    enableCors(&w)
    fmt.Fprintf(w, "This is a simple page for GopherCon Brasil 2019")
}

// This function take the parameters passed on the API and insert on
// the database
func Graph(w http.ResponseWriter, r *http.Request) {
    enableCors(&w)
    

}
