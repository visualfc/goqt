// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QParallelAnimationGroup : QParallelAnimationGroup
type QParallelAnimationGroup struct {
	QAnimationGroup
}
//QParallelAnimationGroup::QParallelAnimationGroup()	
func NewParallelAnimationGroup() *QParallelAnimationGroup {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),317000,317102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QParallelAnimationGroup{}
	_p.SetDriver(__rv,317000,false)
	return _p
} 
//QParallelAnimationGroup::QParallelAnimationGroup(QObject*)	
func NewParallelAnimationGroupWithParent(parent QObjectInterface) *QParallelAnimationGroup {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),317000,317103,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QParallelAnimationGroup{}
	_p.SetDriver(__rv,317000,false)
	return _p
} 
//QParallelAnimationGroup::duration()
func (q *QParallelAnimationGroup) Duration() int {
	var __rv int
	q.Drv(317000,317104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QParallelAnimationGroup::event(QEvent*)
func (q *QParallelAnimationGroup) Event(event *QEvent) bool {
	var __rv bool
	q.Drv(317000,317105,Native(event),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QParallelAnimationGroup::updateCurrentTime(int)
func (q *QParallelAnimationGroup) UpdateCurrentTime(currentTime int)  {
	q.Drv(317000,317106,unsafe.Pointer(&currentTime),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QParallelAnimationGroup::updateDirection(QAbstractAnimation::Direction)
func (q *QParallelAnimationGroup) UpdateDirection(direction QAbstractAnimation_Direction)  {
	q.Drv(317000,317107,unsafe.Pointer(&direction),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QParallelAnimationGroup::updateState(QAbstractAnimation::State,QAbstractAnimation::State)
func (q *QParallelAnimationGroup) UpdateState(newState QAbstractAnimation_State,oldState QAbstractAnimation_State)  {
	q.Drv(317000,317108,unsafe.Pointer(&newState),unsafe.Pointer(&oldState),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
