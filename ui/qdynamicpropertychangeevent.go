// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QDynamicPropertyChangeEvent : QDynamicPropertyChangeEvent
type QDynamicPropertyChangeEvent struct {
	QEvent
}
//QDynamicPropertyChangeEvent::QDynamicPropertyChangeEvent(QByteArray const&)	
func NewDynamicPropertyChangeEvent(name []byte) *QDynamicPropertyChangeEvent {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),28000,28102,unsafe.Pointer(&name),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QDynamicPropertyChangeEvent{}
	_p.SetDriver(__rv,28000,true)
	return _p
} 
//QDynamicPropertyChangeEvent::propertyName()
func (q *QDynamicPropertyChangeEvent) PropertyName() []byte {
	var __rv []byte
	q.Drv(28000,28103,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
