// Copyright © 2015 Gerald Rosenberg.
// Use of this source code is governed by a BSD-style
// license that can be found in the License.md file.
//
package strings

import (
	"sort"

	. "github.com/grosenberg/collections/list"
)

/////////////////////////////////////////////////////////////
// ArrayList type-specific implementations

// Create an ArrayList.
func NewArrayList() *ArrayList {

	comp := func(a, b string) int {
		if a > b {
			return 1
		} else if a < b {
			return -1
		}
		return 0
	}

	return &ArrayList{
		Data: make([]string, 0),
		comp: comp,
	}
}

// Add appends the given values to the receiver
func (a *ArrayList) Add(val ...string) {
	a.Lock()
	defer a.Unlock()

	a.Add_L(-1, val)
}

// Remove deletes the given values from the receiver
func (a *ArrayList) Remove(val ...string) {
	a.Lock()
	defer a.Unlock()

	RemoveByV_L(a, val)
}

// Insert inserts the given values at the given position within the receiver
func (a *ArrayList) Insert(pos int, val ...string) {
	a.Lock()
	defer a.Unlock()

	a.Add_L(pos, val)
}

// Delete removes the values at the given positions within the receiver
func (a *ArrayList) Delete(pos ...int) {
	a.Lock()
	defer a.Unlock()

	a.Remove_L(pos)
}

// Get gets the value at the given position within the receiver
func (a *ArrayList) Get(pos int) string {
	a.Lock()
	defer a.Unlock()

	return a.Get_L(pos).(string)
}

// Set sets the value at the given position within the receiver. The
// prior value at the given position is returned.
func (a *ArrayList) Set(pos int, val string) string {
	a.Lock()
	defer a.Unlock()

	return a.Set_L(pos, val).(string)
}

// Contains returns true iff the receiver contains the given value
func (a *ArrayList) Contains(str string) bool {
	a.Lock()
	defer a.Unlock()

	return Contains_L(a, str)
}

// IndexOf returns the index of the first instance of the given value
// within the receiver or -1 if the value is not present.
func (a *ArrayList) IndexOf(str string) int {
	a.Lock()
	defer a.Unlock()

	return KeyOf_L(a, str).(int)
}

// Iterator returns a slice containing a snapshot copy of the
// values of the receiver
func (a *ArrayList) Iterator() []string {
	a.Lock()
	defer a.Unlock()

	return Iterator_L(a).([]string)
}

// Sort sorts the internal ordering of the receiver. The sort order can be
// changed by providing a suitable comparison function. If the parameter is
// nil, the most recently set comparison function is used. The default sort
// order is alphanumeric ascending.
func (a *ArrayList) Sort(comp func(a, b string) int) {
	a.Lock()
	defer a.Unlock()

	if comp != nil {
		a.comp = comp
	}
	if a.comp != nil { // should not happen
		sort.Sort(a)
	}
}

// Clear removes all elements from the receiver
func (a *ArrayList) Clear() {
	a.Lock()
	defer a.Unlock()

	Clear_L(a)
}

// Size returns the number of elements in the receiver
func (a *ArrayList) Size() int {
	return Len_L(a)
}

// Empty returns true iff the receiver containes no elements
func (a *ArrayList) Empty() bool {
	return Len_L(a) == 0
}

// Clone returns a shallow copy of the receiver
func (a *ArrayList) Clone() *ArrayList {
	a.Lock()
	defer a.Unlock()

	return Copy_L(a).(*ArrayList)
}
