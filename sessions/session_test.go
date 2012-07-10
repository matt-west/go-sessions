// Copyright 2012 Matt West and Contributors. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package sessions

import (
	"testing"
)

// GetVar should retrieve a session variable
func TestGetVar(t *testing.T) {
	// Setup
	ss := NewSessionStore()
	s := ss.NewSession()
	s.SetVar("foo", "bar")

	// Test
	x := s.GetVar("foo")

	if x != "bar" {
		t.Errorf("GetVar('foo') = %v, want 'bar'", x)
	}
}

// SetVar should set a session variable
func TestSetVar(t *testing.T) {
	// Setup
	ss := NewSessionStore()
	s := ss.NewSession()
	s.SetVar("foo", "bar")

	// Test
	val, ok := s.Values["foo"]

	if !ok {
		t.Error("Key not created")
	}

	if val != "bar" {
		t.Errorf("Saved value does not match: expected %v, got '%v'", "'bar'", val)
	}
}

// DestroyVar should remove a session variable
func TestDestroyVar(t *testing.T) {
	// Setup
	ss := NewSessionStore()
	s := ss.NewSession()
	s.SetVar("foo", "bar")

	if s.GetVar("foo") != "bar" {
		t.Error("Test Setup Failed: Unable to set/get session variable")
	}

	// Test
	s.DestroyVar("foo")
	if s.GetVar("foo") != nil {
		t.Error("Key not deleted")
	}
}
