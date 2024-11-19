package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Error", 405)
	}
	inputname := r.URL.Query().Get("name")
	w.Write([]byte("Hello," + inputname))
}

func main() {
	http.HandleFunc("/api/user", handler)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("ошибка запуска сервера")
	}
}
