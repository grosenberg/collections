// Copyright © 2015 Gerald Rosenberg.
// Use of this source code is governed by a BSD-style
// license that can be found in the License.md file.
//
package strings

import . "github.com/grosenberg/collections/list"

/////////////////////////////////////////////////////////////
// Set type-specific implementations

// NewSet creates a new Set.
func NewSet() *Set {
	return &Set{
		Data: make([]string, 0),
	}
}

// Add adds the given values to the receiver
func (s *Set) Add(val ...string) {
	s.Lock()
	defer s.Unlock()

	s.Add_L(-1, val)
}

// Remove removes the given values from the receiver
func (s *Set) Remove(val ...string) {
	s.Lock()
	defer s.Unlock()

	RemoveByV_L(s, val)
}

// Contains returns true iff the receiver contains the given value
func (s *Set) Contains(str string) bool {
	s.Lock()
	defer s.Unlock()

	return Contains_L(s, str)
}

// Iterator returns a slice containing a snapshot copy of the
// values of the receiver
func (s *Set) Iterator() []string {
	s.Lock()
	defer s.Unlock()

	return Iterator_L(s).([]string)
}

// Clear removes all elements from the receiver
func (s *Set) Clear() {
	s.Lock()
	defer s.Unlock()

	Clear_L(s)
}

// Size returns the number of elements in the receiver
func (s *Set) Size() int {
	return Len_L(s)
}

// Empty returns true iff the receiver containes no elements
func (s *Set) Empty() bool {
	return Len_L(s) == 0
}

// Clone returns a shallow copy of the receiver
func (s *Set) Clone() *Set {
	s.Lock()
	defer s.Unlock()

	return Copy_L(s).(*Set)
}
