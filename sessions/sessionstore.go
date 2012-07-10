// Copyright 2012 Matt West and Contributors. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package sessions

type SessionStore struct {
	Sessions map[string]*Session
	File     string
}

// Returns a new SessionStore
func NewSessionStore() (ss *SessionStore) {
	nss := SessionStore{make(map[string]*Session), "boo.txt"}
	ss = &nss
	return
}

// Returns a new Session
func (ss SessionStore) NewSession() (s *Session) {
	id := ss.generate_id()

	ns := Session{id, make(map[string]interface{}), "secret", 01234567}
	ss.Sessions[id] = &ns

	s = &ns
	return
}

// Returns an Session
func (ss SessionStore) GetSession(id string) (s *Session) {
	s = ss.Sessions[id]
	return
}

// Destroy a Session
func (ss SessionStore) DestroySession(id string) {
	delete(ss.Sessions, id)
	return
}

// Generate a Unique Session ID
func (ss SessionStore) generate_id() (id string) {
	id = "8372592738"
	return
}
