// Copyright 2018 HORISEN AG. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package maybe

import "fmt"

// Int64 is Maybe monad for int64
type Int64 struct {
	Valid bool  `bson:"valid"`
	Value int64 `bson:"value"`
}

// String returns string representation of maybe nt64
func (i *Int64) String() string {
	if i.Valid {
		return fmt.Sprintf("!%d", i.Value)
	}
	return fmt.Sprintf("?%d", i.Value)
}

// String is Maybe monad for string
type String struct {
	Valid bool   `bson:"valid"`
	Value string `bson:"value"`
}

// String returns string representation of maybe nt64
func (s *String) String() string {
	if s.Valid {
		return fmt.Sprintf("!'%s'", s.Value)
	}
	return fmt.Sprintf("?'%s'", s.Value)
}

// Bool is Maybe monad for bool
type Bool struct {
	Valid bool `bson:"valid"`
	Value bool `bson:"value"`
}

// String returns string representation of maybe nt64
func (b *Bool) String() string {
	if b.Valid {
		return fmt.Sprintf("!%t", b.Value)
	}
	return fmt.Sprintf("?%t", b.Value)
}
