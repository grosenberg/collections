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

func TestListInsert(t *testing.T) {
	fmt.Println("ArrayList Insert")

	list = strings.NewArrayList()
	list.Add(arg1, arg2, arg3)

	v := arg4
	list.Insert(0, v)
	if list.Size() != 4 {
		t.Errorf("Wrong size %v != 4\n", list.Size())
	}

	r := list.Get(0)
	if r != v {
		t.Errorf("Insert: failed insert at 0\n")
	}

	v = arg4
	list.Insert(2, v)
	if list.Size() != 5 {
		t.Errorf("Wrong size %v != 5\n", list.Size())
	}

	r = list.Get(2)
	if r != v {
		t.Errorf("Insert: failed insert at 2\n")
	}

	v = arg4
	list.Insert(5, v)
	if list.Size() != 6 {
		t.Errorf("Wrong size %v != 6\n", list.Size())
	}

	r = list.Get(5)
	if r != v {
		t.Errorf("Insert: failed insert at 5\n")
	}

	v = arg2
	list.Insert(99, v)
	if list.Size() != 7 {
		t.Errorf("Wrong size %v != 7\n", list.Size())
	}

	r = list.Get(6)
	if r != v {
		t.Errorf("Insert: failed insert at 99\n")
	}
}

func TestListDelete(t *testing.T) {
	fmt.Println("ArrayList Delete")

	list = strings.NewArrayList()
	list.Add(arg1, arg2, arg3)

	list.Delete(5)
	if list.Size() != 3 {
		t.Errorf("Wrong size %v != 3\n", list.Size())
	}

	list.Delete(1)
	if list.Size() != 2 {
		t.Errorf("Wrong size %v != 2\n", list.Size())
	}

	r := list.Get(1)

	if r == arg2 {
		t.Errorf("Delete failed: still contains %v", r)
	}
}

func TestListSet(t *testing.T) {
	fmt.Println("ArrayList Set")

	list = strings.NewArrayList()
	list.Add(arg1, arg2, arg3)

	v := "NewValue"
	r := list.Set(1, v)
	s := list.Get(1)
	if s != v {
		t.Errorf("Set failed: %v != %v\n", s, v)
	}
	if r != arg2 {
		t.Errorf("Set return failed: %v != %v\n", arg2, r)
	}
}

func TestListIndexOf(t *testing.T) {
	fmt.Println("ArrayList IndexOf")

	list = strings.NewArrayList()
	list.Add(arg1, arg2, arg3)

	i := list.IndexOf(arg2)
	j := list.IndexOf(arg4)

	if i != 1 {
		t.Errorf("Wrong index: %v != 1\n", i)
	}
	if j != -1 {
		t.Errorf("Wrong index: %v != -1\n", j)
	}
}
