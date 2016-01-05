// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"syscall"
	"unsafe"
)

import "C"

var (
	proc_qtdrv_addr uintptr = 0
)

//export qtdrv
func qtdrv(p unsafe.Pointer, typeid, funcid C.int, p1, p2, p3, p4, p5, p6, p7, p8, p9, p10, p11, p12 unsafe.Pointer) C.int {
	r0, _, _ := syscall.Syscall15(proc_qtdrv_addr, 15, uintptr(p), uintptr(typeid), uintptr(funcid), uintptr(p1), uintptr(p2), uintptr(p3), uintptr(p4), uintptr(p5), uintptr(p6), uintptr(p7), uintptr(p8), uintptr(p9), uintptr(p10), uintptr(p11), uintptr(p12))
	return C.int(r0)
}

func load_qtdrv_proc() (uintptr, error) {
	lib, err := syscall.LoadDLL("qtdrv.ui.dll")
	if err != nil {
		return 0, err
	}
	proc, err := lib.FindProc("qtdrv")
	if err != nil {
		return 0, err
	}
	return proc.Addr(), nil
}

func init() {
	proc_qtdrv_addr, qtdrv_init_error = load_qtdrv_proc()
	if qtdrv_init_error != nil {
		return
	}
	cdrv_init()
}
