package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func checkStatus(url string, ch chan Result) {
	start := time.Now()
	resp, err := http.Get(url)
	if err == nil {
		defer resp.Body.Close()
	}
	duration := time.Since(start)
	result := Result{
		URL:      url,
		Status:   resp.Status,
		Duration: duration.String(),
	}
	if err != nil {
		result.ErrMessage = err.Error()
	}
	ch <- result
}

type Result struct {
	URL        string `json:"url"`
	Status     string `json:"status"`
	Duration   string `json:"duration"`
	ErrMessage string `json:"errMessage"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	urls := r.URL.Query()["url"]
	if len(urls) == 0 {
		http.Error(w, "Requires list of urls in Query param", http.StatusBadRequest)
		return
	}
	ch := make(chan Result)
	results := make([]Result, 0, len(urls))
	for _, url := range urls {
		go checkStatus(url, ch)
	}
	for i := 0; i < len(urls); i++ {
		res := <-ch
		results = append(results, res)
	}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(results)
	if err != nil {
		http.Error(w, "Failed to encode message", http.StatusInternalServerError)
		return
	}

}
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Server is running")
	})
	http.HandleFunc("/status", handler)
	fmt.Println("Server running on 8000")
	http.ListenAndServe(":8000", nil)
}
