// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QToolBox : QToolBox
type QToolBox struct {
	QFrame
}
func (q *QToolBox) OnCurrentChanged(fn func(int)) uintptr {
	var __rv uintptr
	q.Drv(382000,382102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QToolBox::QToolBox()	
func NewToolBox() *QToolBox {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),382000,382103,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QToolBox{}
	_p.SetDriver(__rv,382000,false)
	return _p
} 
//QToolBox::QToolBox(QWidget*,QFlags<Qt::WindowType>)	
func NewToolBoxWithParentFlags(parent QWidgetInterface,f Qt_WindowType) *QToolBox {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),382000,382104,Native(parent),unsafe.Pointer(&f),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QToolBox{}
	_p.SetDriver(__rv,382000,false)
	return _p
} 
//QToolBox::addItem(QWidget*,QString const&)
func (q *QToolBox) AddItemWithWidgetText(widget QWidgetInterface,text string) int {
	var __rv int
	q.Drv(382000,382105,Native(widget),unsafe.Pointer(&text),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QToolBox::addItem(QWidget*,QIcon const&,QString const&)
func (q *QToolBox) AddItemWithWidgetIconText(widget QWidgetInterface,icon *QIcon,text string) int {
	var __rv int
	q.Drv(382000,382106,Native(widget),Native(icon),unsafe.Pointer(&text),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QToolBox::changeEvent(QEvent*)
func (q *QToolBox) ChangeEvent(value *QEvent)  {
	q.Drv(382000,382107,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QToolBox::count()
func (q *QToolBox) Count() int {
	var __rv int
	q.Drv(382000,382108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QToolBox::currentIndex()
func (q *QToolBox) CurrentIndex() int {
	var __rv int
	q.Drv(382000,382109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QToolBox::currentWidget()
func (q *QToolBox) CurrentWidget() *QWidget {
	var __rv uintptr
	q.Drv(382000,382110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QToolBox::event(QEvent*)
func (q *QToolBox) Event(e *QEvent) bool {
	var __rv bool
	q.Drv(382000,382111,Native(e),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QToolBox::indexOf(QWidget*)
func (q *QToolBox) IndexOf(widget QWidgetInterface) int {
	var __rv int
	q.Drv(382000,382112,Native(widget),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QToolBox::insertItem(int,QWidget*,QString const&)
func (q *QToolBox) InsertItemWithIndexWidgetText(index int,widget QWidgetInterface,text string) int {
	var __rv int
	q.Drv(382000,382113,unsafe.Pointer(&index),Native(widget),unsafe.Pointer(&text),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QToolBox::insertItem(int,QWidget*,QIcon const&,QString const&)
func (q *QToolBox) InsertItemWithIndexWidgetIconText(index int,widget QWidgetInterface,icon *QIcon,text string) int {
	var __rv int
	q.Drv(382000,382114,unsafe.Pointer(&index),Native(widget),Native(icon),unsafe.Pointer(&text),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QToolBox::isItemEnabled(int)
func (q *QToolBox) IsItemEnabled(index int) bool {
	var __rv bool
	q.Drv(382000,382115,unsafe.Pointer(&index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QToolBox::itemIcon(int)
func (q *QToolBox) ItemIcon(index int) *QIcon {
	var __rv uintptr
	q.Drv(382000,382116,unsafe.Pointer(&index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QIcon{}
	_rp.SetDriver(__rv,51000,true)
	return _rp
}	
//QToolBox::itemInserted(int)
func (q *QToolBox) ItemInserted(index int)  {
	q.Drv(382000,382117,unsafe.Pointer(&index),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QToolBox::itemRemoved(int)
func (q *QToolBox) ItemRemoved(index int)  {
	q.Drv(382000,382118,unsafe.Pointer(&index),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QToolBox::itemText(int)
func (q *QToolBox) ItemText(index int) string {
	var __rv string
	q.Drv(382000,382119,unsafe.Pointer(&index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QToolBox::itemToolTip(int)
func (q *QToolBox) ItemToolTip(index int) string {
	var __rv string
	q.Drv(382000,382120,unsafe.Pointer(&index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QToolBox::removeItem(int)
func (q *QToolBox) RemoveItem(index int)  {
	q.Drv(382000,382121,unsafe.Pointer(&index),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QToolBox::setCurrentIndex(int)
func (q *QToolBox) SetCurrentIndex(index int)  {
	q.Drv(382000,382122,unsafe.Pointer(&index),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QToolBox::setCurrentWidget(QWidget*)
func (q *QToolBox) SetCurrentWidget(widget QWidgetInterface)  {
	q.Drv(382000,382123,Native(widget),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QToolBox::setItemEnabled(int,bool)
func (q *QToolBox) SetItemEnabled(index int,enabled bool)  {
	q.Drv(382000,382124,unsafe.Pointer(&index),unsafe.Pointer(&enabled),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QToolBox::setItemIcon(int,QIcon const&)
func (q *QToolBox) SetItemIcon(index int,icon *QIcon)  {
	q.Drv(382000,382125,unsafe.Pointer(&index),Native(icon),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QToolBox::setItemText(int,QString const&)
func (q *QToolBox) SetItemText(index int,text string)  {
	q.Drv(382000,382126,unsafe.Pointer(&index),unsafe.Pointer(&text),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QToolBox::setItemToolTip(int,QString const&)
func (q *QToolBox) SetItemToolTip(index int,toolTip string)  {
	q.Drv(382000,382127,unsafe.Pointer(&index),unsafe.Pointer(&toolTip),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QToolBox::showEvent(QShowEvent*)
func (q *QToolBox) ShowEvent(e *QShowEvent)  {
	q.Drv(382000,382128,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QToolBox::widget(int)
func (q *QToolBox) Widget(index int) *QWidget {
	var __rv uintptr
	q.Drv(382000,382129,unsafe.Pointer(&index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
