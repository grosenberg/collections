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

type ArrayList struct {
	sync.Mutex
	Data    []string              // data store
	keyType int                   // key type
	valType string                // value type
	comp    func(a, b string) int // dynamic sorter function
}

/////////////////////////////////////////////////////////////
// Compile-time implementation prover
var _ L = &ArrayList{}

/////////////////////////////////////////////////////////////
// Type-specific support for generic implementation

func (a *ArrayList) New_L() L {
	return NewArrayList()
}

func (a *ArrayList) TypeK_L() K {
	return a.keyType
}

func (a *ArrayList) TypeV_L() V {
	return a.valType
}

// For a slice, a negative pos k means append at end
func (a *ArrayList) Add_L(k K, v V) {
	i := k.(int)
	if i >= 0 && i <= len(a.Data)-1 {
		tmp := make([]string, 0)
		tmp = generic.Append(tmp, v).Interface().([]string)
		tmp = append(tmp, a.Data[i:]...)
		a.Data = append(a.Data[0:i], tmp...)
	} else {
		a.Data = generic.Append(a.Data, v).Interface().([]string)
	}
}

func (a *ArrayList) Remove_L(k K) {
	keys := k.([]int)
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	for _, key := range keys {
		if key >= 0 && key < len(a.Data) {
			a.Data, a.Data[len(a.Data)-1] = append(a.Data[:key], a.Data[key+1:]...), ""
		}
	}
}

func (a *ArrayList) Keys_L() K {
	ks := make([]int, 0)
	for k, _ := range a.Data {
		ks = append(ks, k)
	}
	return ks
}

func (a *ArrayList) Values_L() V {
	return a.Data
}

func (a *ArrayList) Get_L(k K) V {
	key := k.(int)
	return a.Data[key]
}

func (a *ArrayList) Set_L(k K, v V) V {
	key := k.(int)
	ret := a.Data[key] // existing value
	a.Data[key] = v.(string)
	return ret
}
