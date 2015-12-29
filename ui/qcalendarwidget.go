// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QCalendarWidget_HorizontalHeaderFormat - QCalendarWidget::HorizontalHeaderFormat
type QCalendarWidget_HorizontalHeaderFormat uint32
const (
	QCalendarWidget_NoHorizontalHeader QCalendarWidget_HorizontalHeaderFormat = 0
	QCalendarWidget_SingleLetterDayNames QCalendarWidget_HorizontalHeaderFormat = 1
	QCalendarWidget_ShortDayNames QCalendarWidget_HorizontalHeaderFormat = 2
	QCalendarWidget_LongDayNames QCalendarWidget_HorizontalHeaderFormat = 3
)
//enum QCalendarWidget_VerticalHeaderFormat - QCalendarWidget::VerticalHeaderFormat
type QCalendarWidget_VerticalHeaderFormat uint32
const (
	QCalendarWidget_NoVerticalHeader QCalendarWidget_VerticalHeaderFormat = 0
	QCalendarWidget_ISOWeekNumbers QCalendarWidget_VerticalHeaderFormat = 1
)
//enum QCalendarWidget_SelectionMode - QCalendarWidget::SelectionMode
type QCalendarWidget_SelectionMode uint32
const (
	QCalendarWidget_NoSelection QCalendarWidget_SelectionMode = 0
	QCalendarWidget_SingleSelection QCalendarWidget_SelectionMode = 1
)
//struct QCalendarWidget : QCalendarWidget
type QCalendarWidget struct {
	QWidget
}
func (q *QCalendarWidget) OnClicked(fn func(*QDate)) uintptr {
	var __rv uintptr
	q.Drv(212000,212102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QCalendarWidget) OnActivated(fn func(*QDate)) uintptr {
	var __rv uintptr
	q.Drv(212000,212103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QCalendarWidget) OnCurrentPageChanged(fn func(int,int)) uintptr {
	var __rv uintptr
	q.Drv(212000,212104,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QCalendarWidget) OnSelectionChanged(fn func()) uintptr {
	var __rv uintptr
	q.Drv(212000,212105,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QCalendarWidget::QCalendarWidget()	
func NewCalendarWidget() *QCalendarWidget {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),212000,212106,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QCalendarWidget{}
	_p.SetDriver(__rv,212000,false)
	return _p
} 
//QCalendarWidget::QCalendarWidget(QWidget*)	
func NewCalendarWidgetWithParent(parent QWidgetInterface) *QCalendarWidget {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),212000,212107,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QCalendarWidget{}
	_p.SetDriver(__rv,212000,false)
	return _p
} 
//QCalendarWidget::dateEditAcceptDelay()
func (q *QCalendarWidget) DateEditAcceptDelay() int {
	var __rv int
	q.Drv(212000,212108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QCalendarWidget::dateTextFormat()
func (q *QCalendarWidget) DateTextFormat() map[*QDate]*QTextCharFormat {
	__rv := make(map[*QDate]*QTextCharFormat)
	q.Drv(212000,212109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QCalendarWidget::dateTextFormat(QDate const&)
func (q *QCalendarWidget) DateTextFormatWithDate(date *QDate) *QTextCharFormat {
	var __rv uintptr
	q.Drv(212000,212110,Native(date),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextCharFormat{}
	_rp.SetDriver(__rv,142000,true)
	return _rp
}	
//QCalendarWidget::event(QEvent*)
func (q *QCalendarWidget) Event(event *QEvent) bool {
	var __rv bool
	q.Drv(212000,212111,Native(event),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QCalendarWidget::eventFilter(QObject*,QEvent*)
func (q *QCalendarWidget) EventFilter(watched QObjectInterface,event *QEvent) bool {
	var __rv bool
	q.Drv(212000,212112,Native(watched),Native(event),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QCalendarWidget::firstDayOfWeek()
func (q *QCalendarWidget) FirstDayOfWeek() Qt_DayOfWeek {
	var __rv Qt_DayOfWeek
	q.Drv(212000,212113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QCalendarWidget::headerTextFormat()
func (q *QCalendarWidget) HeaderTextFormat() *QTextCharFormat {
	var __rv uintptr
	q.Drv(212000,212114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextCharFormat{}
	_rp.SetDriver(__rv,142000,true)
	return _rp
}	
//QCalendarWidget::horizontalHeaderFormat()
func (q *QCalendarWidget) HorizontalHeaderFormat() QCalendarWidget_HorizontalHeaderFormat {
	var __rv QCalendarWidget_HorizontalHeaderFormat
	q.Drv(212000,212115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QCalendarWidget::isDateEditEnabled()
func (q *QCalendarWidget) IsDateEditEnabled() bool {
	var __rv bool
	q.Drv(212000,212116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QCalendarWidget::isGridVisible()
func (q *QCalendarWidget) IsGridVisible() bool {
	var __rv bool
	q.Drv(212000,212117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QCalendarWidget::isNavigationBarVisible()
func (q *QCalendarWidget) IsNavigationBarVisible() bool {
	var __rv bool
	q.Drv(212000,212118,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QCalendarWidget::keyPressEvent(QKeyEvent*)
func (q *QCalendarWidget) KeyPressEvent(event *QKeyEvent)  {
	q.Drv(212000,212119,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCalendarWidget::maximumDate()
func (q *QCalendarWidget) MaximumDate() *QDate {
	var __rv uintptr
	q.Drv(212000,212120,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QDate{}
	_rp.SetDriver(__rv,19000,true)
	return _rp
}	
//QCalendarWidget::minimumDate()
func (q *QCalendarWidget) MinimumDate() *QDate {
	var __rv uintptr
	q.Drv(212000,212121,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QDate{}
	_rp.SetDriver(__rv,19000,true)
	return _rp
}	
//QCalendarWidget::minimumSizeHint()
func (q *QCalendarWidget) MinimumSizeHint() *QSize {
	var __rv uintptr
	q.Drv(212000,212122,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QCalendarWidget::monthShown()
func (q *QCalendarWidget) MonthShown() int {
	var __rv int
	q.Drv(212000,212123,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QCalendarWidget::mousePressEvent(QMouseEvent*)
func (q *QCalendarWidget) MousePressEvent(event *QMouseEvent)  {
	q.Drv(212000,212124,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCalendarWidget::paintCell(QPainter*,QRect const&,QDate const&)
func (q *QCalendarWidget) PaintCell(painter *QPainter,rect *QRect,date *QDate)  {
	q.Drv(212000,212125,Native(painter),Native(rect),Native(date),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCalendarWidget::resizeEvent(QResizeEvent*)
func (q *QCalendarWidget) ResizeEvent(event *QResizeEvent)  {
	q.Drv(212000,212126,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCalendarWidget::selectedDate()
func (q *QCalendarWidget) SelectedDate() *QDate {
	var __rv uintptr
	q.Drv(212000,212127,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QDate{}
	_rp.SetDriver(__rv,19000,true)
	return _rp
}	
//QCalendarWidget::selectionMode()
func (q *QCalendarWidget) SelectionMode() QCalendarWidget_SelectionMode {
	var __rv QCalendarWidget_SelectionMode
	q.Drv(212000,212128,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QCalendarWidget::setCurrentPage(int,int)
func (q *QCalendarWidget) SetCurrentPage(year int,month int)  {
	q.Drv(212000,212129,unsafe.Pointer(&year),unsafe.Pointer(&month),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCalendarWidget::setDateEditAcceptDelay(int)
func (q *QCalendarWidget) SetDateEditAcceptDelay(delay int)  {
	q.Drv(212000,212130,unsafe.Pointer(&delay),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCalendarWidget::setDateEditEnabled(bool)
func (q *QCalendarWidget) SetDateEditEnabled(enable bool)  {
	q.Drv(212000,212131,unsafe.Pointer(&enable),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCalendarWidget::setDateRange(QDate const&,QDate const&)
func (q *QCalendarWidget) SetDateRange(min *QDate,max *QDate)  {
	q.Drv(212000,212132,Native(min),Native(max),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCalendarWidget::setDateTextFormat(QDate const&,QTextCharFormat const&)
func (q *QCalendarWidget) SetDateTextFormat(date *QDate,format *QTextCharFormat)  {
	q.Drv(212000,212133,Native(date),Native(format),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCalendarWidget::setFirstDayOfWeek(Qt::DayOfWeek)
func (q *QCalendarWidget) SetFirstDayOfWeek(dayOfWeek Qt_DayOfWeek)  {
	q.Drv(212000,212134,unsafe.Pointer(&dayOfWeek),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCalendarWidget::setGridVisible(bool)
func (q *QCalendarWidget) SetGridVisible(show bool)  {
	q.Drv(212000,212135,unsafe.Pointer(&show),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCalendarWidget::setHeaderTextFormat(QTextCharFormat const&)
func (q *QCalendarWidget) SetHeaderTextFormat(format *QTextCharFormat)  {
	q.Drv(212000,212136,Native(format),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCalendarWidget::setHorizontalHeaderFormat(QCalendarWidget::HorizontalHeaderFormat)
func (q *QCalendarWidget) SetHorizontalHeaderFormat(format QCalendarWidget_HorizontalHeaderFormat)  {
	q.Drv(212000,212137,unsafe.Pointer(&format),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCalendarWidget::setMaximumDate(QDate const&)
func (q *QCalendarWidget) SetMaximumDate(date *QDate)  {
	q.Drv(212000,212138,Native(date),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCalendarWidget::setMinimumDate(QDate const&)
func (q *QCalendarWidget) SetMinimumDate(date *QDate)  {
	q.Drv(212000,212139,Native(date),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCalendarWidget::setNavigationBarVisible(bool)
func (q *QCalendarWidget) SetNavigationBarVisible(visible bool)  {
	q.Drv(212000,212140,unsafe.Pointer(&visible),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCalendarWidget::setSelectedDate(QDate const&)
func (q *QCalendarWidget) SetSelectedDate(date *QDate)  {
	q.Drv(212000,212141,Native(date),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCalendarWidget::setSelectionMode(QCalendarWidget::SelectionMode)
func (q *QCalendarWidget) SetSelectionMode(mode QCalendarWidget_SelectionMode)  {
	q.Drv(212000,212142,unsafe.Pointer(&mode),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCalendarWidget::setVerticalHeaderFormat(QCalendarWidget::VerticalHeaderFormat)
func (q *QCalendarWidget) SetVerticalHeaderFormat(format QCalendarWidget_VerticalHeaderFormat)  {
	q.Drv(212000,212143,unsafe.Pointer(&format),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCalendarWidget::setWeekdayTextFormat(Qt::DayOfWeek,QTextCharFormat const&)
func (q *QCalendarWidget) SetWeekdayTextFormat(dayOfWeek Qt_DayOfWeek,format *QTextCharFormat)  {
	q.Drv(212000,212144,unsafe.Pointer(&dayOfWeek),Native(format),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCalendarWidget::showNextMonth()
func (q *QCalendarWidget) ShowNextMonth()  {
	q.Drv(212000,212145,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCalendarWidget::showNextYear()
func (q *QCalendarWidget) ShowNextYear()  {
	q.Drv(212000,212146,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCalendarWidget::showPreviousMonth()
func (q *QCalendarWidget) ShowPreviousMonth()  {
	q.Drv(212000,212147,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCalendarWidget::showPreviousYear()
func (q *QCalendarWidget) ShowPreviousYear()  {
	q.Drv(212000,212148,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCalendarWidget::showSelectedDate()
func (q *QCalendarWidget) ShowSelectedDate()  {
	q.Drv(212000,212149,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCalendarWidget::showToday()
func (q *QCalendarWidget) ShowToday()  {
	q.Drv(212000,212150,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCalendarWidget::sizeHint()
func (q *QCalendarWidget) SizeHint() *QSize {
	var __rv uintptr
	q.Drv(212000,212151,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QCalendarWidget::updateCell(QDate const&)
func (q *QCalendarWidget) UpdateCell(date *QDate)  {
	q.Drv(212000,212152,Native(date),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCalendarWidget::updateCells()
func (q *QCalendarWidget) UpdateCells()  {
	q.Drv(212000,212153,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCalendarWidget::verticalHeaderFormat()
func (q *QCalendarWidget) VerticalHeaderFormat() QCalendarWidget_VerticalHeaderFormat {
	var __rv QCalendarWidget_VerticalHeaderFormat
	q.Drv(212000,212154,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QCalendarWidget::weekdayTextFormat(Qt::DayOfWeek)
func (q *QCalendarWidget) WeekdayTextFormat(dayOfWeek Qt_DayOfWeek) *QTextCharFormat {
	var __rv uintptr
	q.Drv(212000,212155,unsafe.Pointer(&dayOfWeek),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextCharFormat{}
	_rp.SetDriver(__rv,142000,true)
	return _rp
}	
//QCalendarWidget::yearShown()
func (q *QCalendarWidget) YearShown() int {
	var __rv int
	q.Drv(212000,212156,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
