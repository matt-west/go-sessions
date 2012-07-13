// Copyright 2012 Matt West and Contributors. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"./sessions"
	"log"
	"net/http"
)

// Create a new session store
// (do this once, when your server starts up)
var ss = sessions.NewSessionStore()

func main() {

	// // Set some Session Variables
	// s.SetVar("boo", "hello")
	// s.SetVar("int", 982)
	// // Get the Session Variables
	// fmt.Println(s.GetVar("boo"))
	// fmt.Println(s.GetVar("int"))

	// // Get the Session
	// ns := ss.GetSession(s.ID)
	// // Check that the variables persisted
	// fmt.Println(ns.GetVar("boo"))
	// fmt.Println(ns.GetVar("int"))

	// // Destroy a Session Variable
	// ns.DestroyVar("boo")
	// fmt.Println(ns.GetVar("boo")) // <nil>
	// fmt.Println(s.GetVar("boo"))  // <nil>

	// // Destroy the Session
	// ss.DestroySession(ns.ID)
	// fmt.Println(ns.GetVar("int")) // <nil>
	// fmt.Println(s.GetVar("int"))  // <nil>

	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":9981", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Get Session will either return a current session or create a new one.
	s := ss.GetSession(w, r)

	s.SetVar("foo", "bar")
}
