// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QDragMoveEvent : QDragMoveEvent
type QDragMoveEvent struct {
	QDropEvent
}
//QDragMoveEvent::accept()
func (q *QDragMoveEvent) Accept()  {
	q.Drv(26000,26102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDragMoveEvent::accept(QRect const&)
func (q *QDragMoveEvent) AcceptWithRect(r *QRect)  {
	q.Drv(26000,26103,Native(r),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDragMoveEvent::answerRect()
func (q *QDragMoveEvent) AnswerRect() *QRect {
	var __rv uintptr
	q.Drv(26000,26104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QDragMoveEvent::ignore()
func (q *QDragMoveEvent) Ignore()  {
	q.Drv(26000,26105,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDragMoveEvent::ignore(QRect const&)
func (q *QDragMoveEvent) IgnoreWithRect(r *QRect)  {
	q.Drv(26000,26106,Native(r),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
