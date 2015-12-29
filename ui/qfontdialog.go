// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QFontDialog_FontDialogOption - QFontDialog::FontDialogOption
type QFontDialog_FontDialogOption uint32
const (
	QFontDialog_NoButtons QFontDialog_FontDialogOption = 0x00000001
	QFontDialog_DontUseNativeDialog QFontDialog_FontDialogOption = 0x00000002
)
//struct QFontDialog : QFontDialog
type QFontDialog struct {
	QDialog
}
func (q *QFontDialog) OnFontSelected(fn func(*QFont)) uintptr {
	var __rv uintptr
	q.Drv(243000,243102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QFontDialog) OnCurrentFontChanged(fn func(*QFont)) uintptr {
	var __rv uintptr
	q.Drv(243000,243103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QFontDialog::QFontDialog()	
func NewFontDialog() *QFontDialog {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),243000,243104,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QFontDialog{}
	_p.SetDriver(__rv,243000,false)
	return _p
} 
//QFontDialog::QFontDialog(QWidget*)	
func NewFontDialogWithParent(parent QWidgetInterface) *QFontDialog {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),243000,243105,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QFontDialog{}
	_p.SetDriver(__rv,243000,false)
	return _p
} 
//QFontDialog::QFontDialog(QFont const&,QWidget*)	
func NewFontDialogWithInitialParent(initial *QFont,parent QWidgetInterface) *QFontDialog {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),243000,243106,Native(initial),Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QFontDialog{}
	_p.SetDriver(__rv,243000,false)
	return _p
} 
//QFontDialog::changeEvent(QEvent*)
func (q *QFontDialog) ChangeEvent(event *QEvent)  {
	q.Drv(243000,243107,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFontDialog::currentFont()
func (q *QFontDialog) CurrentFont() *QFont {
	var __rv uintptr
	q.Drv(243000,243108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QFont{}
	_rp.SetDriver(__rv,37000,true)
	return _rp
}	
//QFontDialog::done(int)
func (q *QFontDialog) Done(result int)  {
	q.Drv(243000,243109,unsafe.Pointer(&result),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFontDialog::getFont(bool*)	
func QFontDialogGetFont(ok *bool) *QFont {
	var __rv uintptr
	DirectQtDrv(nil,243000,243110,unsafe.Pointer(&ok),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QFont{}
	_rp.SetDriver(__rv,37000,true)
	return _rp
}	
//QFontDialog::getFont(bool*)
func (q *QFontDialog) GetFont(ok *bool) *QFont {
	var __rv uintptr
	q.Drv(243000,243110,unsafe.Pointer(&ok),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QFont{}
	_rp.SetDriver(__rv,37000,true)
	return _rp
}	
//QFontDialog::getFont(bool*,QWidget*)	
func QFontDialogGetFontWithOkParent(ok *bool,parent QWidgetInterface) *QFont {
	var __rv uintptr
	DirectQtDrv(nil,243000,243111,unsafe.Pointer(&ok),Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QFont{}
	_rp.SetDriver(__rv,37000,true)
	return _rp
}	
//QFontDialog::getFont(bool*,QWidget*)
func (q *QFontDialog) GetFontWithOkParent(ok *bool,parent QWidgetInterface) *QFont {
	var __rv uintptr
	q.Drv(243000,243111,unsafe.Pointer(&ok),Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QFont{}
	_rp.SetDriver(__rv,37000,true)
	return _rp
}	
//QFontDialog::getFont(bool*,QFont const&,QWidget*)	
func QFontDialogGetFontWithOkInitialParent(ok *bool,initial *QFont,parent QWidgetInterface) *QFont {
	var __rv uintptr
	DirectQtDrv(nil,243000,243112,unsafe.Pointer(&ok),Native(initial),Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QFont{}
	_rp.SetDriver(__rv,37000,true)
	return _rp
}	
//QFontDialog::getFont(bool*,QFont const&,QWidget*)
func (q *QFontDialog) GetFontWithOkInitialParent(ok *bool,initial *QFont,parent QWidgetInterface) *QFont {
	var __rv uintptr
	q.Drv(243000,243112,unsafe.Pointer(&ok),Native(initial),Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QFont{}
	_rp.SetDriver(__rv,37000,true)
	return _rp
}	
//QFontDialog::getFont(bool*,QFont const&,QWidget*,QString const&)	
func QFontDialogGetFontWithOkInitialParentTitle(ok *bool,initial *QFont,parent QWidgetInterface,title string) *QFont {
	var __rv uintptr
	DirectQtDrv(nil,243000,243113,unsafe.Pointer(&ok),Native(initial),Native(parent),unsafe.Pointer(&title),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QFont{}
	_rp.SetDriver(__rv,37000,true)
	return _rp
}	
//QFontDialog::getFont(bool*,QFont const&,QWidget*,QString const&)
func (q *QFontDialog) GetFontWithOkInitialParentTitle(ok *bool,initial *QFont,parent QWidgetInterface,title string) *QFont {
	var __rv uintptr
	q.Drv(243000,243113,unsafe.Pointer(&ok),Native(initial),Native(parent),unsafe.Pointer(&title),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QFont{}
	_rp.SetDriver(__rv,37000,true)
	return _rp
}	
//QFontDialog::getFont(bool*,QFont const&,QWidget*,QString const&,QFlags<QFontDialog::FontDialogOption>)	
func QFontDialogGetFontWithOkInitialParentTitleOptions(ok *bool,initial *QFont,parent QWidgetInterface,title string,options QFontDialog_FontDialogOption) *QFont {
	var __rv uintptr
	DirectQtDrv(nil,243000,243114,unsafe.Pointer(&ok),Native(initial),Native(parent),unsafe.Pointer(&title),unsafe.Pointer(&options),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QFont{}
	_rp.SetDriver(__rv,37000,true)
	return _rp
}	
//QFontDialog::getFont(bool*,QFont const&,QWidget*,QString const&,QFlags<QFontDialog::FontDialogOption>)
func (q *QFontDialog) GetFontWithOkInitialParentTitleOptions(ok *bool,initial *QFont,parent QWidgetInterface,title string,options QFontDialog_FontDialogOption) *QFont {
	var __rv uintptr
	q.Drv(243000,243114,unsafe.Pointer(&ok),Native(initial),Native(parent),unsafe.Pointer(&title),unsafe.Pointer(&options),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QFont{}
	_rp.SetDriver(__rv,37000,true)
	return _rp
}	
//QFontDialog::open()
func (q *QFontDialog) Open()  {
	q.Drv(243000,243115,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFontDialog::open(QObject*,char const*)
func (q *QFontDialog) OpenWithObjectMember(receiver QObjectInterface,member string)  {
	q.Drv(243000,243116,Native(receiver),unsafe.Pointer(&member),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFontDialog::options()
func (q *QFontDialog) Options() QFontDialog_FontDialogOption {
	var __rv QFontDialog_FontDialogOption
	q.Drv(243000,243117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontDialog::selectedFont()
func (q *QFontDialog) SelectedFont() *QFont {
	var __rv uintptr
	q.Drv(243000,243118,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QFont{}
	_rp.SetDriver(__rv,37000,true)
	return _rp
}	
//QFontDialog::setCurrentFont(QFont const&)
func (q *QFontDialog) SetCurrentFont(font *QFont)  {
	q.Drv(243000,243119,Native(font),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFontDialog::setOption(QFontDialog::FontDialogOption)
func (q *QFontDialog) SetOption(option QFontDialog_FontDialogOption)  {
	q.Drv(243000,243120,unsafe.Pointer(&option),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFontDialog::setOption(QFontDialog::FontDialogOption,bool)
func (q *QFontDialog) SetOptionWithOptionOn(option QFontDialog_FontDialogOption,on bool)  {
	q.Drv(243000,243121,unsafe.Pointer(&option),unsafe.Pointer(&on),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFontDialog::setOptions(QFlags<QFontDialog::FontDialogOption>)
func (q *QFontDialog) SetOptions(options QFontDialog_FontDialogOption)  {
	q.Drv(243000,243122,unsafe.Pointer(&options),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFontDialog::setVisible(bool)
func (q *QFontDialog) SetVisible(visible bool)  {
	q.Drv(243000,243123,unsafe.Pointer(&visible),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFontDialog::testOption(QFontDialog::FontDialogOption)
func (q *QFontDialog) TestOption(option QFontDialog_FontDialogOption) bool {
	var __rv bool
	q.Drv(243000,243124,unsafe.Pointer(&option),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
