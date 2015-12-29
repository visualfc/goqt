// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QSequentialAnimationGroup : QSequentialAnimationGroup
type QSequentialAnimationGroup struct {
	QAnimationGroup
}
func (q *QSequentialAnimationGroup) OnCurrentAnimationChanged(fn func(*QAbstractAnimation)) uintptr {
	var __rv uintptr
	q.Drv(337000,337102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QSequentialAnimationGroup::QSequentialAnimationGroup()	
func NewSequentialAnimationGroup() *QSequentialAnimationGroup {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),337000,337103,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QSequentialAnimationGroup{}
	_p.SetDriver(__rv,337000,false)
	return _p
} 
//QSequentialAnimationGroup::QSequentialAnimationGroup(QObject*)	
func NewSequentialAnimationGroupWithParent(parent QObjectInterface) *QSequentialAnimationGroup {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),337000,337104,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QSequentialAnimationGroup{}
	_p.SetDriver(__rv,337000,false)
	return _p
} 
//QSequentialAnimationGroup::addPause(int)
func (q *QSequentialAnimationGroup) AddPause(msecs int) *QPauseAnimation {
	var __rv uintptr
	q.Drv(337000,337105,unsafe.Pointer(&msecs),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPauseAnimation{}
	_rp.SetDriver(__rv,318000,false)
	return _rp
}	
//QSequentialAnimationGroup::currentAnimation()
func (q *QSequentialAnimationGroup) CurrentAnimation() *QAbstractAnimation {
	var __rv uintptr
	q.Drv(337000,337106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAbstractAnimation{}
	_rp.SetDriver(__rv,192000,false)
	return _rp
}	
//QSequentialAnimationGroup::duration()
func (q *QSequentialAnimationGroup) Duration() int {
	var __rv int
	q.Drv(337000,337107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSequentialAnimationGroup::event(QEvent*)
func (q *QSequentialAnimationGroup) Event(event *QEvent) bool {
	var __rv bool
	q.Drv(337000,337108,Native(event),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSequentialAnimationGroup::insertPause(int,int)
func (q *QSequentialAnimationGroup) InsertPause(index int,msecs int) *QPauseAnimation {
	var __rv uintptr
	q.Drv(337000,337109,unsafe.Pointer(&index),unsafe.Pointer(&msecs),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPauseAnimation{}
	_rp.SetDriver(__rv,318000,false)
	return _rp
}	
//QSequentialAnimationGroup::updateCurrentTime(int)
func (q *QSequentialAnimationGroup) UpdateCurrentTime(value int)  {
	q.Drv(337000,337110,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSequentialAnimationGroup::updateDirection(QAbstractAnimation::Direction)
func (q *QSequentialAnimationGroup) UpdateDirection(direction QAbstractAnimation_Direction)  {
	q.Drv(337000,337111,unsafe.Pointer(&direction),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSequentialAnimationGroup::updateState(QAbstractAnimation::State,QAbstractAnimation::State)
func (q *QSequentialAnimationGroup) UpdateState(newState QAbstractAnimation_State,oldState QAbstractAnimation_State)  {
	q.Drv(337000,337112,unsafe.Pointer(&newState),unsafe.Pointer(&oldState),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
