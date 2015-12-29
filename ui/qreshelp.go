// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QResHelp : QResHelp
type QResHelp struct {
	BaseDrv
}
//QResHelp::registerResourceData(int,uchar const*,uchar const*,uchar const*)	
func QResHelpRegisterResourceData(version int,tree *byte,name *byte,data *byte) bool {
	var __rv bool
	DirectQtDrv(nil,114000,114102,unsafe.Pointer(&version),unsafe.Pointer(&tree),unsafe.Pointer(&name),unsafe.Pointer(&data),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QResHelp::registerResourceData(int,uchar const*,uchar const*,uchar const*)
func (q *QResHelp) RegisterResourceData(version int,tree *byte,name *byte,data *byte) bool {
	var __rv bool
	q.Drv(114000,114102,unsafe.Pointer(&version),unsafe.Pointer(&tree),unsafe.Pointer(&name),unsafe.Pointer(&data),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QResHelp::unregisterResourceData(int,uchar const*,uchar const*,uchar const*)	
func QResHelpUnregisterResourceData(version int,tree *byte,name *byte,data *byte) bool {
	var __rv bool
	DirectQtDrv(nil,114000,114103,unsafe.Pointer(&version),unsafe.Pointer(&tree),unsafe.Pointer(&name),unsafe.Pointer(&data),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QResHelp::unregisterResourceData(int,uchar const*,uchar const*,uchar const*)
func (q *QResHelp) UnregisterResourceData(version int,tree *byte,name *byte,data *byte) bool {
	var __rv bool
	q.Drv(114000,114103,unsafe.Pointer(&version),unsafe.Pointer(&tree),unsafe.Pointer(&name),unsafe.Pointer(&data),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
