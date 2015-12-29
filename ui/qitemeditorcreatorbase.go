// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QItemEditorCreatorBase : QItemEditorCreatorBase
type QItemEditorCreatorBase struct {
	BaseDrv
}
//QItemEditorCreatorBase::createWidget(QWidget*)
func (q *QItemEditorCreatorBase) CreateWidget(parent QWidgetInterface) *QWidget {
	var __rv uintptr
	q.Drv(60000,60102,Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QItemEditorCreatorBase::valuePropertyName()
func (q *QItemEditorCreatorBase) ValuePropertyName() []byte {
	var __rv []byte
	q.Drv(60000,60103,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
