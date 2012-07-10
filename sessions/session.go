// Copyright 2012 Matt West and Contributors. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package sessions

type Session struct {
	ID      string
	Values  map[string]interface{}
	Secret  string
	Expires int // Unix Timestamp
}

// Retrieves a Session Variable
func (s Session) GetVar(key string) (v interface{}) {
	v = s.Values[key]
	return
}

// Sets a Session Variable
func (s Session) SetVar(key string, value interface{}) {
	s.Values[key] = value
	return
}

// Destroys a Session Variable
func (s Session) DestroyVar(key string) {
	delete(s.Values, key)
	return
}
