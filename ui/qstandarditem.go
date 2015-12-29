// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QStandardItem_ItemType - QStandardItem::ItemType
type QStandardItem_ItemType uint32
const (
	QStandardItem_Type QStandardItem_ItemType = 0
	QStandardItem_UserType QStandardItem_ItemType = 1000
)
//struct QStandardItem : QStandardItem
type QStandardItem struct {
	BaseDrv
}
//QStandardItem::QStandardItem()	
func NewStandardItem() *QStandardItem {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),123000,123102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QStandardItem{}
	_p.SetDriver(__rv,123000,true)
	return _p
} 
//QStandardItem::QStandardItem(QString const&)	
func NewStandardItemWithText(text string) *QStandardItem {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),123000,123103,unsafe.Pointer(&text),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QStandardItem{}
	_p.SetDriver(__rv,123000,true)
	return _p
} 
//QStandardItem::QStandardItem(QIcon const&,QString const&)	
func NewStandardItemWithIconText(icon *QIcon,text string) *QStandardItem {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),123000,123104,Native(icon),unsafe.Pointer(&text),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QStandardItem{}
	_p.SetDriver(__rv,123000,true)
	return _p
} 
//QStandardItem::QStandardItem(int,int)	
func NewStandardItemWithRowsColumns(rows int,columns int) *QStandardItem {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),123000,123105,unsafe.Pointer(&rows),unsafe.Pointer(&columns),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QStandardItem{}
	_p.SetDriver(__rv,123000,true)
	return _p
} 
//QStandardItem::accessibleDescription()
func (q *QStandardItem) AccessibleDescription() string {
	var __rv string
	q.Drv(123000,123106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStandardItem::accessibleText()
func (q *QStandardItem) AccessibleText() string {
	var __rv string
	q.Drv(123000,123107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStandardItem::appendColumn(QList<QStandardItem*> const&)
func (q *QStandardItem) AppendColumn(items []*QStandardItem)  {
	q.Drv(123000,123108,unsafe.Pointer(&items),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStandardItem::appendRow(QList<QStandardItem*> const&)
func (q *QStandardItem) AppendRow(items []*QStandardItem)  {
	q.Drv(123000,123109,unsafe.Pointer(&items),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStandardItem::appendRow(QStandardItem*)
func (q *QStandardItem) AppendRowWithItem(item *QStandardItem)  {
	q.Drv(123000,123110,Native(item),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStandardItem::appendRows(QList<QStandardItem*> const&)
func (q *QStandardItem) AppendRows(items []*QStandardItem)  {
	q.Drv(123000,123111,unsafe.Pointer(&items),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStandardItem::background()
func (q *QStandardItem) Background() *QBrush {
	var __rv uintptr
	q.Drv(123000,123112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QBrush{}
	_rp.SetDriver(__rv,9000,true)
	return _rp
}	
//QStandardItem::checkState()
func (q *QStandardItem) CheckState() Qt_CheckState {
	var __rv Qt_CheckState
	q.Drv(123000,123113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStandardItem::child(int)
func (q *QStandardItem) Child(row int) *QStandardItem {
	var __rv uintptr
	q.Drv(123000,123114,unsafe.Pointer(&row),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QStandardItem{}
	_rp.SetDriver(__rv,123000,true)
	return _rp
}	
//QStandardItem::child(int,int)
func (q *QStandardItem) ChildWithRowColumn(row int,column int) *QStandardItem {
	var __rv uintptr
	q.Drv(123000,123115,unsafe.Pointer(&row),unsafe.Pointer(&column),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QStandardItem{}
	_rp.SetDriver(__rv,123000,true)
	return _rp
}	
//QStandardItem::clone()
func (q *QStandardItem) Clone() *QStandardItem {
	var __rv uintptr
	q.Drv(123000,123116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QStandardItem{}
	_rp.SetDriver(__rv,123000,true)
	return _rp
}	
//QStandardItem::column()
func (q *QStandardItem) Column() int {
	var __rv int
	q.Drv(123000,123117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStandardItem::columnCount()
func (q *QStandardItem) ColumnCount() int {
	var __rv int
	q.Drv(123000,123118,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStandardItem::data()
func (q *QStandardItem) Data() *QVariant {
	var __rv uintptr
	q.Drv(123000,123119,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVariant{}
	_rp.SetDriver(__rv,182000,true)
	return _rp
}	
//QStandardItem::data(int)
func (q *QStandardItem) DataWithRole(role int) *QVariant {
	var __rv uintptr
	q.Drv(123000,123120,unsafe.Pointer(&role),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVariant{}
	_rp.SetDriver(__rv,182000,true)
	return _rp
}	
//QStandardItem::flags()
func (q *QStandardItem) Flags() Qt_ItemFlag {
	var __rv Qt_ItemFlag
	q.Drv(123000,123121,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStandardItem::font()
func (q *QStandardItem) Font() *QFont {
	var __rv uintptr
	q.Drv(123000,123122,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QFont{}
	_rp.SetDriver(__rv,37000,true)
	return _rp
}	
//QStandardItem::foreground()
func (q *QStandardItem) Foreground() *QBrush {
	var __rv uintptr
	q.Drv(123000,123123,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QBrush{}
	_rp.SetDriver(__rv,9000,true)
	return _rp
}	
//QStandardItem::hasChildren()
func (q *QStandardItem) HasChildren() bool {
	var __rv bool
	q.Drv(123000,123124,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStandardItem::icon()
func (q *QStandardItem) Icon() *QIcon {
	var __rv uintptr
	q.Drv(123000,123125,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QIcon{}
	_rp.SetDriver(__rv,51000,true)
	return _rp
}	
//QStandardItem::index()
func (q *QStandardItem) Index() *QModelIndex {
	var __rv uintptr
	q.Drv(123000,123126,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QModelIndex{}
	_rp.SetDriver(__rv,79000,true)
	return _rp
}	
//QStandardItem::insertColumn(int,QList<QStandardItem*> const&)
func (q *QStandardItem) InsertColumn(column int,items []*QStandardItem)  {
	q.Drv(123000,123127,unsafe.Pointer(&column),unsafe.Pointer(&items),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStandardItem::insertColumns(int,int)
func (q *QStandardItem) InsertColumns(column int,count int)  {
	q.Drv(123000,123128,unsafe.Pointer(&column),unsafe.Pointer(&count),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStandardItem::insertRow(int,QList<QStandardItem*> const&)
func (q *QStandardItem) InsertRowWithRowItems(row int,items []*QStandardItem)  {
	q.Drv(123000,123129,unsafe.Pointer(&row),unsafe.Pointer(&items),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStandardItem::insertRow(int,QStandardItem*)
func (q *QStandardItem) InsertRowWithRowItem(row int,item *QStandardItem)  {
	q.Drv(123000,123130,unsafe.Pointer(&row),Native(item),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStandardItem::insertRows(int,QList<QStandardItem*> const&)
func (q *QStandardItem) InsertRowsWithRowItems(row int,items []*QStandardItem)  {
	q.Drv(123000,123131,unsafe.Pointer(&row),unsafe.Pointer(&items),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStandardItem::insertRows(int,int)
func (q *QStandardItem) InsertRowsWithRowCount(row int,count int)  {
	q.Drv(123000,123132,unsafe.Pointer(&row),unsafe.Pointer(&count),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStandardItem::isCheckable()
func (q *QStandardItem) IsCheckable() bool {
	var __rv bool
	q.Drv(123000,123133,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStandardItem::isDragEnabled()
func (q *QStandardItem) IsDragEnabled() bool {
	var __rv bool
	q.Drv(123000,123134,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStandardItem::isDropEnabled()
func (q *QStandardItem) IsDropEnabled() bool {
	var __rv bool
	q.Drv(123000,123135,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStandardItem::isEditable()
func (q *QStandardItem) IsEditable() bool {
	var __rv bool
	q.Drv(123000,123136,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStandardItem::isEnabled()
func (q *QStandardItem) IsEnabled() bool {
	var __rv bool
	q.Drv(123000,123137,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStandardItem::isSelectable()
func (q *QStandardItem) IsSelectable() bool {
	var __rv bool
	q.Drv(123000,123138,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStandardItem::isTristate()
func (q *QStandardItem) IsTristate() bool {
	var __rv bool
	q.Drv(123000,123139,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStandardItem::model()
func (q *QStandardItem) Model() *QStandardItemModel {
	var __rv uintptr
	q.Drv(123000,123140,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QStandardItemModel{}
	_rp.SetDriver(__rv,352000,false)
	return _rp
}	
//QStandardItem::parent()
func (q *QStandardItem) Parent() *QStandardItem {
	var __rv uintptr
	q.Drv(123000,123141,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QStandardItem{}
	_rp.SetDriver(__rv,123000,true)
	return _rp
}	
//QStandardItem::removeColumn(int)
func (q *QStandardItem) RemoveColumn(column int)  {
	q.Drv(123000,123142,unsafe.Pointer(&column),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStandardItem::removeColumns(int,int)
func (q *QStandardItem) RemoveColumns(column int,count int)  {
	q.Drv(123000,123143,unsafe.Pointer(&column),unsafe.Pointer(&count),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStandardItem::removeRow(int)
func (q *QStandardItem) RemoveRow(row int)  {
	q.Drv(123000,123144,unsafe.Pointer(&row),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStandardItem::removeRows(int,int)
func (q *QStandardItem) RemoveRows(row int,count int)  {
	q.Drv(123000,123145,unsafe.Pointer(&row),unsafe.Pointer(&count),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStandardItem::row()
func (q *QStandardItem) Row() int {
	var __rv int
	q.Drv(123000,123146,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStandardItem::rowCount()
func (q *QStandardItem) RowCount() int {
	var __rv int
	q.Drv(123000,123147,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStandardItem::setAccessibleDescription(QString const&)
func (q *QStandardItem) SetAccessibleDescription(accessibleDescription string)  {
	q.Drv(123000,123148,unsafe.Pointer(&accessibleDescription),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStandardItem::setAccessibleText(QString const&)
func (q *QStandardItem) SetAccessibleText(accessibleText string)  {
	q.Drv(123000,123149,unsafe.Pointer(&accessibleText),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStandardItem::setBackground(QBrush const&)
func (q *QStandardItem) SetBackground(brush *QBrush)  {
	q.Drv(123000,123150,Native(brush),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStandardItem::setCheckState(Qt::CheckState)
func (q *QStandardItem) SetCheckState(checkState Qt_CheckState)  {
	q.Drv(123000,123151,unsafe.Pointer(&checkState),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStandardItem::setCheckable(bool)
func (q *QStandardItem) SetCheckable(checkable bool)  {
	q.Drv(123000,123152,unsafe.Pointer(&checkable),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStandardItem::setChild(int,QStandardItem*)
func (q *QStandardItem) SetChildWithRowItem(row int,item *QStandardItem)  {
	q.Drv(123000,123153,unsafe.Pointer(&row),Native(item),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStandardItem::setChild(int,int,QStandardItem*)
func (q *QStandardItem) SetChildWithRowColumnItem(row int,column int,item *QStandardItem)  {
	q.Drv(123000,123154,unsafe.Pointer(&row),unsafe.Pointer(&column),Native(item),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStandardItem::setColumnCount(int)
func (q *QStandardItem) SetColumnCount(columns int)  {
	q.Drv(123000,123155,unsafe.Pointer(&columns),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStandardItem::setData(QVariant const&)
func (q *QStandardItem) SetData(value *QVariant)  {
	q.Drv(123000,123156,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStandardItem::setData(QVariant const&,int)
func (q *QStandardItem) SetDataWithValueRole(value *QVariant,role int)  {
	q.Drv(123000,123157,Native(value),unsafe.Pointer(&role),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStandardItem::setDragEnabled(bool)
func (q *QStandardItem) SetDragEnabled(dragEnabled bool)  {
	q.Drv(123000,123158,unsafe.Pointer(&dragEnabled),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStandardItem::setDropEnabled(bool)
func (q *QStandardItem) SetDropEnabled(dropEnabled bool)  {
	q.Drv(123000,123159,unsafe.Pointer(&dropEnabled),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStandardItem::setEditable(bool)
func (q *QStandardItem) SetEditable(editable bool)  {
	q.Drv(123000,123160,unsafe.Pointer(&editable),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStandardItem::setEnabled(bool)
func (q *QStandardItem) SetEnabled(enabled bool)  {
	q.Drv(123000,123161,unsafe.Pointer(&enabled),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStandardItem::setFlags(QFlags<Qt::ItemFlag>)
func (q *QStandardItem) SetFlags(flags Qt_ItemFlag)  {
	q.Drv(123000,123162,unsafe.Pointer(&flags),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStandardItem::setFont(QFont const&)
func (q *QStandardItem) SetFont(font *QFont)  {
	q.Drv(123000,123163,Native(font),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStandardItem::setForeground(QBrush const&)
func (q *QStandardItem) SetForeground(brush *QBrush)  {
	q.Drv(123000,123164,Native(brush),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStandardItem::setIcon(QIcon const&)
func (q *QStandardItem) SetIcon(icon *QIcon)  {
	q.Drv(123000,123165,Native(icon),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStandardItem::setRowCount(int)
func (q *QStandardItem) SetRowCount(rows int)  {
	q.Drv(123000,123166,unsafe.Pointer(&rows),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStandardItem::setSelectable(bool)
func (q *QStandardItem) SetSelectable(selectable bool)  {
	q.Drv(123000,123167,unsafe.Pointer(&selectable),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStandardItem::setSizeHint(QSize const&)
func (q *QStandardItem) SetSizeHint(sizeHint *QSize)  {
	q.Drv(123000,123168,Native(sizeHint),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStandardItem::setStatusTip(QString const&)
func (q *QStandardItem) SetStatusTip(statusTip string)  {
	q.Drv(123000,123169,unsafe.Pointer(&statusTip),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStandardItem::setText(QString const&)
func (q *QStandardItem) SetText(text string)  {
	q.Drv(123000,123170,unsafe.Pointer(&text),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStandardItem::setTextAlignment(QFlags<Qt::AlignmentFlag>)
func (q *QStandardItem) SetTextAlignment(textAlignment Qt_AlignmentFlag)  {
	q.Drv(123000,123171,unsafe.Pointer(&textAlignment),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStandardItem::setToolTip(QString const&)
func (q *QStandardItem) SetToolTip(toolTip string)  {
	q.Drv(123000,123172,unsafe.Pointer(&toolTip),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStandardItem::setTristate(bool)
func (q *QStandardItem) SetTristate(tristate bool)  {
	q.Drv(123000,123173,unsafe.Pointer(&tristate),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStandardItem::setWhatsThis(QString const&)
func (q *QStandardItem) SetWhatsThis(whatsThis string)  {
	q.Drv(123000,123174,unsafe.Pointer(&whatsThis),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStandardItem::sizeHint()
func (q *QStandardItem) SizeHint() *QSize {
	var __rv uintptr
	q.Drv(123000,123175,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QStandardItem::sortChildren(int)
func (q *QStandardItem) SortChildren(column int)  {
	q.Drv(123000,123176,unsafe.Pointer(&column),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStandardItem::sortChildren(int,Qt::SortOrder)
func (q *QStandardItem) SortChildrenWithColumnOrder(column int,order Qt_SortOrder)  {
	q.Drv(123000,123177,unsafe.Pointer(&column),unsafe.Pointer(&order),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStandardItem::statusTip()
func (q *QStandardItem) StatusTip() string {
	var __rv string
	q.Drv(123000,123178,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStandardItem::takeChild(int)
func (q *QStandardItem) TakeChild(row int) *QStandardItem {
	var __rv uintptr
	q.Drv(123000,123179,unsafe.Pointer(&row),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QStandardItem{}
	_rp.SetDriver(__rv,123000,true)
	return _rp
}	
//QStandardItem::takeChild(int,int)
func (q *QStandardItem) TakeChildWithRowColumn(row int,column int) *QStandardItem {
	var __rv uintptr
	q.Drv(123000,123180,unsafe.Pointer(&row),unsafe.Pointer(&column),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QStandardItem{}
	_rp.SetDriver(__rv,123000,true)
	return _rp
}	
//QStandardItem::takeColumn(int)
func (q *QStandardItem) TakeColumn(column int) []*QStandardItem {
	var __rv []*QStandardItem
	q.Drv(123000,123181,unsafe.Pointer(&column),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStandardItem::takeRow(int)
func (q *QStandardItem) TakeRow(row int) []*QStandardItem {
	var __rv []*QStandardItem
	q.Drv(123000,123182,unsafe.Pointer(&row),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStandardItem::text()
func (q *QStandardItem) Text() string {
	var __rv string
	q.Drv(123000,123183,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStandardItem::textAlignment()
func (q *QStandardItem) TextAlignment() Qt_AlignmentFlag {
	var __rv Qt_AlignmentFlag
	q.Drv(123000,123184,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStandardItem::toolTip()
func (q *QStandardItem) ToolTip() string {
	var __rv string
	q.Drv(123000,123185,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStandardItem::type()
func (q *QStandardItem) Type() int {
	var __rv int
	q.Drv(123000,123186,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStandardItem::whatsThis()
func (q *QStandardItem) WhatsThis() string {
	var __rv string
	q.Drv(123000,123187,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
