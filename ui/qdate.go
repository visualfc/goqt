// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QDate_MonthNameType - QDate::MonthNameType
type QDate_MonthNameType uint32
const (
	QDate_DateFormat QDate_MonthNameType = 0
	QDate_StandaloneFormat QDate_MonthNameType = 0
)
//struct QDate : QDate
type QDate struct {
	BaseDrv
}
//QDate::QDate()	
func NewDate() *QDate {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),19000,19102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QDate{}
	_p.SetDriver(__rv,19000,true)
	return _p
} 
//QDate::QDate(QDate const&)	
func NewDateCopy(other *QDate) *QDate {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),19000,19103,Native(other),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QDate{}
	_p.SetDriver(__rv,19000,true)
	return _p
} 
//QDate::QDate(int,int,int)	
func NewDateWithYMD(y int,m int,d int) *QDate {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),19000,19104,unsafe.Pointer(&y),unsafe.Pointer(&m),unsafe.Pointer(&d),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QDate{}
	_p.SetDriver(__rv,19000,true)
	return _p
} 
//QDate::addDays(int)
func (q *QDate) AddDays(days int) *QDate {
	var __rv uintptr
	q.Drv(19000,19105,unsafe.Pointer(&days),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QDate{}
	_rp.SetDriver(__rv,19000,true)
	return _rp
}	
//QDate::addMonths(int)
func (q *QDate) AddMonths(months int) *QDate {
	var __rv uintptr
	q.Drv(19000,19106,unsafe.Pointer(&months),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QDate{}
	_rp.SetDriver(__rv,19000,true)
	return _rp
}	
//QDate::addYears(int)
func (q *QDate) AddYears(years int) *QDate {
	var __rv uintptr
	q.Drv(19000,19107,unsafe.Pointer(&years),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QDate{}
	_rp.SetDriver(__rv,19000,true)
	return _rp
}	
//QDate::currentDate()	
func QDateCurrentDate() *QDate {
	var __rv uintptr
	DirectQtDrv(nil,19000,19108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QDate{}
	_rp.SetDriver(__rv,19000,true)
	return _rp
}	
//QDate::currentDate()
func (q *QDate) CurrentDate() *QDate {
	var __rv uintptr
	q.Drv(19000,19108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QDate{}
	_rp.SetDriver(__rv,19000,true)
	return _rp
}	
//QDate::day()
func (q *QDate) Day() int {
	var __rv int
	q.Drv(19000,19109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDate::dayOfWeek()
func (q *QDate) DayOfWeek() int {
	var __rv int
	q.Drv(19000,19110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDate::dayOfYear()
func (q *QDate) DayOfYear() int {
	var __rv int
	q.Drv(19000,19111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDate::daysInMonth()
func (q *QDate) DaysInMonth() int {
	var __rv int
	q.Drv(19000,19112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDate::daysInYear()
func (q *QDate) DaysInYear() int {
	var __rv int
	q.Drv(19000,19113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDate::daysTo(QDate const&)
func (q *QDate) DaysTo(value *QDate) int {
	var __rv int
	q.Drv(19000,19114,Native(value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDate::fromJulianDay(int)	
func QDateFromJulianDay(jd int) *QDate {
	var __rv uintptr
	DirectQtDrv(nil,19000,19115,unsafe.Pointer(&jd),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QDate{}
	_rp.SetDriver(__rv,19000,true)
	return _rp
}	
//QDate::fromJulianDay(int)
func (q *QDate) FromJulianDay(jd int) *QDate {
	var __rv uintptr
	q.Drv(19000,19115,unsafe.Pointer(&jd),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QDate{}
	_rp.SetDriver(__rv,19000,true)
	return _rp
}	
//QDate::fromString(QString const&)	
func QDateFromString(s string) *QDate {
	var __rv uintptr
	DirectQtDrv(nil,19000,19116,unsafe.Pointer(&s),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QDate{}
	_rp.SetDriver(__rv,19000,true)
	return _rp
}	
//QDate::fromString(QString const&)
func (q *QDate) FromString(s string) *QDate {
	var __rv uintptr
	q.Drv(19000,19116,unsafe.Pointer(&s),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QDate{}
	_rp.SetDriver(__rv,19000,true)
	return _rp
}	
//QDate::fromString(QString const&,QString const&)	
func QDateFromStringWithTextFormat(s string,format string) *QDate {
	var __rv uintptr
	DirectQtDrv(nil,19000,19117,unsafe.Pointer(&s),unsafe.Pointer(&format),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QDate{}
	_rp.SetDriver(__rv,19000,true)
	return _rp
}	
//QDate::fromString(QString const&,QString const&)
func (q *QDate) FromStringWithTextFormat(s string,format string) *QDate {
	var __rv uintptr
	q.Drv(19000,19117,unsafe.Pointer(&s),unsafe.Pointer(&format),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QDate{}
	_rp.SetDriver(__rv,19000,true)
	return _rp
}	
//QDate::fromString(QString const&,Qt::DateFormat)	
func QDateFromStringWithTextDateformat(s string,f Qt_DateFormat) *QDate {
	var __rv uintptr
	DirectQtDrv(nil,19000,19118,unsafe.Pointer(&s),unsafe.Pointer(&f),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QDate{}
	_rp.SetDriver(__rv,19000,true)
	return _rp
}	
//QDate::fromString(QString const&,Qt::DateFormat)
func (q *QDate) FromStringWithTextDateformat(s string,f Qt_DateFormat) *QDate {
	var __rv uintptr
	q.Drv(19000,19118,unsafe.Pointer(&s),unsafe.Pointer(&f),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QDate{}
	_rp.SetDriver(__rv,19000,true)
	return _rp
}	
//QDate::getDate(int*,int*,int*)
func (q *QDate) GetDate(year *int,month *int,day *int)  {
	q.Drv(19000,19119,unsafe.Pointer(&year),unsafe.Pointer(&month),unsafe.Pointer(&day),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDate::isLeapYear(int)	
func QDateIsLeapYear(year int) bool {
	var __rv bool
	DirectQtDrv(nil,19000,19120,unsafe.Pointer(&year),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDate::isLeapYear(int)
func (q *QDate) IsLeapYear(year int) bool {
	var __rv bool
	q.Drv(19000,19120,unsafe.Pointer(&year),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDate::isNull()
func (q *QDate) IsNull() bool {
	var __rv bool
	q.Drv(19000,19121,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDate::isValid()
func (q *QDate) IsValid() bool {
	var __rv bool
	q.Drv(19000,19122,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDate::isValid(int,int,int)	
func QDateIsValidWithYMD(y int,m int,d int) bool {
	var __rv bool
	DirectQtDrv(nil,19000,19123,unsafe.Pointer(&y),unsafe.Pointer(&m),unsafe.Pointer(&d),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDate::isValid(int,int,int)
func (q *QDate) IsValidWithYMD(y int,m int,d int) bool {
	var __rv bool
	q.Drv(19000,19123,unsafe.Pointer(&y),unsafe.Pointer(&m),unsafe.Pointer(&d),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDate::longDayName(int)	
func QDateLongDayName(weekday int) string {
	var __rv string
	DirectQtDrv(nil,19000,19124,unsafe.Pointer(&weekday),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDate::longDayName(int)
func (q *QDate) LongDayName(weekday int) string {
	var __rv string
	q.Drv(19000,19124,unsafe.Pointer(&weekday),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDate::longDayName(int,QDate::MonthNameType)	
func QDateLongDayNameWithWeekdayType(weekday int,_type QDate_MonthNameType) string {
	var __rv string
	DirectQtDrv(nil,19000,19125,unsafe.Pointer(&weekday),unsafe.Pointer(&_type),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDate::longDayName(int,QDate::MonthNameType)
func (q *QDate) LongDayNameWithWeekdayType(weekday int,_type QDate_MonthNameType) string {
	var __rv string
	q.Drv(19000,19125,unsafe.Pointer(&weekday),unsafe.Pointer(&_type),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDate::longMonthName(int)	
func QDateLongMonthName(month int) string {
	var __rv string
	DirectQtDrv(nil,19000,19126,unsafe.Pointer(&month),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDate::longMonthName(int)
func (q *QDate) LongMonthName(month int) string {
	var __rv string
	q.Drv(19000,19126,unsafe.Pointer(&month),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDate::longMonthName(int,QDate::MonthNameType)	
func QDateLongMonthNameWithMonthType(month int,_type QDate_MonthNameType) string {
	var __rv string
	DirectQtDrv(nil,19000,19127,unsafe.Pointer(&month),unsafe.Pointer(&_type),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDate::longMonthName(int,QDate::MonthNameType)
func (q *QDate) LongMonthNameWithMonthType(month int,_type QDate_MonthNameType) string {
	var __rv string
	q.Drv(19000,19127,unsafe.Pointer(&month),unsafe.Pointer(&_type),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDate::month()
func (q *QDate) Month() int {
	var __rv int
	q.Drv(19000,19128,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDate::setDate(int,int,int)
func (q *QDate) SetDate(year int,month int,day int) bool {
	var __rv bool
	q.Drv(19000,19129,unsafe.Pointer(&year),unsafe.Pointer(&month),unsafe.Pointer(&day),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDate::shortDayName(int)	
func QDateShortDayName(weekday int) string {
	var __rv string
	DirectQtDrv(nil,19000,19130,unsafe.Pointer(&weekday),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDate::shortDayName(int)
func (q *QDate) ShortDayName(weekday int) string {
	var __rv string
	q.Drv(19000,19130,unsafe.Pointer(&weekday),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDate::shortDayName(int,QDate::MonthNameType)	
func QDateShortDayNameWithWeekdayType(weekday int,_type QDate_MonthNameType) string {
	var __rv string
	DirectQtDrv(nil,19000,19131,unsafe.Pointer(&weekday),unsafe.Pointer(&_type),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDate::shortDayName(int,QDate::MonthNameType)
func (q *QDate) ShortDayNameWithWeekdayType(weekday int,_type QDate_MonthNameType) string {
	var __rv string
	q.Drv(19000,19131,unsafe.Pointer(&weekday),unsafe.Pointer(&_type),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDate::shortMonthName(int)	
func QDateShortMonthName(month int) string {
	var __rv string
	DirectQtDrv(nil,19000,19132,unsafe.Pointer(&month),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDate::shortMonthName(int)
func (q *QDate) ShortMonthName(month int) string {
	var __rv string
	q.Drv(19000,19132,unsafe.Pointer(&month),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDate::shortMonthName(int,QDate::MonthNameType)	
func QDateShortMonthNameWithMonthType(month int,_type QDate_MonthNameType) string {
	var __rv string
	DirectQtDrv(nil,19000,19133,unsafe.Pointer(&month),unsafe.Pointer(&_type),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDate::shortMonthName(int,QDate::MonthNameType)
func (q *QDate) ShortMonthNameWithMonthType(month int,_type QDate_MonthNameType) string {
	var __rv string
	q.Drv(19000,19133,unsafe.Pointer(&month),unsafe.Pointer(&_type),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDate::toJulianDay()
func (q *QDate) ToJulianDay() int {
	var __rv int
	q.Drv(19000,19134,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDate::toString()
func (q *QDate) ToString() string {
	var __rv string
	q.Drv(19000,19135,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDate::toString(QString const&)
func (q *QDate) ToStringWithFormat(format string) string {
	var __rv string
	q.Drv(19000,19136,unsafe.Pointer(&format),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDate::toString(Qt::DateFormat)
func (q *QDate) ToStringWithDateformat(f Qt_DateFormat) string {
	var __rv string
	q.Drv(19000,19137,unsafe.Pointer(&f),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDate::weekNumber()
func (q *QDate) WeekNumber() int {
	var __rv int
	q.Drv(19000,19138,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDate::weekNumber(int*)
func (q *QDate) WeekNumberWithYearnum(yearNum *int) int {
	var __rv int
	q.Drv(19000,19139,unsafe.Pointer(&yearNum),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDate::year()
func (q *QDate) Year() int {
	var __rv int
	q.Drv(19000,19140,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
