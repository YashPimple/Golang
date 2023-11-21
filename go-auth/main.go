package main

import (
	"net/http"

	"github.com/YashPimple/go-auth/controllers"
)

// HandlerFunc is a custom type that satisfies http.Handler for a given function signature.
type HandlerFunc func(http.ResponseWriter, *http.Request)

// ServeHTTP implements the http.Handler interface for HandlerFunc.
func (h HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h(w, r)
}

func main() {
	fs := http.FileServer(http.Dir("public"))

	// Wrap your functions with the custom HandlerFunc type
	signinHandler := HandlerFunc(controllers.Signin)
	callbackHandler := HandlerFunc(controllers.Callback)

	http.Handle("/", fs)
	http.Handle("/signin", signinHandler)
	http.Handle("/callback", callbackHandler)

	http.ListenAndServe(":3000", nil)
}
