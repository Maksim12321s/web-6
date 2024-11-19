package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var counter int = 0

type kol struct {
	Count int `json:"count"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]int{"count": counter})

	case "POST":
		var c kol
		err := json.NewDecoder(r.Body).Decode(&c)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		fmt.Println(c.Count)
		counter += c.Count
	}
}
func main() {
	http.HandleFunc("/count", handler)

	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера")
	}
}
