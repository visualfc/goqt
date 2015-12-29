// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QDateTimeEdit_Section - QDateTimeEdit::Section
type QDateTimeEdit_Section uint32
const (
	QDateTimeEdit_NoSection QDateTimeEdit_Section = 0x0000
	QDateTimeEdit_AmPmSection QDateTimeEdit_Section = 0x0001
	QDateTimeEdit_MSecSection QDateTimeEdit_Section = 0x0002
	QDateTimeEdit_SecondSection QDateTimeEdit_Section = 0x0004
	QDateTimeEdit_MinuteSection QDateTimeEdit_Section = 0x0008
	QDateTimeEdit_HourSection QDateTimeEdit_Section = 0x0010
	QDateTimeEdit_DaySection QDateTimeEdit_Section = 0x0100
	QDateTimeEdit_MonthSection QDateTimeEdit_Section = 0x0200
	QDateTimeEdit_YearSection QDateTimeEdit_Section = 0x0400
	QDateTimeEdit_TimeSections_Mask QDateTimeEdit_Section = QDateTimeEdit_AmPmSection|QDateTimeEdit_MSecSection|QDateTimeEdit_SecondSection|QDateTimeEdit_MinuteSection|QDateTimeEdit_HourSection
	QDateTimeEdit_DateSections_Mask QDateTimeEdit_Section = QDateTimeEdit_DaySection|QDateTimeEdit_MonthSection|QDateTimeEdit_YearSection
)
//struct QDateTimeEdit : QDateTimeEdit
type QDateTimeEdit struct {
	QAbstractSpinBox
}
func (q *QDateTimeEdit) OnDateTimeChanged(fn func(*QDateTime)) uintptr {
	var __rv uintptr
	q.Drv(223000,223102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QDateTimeEdit) OnDateChanged(fn func(*QDate)) uintptr {
	var __rv uintptr
	q.Drv(223000,223103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QDateTimeEdit) OnTimeChanged(fn func(*QTime)) uintptr {
	var __rv uintptr
	q.Drv(223000,223104,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QDateTimeEdit::QDateTimeEdit()	
func NewDateTimeEdit() *QDateTimeEdit {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),223000,223105,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QDateTimeEdit{}
	_p.SetDriver(__rv,223000,false)
	return _p
} 
//QDateTimeEdit::QDateTimeEdit(QWidget*)	
func NewDateTimeEditWithParent(parent QWidgetInterface) *QDateTimeEdit {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),223000,223106,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QDateTimeEdit{}
	_p.SetDriver(__rv,223000,false)
	return _p
} 
//QDateTimeEdit::QDateTimeEdit(QDate const&,QWidget*)	
func NewDateTimeEditWithDParent(d *QDate,parent QWidgetInterface) *QDateTimeEdit {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),223000,223107,Native(d),Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QDateTimeEdit{}
	_p.SetDriver(__rv,223000,false)
	return _p
} 
//QDateTimeEdit::QDateTimeEdit(QDateTime const&,QWidget*)	
func NewDateTimeEditWithDtParent(dt *QDateTime,parent QWidgetInterface) *QDateTimeEdit {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),223000,223108,Native(dt),Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QDateTimeEdit{}
	_p.SetDriver(__rv,223000,false)
	return _p
} 
//QDateTimeEdit::QDateTimeEdit(QTime const&,QWidget*)	
func NewDateTimeEditWithTParent(t *QTime,parent QWidgetInterface) *QDateTimeEdit {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),223000,223109,Native(t),Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QDateTimeEdit{}
	_p.SetDriver(__rv,223000,false)
	return _p
} 
//QDateTimeEdit::calendarPopup()
func (q *QDateTimeEdit) CalendarPopup() bool {
	var __rv bool
	q.Drv(223000,223110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDateTimeEdit::calendarWidget()
func (q *QDateTimeEdit) CalendarWidget() *QCalendarWidget {
	var __rv uintptr
	q.Drv(223000,223111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QCalendarWidget{}
	_rp.SetDriver(__rv,212000,false)
	return _rp
}	
//QDateTimeEdit::clear()
func (q *QDateTimeEdit) Clear()  {
	q.Drv(223000,223112,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDateTimeEdit::clearMaximumDate()
func (q *QDateTimeEdit) ClearMaximumDate()  {
	q.Drv(223000,223113,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDateTimeEdit::clearMaximumDateTime()
func (q *QDateTimeEdit) ClearMaximumDateTime()  {
	q.Drv(223000,223114,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDateTimeEdit::clearMaximumTime()
func (q *QDateTimeEdit) ClearMaximumTime()  {
	q.Drv(223000,223115,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDateTimeEdit::clearMinimumDate()
func (q *QDateTimeEdit) ClearMinimumDate()  {
	q.Drv(223000,223116,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDateTimeEdit::clearMinimumDateTime()
func (q *QDateTimeEdit) ClearMinimumDateTime()  {
	q.Drv(223000,223117,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDateTimeEdit::clearMinimumTime()
func (q *QDateTimeEdit) ClearMinimumTime()  {
	q.Drv(223000,223118,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDateTimeEdit::currentSection()
func (q *QDateTimeEdit) CurrentSection() QDateTimeEdit_Section {
	var __rv QDateTimeEdit_Section
	q.Drv(223000,223119,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDateTimeEdit::currentSectionIndex()
func (q *QDateTimeEdit) CurrentSectionIndex() int {
	var __rv int
	q.Drv(223000,223120,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDateTimeEdit::date()
func (q *QDateTimeEdit) Date() *QDate {
	var __rv uintptr
	q.Drv(223000,223121,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QDate{}
	_rp.SetDriver(__rv,19000,true)
	return _rp
}	
//QDateTimeEdit::dateTime()
func (q *QDateTimeEdit) DateTime() *QDateTime {
	var __rv uintptr
	q.Drv(223000,223122,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QDateTime{}
	_rp.SetDriver(__rv,20000,true)
	return _rp
}	
//QDateTimeEdit::dateTimeFromText(QString const&)
func (q *QDateTimeEdit) DateTimeFromText(text string) *QDateTime {
	var __rv uintptr
	q.Drv(223000,223123,unsafe.Pointer(&text),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QDateTime{}
	_rp.SetDriver(__rv,20000,true)
	return _rp
}	
//QDateTimeEdit::displayFormat()
func (q *QDateTimeEdit) DisplayFormat() string {
	var __rv string
	q.Drv(223000,223124,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDateTimeEdit::displayedSections()
func (q *QDateTimeEdit) DisplayedSections() QDateTimeEdit_Section {
	var __rv QDateTimeEdit_Section
	q.Drv(223000,223125,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDateTimeEdit::event(QEvent*)
func (q *QDateTimeEdit) Event(event *QEvent) bool {
	var __rv bool
	q.Drv(223000,223126,Native(event),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDateTimeEdit::fixup(QString&)
func (q *QDateTimeEdit) Fixup(input *string)  {
	q.Drv(223000,223127,unsafe.Pointer(&input),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDateTimeEdit::focusInEvent(QFocusEvent*)
func (q *QDateTimeEdit) FocusInEvent(event *QFocusEvent)  {
	q.Drv(223000,223128,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDateTimeEdit::focusNextPrevChild(bool)
func (q *QDateTimeEdit) FocusNextPrevChild(next bool) bool {
	var __rv bool
	q.Drv(223000,223129,unsafe.Pointer(&next),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDateTimeEdit::keyPressEvent(QKeyEvent*)
func (q *QDateTimeEdit) KeyPressEvent(event *QKeyEvent)  {
	q.Drv(223000,223130,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDateTimeEdit::maximumDate()
func (q *QDateTimeEdit) MaximumDate() *QDate {
	var __rv uintptr
	q.Drv(223000,223131,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QDate{}
	_rp.SetDriver(__rv,19000,true)
	return _rp
}	
//QDateTimeEdit::maximumDateTime()
func (q *QDateTimeEdit) MaximumDateTime() *QDateTime {
	var __rv uintptr
	q.Drv(223000,223132,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QDateTime{}
	_rp.SetDriver(__rv,20000,true)
	return _rp
}	
//QDateTimeEdit::maximumTime()
func (q *QDateTimeEdit) MaximumTime() *QTime {
	var __rv uintptr
	q.Drv(223000,223133,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTime{}
	_rp.SetDriver(__rv,170000,true)
	return _rp
}	
//QDateTimeEdit::minimumDate()
func (q *QDateTimeEdit) MinimumDate() *QDate {
	var __rv uintptr
	q.Drv(223000,223134,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QDate{}
	_rp.SetDriver(__rv,19000,true)
	return _rp
}	
//QDateTimeEdit::minimumDateTime()
func (q *QDateTimeEdit) MinimumDateTime() *QDateTime {
	var __rv uintptr
	q.Drv(223000,223135,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QDateTime{}
	_rp.SetDriver(__rv,20000,true)
	return _rp
}	
//QDateTimeEdit::minimumTime()
func (q *QDateTimeEdit) MinimumTime() *QTime {
	var __rv uintptr
	q.Drv(223000,223136,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTime{}
	_rp.SetDriver(__rv,170000,true)
	return _rp
}	
//QDateTimeEdit::mousePressEvent(QMouseEvent*)
func (q *QDateTimeEdit) MousePressEvent(event *QMouseEvent)  {
	q.Drv(223000,223137,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDateTimeEdit::paintEvent(QPaintEvent*)
func (q *QDateTimeEdit) PaintEvent(event *QPaintEvent)  {
	q.Drv(223000,223138,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDateTimeEdit::sectionAt(int)
func (q *QDateTimeEdit) SectionAt(index int) QDateTimeEdit_Section {
	var __rv QDateTimeEdit_Section
	q.Drv(223000,223139,unsafe.Pointer(&index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDateTimeEdit::sectionCount()
func (q *QDateTimeEdit) SectionCount() int {
	var __rv int
	q.Drv(223000,223140,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDateTimeEdit::sectionText(QDateTimeEdit::Section)
func (q *QDateTimeEdit) SectionText(section QDateTimeEdit_Section) string {
	var __rv string
	q.Drv(223000,223141,unsafe.Pointer(&section),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDateTimeEdit::setCalendarPopup(bool)
func (q *QDateTimeEdit) SetCalendarPopup(enable bool)  {
	q.Drv(223000,223142,unsafe.Pointer(&enable),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDateTimeEdit::setCalendarWidget(QCalendarWidget*)
func (q *QDateTimeEdit) SetCalendarWidget(calendarWidget *QCalendarWidget)  {
	q.Drv(223000,223143,Native(calendarWidget),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDateTimeEdit::setCurrentSection(QDateTimeEdit::Section)
func (q *QDateTimeEdit) SetCurrentSection(section QDateTimeEdit_Section)  {
	q.Drv(223000,223144,unsafe.Pointer(&section),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDateTimeEdit::setCurrentSectionIndex(int)
func (q *QDateTimeEdit) SetCurrentSectionIndex(index int)  {
	q.Drv(223000,223145,unsafe.Pointer(&index),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDateTimeEdit::setDate(QDate const&)
func (q *QDateTimeEdit) SetDate(date *QDate)  {
	q.Drv(223000,223146,Native(date),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDateTimeEdit::setDateRange(QDate const&,QDate const&)
func (q *QDateTimeEdit) SetDateRange(min *QDate,max *QDate)  {
	q.Drv(223000,223147,Native(min),Native(max),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDateTimeEdit::setDateTime(QDateTime const&)
func (q *QDateTimeEdit) SetDateTime(dateTime *QDateTime)  {
	q.Drv(223000,223148,Native(dateTime),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDateTimeEdit::setDateTimeRange(QDateTime const&,QDateTime const&)
func (q *QDateTimeEdit) SetDateTimeRange(min *QDateTime,max *QDateTime)  {
	q.Drv(223000,223149,Native(min),Native(max),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDateTimeEdit::setDisplayFormat(QString const&)
func (q *QDateTimeEdit) SetDisplayFormat(format string)  {
	q.Drv(223000,223150,unsafe.Pointer(&format),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDateTimeEdit::setMaximumDate(QDate const&)
func (q *QDateTimeEdit) SetMaximumDate(max *QDate)  {
	q.Drv(223000,223151,Native(max),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDateTimeEdit::setMaximumDateTime(QDateTime const&)
func (q *QDateTimeEdit) SetMaximumDateTime(dt *QDateTime)  {
	q.Drv(223000,223152,Native(dt),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDateTimeEdit::setMaximumTime(QTime const&)
func (q *QDateTimeEdit) SetMaximumTime(max *QTime)  {
	q.Drv(223000,223153,Native(max),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDateTimeEdit::setMinimumDate(QDate const&)
func (q *QDateTimeEdit) SetMinimumDate(min *QDate)  {
	q.Drv(223000,223154,Native(min),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDateTimeEdit::setMinimumDateTime(QDateTime const&)
func (q *QDateTimeEdit) SetMinimumDateTime(dt *QDateTime)  {
	q.Drv(223000,223155,Native(dt),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDateTimeEdit::setMinimumTime(QTime const&)
func (q *QDateTimeEdit) SetMinimumTime(min *QTime)  {
	q.Drv(223000,223156,Native(min),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDateTimeEdit::setSelectedSection(QDateTimeEdit::Section)
func (q *QDateTimeEdit) SetSelectedSection(section QDateTimeEdit_Section)  {
	q.Drv(223000,223157,unsafe.Pointer(&section),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDateTimeEdit::setTime(QTime const&)
func (q *QDateTimeEdit) SetTime(time *QTime)  {
	q.Drv(223000,223158,Native(time),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDateTimeEdit::setTimeRange(QTime const&,QTime const&)
func (q *QDateTimeEdit) SetTimeRange(min *QTime,max *QTime)  {
	q.Drv(223000,223159,Native(min),Native(max),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDateTimeEdit::setTimeSpec(Qt::TimeSpec)
func (q *QDateTimeEdit) SetTimeSpec(spec Qt_TimeSpec)  {
	q.Drv(223000,223160,unsafe.Pointer(&spec),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDateTimeEdit::sizeHint()
func (q *QDateTimeEdit) SizeHint() *QSize {
	var __rv uintptr
	q.Drv(223000,223161,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QDateTimeEdit::stepBy(int)
func (q *QDateTimeEdit) StepBy(steps int)  {
	q.Drv(223000,223162,unsafe.Pointer(&steps),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDateTimeEdit::stepEnabled()
func (q *QDateTimeEdit) StepEnabled() QAbstractSpinBox_StepEnabledFlag {
	var __rv QAbstractSpinBox_StepEnabledFlag
	q.Drv(223000,223163,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDateTimeEdit::textFromDateTime(QDateTime const&)
func (q *QDateTimeEdit) TextFromDateTime(dt *QDateTime) string {
	var __rv string
	q.Drv(223000,223164,Native(dt),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDateTimeEdit::time()
func (q *QDateTimeEdit) Time() *QTime {
	var __rv uintptr
	q.Drv(223000,223165,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTime{}
	_rp.SetDriver(__rv,170000,true)
	return _rp
}	
//QDateTimeEdit::timeSpec()
func (q *QDateTimeEdit) TimeSpec() Qt_TimeSpec {
	var __rv Qt_TimeSpec
	q.Drv(223000,223166,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDateTimeEdit::validate(QString&,int&)
func (q *QDateTimeEdit) Validate(input *string,pos *int) QValidator_State {
	var __rv QValidator_State
	q.Drv(223000,223167,unsafe.Pointer(&input),unsafe.Pointer(&pos),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDateTimeEdit::wheelEvent(QWheelEvent*)
func (q *QDateTimeEdit) WheelEvent(event *QWheelEvent)  {
	q.Drv(223000,223168,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
