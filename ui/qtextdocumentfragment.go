// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QTextDocumentFragment : QTextDocumentFragment
type QTextDocumentFragment struct {
	BaseDrv
}
//QTextDocumentFragment::QTextDocumentFragment()	
func NewTextDocumentFragment() *QTextDocumentFragment {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),147000,147102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTextDocumentFragment{}
	_p.SetDriver(__rv,147000,true)
	return _p
} 
//QTextDocumentFragment::QTextDocumentFragment(QTextCursor const&)	
func NewTextDocumentFragmentWithRange(_range *QTextCursor) *QTextDocumentFragment {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),147000,147103,Native(_range),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTextDocumentFragment{}
	_p.SetDriver(__rv,147000,true)
	return _p
} 
//QTextDocumentFragment::QTextDocumentFragment(QTextDocument const*)	
func NewTextDocumentFragmentWithDocument(document *QTextDocument) *QTextDocumentFragment {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),147000,147104,Native(document),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTextDocumentFragment{}
	_p.SetDriver(__rv,147000,true)
	return _p
} 
//QTextDocumentFragment::QTextDocumentFragment(QTextDocumentFragment const&)	
func NewTextDocumentFragmentCopy(rhs *QTextDocumentFragment) *QTextDocumentFragment {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),147000,147105,Native(rhs),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTextDocumentFragment{}
	_p.SetDriver(__rv,147000,true)
	return _p
} 
//QTextDocumentFragment::fromHtml(QString const&)	
func QTextDocumentFragmentFromHtml(html string) *QTextDocumentFragment {
	var __rv uintptr
	DirectQtDrv(nil,147000,147106,unsafe.Pointer(&html),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextDocumentFragment{}
	_rp.SetDriver(__rv,147000,true)
	return _rp
}	
//QTextDocumentFragment::fromHtml(QString const&)
func (q *QTextDocumentFragment) FromHtml(html string) *QTextDocumentFragment {
	var __rv uintptr
	q.Drv(147000,147106,unsafe.Pointer(&html),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextDocumentFragment{}
	_rp.SetDriver(__rv,147000,true)
	return _rp
}	
//QTextDocumentFragment::fromHtml(QString const&,QTextDocument const*)	
func QTextDocumentFragmentFromHtmlWithHtmlResourceprovider(html string,resourceProvider *QTextDocument) *QTextDocumentFragment {
	var __rv uintptr
	DirectQtDrv(nil,147000,147107,unsafe.Pointer(&html),Native(resourceProvider),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextDocumentFragment{}
	_rp.SetDriver(__rv,147000,true)
	return _rp
}	
//QTextDocumentFragment::fromHtml(QString const&,QTextDocument const*)
func (q *QTextDocumentFragment) FromHtmlWithHtmlResourceprovider(html string,resourceProvider *QTextDocument) *QTextDocumentFragment {
	var __rv uintptr
	q.Drv(147000,147107,unsafe.Pointer(&html),Native(resourceProvider),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextDocumentFragment{}
	_rp.SetDriver(__rv,147000,true)
	return _rp
}	
//QTextDocumentFragment::fromPlainText(QString const&)	
func QTextDocumentFragmentFromPlainText(plainText string) *QTextDocumentFragment {
	var __rv uintptr
	DirectQtDrv(nil,147000,147108,unsafe.Pointer(&plainText),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextDocumentFragment{}
	_rp.SetDriver(__rv,147000,true)
	return _rp
}	
//QTextDocumentFragment::fromPlainText(QString const&)
func (q *QTextDocumentFragment) FromPlainText(plainText string) *QTextDocumentFragment {
	var __rv uintptr
	q.Drv(147000,147108,unsafe.Pointer(&plainText),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextDocumentFragment{}
	_rp.SetDriver(__rv,147000,true)
	return _rp
}	
//QTextDocumentFragment::isEmpty()
func (q *QTextDocumentFragment) IsEmpty() bool {
	var __rv bool
	q.Drv(147000,147109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextDocumentFragment::toHtml()
func (q *QTextDocumentFragment) ToHtml() string {
	var __rv string
	q.Drv(147000,147110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextDocumentFragment::toHtml(QByteArray const&)
func (q *QTextDocumentFragment) ToHtmlWithEncoding(encoding []byte) string {
	var __rv string
	q.Drv(147000,147111,unsafe.Pointer(&encoding),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextDocumentFragment::toPlainText()
func (q *QTextDocumentFragment) ToPlainText() string {
	var __rv string
	q.Drv(147000,147112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
