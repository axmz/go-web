package main

import (
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/cat.jpg", cat)
	http.ListenAndServe(":8080", nil)
}

func cat(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("cat.jpg")
	if err != nil {
		http.Error(w, "Not found", 404)
		return
	}

	fi, err := f.Stat()
	if err != nil {
		http.Error(w, "Not found", 404)
		return
	}

	http.ServeContent(w, r, f.Name(), fi.ModTime(), f)
}
