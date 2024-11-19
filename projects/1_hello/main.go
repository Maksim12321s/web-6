package main

// здесь надо написать код
import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello,web!"))
}

func main() {
	http.HandleFunc("/get", handler)

	err := http.ListenAndServe(":8083", nil)
	if err != nil {
		fmt.Println("ошибка запуска сервера")
	}
}
