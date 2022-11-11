// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// 2022 - Shrinkage by iDigitalFlame

package runtime

import "unsafe"

type hex uint64

func gwrite(b []byte) {
	if len(b) == 0 {
		return
	}
	g := getg()
	if g == nil || g.writebuf == nil || g.m.dying > 0 {
		writeErr(b)
		return
	}
	n := copy(g.writebuf[len(g.writebuf):cap(g.writebuf)], b)
	g.writebuf = g.writebuf[:len(g.writebuf)+n]
}
func printstring(s string) {
	var (
		b []byte
		v = stringStructOf(&s)
		r = (*slice)(unsafe.Pointer(&b))
	)
	r.len, r.cap, r.array = v.len, v.len, v.str
	gwrite(b)
}

func printsp()                                            {}
func printnl()                                            {}
func printlock()                                          {}
func printunlock()                                        {}
func printbool(_ bool)                                    {}
func printint(_ int64)                                    {}
func printhex(_ uint64)                                   {}
func printuint(_ uint64)                                  {}
func printeface(_ eface)                                  {}
func printiface(_ iface)                                  {}
func printslice(_ []byte)                                 {}
func printfloat(_ float64)                                {}
func printuintptr(_ uintptr)                              {}
func printcomplex(_ complex128)                           {}
func printpointer(_ unsafe.Pointer)                       {}
func hexdumpWords(_, _ uintptr, _type func(uintptr) byte) {}
