package gopherconbr2019

import (
    "fmt"
    "net/http"
)

// This is a simple package used just to call the API
func main() {
    fmt.Println("starting the API...")
    http.HandleFunc("/", api.Index_handler)
    http.ListenAndServe(":8000", nil)
}
