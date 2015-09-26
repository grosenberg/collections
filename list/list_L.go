// Copyright © 2015 Gerald Rosenberg.
// Use of this source code is governed by a BSD-style
// license that can be found in the License.md file.
//
package list

import (
	"reflect"
	"sync"

	"github.com/grosenberg/generic"
)

/////////////////////////////////////////////////////////////
// Generic interfaces

type L interface {
	sync.Locker

	New_L() L
	TypeK_L() K
	TypeV_L() V

	Add_L(K, V)
	Remove_L(K)

	Keys_L() K
	Values_L() V
	Get_L(K) V
	Set_L(K, V) V
}

type K interface{}
type V interface{}

/////////////////////////////////////////////////////////////
// Common generic implementations

func Len_L(list L) int {
	return generic.ValueOf(list.Values_L()).Len()
}

func Clear_L(list L) {
	keys := list.Keys_L()
	list.Remove_L(keys)
}

func Copy_L(list L) L {
	c := list.New_L()
	c.Add_L(-1, list.Values_L())
	return c
}

func Contains_L(list L, val V) bool {
	ks := generic.ValueOf(list.Keys_L())
	for i := 0; i < ks.Len(); i++ {
		key := ks.Index(i).Interface()
		data := list.Get_L(key)
		if val == data {
			return true
		}
	}
	return false
}

func KeyOf_L(list L, val V) K {
	ks := generic.ValueOf(list.Keys_L())
	for i := 0; i < ks.Len(); i++ {
		key := ks.Index(i).Interface()
		data := list.Get_L(key)
		if val == data {
			return key
		}
	}

	// need to think about this
	switch generic.TypeOf(list.TypeK_L()).Kind() {
	case reflect.Int, reflect.Int64:
		return -1
	case reflect.String:
		return ""
	default:
		return nil
	}
}

func KeySet_L(list L) K {
	k := generic.MakeSlice(list.Keys_L()).Interface()
	return generic.Append(k, list.Keys_L()).Interface()
}

func Iterator_L(list L) V {
	v := generic.MakeSlice(list.Values_L()).Interface()
	return generic.Append(v, list.Values_L()).Interface()
}

func RemoveByK_L(list L, k K) {
	list.Remove_L(k)
}

func RemoveByV_L(list L, v V) {
	keys := generic.MakeSlice(list.TypeK_L()).Interface()
	vs := generic.ValueOf(v)
	for i := 0; i < vs.Len(); i++ {
		val := vs.Index(i).Interface()
		key := KeyOf_L(list, val)
		if key != nil {
			keys = generic.Append(keys, key).Interface()
		}
	}
	list.Remove_L(keys)
}
