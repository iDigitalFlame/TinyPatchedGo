// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// 2022 - Shrinkage by iDigitalFlame

package runtime

import "unsafe"

type hex uint64

func printsp() {}
func printnl() {}
func printlock() {}
func printunlock() {}
func gwrite(_ []byte) {}
func printbool(_ bool) {}
func printint(_ int64) {}
func printhex(_ uint64) {}
func printuint(_ uint64) {}
func printeface(_ eface) {}
func printiface(_ iface) {}
func printslice(_ []byte) {}
func printfloat(_ float64) {}
func printstring(_ string) {}
func printuintptr(_ uintptr) {}
func printcomplex(_ complex128) {}
func printpointer(_ unsafe.Pointer) {}
func hexdumpWords(_, _ uintptr, _ func(uintptr) byte) {}
