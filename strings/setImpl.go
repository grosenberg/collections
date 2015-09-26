// Copyright © 2015 Gerald Rosenberg.
// Use of this source code is governed by a BSD-style
// license that can be found in the License.md file.
//
package strings

import (
	"sort"
	"sync"

	. "github.com/grosenberg/collections/list"
	"github.com/grosenberg/generic"
)

/////////////////////////////////////////////////////////////
// Data type specific type

type Set struct {
	sync.Mutex
	Data    []string // data store
	keyType int      // key type
	valType string   // value type
}

/////////////////////////////////////////////////////////////
// Compile-time implementation prover
var _ L = &Set{}

/////////////////////////////////////////////////////////////
// Type-specific support for generic implementation

func (s *Set) New_L() L {
	return NewSet()
}

func (s *Set) TypeK_L() K {
	return s.keyType
}

func (s *Set) TypeV_L() V {
	return s.valType
}

// For a slice, a negative pos k means append at end
func (s *Set) Add_L(k K, v V) {
	i := k.(int)
	if i >= 0 && i <= len(s.Data)-1 {
		tmp := make([]string, 0)
		tmp = generic.Append(tmp, v).Interface().([]string)
		tmp = append(tmp, s.Data[i:]...)
		s.Data = append(s.Data[0:i], tmp...)
	} else {
		s.Data = generic.Append(s.Data, v).Interface().([]string)
	}
}

func (s *Set) Remove_L(k K) {
	keys := k.([]int)
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	for _, key := range keys {
		if key >= 0 && key < len(s.Data) {
			s.Data, s.Data[len(s.Data)-1] = append(s.Data[:key], s.Data[key+1:]...), ""
		}
	}
}

func (s *Set) Keys_L() K {
	ks := make([]int, 0)
	for k, _ := range s.Data {
		ks = append(ks, k)
	}
	return ks
}

func (s *Set) Values_L() V {
	return s.Data
}

func (s *Set) Get_L(k K) V {
	key := k.(int)
	return s.Data[key]
}

func (s *Set) Set_L(k K, v V) V {
	key := k.(int)
	ret := s.Data[key] // existing value
	s.Data[key] = v.(string)
	return ret
}
