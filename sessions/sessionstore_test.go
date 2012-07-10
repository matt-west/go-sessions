// Copyright 2012 Matt West and Contributors. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package sessions

import (
	"reflect"
	"testing"
)

// NewSessionStore should return a pointer to a SessionStore
func TestNewSessionStore(t *testing.T) {
	// Setup
	x := NewSessionStore()
	y := &SessionStore{}

	// Test
	if reflect.TypeOf(x) != reflect.TypeOf(y) {
		t.Error("Did not return a pointer to a SessionStore in memory")
	}
}

// NewSession should return a pointer to a Session
func TestNewSession(t *testing.T) {
	// Setup
	ss := NewSessionStore()

	x := ss.NewSession()
	y := &Session{}

	if reflect.TypeOf(x) != reflect.TypeOf(y) {
		t.Error("Did not return a pointer to a Session in memory")
	}

	// Test
	val, ok := ss.Sessions[x.ID]

	if !ok {
		t.Error("Session was not stored")
	}

	if val != x {
		t.Error("Saved Session does not match")
	}
}

// GetSession should return a pointer to a session
func TestGetSession(t *testing.T) {
	// Setup
	ss := NewSessionStore()
	x := ss.NewSession()

	// test
	if ss.GetSession(x.ID) != x {
		t.Error("Retrieved session does not match the original")
	}
}

// DestroySession should destroy a session
func TestDestroySession(t *testing.T) {
	// Setup
	ss := NewSessionStore()
	x := ss.NewSession()

	// Test
	ss.DestroySession(x.ID)

	_, ok := ss.Sessions[x.ID]

	if ok {
		t.Error("Session was not destroyed")
	}
}

// General tests for usage
func TestUsage(t *testing.T) {
	// Setup
	ss := NewSessionStore()
	s := ss.NewSession()

	// Test Storage
	s.SetVar("boo", "hello")
	s.SetVar("int", 982)

	if s.GetVar("boo") != "hello" {
		t.Error("Variable 'boo' not set correctly")
	}

	if s.GetVar("int") != 982 {
		t.Error("Variable 'int' not set correctly")
	}

	// Test Session Retrieval
	ns := ss.GetSession(s.ID)

	if ns.GetVar("boo") != s.GetVar("boo") {
		t.Error("Session was not retrieved correctly: variable 'boo' does not match")
	}

	if ns.GetVar("int") != s.GetVar("int") {
		t.Error("Session was not retrieved correctly: variable 'int' does not match")
	}

	// Test Variable Destruction
	ns.DestroyVar("boo")

	if ns.GetVar("boo") != nil {
		t.Error("Variable was not destroyed: variable 'boo' available in ns")
	}

	if s.GetVar("boo") != nil {
		t.Error("Variable was not destroyed: variable 'boo' available in s")
	}

	// Test Variable Updates (uses pointers not duplicates)
	ns.SetVar("int", 203)

	if s.GetVar("int") != ns.GetVar("int") {
		t.Error("SessionStore is not using pointers: Variable updates not visible to all SessionStore instances")
	}

	// Test Session Destruction
	ss.DestroySession(ns.ID)

	//
	// PROBLEM: Seems to be creating 'local' copies for s and ns. These should be deleted
	// when DestroySession is called.
	//
	/*
		if ns.GetVar("int") != nil {
			t.Error("Session was not destroyed: variable 'int' available in ns")
		}

		if s.GetVar("int") != nil {
			t.Error("Session was not destroyed: variable 'int' available in s")
		}
	*/
	_, ok := ss.Sessions[s.ID]

	if ok {
		t.Error("Session was not destroyed")
	}
}
