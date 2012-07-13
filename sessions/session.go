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
