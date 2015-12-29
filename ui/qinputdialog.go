// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QInputDialog_InputMode - QInputDialog::InputMode
type QInputDialog_InputMode uint32
const (
	QInputDialog_TextInput QInputDialog_InputMode = 0
	QInputDialog_IntInput QInputDialog_InputMode = 1
	QInputDialog_DoubleInput QInputDialog_InputMode = 2
)
//enum QInputDialog_InputDialogOption - QInputDialog::InputDialogOption
type QInputDialog_InputDialogOption uint32
const (
	QInputDialog_NoButtons QInputDialog_InputDialogOption = 0x00000001
	QInputDialog_UseListViewForComboBoxItems QInputDialog_InputDialogOption = 0x00000002
)
//struct QInputDialog : QInputDialog
type QInputDialog struct {
	QDialog
}
func (q *QInputDialog) OnDoubleValueSelected(fn func(float64)) uintptr {
	var __rv uintptr
	q.Drv(293000,293102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QInputDialog) OnDoubleValueChanged(fn func(float64)) uintptr {
	var __rv uintptr
	q.Drv(293000,293103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QInputDialog) OnTextValueChanged(fn func(string)) uintptr {
	var __rv uintptr
	q.Drv(293000,293104,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QInputDialog) OnIntValueChanged(fn func(int)) uintptr {
	var __rv uintptr
	q.Drv(293000,293105,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QInputDialog) OnIntValueSelected(fn func(int)) uintptr {
	var __rv uintptr
	q.Drv(293000,293106,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QInputDialog) OnTextValueSelected(fn func(string)) uintptr {
	var __rv uintptr
	q.Drv(293000,293107,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QInputDialog::QInputDialog()	
func NewInputDialog() *QInputDialog {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),293000,293108,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QInputDialog{}
	_p.SetDriver(__rv,293000,false)
	return _p
} 
//QInputDialog::QInputDialog(QWidget*,QFlags<Qt::WindowType>)	
func NewInputDialogWithParentFlags(parent QWidgetInterface,flags Qt_WindowType) *QInputDialog {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),293000,293109,Native(parent),unsafe.Pointer(&flags),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QInputDialog{}
	_p.SetDriver(__rv,293000,false)
	return _p
} 
//QInputDialog::cancelButtonText()
func (q *QInputDialog) CancelButtonText() string {
	var __rv string
	q.Drv(293000,293110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QInputDialog::comboBoxItems()
func (q *QInputDialog) ComboBoxItems() []string {
	var __rv []string
	q.Drv(293000,293111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QInputDialog::done(int)
func (q *QInputDialog) Done(result int)  {
	q.Drv(293000,293112,unsafe.Pointer(&result),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QInputDialog::doubleDecimals()
func (q *QInputDialog) DoubleDecimals() int {
	var __rv int
	q.Drv(293000,293113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QInputDialog::doubleMaximum()
func (q *QInputDialog) DoubleMaximum() float64 {
	var __rv float64
	q.Drv(293000,293114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QInputDialog::doubleMinimum()
func (q *QInputDialog) DoubleMinimum() float64 {
	var __rv float64
	q.Drv(293000,293115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QInputDialog::doubleValue()
func (q *QInputDialog) DoubleValue() float64 {
	var __rv float64
	q.Drv(293000,293116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QInputDialog::getDouble(QWidget*,QString const&,QString const&,double,double,double,int,bool*,QFlags<Qt::WindowType>)	
func QInputDialogGetDouble(parent QWidgetInterface,title string,label string,value float64,minValue float64,maxValue float64,decimals int,ok *bool,flags Qt_WindowType) float64 {
	var __rv float64
	DirectQtDrv(nil,293000,293117,Native(parent),unsafe.Pointer(&title),unsafe.Pointer(&label),unsafe.Pointer(&value),unsafe.Pointer(&minValue),unsafe.Pointer(&maxValue),unsafe.Pointer(&decimals),unsafe.Pointer(&ok),unsafe.Pointer(&flags),unsafe.Pointer(&__rv),nil,nil)
	return __rv
}	
//QInputDialog::getDouble(QWidget*,QString const&,QString const&,double,double,double,int,bool*,QFlags<Qt::WindowType>)
func (q *QInputDialog) GetDouble(parent QWidgetInterface,title string,label string,value float64,minValue float64,maxValue float64,decimals int,ok *bool,flags Qt_WindowType) float64 {
	var __rv float64
	q.Drv(293000,293117,Native(parent),unsafe.Pointer(&title),unsafe.Pointer(&label),unsafe.Pointer(&value),unsafe.Pointer(&minValue),unsafe.Pointer(&maxValue),unsafe.Pointer(&decimals),unsafe.Pointer(&ok),unsafe.Pointer(&flags),unsafe.Pointer(&__rv),nil,nil)
	return __rv
}	
//QInputDialog::getInt(QWidget*,QString const&,QString const&,int,int,int,int,bool*,QFlags<Qt::WindowType>)	
func QInputDialogGetInt(parent QWidgetInterface,title string,label string,value int,minValue int,maxValue int,step int,ok *bool,flags Qt_WindowType) int {
	var __rv int
	DirectQtDrv(nil,293000,293118,Native(parent),unsafe.Pointer(&title),unsafe.Pointer(&label),unsafe.Pointer(&value),unsafe.Pointer(&minValue),unsafe.Pointer(&maxValue),unsafe.Pointer(&step),unsafe.Pointer(&ok),unsafe.Pointer(&flags),unsafe.Pointer(&__rv),nil,nil)
	return __rv
}	
//QInputDialog::getInt(QWidget*,QString const&,QString const&,int,int,int,int,bool*,QFlags<Qt::WindowType>)
func (q *QInputDialog) GetInt(parent QWidgetInterface,title string,label string,value int,minValue int,maxValue int,step int,ok *bool,flags Qt_WindowType) int {
	var __rv int
	q.Drv(293000,293118,Native(parent),unsafe.Pointer(&title),unsafe.Pointer(&label),unsafe.Pointer(&value),unsafe.Pointer(&minValue),unsafe.Pointer(&maxValue),unsafe.Pointer(&step),unsafe.Pointer(&ok),unsafe.Pointer(&flags),unsafe.Pointer(&__rv),nil,nil)
	return __rv
}	
//QInputDialog::getItem(QWidget*,QString const&,QString const&,QStringList const&,int,bool,bool*,QFlags<Qt::WindowType>)	
func QInputDialogGetItem(parent QWidgetInterface,title string,label string,items []string,current int,editable bool,ok *bool,flags Qt_WindowType) string {
	var __rv string
	DirectQtDrv(nil,293000,293119,Native(parent),unsafe.Pointer(&title),unsafe.Pointer(&label),unsafe.Pointer(&items),unsafe.Pointer(&current),unsafe.Pointer(&editable),unsafe.Pointer(&ok),unsafe.Pointer(&flags),unsafe.Pointer(&__rv),nil,nil,nil)
	return __rv
}	
//QInputDialog::getItem(QWidget*,QString const&,QString const&,QStringList const&,int,bool,bool*,QFlags<Qt::WindowType>)
func (q *QInputDialog) GetItem(parent QWidgetInterface,title string,label string,items []string,current int,editable bool,ok *bool,flags Qt_WindowType) string {
	var __rv string
	q.Drv(293000,293119,Native(parent),unsafe.Pointer(&title),unsafe.Pointer(&label),unsafe.Pointer(&items),unsafe.Pointer(&current),unsafe.Pointer(&editable),unsafe.Pointer(&ok),unsafe.Pointer(&flags),unsafe.Pointer(&__rv),nil,nil,nil)
	return __rv
}	
//QInputDialog::getText(QWidget*,QString const&,QString const&,QLineEdit::EchoMode,QString const&,bool*,QFlags<Qt::WindowType>)	
func QInputDialogGetText(parent QWidgetInterface,title string,label string,echo QLineEdit_EchoMode,text string,ok *bool,flags Qt_WindowType) string {
	var __rv string
	DirectQtDrv(nil,293000,293120,Native(parent),unsafe.Pointer(&title),unsafe.Pointer(&label),unsafe.Pointer(&echo),unsafe.Pointer(&text),unsafe.Pointer(&ok),unsafe.Pointer(&flags),unsafe.Pointer(&__rv),nil,nil,nil,nil)
	return __rv
}	
//QInputDialog::getText(QWidget*,QString const&,QString const&,QLineEdit::EchoMode,QString const&,bool*,QFlags<Qt::WindowType>)
func (q *QInputDialog) GetText(parent QWidgetInterface,title string,label string,echo QLineEdit_EchoMode,text string,ok *bool,flags Qt_WindowType) string {
	var __rv string
	q.Drv(293000,293120,Native(parent),unsafe.Pointer(&title),unsafe.Pointer(&label),unsafe.Pointer(&echo),unsafe.Pointer(&text),unsafe.Pointer(&ok),unsafe.Pointer(&flags),unsafe.Pointer(&__rv),nil,nil,nil,nil)
	return __rv
}	
//QInputDialog::inputMode()
func (q *QInputDialog) InputMode() QInputDialog_InputMode {
	var __rv QInputDialog_InputMode
	q.Drv(293000,293121,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QInputDialog::intMaximum()
func (q *QInputDialog) IntMaximum() int {
	var __rv int
	q.Drv(293000,293122,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QInputDialog::intMinimum()
func (q *QInputDialog) IntMinimum() int {
	var __rv int
	q.Drv(293000,293123,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QInputDialog::intStep()
func (q *QInputDialog) IntStep() int {
	var __rv int
	q.Drv(293000,293124,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QInputDialog::intValue()
func (q *QInputDialog) IntValue() int {
	var __rv int
	q.Drv(293000,293125,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QInputDialog::isComboBoxEditable()
func (q *QInputDialog) IsComboBoxEditable() bool {
	var __rv bool
	q.Drv(293000,293126,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QInputDialog::labelText()
func (q *QInputDialog) LabelText() string {
	var __rv string
	q.Drv(293000,293127,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QInputDialog::minimumSizeHint()
func (q *QInputDialog) MinimumSizeHint() *QSize {
	var __rv uintptr
	q.Drv(293000,293128,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QInputDialog::okButtonText()
func (q *QInputDialog) OkButtonText() string {
	var __rv string
	q.Drv(293000,293129,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QInputDialog::open()
func (q *QInputDialog) Open()  {
	q.Drv(293000,293130,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QInputDialog::open(QObject*,char const*)
func (q *QInputDialog) OpenWithObjectMember(receiver QObjectInterface,member string)  {
	q.Drv(293000,293131,Native(receiver),unsafe.Pointer(&member),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QInputDialog::options()
func (q *QInputDialog) Options() QInputDialog_InputDialogOption {
	var __rv QInputDialog_InputDialogOption
	q.Drv(293000,293132,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QInputDialog::setCancelButtonText(QString const&)
func (q *QInputDialog) SetCancelButtonText(text string)  {
	q.Drv(293000,293133,unsafe.Pointer(&text),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QInputDialog::setComboBoxEditable(bool)
func (q *QInputDialog) SetComboBoxEditable(editable bool)  {
	q.Drv(293000,293134,unsafe.Pointer(&editable),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QInputDialog::setComboBoxItems(QStringList const&)
func (q *QInputDialog) SetComboBoxItems(items []string)  {
	q.Drv(293000,293135,unsafe.Pointer(&items),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QInputDialog::setDoubleDecimals(int)
func (q *QInputDialog) SetDoubleDecimals(decimals int)  {
	q.Drv(293000,293136,unsafe.Pointer(&decimals),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QInputDialog::setDoubleMaximum(double)
func (q *QInputDialog) SetDoubleMaximum(max float64)  {
	q.Drv(293000,293137,unsafe.Pointer(&max),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QInputDialog::setDoubleMinimum(double)
func (q *QInputDialog) SetDoubleMinimum(min float64)  {
	q.Drv(293000,293138,unsafe.Pointer(&min),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QInputDialog::setDoubleRange(double,double)
func (q *QInputDialog) SetDoubleRange(min float64,max float64)  {
	q.Drv(293000,293139,unsafe.Pointer(&min),unsafe.Pointer(&max),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QInputDialog::setDoubleValue(double)
func (q *QInputDialog) SetDoubleValue(value float64)  {
	q.Drv(293000,293140,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QInputDialog::setInputMode(QInputDialog::InputMode)
func (q *QInputDialog) SetInputMode(mode QInputDialog_InputMode)  {
	q.Drv(293000,293141,unsafe.Pointer(&mode),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QInputDialog::setIntMaximum(int)
func (q *QInputDialog) SetIntMaximum(max int)  {
	q.Drv(293000,293142,unsafe.Pointer(&max),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QInputDialog::setIntMinimum(int)
func (q *QInputDialog) SetIntMinimum(min int)  {
	q.Drv(293000,293143,unsafe.Pointer(&min),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QInputDialog::setIntRange(int,int)
func (q *QInputDialog) SetIntRange(min int,max int)  {
	q.Drv(293000,293144,unsafe.Pointer(&min),unsafe.Pointer(&max),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QInputDialog::setIntStep(int)
func (q *QInputDialog) SetIntStep(step int)  {
	q.Drv(293000,293145,unsafe.Pointer(&step),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QInputDialog::setIntValue(int)
func (q *QInputDialog) SetIntValue(value int)  {
	q.Drv(293000,293146,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QInputDialog::setLabelText(QString const&)
func (q *QInputDialog) SetLabelText(text string)  {
	q.Drv(293000,293147,unsafe.Pointer(&text),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QInputDialog::setOkButtonText(QString const&)
func (q *QInputDialog) SetOkButtonText(text string)  {
	q.Drv(293000,293148,unsafe.Pointer(&text),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QInputDialog::setOption(QInputDialog::InputDialogOption)
func (q *QInputDialog) SetOption(option QInputDialog_InputDialogOption)  {
	q.Drv(293000,293149,unsafe.Pointer(&option),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QInputDialog::setOption(QInputDialog::InputDialogOption,bool)
func (q *QInputDialog) SetOptionWithOptionOn(option QInputDialog_InputDialogOption,on bool)  {
	q.Drv(293000,293150,unsafe.Pointer(&option),unsafe.Pointer(&on),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QInputDialog::setOptions(QFlags<QInputDialog::InputDialogOption>)
func (q *QInputDialog) SetOptions(options QInputDialog_InputDialogOption)  {
	q.Drv(293000,293151,unsafe.Pointer(&options),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QInputDialog::setTextEchoMode(QLineEdit::EchoMode)
func (q *QInputDialog) SetTextEchoMode(mode QLineEdit_EchoMode)  {
	q.Drv(293000,293152,unsafe.Pointer(&mode),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QInputDialog::setTextValue(QString const&)
func (q *QInputDialog) SetTextValue(text string)  {
	q.Drv(293000,293153,unsafe.Pointer(&text),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QInputDialog::setVisible(bool)
func (q *QInputDialog) SetVisible(visible bool)  {
	q.Drv(293000,293154,unsafe.Pointer(&visible),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QInputDialog::sizeHint()
func (q *QInputDialog) SizeHint() *QSize {
	var __rv uintptr
	q.Drv(293000,293155,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QInputDialog::testOption(QInputDialog::InputDialogOption)
func (q *QInputDialog) TestOption(option QInputDialog_InputDialogOption) bool {
	var __rv bool
	q.Drv(293000,293156,unsafe.Pointer(&option),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QInputDialog::textEchoMode()
func (q *QInputDialog) TextEchoMode() QLineEdit_EchoMode {
	var __rv QLineEdit_EchoMode
	q.Drv(293000,293157,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QInputDialog::textValue()
func (q *QInputDialog) TextValue() string {
	var __rv string
	q.Drv(293000,293158,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
