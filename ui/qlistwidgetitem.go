// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QListWidgetItem_ItemType - QListWidgetItem::ItemType
type QListWidgetItem_ItemType uint32
const (
	QListWidgetItem_Type QListWidgetItem_ItemType = 0
	QListWidgetItem_UserType QListWidgetItem_ItemType = 1000
)
//struct QListWidgetItem : QListWidgetItem
type QListWidgetItem struct {
	BaseDrv
}
//QListWidgetItem::QListWidgetItem()	
func NewListWidgetItem() *QListWidgetItem {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),71000,71102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QListWidgetItem{}
	_p.SetDriver(__rv,71000,true)
	return _p
} 
//QListWidgetItem::QListWidgetItem(QListWidgetItem const&)	
func NewListWidgetItemCopy(other *QListWidgetItem) *QListWidgetItem {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),71000,71103,Native(other),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QListWidgetItem{}
	_p.SetDriver(__rv,71000,true)
	return _p
} 
//QListWidgetItem::QListWidgetItem(QListWidget*,int)	
func NewListWidgetItemWithViewType(view *QListWidget,_type int) *QListWidgetItem {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),71000,71104,Native(view),unsafe.Pointer(&_type),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QListWidgetItem{}
	_p.SetDriver(__rv,71000,true)
	return _p
} 
//QListWidgetItem::QListWidgetItem(QString const&,QListWidget*,int)	
func NewListWidgetItemWithTextViewType(text string,view *QListWidget,_type int) *QListWidgetItem {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),71000,71105,unsafe.Pointer(&text),Native(view),unsafe.Pointer(&_type),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QListWidgetItem{}
	_p.SetDriver(__rv,71000,true)
	return _p
} 
//QListWidgetItem::QListWidgetItem(QIcon const&,QString const&,QListWidget*,int)	
func NewListWidgetItemWithIconTextViewType(icon *QIcon,text string,view *QListWidget,_type int) *QListWidgetItem {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),71000,71106,Native(icon),unsafe.Pointer(&text),Native(view),unsafe.Pointer(&_type),nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QListWidgetItem{}
	_p.SetDriver(__rv,71000,true)
	return _p
} 
//QListWidgetItem::background()
func (q *QListWidgetItem) Background() *QBrush {
	var __rv uintptr
	q.Drv(71000,71107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QBrush{}
	_rp.SetDriver(__rv,9000,true)
	return _rp
}	
//QListWidgetItem::backgroundColor()
func (q *QListWidgetItem) BackgroundColor() *QColor {
	var __rv uintptr
	q.Drv(71000,71108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QColor{}
	_rp.SetDriver(__rv,13000,true)
	return _rp
}	
//QListWidgetItem::checkState()
func (q *QListWidgetItem) CheckState() Qt_CheckState {
	var __rv Qt_CheckState
	q.Drv(71000,71109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QListWidgetItem::clone()
func (q *QListWidgetItem) Clone() *QListWidgetItem {
	var __rv uintptr
	q.Drv(71000,71110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QListWidgetItem{}
	_rp.SetDriver(__rv,71000,true)
	return _rp
}	
//QListWidgetItem::data(int)
func (q *QListWidgetItem) Data(role int) *QVariant {
	var __rv uintptr
	q.Drv(71000,71111,unsafe.Pointer(&role),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVariant{}
	_rp.SetDriver(__rv,182000,true)
	return _rp
}	
//QListWidgetItem::flags()
func (q *QListWidgetItem) Flags() Qt_ItemFlag {
	var __rv Qt_ItemFlag
	q.Drv(71000,71112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QListWidgetItem::font()
func (q *QListWidgetItem) Font() *QFont {
	var __rv uintptr
	q.Drv(71000,71113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QFont{}
	_rp.SetDriver(__rv,37000,true)
	return _rp
}	
//QListWidgetItem::foreground()
func (q *QListWidgetItem) Foreground() *QBrush {
	var __rv uintptr
	q.Drv(71000,71114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QBrush{}
	_rp.SetDriver(__rv,9000,true)
	return _rp
}	
//QListWidgetItem::icon()
func (q *QListWidgetItem) Icon() *QIcon {
	var __rv uintptr
	q.Drv(71000,71115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QIcon{}
	_rp.SetDriver(__rv,51000,true)
	return _rp
}	
//QListWidgetItem::isHidden()
func (q *QListWidgetItem) IsHidden() bool {
	var __rv bool
	q.Drv(71000,71116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QListWidgetItem::isSelected()
func (q *QListWidgetItem) IsSelected() bool {
	var __rv bool
	q.Drv(71000,71117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QListWidgetItem::listWidget()
func (q *QListWidgetItem) ListWidget() *QListWidget {
	var __rv uintptr
	q.Drv(71000,71118,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QListWidget{}
	_rp.SetDriver(__rv,304000,false)
	return _rp
}	
//QListWidgetItem::setBackground(QBrush const&)
func (q *QListWidgetItem) SetBackground(brush *QBrush)  {
	q.Drv(71000,71119,Native(brush),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListWidgetItem::setBackgroundColor(QColor const&)
func (q *QListWidgetItem) SetBackgroundColor(color *QColor)  {
	q.Drv(71000,71120,Native(color),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListWidgetItem::setCheckState(Qt::CheckState)
func (q *QListWidgetItem) SetCheckState(state Qt_CheckState)  {
	q.Drv(71000,71121,unsafe.Pointer(&state),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListWidgetItem::setData(int,QVariant const&)
func (q *QListWidgetItem) SetData(role int,value *QVariant)  {
	q.Drv(71000,71122,unsafe.Pointer(&role),Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListWidgetItem::setFlags(QFlags<Qt::ItemFlag>)
func (q *QListWidgetItem) SetFlags(flags Qt_ItemFlag)  {
	q.Drv(71000,71123,unsafe.Pointer(&flags),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListWidgetItem::setFont(QFont const&)
func (q *QListWidgetItem) SetFont(font *QFont)  {
	q.Drv(71000,71124,Native(font),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListWidgetItem::setForeground(QBrush const&)
func (q *QListWidgetItem) SetForeground(brush *QBrush)  {
	q.Drv(71000,71125,Native(brush),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListWidgetItem::setHidden(bool)
func (q *QListWidgetItem) SetHidden(hide bool)  {
	q.Drv(71000,71126,unsafe.Pointer(&hide),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListWidgetItem::setIcon(QIcon const&)
func (q *QListWidgetItem) SetIcon(icon *QIcon)  {
	q.Drv(71000,71127,Native(icon),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListWidgetItem::setSelected(bool)
func (q *QListWidgetItem) SetSelected(_select bool)  {
	q.Drv(71000,71128,unsafe.Pointer(&_select),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListWidgetItem::setSizeHint(QSize const&)
func (q *QListWidgetItem) SetSizeHint(size *QSize)  {
	q.Drv(71000,71129,Native(size),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListWidgetItem::setStatusTip(QString const&)
func (q *QListWidgetItem) SetStatusTip(statusTip string)  {
	q.Drv(71000,71130,unsafe.Pointer(&statusTip),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListWidgetItem::setText(QString const&)
func (q *QListWidgetItem) SetText(text string)  {
	q.Drv(71000,71131,unsafe.Pointer(&text),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListWidgetItem::setTextAlignment(int)
func (q *QListWidgetItem) SetTextAlignment(alignment int)  {
	q.Drv(71000,71132,unsafe.Pointer(&alignment),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListWidgetItem::setTextColor(QColor const&)
func (q *QListWidgetItem) SetTextColor(color *QColor)  {
	q.Drv(71000,71133,Native(color),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListWidgetItem::setToolTip(QString const&)
func (q *QListWidgetItem) SetToolTip(toolTip string)  {
	q.Drv(71000,71134,unsafe.Pointer(&toolTip),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListWidgetItem::setWhatsThis(QString const&)
func (q *QListWidgetItem) SetWhatsThis(whatsThis string)  {
	q.Drv(71000,71135,unsafe.Pointer(&whatsThis),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListWidgetItem::sizeHint()
func (q *QListWidgetItem) SizeHint() *QSize {
	var __rv uintptr
	q.Drv(71000,71136,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QListWidgetItem::statusTip()
func (q *QListWidgetItem) StatusTip() string {
	var __rv string
	q.Drv(71000,71137,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QListWidgetItem::text()
func (q *QListWidgetItem) Text() string {
	var __rv string
	q.Drv(71000,71138,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QListWidgetItem::textAlignment()
func (q *QListWidgetItem) TextAlignment() int {
	var __rv int
	q.Drv(71000,71139,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QListWidgetItem::textColor()
func (q *QListWidgetItem) TextColor() *QColor {
	var __rv uintptr
	q.Drv(71000,71140,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QColor{}
	_rp.SetDriver(__rv,13000,true)
	return _rp
}	
//QListWidgetItem::toolTip()
func (q *QListWidgetItem) ToolTip() string {
	var __rv string
	q.Drv(71000,71141,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QListWidgetItem::type()
func (q *QListWidgetItem) Type() int {
	var __rv int
	q.Drv(71000,71142,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QListWidgetItem::whatsThis()
func (q *QListWidgetItem) WhatsThis() string {
	var __rv string
	q.Drv(71000,71143,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
