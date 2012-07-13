// Copyright 2012 Matt West and Contributors. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package sessions

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"reflect"
	"time"
)

type SessionStore struct {
	Sessions map[string]*Session
	File     string
}

// Returns a new SessionStore
func NewSessionStore() (ss *SessionStore) {
	nss := SessionStore{make(map[string]*Session), "boo.txt"} // TODO
	ss = &nss
	return
}

// Returns a new Session
func (ss *SessionStore) NewSession(w http.ResponseWriter) (s *Session) {
	id, _ := ss.generate_id()

	ns := Session{id, make(map[string]interface{}), "secret", time.Now()} // TODO
	ss.Sessions[id] = &ns

	// Need to Set a Cookie Here that stores the session ID on the client
	// TODO: Put more info in this cookie, at the moment it is just a PoC
	cookie := http.Cookie{
		Name:  "sid",
		Value: id,
	}

	http.SetCookie(w, &cookie)

	s = &ns
	return
}

// Returns an Session
func (ss *SessionStore) GetSession(w http.ResponseWriter, r *http.Request) (s *Session) {
	// Try and Retrieve a Session First
	cookie, err := r.Cookie("sid")

	if err == nil {
		// Check that the Cookie can be found
		if reflect.TypeOf(cookie) != reflect.TypeOf(http.ErrNoCookie) {
			// Check that the session still exists on the server side
			_, ok := ss.Sessions[cookie.Value]

			if ok {
				s = ss.Sessions[cookie.Value]
			}
		}
	}

	// If no Session, create a new one.
	if s == nil {
		s = ss.NewSession(w)
	}

	return
}

// Destroy a Session
func (ss *SessionStore) DestroySession(id string) {
	delete(ss.Sessions, id)
	return
}

// Generate a Unique Session ID
func (ss SessionStore) generate_id() (string, error) {
	// Following code from: http://www.ashishbanerjee.com/home/go/go-generate-uuid
	uuid := make([]byte, 16)
	n, err := rand.Read(uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}

	// TODO: verify the two lines implement RFC 4122 correctly
	uuid[8] = 0x80 // variant bits see page 5
	uuid[4] = 0x40 // version 4 Pseudo Random, see page 7

	return hex.EncodeToString(uuid), nil
}

// TODO: Save the SessionStore to a File
func Save(filename string) error {
	return nil
}

// TODO: Load a SessionStore from a File
func Load(filename string) error {
	return nil
}
