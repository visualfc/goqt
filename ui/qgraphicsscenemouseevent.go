// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QGraphicsSceneMouseEvent : QGraphicsSceneMouseEvent
type QGraphicsSceneMouseEvent struct {
	QGraphicsSceneEvent
}
//QGraphicsSceneMouseEvent::QGraphicsSceneMouseEvent()	
func NewGraphicsSceneMouseEvent() *QGraphicsSceneMouseEvent {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),278000,278102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGraphicsSceneMouseEvent{}
	_p.SetDriver(__rv,278000,true)
	return _p
} 
//QGraphicsSceneMouseEvent::QGraphicsSceneMouseEvent(QEvent::Type)	
func NewGraphicsSceneMouseEventWithType(_type QEvent_Type) *QGraphicsSceneMouseEvent {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),278000,278103,unsafe.Pointer(&_type),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGraphicsSceneMouseEvent{}
	_p.SetDriver(__rv,278000,true)
	return _p
} 
//QGraphicsSceneMouseEvent::button()
func (q *QGraphicsSceneMouseEvent) Button() Qt_MouseButton {
	var __rv Qt_MouseButton
	q.Drv(278000,278104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsSceneMouseEvent::buttonDownPos(Qt::MouseButton)
func (q *QGraphicsSceneMouseEvent) ButtonDownPos(button Qt_MouseButton) *QPointF {
	var __rv uintptr
	q.Drv(278000,278105,unsafe.Pointer(&button),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QGraphicsSceneMouseEvent::buttonDownScenePos(Qt::MouseButton)
func (q *QGraphicsSceneMouseEvent) ButtonDownScenePos(button Qt_MouseButton) *QPointF {
	var __rv uintptr
	q.Drv(278000,278106,unsafe.Pointer(&button),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QGraphicsSceneMouseEvent::buttonDownScreenPos(Qt::MouseButton)
func (q *QGraphicsSceneMouseEvent) ButtonDownScreenPos(button Qt_MouseButton) *QPoint {
	var __rv uintptr
	q.Drv(278000,278107,unsafe.Pointer(&button),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPoint{}
	_rp.SetDriver(__rv,99000,true)
	return _rp
}	
//QGraphicsSceneMouseEvent::buttons()
func (q *QGraphicsSceneMouseEvent) Buttons() Qt_MouseButton {
	var __rv Qt_MouseButton
	q.Drv(278000,278108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsSceneMouseEvent::lastPos()
func (q *QGraphicsSceneMouseEvent) LastPos() *QPointF {
	var __rv uintptr
	q.Drv(278000,278109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QGraphicsSceneMouseEvent::lastScenePos()
func (q *QGraphicsSceneMouseEvent) LastScenePos() *QPointF {
	var __rv uintptr
	q.Drv(278000,278110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QGraphicsSceneMouseEvent::lastScreenPos()
func (q *QGraphicsSceneMouseEvent) LastScreenPos() *QPoint {
	var __rv uintptr
	q.Drv(278000,278111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPoint{}
	_rp.SetDriver(__rv,99000,true)
	return _rp
}	
//QGraphicsSceneMouseEvent::modifiers()
func (q *QGraphicsSceneMouseEvent) Modifiers() Qt_KeyboardModifier {
	var __rv Qt_KeyboardModifier
	q.Drv(278000,278112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsSceneMouseEvent::pos()
func (q *QGraphicsSceneMouseEvent) Pos() *QPointF {
	var __rv uintptr
	q.Drv(278000,278113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QGraphicsSceneMouseEvent::scenePos()
func (q *QGraphicsSceneMouseEvent) ScenePos() *QPointF {
	var __rv uintptr
	q.Drv(278000,278114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QGraphicsSceneMouseEvent::screenPos()
func (q *QGraphicsSceneMouseEvent) ScreenPos() *QPoint {
	var __rv uintptr
	q.Drv(278000,278115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPoint{}
	_rp.SetDriver(__rv,99000,true)
	return _rp
}	
//QGraphicsSceneMouseEvent::setButton(Qt::MouseButton)
func (q *QGraphicsSceneMouseEvent) SetButton(button Qt_MouseButton)  {
	q.Drv(278000,278116,unsafe.Pointer(&button),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsSceneMouseEvent::setButtonDownPos(Qt::MouseButton,QPointF const&)
func (q *QGraphicsSceneMouseEvent) SetButtonDownPos(button Qt_MouseButton,pos *QPointF)  {
	q.Drv(278000,278117,unsafe.Pointer(&button),Native(pos),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsSceneMouseEvent::setButtonDownScenePos(Qt::MouseButton,QPointF const&)
func (q *QGraphicsSceneMouseEvent) SetButtonDownScenePos(button Qt_MouseButton,pos *QPointF)  {
	q.Drv(278000,278118,unsafe.Pointer(&button),Native(pos),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsSceneMouseEvent::setButtonDownScreenPos(Qt::MouseButton,QPoint const&)
func (q *QGraphicsSceneMouseEvent) SetButtonDownScreenPos(button Qt_MouseButton,pos *QPoint)  {
	q.Drv(278000,278119,unsafe.Pointer(&button),Native(pos),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsSceneMouseEvent::setButtons(QFlags<Qt::MouseButton>)
func (q *QGraphicsSceneMouseEvent) SetButtons(buttons Qt_MouseButton)  {
	q.Drv(278000,278120,unsafe.Pointer(&buttons),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsSceneMouseEvent::setLastPos(QPointF const&)
func (q *QGraphicsSceneMouseEvent) SetLastPos(pos *QPointF)  {
	q.Drv(278000,278121,Native(pos),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsSceneMouseEvent::setLastScenePos(QPointF const&)
func (q *QGraphicsSceneMouseEvent) SetLastScenePos(pos *QPointF)  {
	q.Drv(278000,278122,Native(pos),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsSceneMouseEvent::setLastScreenPos(QPoint const&)
func (q *QGraphicsSceneMouseEvent) SetLastScreenPos(pos *QPoint)  {
	q.Drv(278000,278123,Native(pos),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsSceneMouseEvent::setModifiers(QFlags<Qt::KeyboardModifier>)
func (q *QGraphicsSceneMouseEvent) SetModifiers(modifiers Qt_KeyboardModifier)  {
	q.Drv(278000,278124,unsafe.Pointer(&modifiers),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsSceneMouseEvent::setPos(QPointF const&)
func (q *QGraphicsSceneMouseEvent) SetPos(pos *QPointF)  {
	q.Drv(278000,278125,Native(pos),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsSceneMouseEvent::setScenePos(QPointF const&)
func (q *QGraphicsSceneMouseEvent) SetScenePos(pos *QPointF)  {
	q.Drv(278000,278126,Native(pos),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsSceneMouseEvent::setScreenPos(QPoint const&)
func (q *QGraphicsSceneMouseEvent) SetScreenPos(pos *QPoint)  {
	q.Drv(278000,278127,Native(pos),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
