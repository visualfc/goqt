// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QGraphicsSceneDragDropEvent : QGraphicsSceneDragDropEvent
type QGraphicsSceneDragDropEvent struct {
	QGraphicsSceneEvent
}
//QGraphicsSceneDragDropEvent::QGraphicsSceneDragDropEvent()	
func NewGraphicsSceneDragDropEvent() *QGraphicsSceneDragDropEvent {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),274000,274102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGraphicsSceneDragDropEvent{}
	_p.SetDriver(__rv,274000,true)
	return _p
} 
//QGraphicsSceneDragDropEvent::QGraphicsSceneDragDropEvent(QEvent::Type)	
func NewGraphicsSceneDragDropEventWithType(_type QEvent_Type) *QGraphicsSceneDragDropEvent {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),274000,274103,unsafe.Pointer(&_type),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGraphicsSceneDragDropEvent{}
	_p.SetDriver(__rv,274000,true)
	return _p
} 
//QGraphicsSceneDragDropEvent::acceptProposedAction()
func (q *QGraphicsSceneDragDropEvent) AcceptProposedAction()  {
	q.Drv(274000,274104,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsSceneDragDropEvent::buttons()
func (q *QGraphicsSceneDragDropEvent) Buttons() Qt_MouseButton {
	var __rv Qt_MouseButton
	q.Drv(274000,274105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsSceneDragDropEvent::dropAction()
func (q *QGraphicsSceneDragDropEvent) DropAction() Qt_DropAction {
	var __rv Qt_DropAction
	q.Drv(274000,274106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsSceneDragDropEvent::mimeData()
func (q *QGraphicsSceneDragDropEvent) MimeData() *QMimeData {
	var __rv uintptr
	q.Drv(274000,274107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QMimeData{}
	_rp.SetDriver(__rv,311000,false)
	return _rp
}	
//QGraphicsSceneDragDropEvent::modifiers()
func (q *QGraphicsSceneDragDropEvent) Modifiers() Qt_KeyboardModifier {
	var __rv Qt_KeyboardModifier
	q.Drv(274000,274108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsSceneDragDropEvent::pos()
func (q *QGraphicsSceneDragDropEvent) Pos() *QPointF {
	var __rv uintptr
	q.Drv(274000,274109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QGraphicsSceneDragDropEvent::possibleActions()
func (q *QGraphicsSceneDragDropEvent) PossibleActions() Qt_DropAction {
	var __rv Qt_DropAction
	q.Drv(274000,274110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsSceneDragDropEvent::proposedAction()
func (q *QGraphicsSceneDragDropEvent) ProposedAction() Qt_DropAction {
	var __rv Qt_DropAction
	q.Drv(274000,274111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsSceneDragDropEvent::scenePos()
func (q *QGraphicsSceneDragDropEvent) ScenePos() *QPointF {
	var __rv uintptr
	q.Drv(274000,274112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QGraphicsSceneDragDropEvent::screenPos()
func (q *QGraphicsSceneDragDropEvent) ScreenPos() *QPoint {
	var __rv uintptr
	q.Drv(274000,274113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPoint{}
	_rp.SetDriver(__rv,99000,true)
	return _rp
}	
//QGraphicsSceneDragDropEvent::setButtons(QFlags<Qt::MouseButton>)
func (q *QGraphicsSceneDragDropEvent) SetButtons(buttons Qt_MouseButton)  {
	q.Drv(274000,274114,unsafe.Pointer(&buttons),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsSceneDragDropEvent::setDropAction(Qt::DropAction)
func (q *QGraphicsSceneDragDropEvent) SetDropAction(action Qt_DropAction)  {
	q.Drv(274000,274115,unsafe.Pointer(&action),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsSceneDragDropEvent::setMimeData(QMimeData const*)
func (q *QGraphicsSceneDragDropEvent) SetMimeData(data *QMimeData)  {
	q.Drv(274000,274116,Native(data),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsSceneDragDropEvent::setModifiers(QFlags<Qt::KeyboardModifier>)
func (q *QGraphicsSceneDragDropEvent) SetModifiers(modifiers Qt_KeyboardModifier)  {
	q.Drv(274000,274117,unsafe.Pointer(&modifiers),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsSceneDragDropEvent::setPos(QPointF const&)
func (q *QGraphicsSceneDragDropEvent) SetPos(pos *QPointF)  {
	q.Drv(274000,274118,Native(pos),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsSceneDragDropEvent::setPossibleActions(QFlags<Qt::DropAction>)
func (q *QGraphicsSceneDragDropEvent) SetPossibleActions(actions Qt_DropAction)  {
	q.Drv(274000,274119,unsafe.Pointer(&actions),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsSceneDragDropEvent::setProposedAction(Qt::DropAction)
func (q *QGraphicsSceneDragDropEvent) SetProposedAction(action Qt_DropAction)  {
	q.Drv(274000,274120,unsafe.Pointer(&action),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsSceneDragDropEvent::setScenePos(QPointF const&)
func (q *QGraphicsSceneDragDropEvent) SetScenePos(pos *QPointF)  {
	q.Drv(274000,274121,Native(pos),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsSceneDragDropEvent::setScreenPos(QPoint const&)
func (q *QGraphicsSceneDragDropEvent) SetScreenPos(pos *QPoint)  {
	q.Drv(274000,274122,Native(pos),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsSceneDragDropEvent::setSource(QWidget*)
func (q *QGraphicsSceneDragDropEvent) SetSource(source QWidgetInterface)  {
	q.Drv(274000,274123,Native(source),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsSceneDragDropEvent::source()
func (q *QGraphicsSceneDragDropEvent) Source() *QWidget {
	var __rv uintptr
	q.Drv(274000,274124,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
