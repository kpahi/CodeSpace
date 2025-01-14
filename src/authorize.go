//
// Author: Vinhthuy Phan, 2018
//
package main

import (
	"fmt"
	"net/http"
	"strconv"
)

//-----------------------------------------------------------------
// Authorize localhost
//-----------------------------------------------------------------
func AuthorizeLocalhost(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Host == "localhost:8080" {
			fn(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Println("Unauthorized access: host is not local.")
			fmt.Fprint(w, "Unauthorized access: host is not local.")
		}
	}
}

//-----------------------------------------------------------------
func Authorize(fn func(http.ResponseWriter, *http.Request, string, int), userRole string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		givenRole := r.FormValue("role")
		if userRole != "" && userRole != givenRole {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Println("Unauthorized access:", r.FormValue("name"))
			fmt.Fprint(w, "Unauthorized access. Please register again.")
			return
		}
		uid, err := strconv.Atoi(r.FormValue("uid"))
		if err == nil {
			ok := false
			var password string
			if givenRole == "teacher" {
				password, ok = Teacher[uid]
				if ok && password != r.FormValue("password") {
					ok = false
				}
				if !ok {
					c, err := r.Cookie("session_token")
					if err != nil {
						ok = false
						fmt.Println("No Token")
					} else {
						sessionToken := c.Value
						userSession, exists := sessions[sessionToken]
						if !exists || userSession.isExpired() {
							ok = false
						} else {
							ok = true
						}
					}
				}
			} else {
				_, ok = Students[uid]
				if !ok {
					ok = load_and_authorize_student(uid, r.FormValue("password"))
				} else if Students[uid].Password != r.FormValue("password") {
					ok = false
				}
			}
			if ok {
				fn(w, r, r.FormValue("name"), uid)
				return
			}
		}
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Println("Unauthorized access:", r.FormValue("name"))
		fmt.Fprint(w, "Unauthorized access. Please register again.")
	}
}

//-----------------------------------------------------------------
