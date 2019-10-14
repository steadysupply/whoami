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

    fmt.Fprintf(os.Stdout, "port :%s\n", port)
    hostname, _ := os.Hostname()
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

        // just to demo handling json in go
        response := map[string]string{
            "name": os.Getenv("NAME"),
            "hostname": hostname,
            "build": "1",
        }
        response_bytes, _ := json.Marshal(response)
        fmt.Fprintf(os.Stdout, "sent %s\n", response_bytes)
 	    w.Write(response_bytes)

    })


    log.Fatal(http.ListenAndServe(":" + port, nil))
}

