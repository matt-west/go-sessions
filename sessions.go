// Copyright 2012 Matt West and Contributors. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"./sessions"
	"fmt"
)

// Create a new session store
// (do this once, when your server starts up)
var ss = sessions.NewSessionStore()

func main() {
	// Create a new Session
	s := ss.NewSession()

	// Set some Session Variables
	s.SetVar("boo", "hello")
	s.SetVar("int", 982)
	// Get the Session Variables
	fmt.Println(s.GetVar("boo"))
	fmt.Println(s.GetVar("int"))

	// Get the Session
	ns := ss.GetSession(s.ID)
	// Check that the variables persisted
	fmt.Println(ns.GetVar("boo"))
	fmt.Println(ns.GetVar("int"))

	// Destroy a Session Variable
	ns.DestroyVar("boo")
	fmt.Println(ns.GetVar("boo")) // <nil>
	fmt.Println(s.GetVar("boo")) // <nil>

	// Destroy the Session
	ss.DestroySession(ns.ID)
	fmt.Println(ns.GetVar("boo")) // <nil>
	fmt.Println(s.GetVar("boo"))  // <nil>
}
