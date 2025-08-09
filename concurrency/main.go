package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

var fortunes = []string{
	"You will be successful",
	"Today will be a great day",
	"You will achieve your dreams",
	"You will marry a good wife",
	"Today you will win",
}

func tellFortune(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "stranger"
	}
	go func(msg string) {
		time.Sleep(3 * time.Second)
		fmt.Println(msg)
	}("Delay completed")
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(fortunes))
	fortune := fortunes[index]
	fmt.Fprintf(w, "Hello %s Your fortune %s", name, fortune)

}
func main() {
	http.HandleFunc("/fortune", tellFortune)
	fmt.Println("Server is running on port 8000")
	http.ListenAndServe(":8000", nil)
}
