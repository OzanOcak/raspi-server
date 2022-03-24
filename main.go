package main

import (
    "fmt"
    "net/http"
)

func main() {
    // Register handler for default route
    http.HandleFunc("/", HelloHandler)

    // Start server listening
    fmt.Println("Listening on port 3000...")
    err := http.ListenAndServe(":3000", nil)
    if err != nil {
        panic(err)
    }

}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
