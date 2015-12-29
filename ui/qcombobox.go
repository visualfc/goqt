// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QComboBox_InsertPolicy - QComboBox::InsertPolicy
type QComboBox_InsertPolicy uint32
const (
	QComboBox_NoInsert QComboBox_InsertPolicy = 0
	QComboBox_InsertAtTop QComboBox_InsertPolicy = 1
	QComboBox_InsertAtCurrent QComboBox_InsertPolicy = 2
	QComboBox_InsertAtBottom QComboBox_InsertPolicy = 3
	QComboBox_InsertAfterCurrent QComboBox_InsertPolicy = 4
	QComboBox_InsertBeforeCurrent QComboBox_InsertPolicy = 5
	QComboBox_InsertAlphabetically QComboBox_InsertPolicy = 6
)
//enum QComboBox_SizeAdjustPolicy - QComboBox::SizeAdjustPolicy
type QComboBox_SizeAdjustPolicy uint32
const (
	QComboBox_AdjustToContents QComboBox_SizeAdjustPolicy = 0
	QComboBox_AdjustToContentsOnFirstShow QComboBox_SizeAdjustPolicy = 1
	QComboBox_AdjustToMinimumContentsLength QComboBox_SizeAdjustPolicy = 2
	QComboBox_AdjustToMinimumContentsLengthWithIcon QComboBox_SizeAdjustPolicy = 3
)
//struct QComboBox : QComboBox
type QComboBox struct {
	QWidget
}
func (q *QComboBox) OnHighlighted(fn func(string)) uintptr {
	var __rv uintptr
	q.Drv(217000,217102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QComboBox) OnHighlightedWithIndex(fn func(int)) uintptr {
	var __rv uintptr
	q.Drv(217000,217103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QComboBox) OnEditTextChanged(fn func(string)) uintptr {
	var __rv uintptr
	q.Drv(217000,217104,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QComboBox) OnActivated(fn func(string)) uintptr {
	var __rv uintptr
	q.Drv(217000,217105,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QComboBox) OnActivatedWithIndex(fn func(int)) uintptr {
	var __rv uintptr
	q.Drv(217000,217106,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QComboBox) OnCurrentIndexChanged(fn func(string)) uintptr {
	var __rv uintptr
	q.Drv(217000,217107,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QComboBox) OnCurrentIndexChangedWithIndex(fn func(int)) uintptr {
	var __rv uintptr
	q.Drv(217000,217108,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QComboBox::QComboBox()	
func NewComboBox() *QComboBox {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),217000,217109,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QComboBox{}
	_p.SetDriver(__rv,217000,false)
	return _p
} 
//QComboBox::QComboBox(QWidget*)	
func NewComboBoxWithParent(parent QWidgetInterface) *QComboBox {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),217000,217110,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QComboBox{}
	_p.SetDriver(__rv,217000,false)
	return _p
} 
//QComboBox::addItem(QString const&)
func (q *QComboBox) AddItem(text string)  {
	q.Drv(217000,217111,unsafe.Pointer(&text),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QComboBox::addItem(QString const&,QVariant const&)
func (q *QComboBox) AddItemWithTextUserdata(text string,userData *QVariant)  {
	q.Drv(217000,217112,unsafe.Pointer(&text),Native(userData),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QComboBox::addItem(QIcon const&,QString const&,QVariant const&)
func (q *QComboBox) AddItemWithIconTextUserdata(icon *QIcon,text string,userData *QVariant)  {
	q.Drv(217000,217113,Native(icon),unsafe.Pointer(&text),Native(userData),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QComboBox::addItems(QStringList const&)
func (q *QComboBox) AddItems(texts []string)  {
	q.Drv(217000,217114,unsafe.Pointer(&texts),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QComboBox::autoCompletion()
func (q *QComboBox) AutoCompletion() bool {
	var __rv bool
	q.Drv(217000,217115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QComboBox::autoCompletionCaseSensitivity()
func (q *QComboBox) AutoCompletionCaseSensitivity() Qt_CaseSensitivity {
	var __rv Qt_CaseSensitivity
	q.Drv(217000,217116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QComboBox::changeEvent(QEvent*)
func (q *QComboBox) ChangeEvent(e *QEvent)  {
	q.Drv(217000,217117,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QComboBox::clear()
func (q *QComboBox) Clear()  {
	q.Drv(217000,217118,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QComboBox::clearEditText()
func (q *QComboBox) ClearEditText()  {
	q.Drv(217000,217119,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QComboBox::completer()
func (q *QComboBox) Completer() *QCompleter {
	var __rv uintptr
	q.Drv(217000,217120,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QCompleter{}
	_rp.SetDriver(__rv,220000,false)
	return _rp
}	
//QComboBox::contextMenuEvent(QContextMenuEvent*)
func (q *QComboBox) ContextMenuEvent(e *QContextMenuEvent)  {
	q.Drv(217000,217121,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QComboBox::count()
func (q *QComboBox) Count() int {
	var __rv int
	q.Drv(217000,217122,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QComboBox::currentIndex()
func (q *QComboBox) CurrentIndex() int {
	var __rv int
	q.Drv(217000,217123,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QComboBox::currentText()
func (q *QComboBox) CurrentText() string {
	var __rv string
	q.Drv(217000,217124,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QComboBox::duplicatesEnabled()
func (q *QComboBox) DuplicatesEnabled() bool {
	var __rv bool
	q.Drv(217000,217125,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QComboBox::event(QEvent*)
func (q *QComboBox) Event(event *QEvent) bool {
	var __rv bool
	q.Drv(217000,217126,Native(event),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QComboBox::findData(QVariant const&)
func (q *QComboBox) FindData(data *QVariant) int {
	var __rv int
	q.Drv(217000,217127,Native(data),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QComboBox::findData(QVariant const&,int,QFlags<Qt::MatchFlag>)
func (q *QComboBox) FindDataWithDataRoleFlags(data *QVariant,role int,flags Qt_MatchFlag) int {
	var __rv int
	q.Drv(217000,217128,Native(data),unsafe.Pointer(&role),unsafe.Pointer(&flags),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QComboBox::findText(QString const&)
func (q *QComboBox) FindText(text string) int {
	var __rv int
	q.Drv(217000,217129,unsafe.Pointer(&text),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QComboBox::findText(QString const&,QFlags<Qt::MatchFlag>)
func (q *QComboBox) FindTextWithTextFlags(text string,flags Qt_MatchFlag) int {
	var __rv int
	q.Drv(217000,217130,unsafe.Pointer(&text),unsafe.Pointer(&flags),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QComboBox::focusInEvent(QFocusEvent*)
func (q *QComboBox) FocusInEvent(e *QFocusEvent)  {
	q.Drv(217000,217131,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QComboBox::focusOutEvent(QFocusEvent*)
func (q *QComboBox) FocusOutEvent(e *QFocusEvent)  {
	q.Drv(217000,217132,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QComboBox::hasFrame()
func (q *QComboBox) HasFrame() bool {
	var __rv bool
	q.Drv(217000,217133,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QComboBox::hideEvent(QHideEvent*)
func (q *QComboBox) HideEvent(e *QHideEvent)  {
	q.Drv(217000,217134,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QComboBox::hidePopup()
func (q *QComboBox) HidePopup()  {
	q.Drv(217000,217135,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QComboBox::iconSize()
func (q *QComboBox) IconSize() *QSize {
	var __rv uintptr
	q.Drv(217000,217136,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QComboBox::inputMethodEvent(QInputMethodEvent*)
func (q *QComboBox) InputMethodEvent(value *QInputMethodEvent)  {
	q.Drv(217000,217137,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QComboBox::inputMethodQuery(Qt::InputMethodQuery)
func (q *QComboBox) InputMethodQuery(value Qt_InputMethodQuery) *QVariant {
	var __rv uintptr
	q.Drv(217000,217138,unsafe.Pointer(&value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVariant{}
	_rp.SetDriver(__rv,182000,true)
	return _rp
}	
//QComboBox::insertItem(int,QString const&,QVariant const&)
func (q *QComboBox) InsertItemWithIndexTextUserdata(index int,text string,userData *QVariant)  {
	q.Drv(217000,217139,unsafe.Pointer(&index),unsafe.Pointer(&text),Native(userData),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QComboBox::insertItem(int,QIcon const&,QString const&,QVariant const&)
func (q *QComboBox) InsertItemWithIndexIconTextUserdata(index int,icon *QIcon,text string,userData *QVariant)  {
	q.Drv(217000,217140,unsafe.Pointer(&index),Native(icon),unsafe.Pointer(&text),Native(userData),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QComboBox::insertItems(int,QStringList const&)
func (q *QComboBox) InsertItems(index int,texts []string)  {
	q.Drv(217000,217141,unsafe.Pointer(&index),unsafe.Pointer(&texts),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QComboBox::insertPolicy()
func (q *QComboBox) InsertPolicy() QComboBox_InsertPolicy {
	var __rv QComboBox_InsertPolicy
	q.Drv(217000,217142,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QComboBox::insertSeparator(int)
func (q *QComboBox) InsertSeparator(index int)  {
	q.Drv(217000,217143,unsafe.Pointer(&index),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QComboBox::isEditable()
func (q *QComboBox) IsEditable() bool {
	var __rv bool
	q.Drv(217000,217144,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QComboBox::itemData(int)
func (q *QComboBox) ItemData(index int) *QVariant {
	var __rv uintptr
	q.Drv(217000,217145,unsafe.Pointer(&index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVariant{}
	_rp.SetDriver(__rv,182000,true)
	return _rp
}	
//QComboBox::itemData(int,int)
func (q *QComboBox) ItemDataWithIndexRole(index int,role int) *QVariant {
	var __rv uintptr
	q.Drv(217000,217146,unsafe.Pointer(&index),unsafe.Pointer(&role),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVariant{}
	_rp.SetDriver(__rv,182000,true)
	return _rp
}	
//QComboBox::itemDelegate()
func (q *QComboBox) ItemDelegate() *QAbstractItemDelegate {
	var __rv uintptr
	q.Drv(217000,217147,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAbstractItemDelegate{}
	_rp.SetDriver(__rv,194000,false)
	return _rp
}	
//QComboBox::itemIcon(int)
func (q *QComboBox) ItemIcon(index int) *QIcon {
	var __rv uintptr
	q.Drv(217000,217148,unsafe.Pointer(&index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QIcon{}
	_rp.SetDriver(__rv,51000,true)
	return _rp
}	
//QComboBox::itemText(int)
func (q *QComboBox) ItemText(index int) string {
	var __rv string
	q.Drv(217000,217149,unsafe.Pointer(&index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QComboBox::keyPressEvent(QKeyEvent*)
func (q *QComboBox) KeyPressEvent(e *QKeyEvent)  {
	q.Drv(217000,217150,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QComboBox::keyReleaseEvent(QKeyEvent*)
func (q *QComboBox) KeyReleaseEvent(e *QKeyEvent)  {
	q.Drv(217000,217151,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QComboBox::lineEdit()
func (q *QComboBox) LineEdit() *QLineEdit {
	var __rv uintptr
	q.Drv(217000,217152,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QLineEdit{}
	_rp.SetDriver(__rv,302000,false)
	return _rp
}	
//QComboBox::maxCount()
func (q *QComboBox) MaxCount() int {
	var __rv int
	q.Drv(217000,217153,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QComboBox::maxVisibleItems()
func (q *QComboBox) MaxVisibleItems() int {
	var __rv int
	q.Drv(217000,217154,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QComboBox::minimumContentsLength()
func (q *QComboBox) MinimumContentsLength() int {
	var __rv int
	q.Drv(217000,217155,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QComboBox::minimumSizeHint()
func (q *QComboBox) MinimumSizeHint() *QSize {
	var __rv uintptr
	q.Drv(217000,217156,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QComboBox::model()
func (q *QComboBox) Model() *QAbstractItemModel {
	var __rv uintptr
	q.Drv(217000,217157,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAbstractItemModel{}
	_rp.SetDriver(__rv,195000,false)
	return _rp
}	
//QComboBox::modelColumn()
func (q *QComboBox) ModelColumn() int {
	var __rv int
	q.Drv(217000,217158,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QComboBox::mousePressEvent(QMouseEvent*)
func (q *QComboBox) MousePressEvent(e *QMouseEvent)  {
	q.Drv(217000,217159,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QComboBox::mouseReleaseEvent(QMouseEvent*)
func (q *QComboBox) MouseReleaseEvent(e *QMouseEvent)  {
	q.Drv(217000,217160,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QComboBox::paintEvent(QPaintEvent*)
func (q *QComboBox) PaintEvent(e *QPaintEvent)  {
	q.Drv(217000,217161,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QComboBox::removeItem(int)
func (q *QComboBox) RemoveItem(index int)  {
	q.Drv(217000,217162,unsafe.Pointer(&index),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QComboBox::resizeEvent(QResizeEvent*)
func (q *QComboBox) ResizeEvent(e *QResizeEvent)  {
	q.Drv(217000,217163,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QComboBox::rootModelIndex()
func (q *QComboBox) RootModelIndex() *QModelIndex {
	var __rv uintptr
	q.Drv(217000,217164,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QModelIndex{}
	_rp.SetDriver(__rv,79000,true)
	return _rp
}	
//QComboBox::setAutoCompletion(bool)
func (q *QComboBox) SetAutoCompletion(enable bool)  {
	q.Drv(217000,217165,unsafe.Pointer(&enable),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QComboBox::setAutoCompletionCaseSensitivity(Qt::CaseSensitivity)
func (q *QComboBox) SetAutoCompletionCaseSensitivity(sensitivity Qt_CaseSensitivity)  {
	q.Drv(217000,217166,unsafe.Pointer(&sensitivity),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QComboBox::setCompleter(QCompleter*)
func (q *QComboBox) SetCompleter(c *QCompleter)  {
	q.Drv(217000,217167,Native(c),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QComboBox::setCurrentIndex(int)
func (q *QComboBox) SetCurrentIndex(index int)  {
	q.Drv(217000,217168,unsafe.Pointer(&index),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QComboBox::setDuplicatesEnabled(bool)
func (q *QComboBox) SetDuplicatesEnabled(enable bool)  {
	q.Drv(217000,217169,unsafe.Pointer(&enable),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QComboBox::setEditText(QString const&)
func (q *QComboBox) SetEditText(text string)  {
	q.Drv(217000,217170,unsafe.Pointer(&text),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QComboBox::setEditable(bool)
func (q *QComboBox) SetEditable(editable bool)  {
	q.Drv(217000,217171,unsafe.Pointer(&editable),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QComboBox::setFrame(bool)
func (q *QComboBox) SetFrame(value bool)  {
	q.Drv(217000,217172,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QComboBox::setIconSize(QSize const&)
func (q *QComboBox) SetIconSize(size *QSize)  {
	q.Drv(217000,217173,Native(size),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QComboBox::setInsertPolicy(QComboBox::InsertPolicy)
func (q *QComboBox) SetInsertPolicy(policy QComboBox_InsertPolicy)  {
	q.Drv(217000,217174,unsafe.Pointer(&policy),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QComboBox::setItemData(int,QVariant const&,int)
func (q *QComboBox) SetItemData(index int,value *QVariant,role int)  {
	q.Drv(217000,217175,unsafe.Pointer(&index),Native(value),unsafe.Pointer(&role),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QComboBox::setItemDelegate(QAbstractItemDelegate*)
func (q *QComboBox) SetItemDelegate(delegate *QAbstractItemDelegate)  {
	q.Drv(217000,217176,Native(delegate),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QComboBox::setItemIcon(int,QIcon const&)
func (q *QComboBox) SetItemIcon(index int,icon *QIcon)  {
	q.Drv(217000,217177,unsafe.Pointer(&index),Native(icon),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QComboBox::setItemText(int,QString const&)
func (q *QComboBox) SetItemText(index int,text string)  {
	q.Drv(217000,217178,unsafe.Pointer(&index),unsafe.Pointer(&text),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QComboBox::setLineEdit(QLineEdit*)
func (q *QComboBox) SetLineEdit(edit *QLineEdit)  {
	q.Drv(217000,217179,Native(edit),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QComboBox::setMaxCount(int)
func (q *QComboBox) SetMaxCount(max int)  {
	q.Drv(217000,217180,unsafe.Pointer(&max),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QComboBox::setMaxVisibleItems(int)
func (q *QComboBox) SetMaxVisibleItems(maxItems int)  {
	q.Drv(217000,217181,unsafe.Pointer(&maxItems),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QComboBox::setMinimumContentsLength(int)
func (q *QComboBox) SetMinimumContentsLength(characters int)  {
	q.Drv(217000,217182,unsafe.Pointer(&characters),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QComboBox::setModel(QAbstractItemModel*)
func (q *QComboBox) SetModel(model QAbstractItemModelInterface)  {
	q.Drv(217000,217183,Native(model),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QComboBox::setModelColumn(int)
func (q *QComboBox) SetModelColumn(visibleColumn int)  {
	q.Drv(217000,217184,unsafe.Pointer(&visibleColumn),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QComboBox::setRootModelIndex(QModelIndex const&)
func (q *QComboBox) SetRootModelIndex(index *QModelIndex)  {
	q.Drv(217000,217185,Native(index),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QComboBox::setSizeAdjustPolicy(QComboBox::SizeAdjustPolicy)
func (q *QComboBox) SetSizeAdjustPolicy(policy QComboBox_SizeAdjustPolicy)  {
	q.Drv(217000,217186,unsafe.Pointer(&policy),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QComboBox::setValidator(QValidator const*)
func (q *QComboBox) SetValidator(v *QValidator)  {
	q.Drv(217000,217187,Native(v),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QComboBox::setView(QAbstractItemView*)
func (q *QComboBox) SetView(itemView *QAbstractItemView)  {
	q.Drv(217000,217188,Native(itemView),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QComboBox::showEvent(QShowEvent*)
func (q *QComboBox) ShowEvent(e *QShowEvent)  {
	q.Drv(217000,217189,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QComboBox::showPopup()
func (q *QComboBox) ShowPopup()  {
	q.Drv(217000,217190,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QComboBox::sizeAdjustPolicy()
func (q *QComboBox) SizeAdjustPolicy() QComboBox_SizeAdjustPolicy {
	var __rv QComboBox_SizeAdjustPolicy
	q.Drv(217000,217191,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QComboBox::sizeHint()
func (q *QComboBox) SizeHint() *QSize {
	var __rv uintptr
	q.Drv(217000,217192,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QComboBox::validator()
func (q *QComboBox) Validator() *QValidator {
	var __rv uintptr
	q.Drv(217000,217193,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QValidator{}
	_rp.SetDriver(__rv,393000,false)
	return _rp
}	
//QComboBox::view()
func (q *QComboBox) View() *QAbstractItemView {
	var __rv uintptr
	q.Drv(217000,217194,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAbstractItemView{}
	_rp.SetDriver(__rv,196000,false)
	return _rp
}	
//QComboBox::wheelEvent(QWheelEvent*)
func (q *QComboBox) WheelEvent(e *QWheelEvent)  {
	q.Drv(217000,217195,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
