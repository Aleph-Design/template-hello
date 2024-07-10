package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// func WriteToConsole(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println("Hit the page")
// 		next.ServeHTTP(w, r)
// 	})
// }

func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	// set info to create a surf token
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.SetSecure,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// loads and save session data for current request, and 
// communicate session token to and from client in a cookie.
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
