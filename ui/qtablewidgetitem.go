// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QTableWidgetItem_ItemType - QTableWidgetItem::ItemType
type QTableWidgetItem_ItemType uint32
const (
	QTableWidgetItem_Type QTableWidgetItem_ItemType = 0
	QTableWidgetItem_UserType QTableWidgetItem_ItemType = 1000
)
//struct QTableWidgetItem : QTableWidgetItem
type QTableWidgetItem struct {
	BaseDrv
}
//QTableWidgetItem::QTableWidgetItem()	
func NewTableWidgetItem() *QTableWidgetItem {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),134000,134102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTableWidgetItem{}
	_p.SetDriver(__rv,134000,true)
	return _p
} 
//QTableWidgetItem::QTableWidgetItem(QTableWidgetItem const&)	
func NewTableWidgetItemCopy(other *QTableWidgetItem) *QTableWidgetItem {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),134000,134103,Native(other),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTableWidgetItem{}
	_p.SetDriver(__rv,134000,true)
	return _p
} 
//QTableWidgetItem::QTableWidgetItem(int)	
func NewTableWidgetItemWithType(_type int) *QTableWidgetItem {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),134000,134104,unsafe.Pointer(&_type),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTableWidgetItem{}
	_p.SetDriver(__rv,134000,true)
	return _p
} 
//QTableWidgetItem::QTableWidgetItem(QString const&,int)	
func NewTableWidgetItemWithTextType(text string,_type int) *QTableWidgetItem {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),134000,134105,unsafe.Pointer(&text),unsafe.Pointer(&_type),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTableWidgetItem{}
	_p.SetDriver(__rv,134000,true)
	return _p
} 
//QTableWidgetItem::QTableWidgetItem(QIcon const&,QString const&,int)	
func NewTableWidgetItemWithIconTextType(icon *QIcon,text string,_type int) *QTableWidgetItem {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),134000,134106,Native(icon),unsafe.Pointer(&text),unsafe.Pointer(&_type),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTableWidgetItem{}
	_p.SetDriver(__rv,134000,true)
	return _p
} 
//QTableWidgetItem::background()
func (q *QTableWidgetItem) Background() *QBrush {
	var __rv uintptr
	q.Drv(134000,134107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QBrush{}
	_rp.SetDriver(__rv,9000,true)
	return _rp
}	
//QTableWidgetItem::backgroundColor()
func (q *QTableWidgetItem) BackgroundColor() *QColor {
	var __rv uintptr
	q.Drv(134000,134108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QColor{}
	_rp.SetDriver(__rv,13000,true)
	return _rp
}	
//QTableWidgetItem::checkState()
func (q *QTableWidgetItem) CheckState() Qt_CheckState {
	var __rv Qt_CheckState
	q.Drv(134000,134109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTableWidgetItem::clone()
func (q *QTableWidgetItem) Clone() *QTableWidgetItem {
	var __rv uintptr
	q.Drv(134000,134110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTableWidgetItem{}
	_rp.SetDriver(__rv,134000,true)
	return _rp
}	
//QTableWidgetItem::column()
func (q *QTableWidgetItem) Column() int {
	var __rv int
	q.Drv(134000,134111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTableWidgetItem::data(int)
func (q *QTableWidgetItem) Data(role int) *QVariant {
	var __rv uintptr
	q.Drv(134000,134112,unsafe.Pointer(&role),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVariant{}
	_rp.SetDriver(__rv,182000,true)
	return _rp
}	
//QTableWidgetItem::flags()
func (q *QTableWidgetItem) Flags() Qt_ItemFlag {
	var __rv Qt_ItemFlag
	q.Drv(134000,134113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTableWidgetItem::font()
func (q *QTableWidgetItem) Font() *QFont {
	var __rv uintptr
	q.Drv(134000,134114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QFont{}
	_rp.SetDriver(__rv,37000,true)
	return _rp
}	
//QTableWidgetItem::foreground()
func (q *QTableWidgetItem) Foreground() *QBrush {
	var __rv uintptr
	q.Drv(134000,134115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QBrush{}
	_rp.SetDriver(__rv,9000,true)
	return _rp
}	
//QTableWidgetItem::icon()
func (q *QTableWidgetItem) Icon() *QIcon {
	var __rv uintptr
	q.Drv(134000,134116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QIcon{}
	_rp.SetDriver(__rv,51000,true)
	return _rp
}	
//QTableWidgetItem::isSelected()
func (q *QTableWidgetItem) IsSelected() bool {
	var __rv bool
	q.Drv(134000,134117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTableWidgetItem::row()
func (q *QTableWidgetItem) Row() int {
	var __rv int
	q.Drv(134000,134118,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTableWidgetItem::setBackground(QBrush const&)
func (q *QTableWidgetItem) SetBackground(brush *QBrush)  {
	q.Drv(134000,134119,Native(brush),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableWidgetItem::setBackgroundColor(QColor const&)
func (q *QTableWidgetItem) SetBackgroundColor(color *QColor)  {
	q.Drv(134000,134120,Native(color),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableWidgetItem::setCheckState(Qt::CheckState)
func (q *QTableWidgetItem) SetCheckState(state Qt_CheckState)  {
	q.Drv(134000,134121,unsafe.Pointer(&state),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableWidgetItem::setData(int,QVariant const&)
func (q *QTableWidgetItem) SetData(role int,value *QVariant)  {
	q.Drv(134000,134122,unsafe.Pointer(&role),Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableWidgetItem::setFlags(QFlags<Qt::ItemFlag>)
func (q *QTableWidgetItem) SetFlags(flags Qt_ItemFlag)  {
	q.Drv(134000,134123,unsafe.Pointer(&flags),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableWidgetItem::setFont(QFont const&)
func (q *QTableWidgetItem) SetFont(font *QFont)  {
	q.Drv(134000,134124,Native(font),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableWidgetItem::setForeground(QBrush const&)
func (q *QTableWidgetItem) SetForeground(brush *QBrush)  {
	q.Drv(134000,134125,Native(brush),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableWidgetItem::setIcon(QIcon const&)
func (q *QTableWidgetItem) SetIcon(icon *QIcon)  {
	q.Drv(134000,134126,Native(icon),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableWidgetItem::setSelected(bool)
func (q *QTableWidgetItem) SetSelected(_select bool)  {
	q.Drv(134000,134127,unsafe.Pointer(&_select),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableWidgetItem::setSizeHint(QSize const&)
func (q *QTableWidgetItem) SetSizeHint(size *QSize)  {
	q.Drv(134000,134128,Native(size),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableWidgetItem::setStatusTip(QString const&)
func (q *QTableWidgetItem) SetStatusTip(statusTip string)  {
	q.Drv(134000,134129,unsafe.Pointer(&statusTip),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableWidgetItem::setText(QString const&)
func (q *QTableWidgetItem) SetText(text string)  {
	q.Drv(134000,134130,unsafe.Pointer(&text),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableWidgetItem::setTextAlignment(int)
func (q *QTableWidgetItem) SetTextAlignment(alignment int)  {
	q.Drv(134000,134131,unsafe.Pointer(&alignment),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableWidgetItem::setTextColor(QColor const&)
func (q *QTableWidgetItem) SetTextColor(color *QColor)  {
	q.Drv(134000,134132,Native(color),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableWidgetItem::setToolTip(QString const&)
func (q *QTableWidgetItem) SetToolTip(toolTip string)  {
	q.Drv(134000,134133,unsafe.Pointer(&toolTip),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableWidgetItem::setWhatsThis(QString const&)
func (q *QTableWidgetItem) SetWhatsThis(whatsThis string)  {
	q.Drv(134000,134134,unsafe.Pointer(&whatsThis),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableWidgetItem::sizeHint()
func (q *QTableWidgetItem) SizeHint() *QSize {
	var __rv uintptr
	q.Drv(134000,134135,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QTableWidgetItem::statusTip()
func (q *QTableWidgetItem) StatusTip() string {
	var __rv string
	q.Drv(134000,134136,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTableWidgetItem::tableWidget()
func (q *QTableWidgetItem) TableWidget() *QTableWidget {
	var __rv uintptr
	q.Drv(134000,134137,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTableWidget{}
	_rp.SetDriver(__rv,366000,false)
	return _rp
}	
//QTableWidgetItem::text()
func (q *QTableWidgetItem) Text() string {
	var __rv string
	q.Drv(134000,134138,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTableWidgetItem::textAlignment()
func (q *QTableWidgetItem) TextAlignment() int {
	var __rv int
	q.Drv(134000,134139,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTableWidgetItem::textColor()
func (q *QTableWidgetItem) TextColor() *QColor {
	var __rv uintptr
	q.Drv(134000,134140,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QColor{}
	_rp.SetDriver(__rv,13000,true)
	return _rp
}	
//QTableWidgetItem::toolTip()
func (q *QTableWidgetItem) ToolTip() string {
	var __rv string
	q.Drv(134000,134141,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTableWidgetItem::type()
func (q *QTableWidgetItem) Type() int {
	var __rv int
	q.Drv(134000,134142,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTableWidgetItem::whatsThis()
func (q *QTableWidgetItem) WhatsThis() string {
	var __rv string
	q.Drv(134000,134143,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
