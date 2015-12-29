// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QInputMethodEventAttribute : QInputMethodEvent::Attribute
type QInputMethodEventAttribute struct {
	BaseDrv
}
//QInputMethodEvent::Attribute::Attribute(QInputMethodEvent::AttributeType,int,int,QVariant)	
func NewInputMethodEventAttribute(t QInputMethodEvent_AttributeType,s int,l int,val *QVariant) *QInputMethodEventAttribute {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),59000,59102,unsafe.Pointer(&t),unsafe.Pointer(&s),unsafe.Pointer(&l),Native(val),nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QInputMethodEventAttribute{}
	_p.SetDriver(__rv,59000,true)
	return _p
} 
