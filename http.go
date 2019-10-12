package main

import (
  "encoding/json"
  "fmt"
  "log"
  "net/http"
  "os"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    fmt.Fprintf(os.Stdout, "Listening on :%s\n", port)
    hostname, _ := os.Hostname()
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

        fmt.Fprintf(os.Stdout, "I’m %s\n", hostname)

        // just to demo handling json in go
        response := map[string]string{
            "name": os.Getenv("NAME"),
            "hostname": hostname,
        }
        response_bytes, _ := json.Marshal(response)
 	    w.Write(response_bytes)

    })


    log.Fatal(http.ListenAndServe(":" + port, nil))
}

