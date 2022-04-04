package main

import (
	"fmt"
	"net/http"
	"time"
)

const port = "8080"

func sleeper(w http.ResponseWriter, r *http.Request) {
	if sleep := r.URL.Query().Get("sleep"); sleep != "" {
		sleeptime, err := time.ParseDuration(sleep)
		if err != nil {
			w.Write([]byte(fmt.Sprintf("Error parsing '%s': %s. Did you use a Go duration (like 1m10s)?", sleep, err)))
			return
		}
		time.Sleep(sleeptime)
		w.Write([]byte(fmt.Sprintf("%v has elapsed. The sleeper must awaken.", sleeptime)))
	} else {
		w.Write([]byte("No sleep interval specified, staying woke."))
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", sleeper)
	http.ListenAndServe(fmt.Sprintf(":%s", port), mux)
}
