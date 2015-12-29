// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QPauseAnimation : QPauseAnimation
type QPauseAnimation struct {
	QAbstractAnimation
}
//QPauseAnimation::QPauseAnimation()	
func NewPauseAnimation() *QPauseAnimation {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),318000,318102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPauseAnimation{}
	_p.SetDriver(__rv,318000,false)
	return _p
} 
//QPauseAnimation::QPauseAnimation(QObject*)	
func NewPauseAnimationWithParent(parent QObjectInterface) *QPauseAnimation {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),318000,318103,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPauseAnimation{}
	_p.SetDriver(__rv,318000,false)
	return _p
} 
//QPauseAnimation::QPauseAnimation(int,QObject*)	
func NewPauseAnimationWithMsecsParent(msecs int,parent QObjectInterface) *QPauseAnimation {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),318000,318104,unsafe.Pointer(&msecs),Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPauseAnimation{}
	_p.SetDriver(__rv,318000,false)
	return _p
} 
//QPauseAnimation::duration()
func (q *QPauseAnimation) Duration() int {
	var __rv int
	q.Drv(318000,318105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPauseAnimation::event(QEvent*)
func (q *QPauseAnimation) Event(e *QEvent) bool {
	var __rv bool
	q.Drv(318000,318106,Native(e),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPauseAnimation::setDuration(int)
func (q *QPauseAnimation) SetDuration(msecs int)  {
	q.Drv(318000,318107,unsafe.Pointer(&msecs),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPauseAnimation::updateCurrentTime(int)
func (q *QPauseAnimation) UpdateCurrentTime(value int)  {
	q.Drv(318000,318108,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
