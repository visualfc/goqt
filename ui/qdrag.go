// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QDrag : QDrag
type QDrag struct {
	QObject
}
func (q *QDrag) OnTargetChanged(fn func(*QWidget)) uintptr {
	var __rv uintptr
	q.Drv(232000,232102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QDrag) OnActionChanged(fn func(Qt_DropAction)) uintptr {
	var __rv uintptr
	q.Drv(232000,232103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QDrag::QDrag(QWidget*)	
func NewDrag(dragSource QWidgetInterface) *QDrag {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),232000,232104,Native(dragSource),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QDrag{}
	_p.SetDriver(__rv,232000,false)
	return _p
} 
//QDrag::exec()
func (q *QDrag) Exec() Qt_DropAction {
	var __rv Qt_DropAction
	q.Drv(232000,232105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDrag::exec(QFlags<Qt::DropAction>)
func (q *QDrag) ExecWithSupportedactions(supportedActions Qt_DropAction) Qt_DropAction {
	var __rv Qt_DropAction
	q.Drv(232000,232106,unsafe.Pointer(&supportedActions),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDrag::exec(QFlags<Qt::DropAction>,Qt::DropAction)
func (q *QDrag) ExecWithSupportedactionsDefaultaction(supportedActions Qt_DropAction,defaultAction Qt_DropAction) Qt_DropAction {
	var __rv Qt_DropAction
	q.Drv(232000,232107,unsafe.Pointer(&supportedActions),unsafe.Pointer(&defaultAction),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDrag::hotSpot()
func (q *QDrag) HotSpot() *QPoint {
	var __rv uintptr
	q.Drv(232000,232108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPoint{}
	_rp.SetDriver(__rv,99000,true)
	return _rp
}	
//QDrag::mimeData()
func (q *QDrag) MimeData() *QMimeData {
	var __rv uintptr
	q.Drv(232000,232109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QMimeData{}
	_rp.SetDriver(__rv,311000,false)
	return _rp
}	
//QDrag::pixmap()
func (q *QDrag) Pixmap() *QPixmap {
	var __rv uintptr
	q.Drv(232000,232110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPixmap{}
	_rp.SetDriver(__rv,96000,true)
	return _rp
}	
//QDrag::setDragCursor(QPixmap const&,Qt::DropAction)
func (q *QDrag) SetDragCursor(cursor *QPixmap,action Qt_DropAction)  {
	q.Drv(232000,232111,Native(cursor),unsafe.Pointer(&action),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDrag::setHotSpot(QPoint const&)
func (q *QDrag) SetHotSpot(hotspot *QPoint)  {
	q.Drv(232000,232112,Native(hotspot),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDrag::setMimeData(QMimeData*)
func (q *QDrag) SetMimeData(data *QMimeData)  {
	q.Drv(232000,232113,Native(data),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDrag::setPixmap(QPixmap const&)
func (q *QDrag) SetPixmap(value *QPixmap)  {
	q.Drv(232000,232114,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDrag::source()
func (q *QDrag) Source() *QWidget {
	var __rv uintptr
	q.Drv(232000,232115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QDrag::start()
func (q *QDrag) Start() Qt_DropAction {
	var __rv Qt_DropAction
	q.Drv(232000,232116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDrag::start(QFlags<Qt::DropAction>)
func (q *QDrag) StartWithSupportedactions(supportedActions Qt_DropAction) Qt_DropAction {
	var __rv Qt_DropAction
	q.Drv(232000,232117,unsafe.Pointer(&supportedActions),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDrag::target()
func (q *QDrag) Target() *QWidget {
	var __rv uintptr
	q.Drv(232000,232118,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
