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
	arg1 = "First"
	arg2 = "Second"
	arg3 = "Third"
	arg4 = "Forth"
)

var list *strings.ArrayList

func TestListEmpty(t *testing.T) {
	fmt.Println("ArrayList Empty")

	list = strings.NewArrayList()

	if !list.Empty() {
		t.Errorf("Empty: should be empty\n")
	}

	list.Add(arg1)

	if list.Empty() {
		t.Errorf("Empty: should not be empty\n")
	}
}

func TestListSize(t *testing.T) {
	fmt.Println("ArrayList Size")

	list = strings.NewArrayList()
	list.Add(arg1)
	list.Add(arg2)

	if list.Size() != 2 {
		t.Errorf("Size failed: wrong size %v != 2\n", list.Size())
	}
}

func TestListClear(t *testing.T) {
	fmt.Println("ArrayList Clear")

	list = strings.NewArrayList()
	list.Add(arg1)
	list.Add(arg2)
	list.Clear()
	if !list.Empty() {
		t.Errorf("Clear failed: Should be empty\n")
	}
}
func TestListClone(t *testing.T) {
	fmt.Println("ArrayList Clone")

	list = strings.NewArrayList()
	list.Add(arg1)
	list.Add(arg2)
	set1 := list.Clone()
	if set1.Size() != 2 {
		t.Errorf("Clone failed: size %v != 2\n", set1.Size())
	}
}

func TestListIterator(t *testing.T) {
	fmt.Println("ArrayList Iterator")

	list = strings.NewArrayList()
	list.Add(arg1, arg2, arg3)
	if list.Size() != 3 {
		t.Errorf("Wrong size %v != 3\n", list.Size())
	}

	found := 0
	expected := []string{arg1, arg2, arg3}
	for _, v := range list.Iterator() {
		for _, x := range expected {
			if x == v {
				found++
				break
			}
		}
	}
	if found != 3 {
		t.Errorf("Iterator failed: %v != %v", list.Iterator(), expected)
	}
}

func TestListContains(t *testing.T) {
	fmt.Println("ArrayList Contains")

	list = strings.NewArrayList()
	list.Add(arg1, arg2, arg3)
	if list.Size() != 3 {
		t.Errorf("Wrong size %v != 3\n", list.Size())
	}

	expected := []string{arg1, arg2, arg3}
	for _, v := range list.Iterator() {
		if !list.Contains(v) {
			t.Errorf("Contains failed: %v not in %v", v, expected)
		}
	}

	v := arg4
	if list.Contains(v) {
		t.Errorf("Contains failed: %v found in %v", v, expected)
	}
}

func TestListRemove(t *testing.T) {
	fmt.Println("ArrayList Remove")

	list = strings.NewArrayList()
	list.Add(arg1, arg2, arg3)
	list.Remove(arg4) // ignores this correctly
	if list.Size() != 3 {
		t.Errorf("Wrong size %v != 3\n", list.Size())
	}

	v := arg2
	list.Remove(v)
	if list.Size() != 2 {
		t.Errorf("Wrong size %v != 2\n", list.Size())
	}
	if list.Contains(v) {
		t.Errorf("Remove failed: still contains %v", v)
	}

	list.Remove(arg1, arg3)
	if !list.Empty() {
		t.Errorf("Empty: should be empty\n")
	}
}
