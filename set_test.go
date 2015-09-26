// Copyright © 2015 Gerald Rosenberg.
// Use of this source code is governed by a BSD-style
// license that can be found in the License.md file.
//
package collections

import (
	"fmt"
	"testing"

	"github.com/grosenberg/collections/strings"
)

const (
	val1 = "First"
	val2 = "Second"
	val3 = "Third"
	val4 = "Forth"
)

var set *strings.Set

func TestEmpty(t *testing.T) {
	fmt.Println("Set Empty")

	set = strings.NewSet()

	if !set.Empty() {
		t.Errorf("Empty: should be empty\n")
	}

	set.Add(val1)

	if set.Empty() {
		t.Errorf("Empty: should not be empty\n")
	}
}

func TestSize(t *testing.T) {
	fmt.Println("Set Size")

	set = strings.NewSet()
	set.Add(val1)
	set.Add(val2)

	if set.Size() != 2 {
		t.Errorf("Size failed: wrong size %v != 2\n", set.Size())
	}
}

func TestClear(t *testing.T) {
	fmt.Println("Set Clear")

	set = strings.NewSet()
	set.Add(val1)
	set.Add(val2)
	set.Clear()
	if !set.Empty() {
		t.Errorf("Clear failed: Should be empty\n")
	}
}
func TestClone(t *testing.T) {
	fmt.Println("Set Clone")

	set = strings.NewSet()
	set.Add(val1)
	set.Add(val2)
	set1 := set.Clone()
	if set1.Size() != 2 {
		t.Errorf("Clone failed: size %v != 2\n", set1.Size())
	}
}

func TestIterator(t *testing.T) {
	fmt.Println("Set Iterator")

	set = strings.NewSet()
	set.Add(val1, val2, val3)
	if set.Size() != 3 {
		t.Errorf("Wrong size %v != 3\n", set.Size())
	}

	found := 0
	expected := []string{val1, val2, val3}
	for _, v := range set.Iterator() {
		for _, x := range expected {
			if x == v {
				found++
				break
			}
		}
	}
	if found != 3 {
		t.Errorf("Iterator failed: %v != %v", set.Iterator(), expected)
	}
}

func TestContains(t *testing.T) {
	fmt.Println("Set Contains")

	set = strings.NewSet()
	set.Add(val1, val2, val3)
	if set.Size() != 3 {
		t.Errorf("Wrong size %v != 3\n", set.Size())
	}

	expected := []string{val1, val2, val3}
	for _, v := range set.Iterator() {
		if !set.Contains(v) {
			t.Errorf("Contains failed: %v not in %v", v, expected)
		}
	}

	v := val4
	if set.Contains(v) {
		t.Errorf("Contains failed: %v found in %v", v, expected)
	}
}

func TestRemove(t *testing.T) {
	fmt.Println("Set Remove")

	set = strings.NewSet()
	set.Add(val1, val2, val3)
	set.Remove(val4) // ignores this correctly
	if set.Size() != 3 {
		t.Errorf("Wrong size %v != 3\n", set.Size())
	}

	v := val2
	set.Remove(v)
	if set.Size() != 2 {
		t.Errorf("Wrong size %v != 2\n", set.Size())
	}
	if set.Contains(v) {
		t.Errorf("Remove failed: still contains %v", v)
	}

	set.Remove(val1, val3)
	if !set.Empty() {
		t.Errorf("Empty: should be empty\n")
	}
}
