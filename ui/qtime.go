// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QTime_TimeFlag - QTime::TimeFlag
type QTime_TimeFlag int32
const (
	QTime_NullTime QTime_TimeFlag = -1
)
//struct QTime : QTime
type QTime struct {
	BaseDrv
}
//QTime::QTime()	
func NewTime() *QTime {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),170000,170102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTime{}
	_p.SetDriver(__rv,170000,true)
	return _p
} 
//QTime::QTime(QTime const&)	
func NewTimeCopy(other *QTime) *QTime {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),170000,170103,Native(other),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTime{}
	_p.SetDriver(__rv,170000,true)
	return _p
} 
//QTime::QTime(int,int,int,int)	
func NewTimeWithHeightMSMs(h int,m int,s int,ms int) *QTime {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),170000,170104,unsafe.Pointer(&h),unsafe.Pointer(&m),unsafe.Pointer(&s),unsafe.Pointer(&ms),nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTime{}
	_p.SetDriver(__rv,170000,true)
	return _p
} 
//QTime::addMSecs(int)
func (q *QTime) AddMSecs(ms int) *QTime {
	var __rv uintptr
	q.Drv(170000,170105,unsafe.Pointer(&ms),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTime{}
	_rp.SetDriver(__rv,170000,true)
	return _rp
}	
//QTime::addSecs(int)
func (q *QTime) AddSecs(secs int) *QTime {
	var __rv uintptr
	q.Drv(170000,170106,unsafe.Pointer(&secs),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTime{}
	_rp.SetDriver(__rv,170000,true)
	return _rp
}	
//QTime::currentTime()	
func QTimeCurrentTime() *QTime {
	var __rv uintptr
	DirectQtDrv(nil,170000,170107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTime{}
	_rp.SetDriver(__rv,170000,true)
	return _rp
}	
//QTime::currentTime()
func (q *QTime) CurrentTime() *QTime {
	var __rv uintptr
	q.Drv(170000,170107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTime{}
	_rp.SetDriver(__rv,170000,true)
	return _rp
}	
//QTime::elapsed()
func (q *QTime) Elapsed() int {
	var __rv int
	q.Drv(170000,170108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTime::fromString(QString const&)	
func QTimeFromString(s string) *QTime {
	var __rv uintptr
	DirectQtDrv(nil,170000,170109,unsafe.Pointer(&s),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTime{}
	_rp.SetDriver(__rv,170000,true)
	return _rp
}	
//QTime::fromString(QString const&)
func (q *QTime) FromString(s string) *QTime {
	var __rv uintptr
	q.Drv(170000,170109,unsafe.Pointer(&s),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTime{}
	_rp.SetDriver(__rv,170000,true)
	return _rp
}	
//QTime::fromString(QString const&,QString const&)	
func QTimeFromStringWithTextFormat(s string,format string) *QTime {
	var __rv uintptr
	DirectQtDrv(nil,170000,170110,unsafe.Pointer(&s),unsafe.Pointer(&format),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTime{}
	_rp.SetDriver(__rv,170000,true)
	return _rp
}	
//QTime::fromString(QString const&,QString const&)
func (q *QTime) FromStringWithTextFormat(s string,format string) *QTime {
	var __rv uintptr
	q.Drv(170000,170110,unsafe.Pointer(&s),unsafe.Pointer(&format),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTime{}
	_rp.SetDriver(__rv,170000,true)
	return _rp
}	
//QTime::fromString(QString const&,Qt::DateFormat)	
func QTimeFromStringWithTextDateformat(s string,f Qt_DateFormat) *QTime {
	var __rv uintptr
	DirectQtDrv(nil,170000,170111,unsafe.Pointer(&s),unsafe.Pointer(&f),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTime{}
	_rp.SetDriver(__rv,170000,true)
	return _rp
}	
//QTime::fromString(QString const&,Qt::DateFormat)
func (q *QTime) FromStringWithTextDateformat(s string,f Qt_DateFormat) *QTime {
	var __rv uintptr
	q.Drv(170000,170111,unsafe.Pointer(&s),unsafe.Pointer(&f),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTime{}
	_rp.SetDriver(__rv,170000,true)
	return _rp
}	
//QTime::hour()
func (q *QTime) Hour() int {
	var __rv int
	q.Drv(170000,170112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTime::isNull()
func (q *QTime) IsNull() bool {
	var __rv bool
	q.Drv(170000,170113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTime::isValid()
func (q *QTime) IsValid() bool {
	var __rv bool
	q.Drv(170000,170114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTime::isValid(int,int,int,int)	
func QTimeIsValidWithHeightMSMs(h int,m int,s int,ms int) bool {
	var __rv bool
	DirectQtDrv(nil,170000,170115,unsafe.Pointer(&h),unsafe.Pointer(&m),unsafe.Pointer(&s),unsafe.Pointer(&ms),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTime::isValid(int,int,int,int)
func (q *QTime) IsValidWithHeightMSMs(h int,m int,s int,ms int) bool {
	var __rv bool
	q.Drv(170000,170115,unsafe.Pointer(&h),unsafe.Pointer(&m),unsafe.Pointer(&s),unsafe.Pointer(&ms),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTime::minute()
func (q *QTime) Minute() int {
	var __rv int
	q.Drv(170000,170116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTime::msec()
func (q *QTime) Msec() int {
	var __rv int
	q.Drv(170000,170117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTime::msecsTo(QTime const&)
func (q *QTime) MsecsTo(value *QTime) int {
	var __rv int
	q.Drv(170000,170118,Native(value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTime::restart()
func (q *QTime) Restart() int {
	var __rv int
	q.Drv(170000,170119,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTime::second()
func (q *QTime) Second() int {
	var __rv int
	q.Drv(170000,170120,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTime::secsTo(QTime const&)
func (q *QTime) SecsTo(value *QTime) int {
	var __rv int
	q.Drv(170000,170121,Native(value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTime::setHMS(int,int,int,int)
func (q *QTime) SetHMS(h int,m int,s int,ms int) bool {
	var __rv bool
	q.Drv(170000,170122,unsafe.Pointer(&h),unsafe.Pointer(&m),unsafe.Pointer(&s),unsafe.Pointer(&ms),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTime::start()
func (q *QTime) Start()  {
	q.Drv(170000,170123,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTime::toString()
func (q *QTime) ToString() string {
	var __rv string
	q.Drv(170000,170124,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTime::toString(QString const&)
func (q *QTime) ToStringWithFormat(format string) string {
	var __rv string
	q.Drv(170000,170125,unsafe.Pointer(&format),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTime::toString(Qt::DateFormat)
func (q *QTime) ToStringWithDateformat(f Qt_DateFormat) string {
	var __rv string
	q.Drv(170000,170126,unsafe.Pointer(&f),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
