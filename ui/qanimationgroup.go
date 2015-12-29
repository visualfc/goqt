// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QAnimationGroup : QAnimationGroup
type QAnimationGroup struct {
	QAbstractAnimation
}
//QAnimationGroup::addAnimation(QAbstractAnimation*)
func (q *QAnimationGroup) AddAnimation(animation *QAbstractAnimation)  {
	q.Drv(209000,209102,Native(animation),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAnimationGroup::animationAt(int)
func (q *QAnimationGroup) AnimationAt(index int) *QAbstractAnimation {
	var __rv uintptr
	q.Drv(209000,209103,unsafe.Pointer(&index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAbstractAnimation{}
	_rp.SetDriver(__rv,192000,false)
	return _rp
}	
//QAnimationGroup::animationCount()
func (q *QAnimationGroup) AnimationCount() int {
	var __rv int
	q.Drv(209000,209104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAnimationGroup::clear()
func (q *QAnimationGroup) Clear()  {
	q.Drv(209000,209105,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAnimationGroup::event(QEvent*)
func (q *QAnimationGroup) Event(event *QEvent) bool {
	var __rv bool
	q.Drv(209000,209106,Native(event),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAnimationGroup::indexOfAnimation(QAbstractAnimation*)
func (q *QAnimationGroup) IndexOfAnimation(animation *QAbstractAnimation) int {
	var __rv int
	q.Drv(209000,209107,Native(animation),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAnimationGroup::insertAnimation(int,QAbstractAnimation*)
func (q *QAnimationGroup) InsertAnimation(index int,animation *QAbstractAnimation)  {
	q.Drv(209000,209108,unsafe.Pointer(&index),Native(animation),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAnimationGroup::removeAnimation(QAbstractAnimation*)
func (q *QAnimationGroup) RemoveAnimation(animation *QAbstractAnimation)  {
	q.Drv(209000,209109,Native(animation),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAnimationGroup::takeAnimation(int)
func (q *QAnimationGroup) TakeAnimation(index int) *QAbstractAnimation {
	var __rv uintptr
	q.Drv(209000,209110,unsafe.Pointer(&index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAbstractAnimation{}
	_rp.SetDriver(__rv,192000,false)
	return _rp
}	
