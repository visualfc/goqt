// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QColorDialog_ColorDialogOption - QColorDialog::ColorDialogOption
type QColorDialog_ColorDialogOption uint32
const (
	QColorDialog_ShowAlphaChannel QColorDialog_ColorDialogOption = 0x00000001
	QColorDialog_NoButtons QColorDialog_ColorDialogOption = 0x00000002
	QColorDialog_DontUseNativeDialog QColorDialog_ColorDialogOption = 0x00000004
)
//struct QColorDialog : QColorDialog
type QColorDialog struct {
	QDialog
}
func (q *QColorDialog) OnCurrentColorChanged(fn func(*QColor)) uintptr {
	var __rv uintptr
	q.Drv(215000,215102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QColorDialog) OnColorSelected(fn func(*QColor)) uintptr {
	var __rv uintptr
	q.Drv(215000,215103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QColorDialog::QColorDialog()	
func NewColorDialog() *QColorDialog {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),215000,215104,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QColorDialog{}
	_p.SetDriver(__rv,215000,false)
	return _p
} 
//QColorDialog::QColorDialog(QWidget*)	
func NewColorDialogWithParent(parent QWidgetInterface) *QColorDialog {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),215000,215105,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QColorDialog{}
	_p.SetDriver(__rv,215000,false)
	return _p
} 
//QColorDialog::QColorDialog(QColor const&,QWidget*)	
func NewColorDialogWithColorParent(initial *QColor,parent QWidgetInterface) *QColorDialog {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),215000,215106,Native(initial),Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QColorDialog{}
	_p.SetDriver(__rv,215000,false)
	return _p
} 
//QColorDialog::changeEvent(QEvent*)
func (q *QColorDialog) ChangeEvent(event *QEvent)  {
	q.Drv(215000,215107,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QColorDialog::currentColor()
func (q *QColorDialog) CurrentColor() *QColor {
	var __rv uintptr
	q.Drv(215000,215108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QColor{}
	_rp.SetDriver(__rv,13000,true)
	return _rp
}	
//QColorDialog::customColor(int)	
func QColorDialogCustomColor(index int) *QColor {
	var __rv uintptr
	DirectQtDrv(nil,215000,215109,unsafe.Pointer(&index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QColor{}
	_rp.SetDriver(__rv,13000,true)
	return _rp
}	
//QColorDialog::customColor(int)
func (q *QColorDialog) CustomColor(index int) *QColor {
	var __rv uintptr
	q.Drv(215000,215109,unsafe.Pointer(&index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QColor{}
	_rp.SetDriver(__rv,13000,true)
	return _rp
}	
//QColorDialog::customCount()	
func QColorDialogCustomCount() int {
	var __rv int
	DirectQtDrv(nil,215000,215110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QColorDialog::customCount()
func (q *QColorDialog) CustomCount() int {
	var __rv int
	q.Drv(215000,215110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QColorDialog::done(int)
func (q *QColorDialog) Done(result int)  {
	q.Drv(215000,215111,unsafe.Pointer(&result),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QColorDialog::getColor()	
func QColorDialogGetColor() *QColor {
	var __rv uintptr
	DirectQtDrv(nil,215000,215112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QColor{}
	_rp.SetDriver(__rv,13000,true)
	return _rp
}	
//QColorDialog::getColor()
func (q *QColorDialog) GetColor() *QColor {
	var __rv uintptr
	q.Drv(215000,215112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QColor{}
	_rp.SetDriver(__rv,13000,true)
	return _rp
}	
//QColorDialog::getColor(QColor const&,QWidget*)	
func QColorDialogGetColorWithColorParent(initial *QColor,parent QWidgetInterface) *QColor {
	var __rv uintptr
	DirectQtDrv(nil,215000,215113,Native(initial),Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QColor{}
	_rp.SetDriver(__rv,13000,true)
	return _rp
}	
//QColorDialog::getColor(QColor const&,QWidget*)
func (q *QColorDialog) GetColorWithColorParent(initial *QColor,parent QWidgetInterface) *QColor {
	var __rv uintptr
	q.Drv(215000,215113,Native(initial),Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QColor{}
	_rp.SetDriver(__rv,13000,true)
	return _rp
}	
//QColorDialog::getColor(QColor const&,QWidget*,QString const&,QFlags<QColorDialog::ColorDialogOption>)	
func QColorDialogGetColorWithColorParentTitleOptions(initial *QColor,parent QWidgetInterface,title string,options QColorDialog_ColorDialogOption) *QColor {
	var __rv uintptr
	DirectQtDrv(nil,215000,215114,Native(initial),Native(parent),unsafe.Pointer(&title),unsafe.Pointer(&options),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QColor{}
	_rp.SetDriver(__rv,13000,true)
	return _rp
}	
//QColorDialog::getColor(QColor const&,QWidget*,QString const&,QFlags<QColorDialog::ColorDialogOption>)
func (q *QColorDialog) GetColorWithColorParentTitleOptions(initial *QColor,parent QWidgetInterface,title string,options QColorDialog_ColorDialogOption) *QColor {
	var __rv uintptr
	q.Drv(215000,215114,Native(initial),Native(parent),unsafe.Pointer(&title),unsafe.Pointer(&options),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QColor{}
	_rp.SetDriver(__rv,13000,true)
	return _rp
}	
//QColorDialog::getRgba()	
func QColorDialogGetRgba() uint {
	var __rv uint
	DirectQtDrv(nil,215000,215115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QColorDialog::getRgba()
func (q *QColorDialog) GetRgba() uint {
	var __rv uint
	q.Drv(215000,215115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QColorDialog::getRgba(unsigned int,bool*,QWidget*)	
func QColorDialogGetRgbaWithRgbaOkParent(rgba uint,ok *bool,parent QWidgetInterface) uint {
	var __rv uint
	DirectQtDrv(nil,215000,215116,unsafe.Pointer(&rgba),unsafe.Pointer(&ok),Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QColorDialog::getRgba(unsigned int,bool*,QWidget*)
func (q *QColorDialog) GetRgbaWithRgbaOkParent(rgba uint,ok *bool,parent QWidgetInterface) uint {
	var __rv uint
	q.Drv(215000,215116,unsafe.Pointer(&rgba),unsafe.Pointer(&ok),Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QColorDialog::open()
func (q *QColorDialog) Open()  {
	q.Drv(215000,215117,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QColorDialog::open(QObject*,char const*)
func (q *QColorDialog) OpenWithObjectMember(receiver QObjectInterface,member string)  {
	q.Drv(215000,215118,Native(receiver),unsafe.Pointer(&member),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QColorDialog::options()
func (q *QColorDialog) Options() QColorDialog_ColorDialogOption {
	var __rv QColorDialog_ColorDialogOption
	q.Drv(215000,215119,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QColorDialog::selectedColor()
func (q *QColorDialog) SelectedColor() *QColor {
	var __rv uintptr
	q.Drv(215000,215120,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QColor{}
	_rp.SetDriver(__rv,13000,true)
	return _rp
}	
//QColorDialog::setCurrentColor(QColor const&)
func (q *QColorDialog) SetCurrentColor(color *QColor)  {
	q.Drv(215000,215121,Native(color),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QColorDialog::setCustomColor(int,QColor)	
func QColorDialogSetCustomColor(index int,color *QColor)  {
	DirectQtDrv(nil,215000,215122,unsafe.Pointer(&index),Native(color),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QColorDialog::setCustomColor(int,QColor)
func (q *QColorDialog) SetCustomColor(index int,color *QColor)  {
	q.Drv(215000,215122,unsafe.Pointer(&index),Native(color),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QColorDialog::setOption(QColorDialog::ColorDialogOption)
func (q *QColorDialog) SetOption(option QColorDialog_ColorDialogOption)  {
	q.Drv(215000,215123,unsafe.Pointer(&option),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QColorDialog::setOption(QColorDialog::ColorDialogOption,bool)
func (q *QColorDialog) SetOptionWithOptionOn(option QColorDialog_ColorDialogOption,on bool)  {
	q.Drv(215000,215124,unsafe.Pointer(&option),unsafe.Pointer(&on),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QColorDialog::setOptions(QFlags<QColorDialog::ColorDialogOption>)
func (q *QColorDialog) SetOptions(options QColorDialog_ColorDialogOption)  {
	q.Drv(215000,215125,unsafe.Pointer(&options),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QColorDialog::setStandardColor(int,QColor)	
func QColorDialogSetStandardColor(index int,color *QColor)  {
	DirectQtDrv(nil,215000,215126,unsafe.Pointer(&index),Native(color),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QColorDialog::setStandardColor(int,QColor)
func (q *QColorDialog) SetStandardColor(index int,color *QColor)  {
	q.Drv(215000,215126,unsafe.Pointer(&index),Native(color),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QColorDialog::setVisible(bool)
func (q *QColorDialog) SetVisible(visible bool)  {
	q.Drv(215000,215127,unsafe.Pointer(&visible),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QColorDialog::testOption(QColorDialog::ColorDialogOption)
func (q *QColorDialog) TestOption(option QColorDialog_ColorDialogOption) bool {
	var __rv bool
	q.Drv(215000,215128,unsafe.Pointer(&option),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
