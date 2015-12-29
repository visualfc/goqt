// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QTreeWidgetItem_ItemType - QTreeWidgetItem::ItemType
type QTreeWidgetItem_ItemType uint32
const (
	QTreeWidgetItem_Type QTreeWidgetItem_ItemType = 0
	QTreeWidgetItem_UserType QTreeWidgetItem_ItemType = 1000
)
//enum QTreeWidgetItem_ChildIndicatorPolicy - QTreeWidgetItem::ChildIndicatorPolicy
type QTreeWidgetItem_ChildIndicatorPolicy uint32
const (
	QTreeWidgetItem_ShowIndicator QTreeWidgetItem_ChildIndicatorPolicy = 0
	QTreeWidgetItem_DontShowIndicator QTreeWidgetItem_ChildIndicatorPolicy = 1
	QTreeWidgetItem_DontShowIndicatorWhenChildless QTreeWidgetItem_ChildIndicatorPolicy = 2
)
//struct QTreeWidgetItem : QTreeWidgetItem
type QTreeWidgetItem struct {
	BaseDrv
}
//QTreeWidgetItem::QTreeWidgetItem()	
func NewTreeWidgetItem() *QTreeWidgetItem {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),177000,177102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTreeWidgetItem{}
	_p.SetDriver(__rv,177000,true)
	return _p
} 
//QTreeWidgetItem::QTreeWidgetItem(QTreeWidgetItem const&)	
func NewTreeWidgetItemCopy(other *QTreeWidgetItem) *QTreeWidgetItem {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),177000,177103,Native(other),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTreeWidgetItem{}
	_p.SetDriver(__rv,177000,true)
	return _p
} 
//QTreeWidgetItem::QTreeWidgetItem(int)	
func NewTreeWidgetItemWithType(_type int) *QTreeWidgetItem {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),177000,177104,unsafe.Pointer(&_type),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTreeWidgetItem{}
	_p.SetDriver(__rv,177000,true)
	return _p
} 
//QTreeWidgetItem::QTreeWidgetItem(QStringList const&,int)	
func NewTreeWidgetItemWithStringsType(strings []string,_type int) *QTreeWidgetItem {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),177000,177105,unsafe.Pointer(&strings),unsafe.Pointer(&_type),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTreeWidgetItem{}
	_p.SetDriver(__rv,177000,true)
	return _p
} 
//QTreeWidgetItem::QTreeWidgetItem(QTreeWidget*,int)	
func NewTreeWidgetItemWithViewType(view *QTreeWidget,_type int) *QTreeWidgetItem {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),177000,177106,Native(view),unsafe.Pointer(&_type),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTreeWidgetItem{}
	_p.SetDriver(__rv,177000,true)
	return _p
} 
//QTreeWidgetItem::QTreeWidgetItem(QTreeWidgetItem*,int)	
func NewTreeWidgetItemWithParentType(parent *QTreeWidgetItem,_type int) *QTreeWidgetItem {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),177000,177107,Native(parent),unsafe.Pointer(&_type),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTreeWidgetItem{}
	_p.SetDriver(__rv,177000,true)
	return _p
} 
//QTreeWidgetItem::QTreeWidgetItem(QTreeWidget*,QStringList const&,int)	
func NewTreeWidgetItemWithViewStringsType(view *QTreeWidget,strings []string,_type int) *QTreeWidgetItem {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),177000,177108,Native(view),unsafe.Pointer(&strings),unsafe.Pointer(&_type),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTreeWidgetItem{}
	_p.SetDriver(__rv,177000,true)
	return _p
} 
//QTreeWidgetItem::QTreeWidgetItem(QTreeWidget*,QTreeWidgetItem*,int)	
func NewTreeWidgetItemWithViewAfterType(view *QTreeWidget,after *QTreeWidgetItem,_type int) *QTreeWidgetItem {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),177000,177109,Native(view),Native(after),unsafe.Pointer(&_type),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTreeWidgetItem{}
	_p.SetDriver(__rv,177000,true)
	return _p
} 
//QTreeWidgetItem::QTreeWidgetItem(QTreeWidgetItem*,QStringList const&,int)	
func NewTreeWidgetItemWithParentStringsType(parent *QTreeWidgetItem,strings []string,_type int) *QTreeWidgetItem {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),177000,177110,Native(parent),unsafe.Pointer(&strings),unsafe.Pointer(&_type),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTreeWidgetItem{}
	_p.SetDriver(__rv,177000,true)
	return _p
} 
//QTreeWidgetItem::QTreeWidgetItem(QTreeWidgetItem*,QTreeWidgetItem*,int)	
func NewTreeWidgetItemWithParentAfterType(parent *QTreeWidgetItem,after *QTreeWidgetItem,_type int) *QTreeWidgetItem {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),177000,177111,Native(parent),Native(after),unsafe.Pointer(&_type),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTreeWidgetItem{}
	_p.SetDriver(__rv,177000,true)
	return _p
} 
//QTreeWidgetItem::addChild(QTreeWidgetItem*)
func (q *QTreeWidgetItem) AddChild(child *QTreeWidgetItem)  {
	q.Drv(177000,177112,Native(child),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeWidgetItem::addChildren(QList<QTreeWidgetItem*> const&)
func (q *QTreeWidgetItem) AddChildren(children []*QTreeWidgetItem)  {
	q.Drv(177000,177113,unsafe.Pointer(&children),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeWidgetItem::background(int)
func (q *QTreeWidgetItem) Background(column int) *QBrush {
	var __rv uintptr
	q.Drv(177000,177114,unsafe.Pointer(&column),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QBrush{}
	_rp.SetDriver(__rv,9000,true)
	return _rp
}	
//QTreeWidgetItem::backgroundColor(int)
func (q *QTreeWidgetItem) BackgroundColor(column int) *QColor {
	var __rv uintptr
	q.Drv(177000,177115,unsafe.Pointer(&column),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QColor{}
	_rp.SetDriver(__rv,13000,true)
	return _rp
}	
//QTreeWidgetItem::checkState(int)
func (q *QTreeWidgetItem) CheckState(column int) Qt_CheckState {
	var __rv Qt_CheckState
	q.Drv(177000,177116,unsafe.Pointer(&column),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTreeWidgetItem::child(int)
func (q *QTreeWidgetItem) Child(index int) *QTreeWidgetItem {
	var __rv uintptr
	q.Drv(177000,177117,unsafe.Pointer(&index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTreeWidgetItem{}
	_rp.SetDriver(__rv,177000,true)
	return _rp
}	
//QTreeWidgetItem::childCount()
func (q *QTreeWidgetItem) ChildCount() int {
	var __rv int
	q.Drv(177000,177118,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTreeWidgetItem::childIndicatorPolicy()
func (q *QTreeWidgetItem) ChildIndicatorPolicy() QTreeWidgetItem_ChildIndicatorPolicy {
	var __rv QTreeWidgetItem_ChildIndicatorPolicy
	q.Drv(177000,177119,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTreeWidgetItem::clone()
func (q *QTreeWidgetItem) Clone() *QTreeWidgetItem {
	var __rv uintptr
	q.Drv(177000,177120,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTreeWidgetItem{}
	_rp.SetDriver(__rv,177000,true)
	return _rp
}	
//QTreeWidgetItem::columnCount()
func (q *QTreeWidgetItem) ColumnCount() int {
	var __rv int
	q.Drv(177000,177121,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTreeWidgetItem::data(int,int)
func (q *QTreeWidgetItem) Data(column int,role int) *QVariant {
	var __rv uintptr
	q.Drv(177000,177122,unsafe.Pointer(&column),unsafe.Pointer(&role),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVariant{}
	_rp.SetDriver(__rv,182000,true)
	return _rp
}	
//QTreeWidgetItem::flags()
func (q *QTreeWidgetItem) Flags() Qt_ItemFlag {
	var __rv Qt_ItemFlag
	q.Drv(177000,177123,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTreeWidgetItem::font(int)
func (q *QTreeWidgetItem) Font(column int) *QFont {
	var __rv uintptr
	q.Drv(177000,177124,unsafe.Pointer(&column),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QFont{}
	_rp.SetDriver(__rv,37000,true)
	return _rp
}	
//QTreeWidgetItem::foreground(int)
func (q *QTreeWidgetItem) Foreground(column int) *QBrush {
	var __rv uintptr
	q.Drv(177000,177125,unsafe.Pointer(&column),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QBrush{}
	_rp.SetDriver(__rv,9000,true)
	return _rp
}	
//QTreeWidgetItem::icon(int)
func (q *QTreeWidgetItem) Icon(column int) *QIcon {
	var __rv uintptr
	q.Drv(177000,177126,unsafe.Pointer(&column),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QIcon{}
	_rp.SetDriver(__rv,51000,true)
	return _rp
}	
//QTreeWidgetItem::indexOfChild(QTreeWidgetItem*)
func (q *QTreeWidgetItem) IndexOfChild(child *QTreeWidgetItem) int {
	var __rv int
	q.Drv(177000,177127,Native(child),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTreeWidgetItem::insertChild(int,QTreeWidgetItem*)
func (q *QTreeWidgetItem) InsertChild(index int,child *QTreeWidgetItem)  {
	q.Drv(177000,177128,unsafe.Pointer(&index),Native(child),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeWidgetItem::insertChildren(int,QList<QTreeWidgetItem*> const&)
func (q *QTreeWidgetItem) InsertChildren(index int,children []*QTreeWidgetItem)  {
	q.Drv(177000,177129,unsafe.Pointer(&index),unsafe.Pointer(&children),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeWidgetItem::isDisabled()
func (q *QTreeWidgetItem) IsDisabled() bool {
	var __rv bool
	q.Drv(177000,177130,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTreeWidgetItem::isExpanded()
func (q *QTreeWidgetItem) IsExpanded() bool {
	var __rv bool
	q.Drv(177000,177131,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTreeWidgetItem::isFirstColumnSpanned()
func (q *QTreeWidgetItem) IsFirstColumnSpanned() bool {
	var __rv bool
	q.Drv(177000,177132,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTreeWidgetItem::isHidden()
func (q *QTreeWidgetItem) IsHidden() bool {
	var __rv bool
	q.Drv(177000,177133,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTreeWidgetItem::isSelected()
func (q *QTreeWidgetItem) IsSelected() bool {
	var __rv bool
	q.Drv(177000,177134,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTreeWidgetItem::parent()
func (q *QTreeWidgetItem) Parent() *QTreeWidgetItem {
	var __rv uintptr
	q.Drv(177000,177135,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTreeWidgetItem{}
	_rp.SetDriver(__rv,177000,true)
	return _rp
}	
//QTreeWidgetItem::removeChild(QTreeWidgetItem*)
func (q *QTreeWidgetItem) RemoveChild(child *QTreeWidgetItem)  {
	q.Drv(177000,177136,Native(child),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeWidgetItem::setBackground(int,QBrush const&)
func (q *QTreeWidgetItem) SetBackground(column int,brush *QBrush)  {
	q.Drv(177000,177137,unsafe.Pointer(&column),Native(brush),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeWidgetItem::setBackgroundColor(int,QColor const&)
func (q *QTreeWidgetItem) SetBackgroundColor(column int,color *QColor)  {
	q.Drv(177000,177138,unsafe.Pointer(&column),Native(color),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeWidgetItem::setCheckState(int,Qt::CheckState)
func (q *QTreeWidgetItem) SetCheckState(column int,state Qt_CheckState)  {
	q.Drv(177000,177139,unsafe.Pointer(&column),unsafe.Pointer(&state),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeWidgetItem::setChildIndicatorPolicy(QTreeWidgetItem::ChildIndicatorPolicy)
func (q *QTreeWidgetItem) SetChildIndicatorPolicy(policy QTreeWidgetItem_ChildIndicatorPolicy)  {
	q.Drv(177000,177140,unsafe.Pointer(&policy),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeWidgetItem::setData(int,int,QVariant const&)
func (q *QTreeWidgetItem) SetData(column int,role int,value *QVariant)  {
	q.Drv(177000,177141,unsafe.Pointer(&column),unsafe.Pointer(&role),Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeWidgetItem::setDisabled(bool)
func (q *QTreeWidgetItem) SetDisabled(disabled bool)  {
	q.Drv(177000,177142,unsafe.Pointer(&disabled),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeWidgetItem::setExpanded(bool)
func (q *QTreeWidgetItem) SetExpanded(expand bool)  {
	q.Drv(177000,177143,unsafe.Pointer(&expand),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeWidgetItem::setFirstColumnSpanned(bool)
func (q *QTreeWidgetItem) SetFirstColumnSpanned(span bool)  {
	q.Drv(177000,177144,unsafe.Pointer(&span),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeWidgetItem::setFlags(QFlags<Qt::ItemFlag>)
func (q *QTreeWidgetItem) SetFlags(flags Qt_ItemFlag)  {
	q.Drv(177000,177145,unsafe.Pointer(&flags),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeWidgetItem::setFont(int,QFont const&)
func (q *QTreeWidgetItem) SetFont(column int,font *QFont)  {
	q.Drv(177000,177146,unsafe.Pointer(&column),Native(font),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeWidgetItem::setForeground(int,QBrush const&)
func (q *QTreeWidgetItem) SetForeground(column int,brush *QBrush)  {
	q.Drv(177000,177147,unsafe.Pointer(&column),Native(brush),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeWidgetItem::setHidden(bool)
func (q *QTreeWidgetItem) SetHidden(hide bool)  {
	q.Drv(177000,177148,unsafe.Pointer(&hide),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeWidgetItem::setIcon(int,QIcon const&)
func (q *QTreeWidgetItem) SetIcon(column int,icon *QIcon)  {
	q.Drv(177000,177149,unsafe.Pointer(&column),Native(icon),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeWidgetItem::setSelected(bool)
func (q *QTreeWidgetItem) SetSelected(_select bool)  {
	q.Drv(177000,177150,unsafe.Pointer(&_select),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeWidgetItem::setSizeHint(int,QSize const&)
func (q *QTreeWidgetItem) SetSizeHint(column int,size *QSize)  {
	q.Drv(177000,177151,unsafe.Pointer(&column),Native(size),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeWidgetItem::setStatusTip(int,QString const&)
func (q *QTreeWidgetItem) SetStatusTip(column int,statusTip string)  {
	q.Drv(177000,177152,unsafe.Pointer(&column),unsafe.Pointer(&statusTip),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeWidgetItem::setText(int,QString const&)
func (q *QTreeWidgetItem) SetText(column int,text string)  {
	q.Drv(177000,177153,unsafe.Pointer(&column),unsafe.Pointer(&text),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeWidgetItem::setTextAlignment(int,int)
func (q *QTreeWidgetItem) SetTextAlignment(column int,alignment int)  {
	q.Drv(177000,177154,unsafe.Pointer(&column),unsafe.Pointer(&alignment),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeWidgetItem::setTextColor(int,QColor const&)
func (q *QTreeWidgetItem) SetTextColor(column int,color *QColor)  {
	q.Drv(177000,177155,unsafe.Pointer(&column),Native(color),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeWidgetItem::setToolTip(int,QString const&)
func (q *QTreeWidgetItem) SetToolTip(column int,toolTip string)  {
	q.Drv(177000,177156,unsafe.Pointer(&column),unsafe.Pointer(&toolTip),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeWidgetItem::setWhatsThis(int,QString const&)
func (q *QTreeWidgetItem) SetWhatsThis(column int,whatsThis string)  {
	q.Drv(177000,177157,unsafe.Pointer(&column),unsafe.Pointer(&whatsThis),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeWidgetItem::sizeHint(int)
func (q *QTreeWidgetItem) SizeHint(column int) *QSize {
	var __rv uintptr
	q.Drv(177000,177158,unsafe.Pointer(&column),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QTreeWidgetItem::sortChildren(int,Qt::SortOrder)
func (q *QTreeWidgetItem) SortChildren(column int,order Qt_SortOrder)  {
	q.Drv(177000,177159,unsafe.Pointer(&column),unsafe.Pointer(&order),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeWidgetItem::statusTip(int)
func (q *QTreeWidgetItem) StatusTip(column int) string {
	var __rv string
	q.Drv(177000,177160,unsafe.Pointer(&column),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTreeWidgetItem::takeChild(int)
func (q *QTreeWidgetItem) TakeChild(index int) *QTreeWidgetItem {
	var __rv uintptr
	q.Drv(177000,177161,unsafe.Pointer(&index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTreeWidgetItem{}
	_rp.SetDriver(__rv,177000,true)
	return _rp
}	
//QTreeWidgetItem::takeChildren()
func (q *QTreeWidgetItem) TakeChildren() []*QTreeWidgetItem {
	var __rv []*QTreeWidgetItem
	q.Drv(177000,177162,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTreeWidgetItem::text(int)
func (q *QTreeWidgetItem) Text(column int) string {
	var __rv string
	q.Drv(177000,177163,unsafe.Pointer(&column),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTreeWidgetItem::textAlignment(int)
func (q *QTreeWidgetItem) TextAlignment(column int) int {
	var __rv int
	q.Drv(177000,177164,unsafe.Pointer(&column),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTreeWidgetItem::textColor(int)
func (q *QTreeWidgetItem) TextColor(column int) *QColor {
	var __rv uintptr
	q.Drv(177000,177165,unsafe.Pointer(&column),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QColor{}
	_rp.SetDriver(__rv,13000,true)
	return _rp
}	
//QTreeWidgetItem::toolTip(int)
func (q *QTreeWidgetItem) ToolTip(column int) string {
	var __rv string
	q.Drv(177000,177166,unsafe.Pointer(&column),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTreeWidgetItem::treeWidget()
func (q *QTreeWidgetItem) TreeWidget() *QTreeWidget {
	var __rv uintptr
	q.Drv(177000,177167,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTreeWidget{}
	_rp.SetDriver(__rv,386000,false)
	return _rp
}	
//QTreeWidgetItem::type()
func (q *QTreeWidgetItem) Type() int {
	var __rv int
	q.Drv(177000,177168,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTreeWidgetItem::whatsThis(int)
func (q *QTreeWidgetItem) WhatsThis(column int) string {
	var __rv string
	q.Drv(177000,177169,unsafe.Pointer(&column),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
