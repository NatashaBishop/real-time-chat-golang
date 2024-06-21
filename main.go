// my-spa-server/main.go
package main

import (
    "net/http"
    "log"
    "path/filepath"
    "os"
)

func main() {
    fs := http.FileServer(http.Dir("./static"))

    http.Handle("/", http.StripPrefix("/", fs))

    // Handle SPA routing
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        // Check if the file exists
        _, err := os.Stat(filepath.Join("./static", r.URL.Path))
        if os.IsNotExist(err) {
            http.ServeFile(w, r, "./static/index.html")
            return
        } else if err != nil {
            log.Fatal(err)
        }

        fs.ServeHTTP(w, r)
    })

    log.Println("Serving on :8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal(err)
    }
}
