func main() {
    router := mux.NewRouter()
    router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {

      
        // an example API handler
        json.NewEncoder(w).Encode(map[string]bool{"ok": true})
    })

    spa := router.PathPrefix("/").Handler(http.StripPrefix("/",http.FileServer(http.Dir("./static/")
    router.PathPrefix("/").Handler(spa)

    srv := &http.Server{
        Handler: router,
        Addr:    "127.0.0.1:8000",
        // Good practice: enforce timeouts for servers you create!
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
    }

    log.Fatal(srv.ListenAndServe())
}
