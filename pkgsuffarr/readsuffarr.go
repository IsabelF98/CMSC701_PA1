// This code comes directly from the GO suffix array package
// Modifications were made to access suffix array easier

// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package suffixarray implements substring search in logarithmic time using
// an in-memory suffix array.

package pkgsuffarr

import (
	"encoding/binary"
	"errors"
	"io"
	"math"
)

// Can change for testing
var maxData32 int = realMaxData32

const realMaxData32 = math.MaxInt32

type MyIndex struct {
	Data []byte
	Sa   Ints // suffix array for data; sa.len() == len(data)
}

// An ints is either an []Int32 or an []int64.
// That is, one of them is empty, and one is the real data.
// The int64 form is used when len(data) > maxData32
type Ints struct {
	Int32 []int32
	int64 []int64
}

func (a *Ints) len() int {
	return len(a.Int32) + len(a.int64)
}

func (a *Ints) set(i int, v int64) {
	if a.Int32 != nil {
		a.Int32[i] = int32(v)
	} else {
		a.int64[i] = v
	}
}

func (a *Ints) slice(i, j int) Ints {
	if a.Int32 != nil {
		return Ints{a.Int32[i:j], nil}
	}
	return Ints{nil, a.int64[i:j]}
}

var errTooBig = errors.New("suffixarray: data too large")

// readInt reads an int x from r using buf to buffer the read and returns x.
func readInt(r io.Reader, buf []byte) (int64, error) {
	_, err := io.ReadFull(r, buf[0:binary.MaxVarintLen64]) // ok to continue with error
	x, _ := binary.Varint(buf)
	return x, err
}

// readSlice reads data[:n] from r and returns n.
// It uses buf to buffer the read.
func readSlice(r io.Reader, buf []byte, data Ints) (n int, err error) {
	// read buffer size
	var size64 int64
	size64, err = readInt(r, buf)
	if err != nil {
		return
	}
	if int64(int(size64)) != size64 || int(size64) < 0 {
		// We never write chunks this big anyway.
		return 0, errTooBig
	}
	size := int(size64)

	// read buffer w/o the size
	if _, err = io.ReadFull(r, buf[binary.MaxVarintLen64:size]); err != nil {
		return
	}

	// decode as many elements as present in buf
	for p := binary.MaxVarintLen64; p < size; n++ {
		x, w := binary.Uvarint(buf[p:])
		data.set(n, int64(x))
		p += w
	}

	return
}

const bufSize = 16 << 10 // reasonable for BenchmarkSaveRestore

// Read reads the index from r into x; x must not be nil.
func (x *MyIndex) Read(r io.Reader) error {
	// buffer for all reads
	buf := make([]byte, bufSize)

	// read length
	n64, err := readInt(r, buf)
	if err != nil {
		return err
	}
	if int64(int(n64)) != n64 || int(n64) < 0 {
		return errTooBig
	}
	n := int(n64)

	// allocate space
	if 2*n < cap(x.Data) || cap(x.Data) < n || x.Sa.Int32 != nil && n > maxData32 || x.Sa.int64 != nil && n <= maxData32 {
		// new data is significantly smaller or larger than
		// existing buffers - allocate new ones
		x.Data = make([]byte, n)
		x.Sa.Int32 = nil
		x.Sa.int64 = nil
		if n <= maxData32 {
			x.Sa.Int32 = make([]int32, n)
		} else {
			x.Sa.int64 = make([]int64, n)
		}
	} else {
		// re-use existing buffers
		x.Data = x.Data[0:n]
		x.Sa = x.Sa.slice(0, n)
	}

	// read data
	if _, err := io.ReadFull(r, x.Data); err != nil {
		return err
	}

	// read index
	sa := x.Sa
	for sa.len() > 0 {
		n, err := readSlice(r, buf, sa)
		if err != nil {
			return err
		}
		sa = sa.slice(n, sa.len())
	}
	return nil
}
