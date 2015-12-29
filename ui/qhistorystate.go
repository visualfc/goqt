// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QHistoryState_HistoryType - QHistoryState::HistoryType
type QHistoryState_HistoryType uint32
const (
	QHistoryState_ShallowHistory QHistoryState_HistoryType = 0
	QHistoryState_DeepHistory QHistoryState_HistoryType = 1
)
//struct QHistoryState : QHistoryState
type QHistoryState struct {
	QAbstractState
}
//QHistoryState::QHistoryState()	
func NewHistoryState() *QHistoryState {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),291000,291102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QHistoryState{}
	_p.SetDriver(__rv,291000,false)
	return _p
} 
//QHistoryState::QHistoryState(QState*)	
func NewHistoryStateWithParent(parent *QState) *QHistoryState {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),291000,291103,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QHistoryState{}
	_p.SetDriver(__rv,291000,false)
	return _p
} 
//QHistoryState::QHistoryState(QHistoryState::HistoryType,QState*)	
func NewHistoryStateWithTypeParent(_type QHistoryState_HistoryType,parent *QState) *QHistoryState {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),291000,291104,unsafe.Pointer(&_type),Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QHistoryState{}
	_p.SetDriver(__rv,291000,false)
	return _p
} 
//QHistoryState::defaultState()
func (q *QHistoryState) DefaultState() *QAbstractState {
	var __rv uintptr
	q.Drv(291000,291105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAbstractState{}
	_rp.SetDriver(__rv,203000,false)
	return _rp
}	
//QHistoryState::event(QEvent*)
func (q *QHistoryState) Event(e *QEvent) bool {
	var __rv bool
	q.Drv(291000,291106,Native(e),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QHistoryState::historyType()
func (q *QHistoryState) HistoryType() QHistoryState_HistoryType {
	var __rv QHistoryState_HistoryType
	q.Drv(291000,291107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QHistoryState::onEntry(QEvent*)
func (q *QHistoryState) OnEntry(event *QEvent)  {
	q.Drv(291000,291108,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QHistoryState::onExit(QEvent*)
func (q *QHistoryState) OnExit(event *QEvent)  {
	q.Drv(291000,291109,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QHistoryState::setDefaultState(QAbstractState*)
func (q *QHistoryState) SetDefaultState(state *QAbstractState)  {
	q.Drv(291000,291110,Native(state),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QHistoryState::setHistoryType(QHistoryState::HistoryType)
func (q *QHistoryState) SetHistoryType(_type QHistoryState_HistoryType)  {
	q.Drv(291000,291111,unsafe.Pointer(&_type),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
