// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QBasicTimer : QBasicTimer
type QBasicTimer struct {
	BaseDrv
}
//QBasicTimer::QBasicTimer()	
func NewBasicTimer() *QBasicTimer {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),7000,7102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QBasicTimer{}
	_p.SetDriver(__rv,7000,true)
	return _p
} 
//QBasicTimer::isActive()
func (q *QBasicTimer) IsActive() bool {
	var __rv bool
	q.Drv(7000,7103,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QBasicTimer::start(int,QObject*)
func (q *QBasicTimer) Start(msec int,obj QObjectInterface)  {
	q.Drv(7000,7104,unsafe.Pointer(&msec),Native(obj),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QBasicTimer::stop()
func (q *QBasicTimer) Stop()  {
	q.Drv(7000,7105,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QBasicTimer::timerId()
func (q *QBasicTimer) TimerId() int {
	var __rv int
	q.Drv(7000,7106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
