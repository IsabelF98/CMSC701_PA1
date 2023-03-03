// This code comes directly from the GO suffix array package
// Modifications were made to access suffix array easier

// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package suffixarray implements substring search in logarithmic time using
// an in-memory suffix array.

package pkgsuffarr

import (
	"bytes"
	"sort"
)

func (a *Ints) get(i int) int64 {
	if a.Int32 != nil {
		return int64(a.Int32[i])
	}
	return a.int64[i]
}

func (x *MyIndex) at(i int) []byte {
	return x.Data[x.Sa.get(i):]
}

// lookupAll returns a slice into the matching region of the index.
// The runtime is O(log(N)*len(s)).
func (x *MyIndex) lookupAll(s []byte) Ints {
	// find matching suffix index range [i:j]
	// find the first index where s would be the prefix
	i := sort.Search(x.Sa.len(), func(i int) bool { return bytes.Compare(x.at(i), s) >= 0 })
	// starting at i, find the first index at which s is not a prefix
	j := i + sort.Search(x.Sa.len()-i, func(j int) bool { return !bytes.HasPrefix(x.at(j+i), s) })
	return x.Sa.slice(i, j)
}

// Lookup returns an unsorted list of at most n indices where the byte string s
// occurs in the indexed data. If n < 0, all occurrences are returned.
// The result is nil if s is empty, s is not found, or n == 0.
// Lookup time is O(log(N)*len(s) + len(result)) where N is the
// size of the indexed data.
func (x *MyIndex) Lookup(s []byte, n int) (result []int) {
	if len(s) > 0 && n != 0 {
		matches := x.lookupAll(s)
		count := matches.len()
		if n < 0 || count < n {
			n = count
		}
		// 0 <= n <= count
		if n > 0 {
			result = make([]int, n)
			if matches.Int32 != nil {
				for i := range result {
					result[i] = int(matches.Int32[i])
				}
			} else {
				for i := range result {
					result[i] = int(matches.int64[i])
				}
			}
		}
	}
	return
}
