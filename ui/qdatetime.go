// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QDateTime : QDateTime
type QDateTime struct {
	BaseDrv
}
//QDateTime::QDateTime()	
func NewDateTime() *QDateTime {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),20000,20102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QDateTime{}
	_p.SetDriver(__rv,20000,true)
	return _p
} 
//QDateTime::QDateTime(QDate const&)	
func NewDateTimeWithDate(value *QDate) *QDateTime {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),20000,20103,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QDateTime{}
	_p.SetDriver(__rv,20000,true)
	return _p
} 
//QDateTime::QDateTime(QDateTime const&)	
func NewDateTimeCopy(other *QDateTime) *QDateTime {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),20000,20104,Native(other),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QDateTime{}
	_p.SetDriver(__rv,20000,true)
	return _p
} 
//QDateTime::QDateTime(QDate const&,QTime const&,Qt::TimeSpec)	
func NewDateTimeWithDateTimeSpec(value2 *QDate,value3 *QTime,spec Qt_TimeSpec) *QDateTime {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),20000,20105,Native(value2),Native(value3),unsafe.Pointer(&spec),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QDateTime{}
	_p.SetDriver(__rv,20000,true)
	return _p
} 
//QDateTime::addDays(int)
func (q *QDateTime) AddDays(days int) *QDateTime {
	var __rv uintptr
	q.Drv(20000,20106,unsafe.Pointer(&days),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QDateTime{}
	_rp.SetDriver(__rv,20000,true)
	return _rp
}	
//QDateTime::addMSecs(qint64)
func (q *QDateTime) AddMSecs(msecs int64) *QDateTime {
	var __rv uintptr
	q.Drv(20000,20107,unsafe.Pointer(&msecs),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QDateTime{}
	_rp.SetDriver(__rv,20000,true)
	return _rp
}	
//QDateTime::addMonths(int)
func (q *QDateTime) AddMonths(months int) *QDateTime {
	var __rv uintptr
	q.Drv(20000,20108,unsafe.Pointer(&months),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QDateTime{}
	_rp.SetDriver(__rv,20000,true)
	return _rp
}	
//QDateTime::addSecs(int)
func (q *QDateTime) AddSecs(secs int) *QDateTime {
	var __rv uintptr
	q.Drv(20000,20109,unsafe.Pointer(&secs),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QDateTime{}
	_rp.SetDriver(__rv,20000,true)
	return _rp
}	
//QDateTime::addYears(int)
func (q *QDateTime) AddYears(years int) *QDateTime {
	var __rv uintptr
	q.Drv(20000,20110,unsafe.Pointer(&years),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QDateTime{}
	_rp.SetDriver(__rv,20000,true)
	return _rp
}	
//QDateTime::currentDateTime()	
func QDateTimeCurrentDateTime() *QDateTime {
	var __rv uintptr
	DirectQtDrv(nil,20000,20111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QDateTime{}
	_rp.SetDriver(__rv,20000,true)
	return _rp
}	
//QDateTime::currentDateTime()
func (q *QDateTime) CurrentDateTime() *QDateTime {
	var __rv uintptr
	q.Drv(20000,20111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QDateTime{}
	_rp.SetDriver(__rv,20000,true)
	return _rp
}	
//QDateTime::currentDateTimeUtc()	
func QDateTimeCurrentDateTimeUtc() *QDateTime {
	var __rv uintptr
	DirectQtDrv(nil,20000,20112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QDateTime{}
	_rp.SetDriver(__rv,20000,true)
	return _rp
}	
//QDateTime::currentDateTimeUtc()
func (q *QDateTime) CurrentDateTimeUtc() *QDateTime {
	var __rv uintptr
	q.Drv(20000,20112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QDateTime{}
	_rp.SetDriver(__rv,20000,true)
	return _rp
}	
//QDateTime::currentMSecsSinceEpoch()	
func QDateTimeCurrentMSecsSinceEpoch() int64 {
	var __rv int64
	DirectQtDrv(nil,20000,20113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDateTime::currentMSecsSinceEpoch()
func (q *QDateTime) CurrentMSecsSinceEpoch() int64 {
	var __rv int64
	q.Drv(20000,20113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDateTime::date()
func (q *QDateTime) Date() *QDate {
	var __rv uintptr
	q.Drv(20000,20114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QDate{}
	_rp.SetDriver(__rv,19000,true)
	return _rp
}	
//QDateTime::daysTo(QDateTime const&)
func (q *QDateTime) DaysTo(value *QDateTime) int {
	var __rv int
	q.Drv(20000,20115,Native(value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDateTime::fromMSecsSinceEpoch(qint64)	
func QDateTimeFromMSecsSinceEpoch(msecs int64) *QDateTime {
	var __rv uintptr
	DirectQtDrv(nil,20000,20116,unsafe.Pointer(&msecs),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QDateTime{}
	_rp.SetDriver(__rv,20000,true)
	return _rp
}	
//QDateTime::fromMSecsSinceEpoch(qint64)
func (q *QDateTime) FromMSecsSinceEpoch(msecs int64) *QDateTime {
	var __rv uintptr
	q.Drv(20000,20116,unsafe.Pointer(&msecs),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QDateTime{}
	_rp.SetDriver(__rv,20000,true)
	return _rp
}	
//QDateTime::fromString(QString const&)	
func QDateTimeFromString(s string) *QDateTime {
	var __rv uintptr
	DirectQtDrv(nil,20000,20117,unsafe.Pointer(&s),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QDateTime{}
	_rp.SetDriver(__rv,20000,true)
	return _rp
}	
//QDateTime::fromString(QString const&)
func (q *QDateTime) FromString(s string) *QDateTime {
	var __rv uintptr
	q.Drv(20000,20117,unsafe.Pointer(&s),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QDateTime{}
	_rp.SetDriver(__rv,20000,true)
	return _rp
}	
//QDateTime::fromString(QString const&,QString const&)	
func QDateTimeFromStringWithTextFormat(s string,format string) *QDateTime {
	var __rv uintptr
	DirectQtDrv(nil,20000,20118,unsafe.Pointer(&s),unsafe.Pointer(&format),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QDateTime{}
	_rp.SetDriver(__rv,20000,true)
	return _rp
}	
//QDateTime::fromString(QString const&,QString const&)
func (q *QDateTime) FromStringWithTextFormat(s string,format string) *QDateTime {
	var __rv uintptr
	q.Drv(20000,20118,unsafe.Pointer(&s),unsafe.Pointer(&format),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QDateTime{}
	_rp.SetDriver(__rv,20000,true)
	return _rp
}	
//QDateTime::fromString(QString const&,Qt::DateFormat)	
func QDateTimeFromStringWithTextDateformat(s string,f Qt_DateFormat) *QDateTime {
	var __rv uintptr
	DirectQtDrv(nil,20000,20119,unsafe.Pointer(&s),unsafe.Pointer(&f),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QDateTime{}
	_rp.SetDriver(__rv,20000,true)
	return _rp
}	
//QDateTime::fromString(QString const&,Qt::DateFormat)
func (q *QDateTime) FromStringWithTextDateformat(s string,f Qt_DateFormat) *QDateTime {
	var __rv uintptr
	q.Drv(20000,20119,unsafe.Pointer(&s),unsafe.Pointer(&f),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QDateTime{}
	_rp.SetDriver(__rv,20000,true)
	return _rp
}	
//QDateTime::fromTime_t(unsigned int)	
func QDateTimeFromTime_t(secsSince1Jan1970UTC uint) *QDateTime {
	var __rv uintptr
	DirectQtDrv(nil,20000,20120,unsafe.Pointer(&secsSince1Jan1970UTC),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QDateTime{}
	_rp.SetDriver(__rv,20000,true)
	return _rp
}	
//QDateTime::fromTime_t(unsigned int)
func (q *QDateTime) FromTime_t(secsSince1Jan1970UTC uint) *QDateTime {
	var __rv uintptr
	q.Drv(20000,20120,unsafe.Pointer(&secsSince1Jan1970UTC),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QDateTime{}
	_rp.SetDriver(__rv,20000,true)
	return _rp
}	
//QDateTime::isNull()
func (q *QDateTime) IsNull() bool {
	var __rv bool
	q.Drv(20000,20121,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDateTime::isValid()
func (q *QDateTime) IsValid() bool {
	var __rv bool
	q.Drv(20000,20122,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDateTime::msecsTo(QDateTime const&)
func (q *QDateTime) MsecsTo(value *QDateTime) int64 {
	var __rv int64
	q.Drv(20000,20123,Native(value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDateTime::secsTo(QDateTime const&)
func (q *QDateTime) SecsTo(value *QDateTime) int {
	var __rv int
	q.Drv(20000,20124,Native(value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDateTime::setDate(QDate const&)
func (q *QDateTime) SetDate(date *QDate)  {
	q.Drv(20000,20125,Native(date),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDateTime::setMSecsSinceEpoch(qint64)
func (q *QDateTime) SetMSecsSinceEpoch(msecs int64)  {
	q.Drv(20000,20126,unsafe.Pointer(&msecs),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDateTime::setTime(QTime const&)
func (q *QDateTime) SetTime(time *QTime)  {
	q.Drv(20000,20127,Native(time),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDateTime::setTimeSpec(Qt::TimeSpec)
func (q *QDateTime) SetTimeSpec(spec Qt_TimeSpec)  {
	q.Drv(20000,20128,unsafe.Pointer(&spec),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDateTime::setTime_t(unsigned int)
func (q *QDateTime) SetTime_t(secsSince1Jan1970UTC uint)  {
	q.Drv(20000,20129,unsafe.Pointer(&secsSince1Jan1970UTC),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDateTime::setUtcOffset(int)
func (q *QDateTime) SetUtcOffset(seconds int)  {
	q.Drv(20000,20130,unsafe.Pointer(&seconds),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDateTime::time()
func (q *QDateTime) Time() *QTime {
	var __rv uintptr
	q.Drv(20000,20131,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTime{}
	_rp.SetDriver(__rv,170000,true)
	return _rp
}	
//QDateTime::timeSpec()
func (q *QDateTime) TimeSpec() Qt_TimeSpec {
	var __rv Qt_TimeSpec
	q.Drv(20000,20132,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDateTime::toLocalTime()
func (q *QDateTime) ToLocalTime() *QDateTime {
	var __rv uintptr
	q.Drv(20000,20133,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QDateTime{}
	_rp.SetDriver(__rv,20000,true)
	return _rp
}	
//QDateTime::toMSecsSinceEpoch()
func (q *QDateTime) ToMSecsSinceEpoch() int64 {
	var __rv int64
	q.Drv(20000,20134,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDateTime::toString()
func (q *QDateTime) ToString() string {
	var __rv string
	q.Drv(20000,20135,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDateTime::toString(QString const&)
func (q *QDateTime) ToStringWithFormat(format string) string {
	var __rv string
	q.Drv(20000,20136,unsafe.Pointer(&format),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDateTime::toString(Qt::DateFormat)
func (q *QDateTime) ToStringWithDateformat(f Qt_DateFormat) string {
	var __rv string
	q.Drv(20000,20137,unsafe.Pointer(&f),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDateTime::toTimeSpec(Qt::TimeSpec)
func (q *QDateTime) ToTimeSpec(spec Qt_TimeSpec) *QDateTime {
	var __rv uintptr
	q.Drv(20000,20138,unsafe.Pointer(&spec),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QDateTime{}
	_rp.SetDriver(__rv,20000,true)
	return _rp
}	
//QDateTime::toTime_t()
func (q *QDateTime) ToTime_t() uint {
	var __rv uint
	q.Drv(20000,20139,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDateTime::toUTC()
func (q *QDateTime) ToUTC() *QDateTime {
	var __rv uintptr
	q.Drv(20000,20140,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QDateTime{}
	_rp.SetDriver(__rv,20000,true)
	return _rp
}	
//QDateTime::utcOffset()
func (q *QDateTime) UtcOffset() int {
	var __rv int
	q.Drv(20000,20141,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
