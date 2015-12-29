// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QTextBlockUserData : QTextBlockUserData
type QTextBlockUserData struct {
	BaseDrv
}
//QTextBlockUserData::QTextBlockUserData()	
func NewTextBlockUserData() *QTextBlockUserData {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),140000,140102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTextBlockUserData{}
	_p.SetDriver(__rv,140000,true)
	return _p
} 
