// Copyright © 2015 Gerald Rosenberg.
// Use of this source code is governed by a BSD-style
// license that can be found in the License.md file.
//
package strings

// sort interface implementation

func (a *ArrayList) Len() int {
	return len(a.Data)
}

func (a *ArrayList) Swap(i, j int) {
	a.Data[i], a.Data[j] = a.Data[j], a.Data[i]
}

func (a *ArrayList) Less(i, j int) bool {
	return a.comp(a.Data[i], a.Data[j]) < 0
}
