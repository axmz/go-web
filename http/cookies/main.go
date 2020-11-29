package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", counter)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func counter(w http.ResponseWriter, r *http.Request) {
	// get cookie
	cookie, err := r.Cookie("counter")
	// if no cookie, set one
	if err == http.ErrNoCookie {
		cookie = &http.Cookie{Name: "counter", Value: "0"}
	}
	// increase counter
	counter, _ := strconv.Atoi(cookie.Value)
	counter++
	cookie.Value = strconv.Itoa(counter)
	// send response
	http.SetCookie(w, cookie)
	fmt.Fprintln(w, counter)
}
