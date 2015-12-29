// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QClipboard_Mode - QClipboard::Mode
type QClipboard_Mode uint32
const (
	QClipboard_Clipboard QClipboard_Mode = 0
	QClipboard_Selection QClipboard_Mode = 1
	QClipboard_FindBuffer QClipboard_Mode = 2
	QClipboard_LastMode QClipboard_Mode = QClipboard_FindBuffer
)
//struct QClipboard : QClipboard
type QClipboard struct {
	QObject
}
func (q *QClipboard) OnFindBufferChanged(fn func()) uintptr {
	var __rv uintptr
	q.Drv(214000,214102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QClipboard) OnSelectionChanged(fn func()) uintptr {
	var __rv uintptr
	q.Drv(214000,214103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QClipboard) OnDataChanged(fn func()) uintptr {
	var __rv uintptr
	q.Drv(214000,214104,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QClipboard) OnChanged(fn func(QClipboard_Mode)) uintptr {
	var __rv uintptr
	q.Drv(214000,214105,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QClipboard::clear()
func (q *QClipboard) Clear()  {
	q.Drv(214000,214106,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QClipboard::clear(QClipboard::Mode)
func (q *QClipboard) ClearWithMode(mode QClipboard_Mode)  {
	q.Drv(214000,214107,unsafe.Pointer(&mode),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QClipboard::event(QEvent*)
func (q *QClipboard) Event(value *QEvent) bool {
	var __rv bool
	q.Drv(214000,214108,Native(value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QClipboard::image()
func (q *QClipboard) Image() *QImage {
	var __rv uintptr
	q.Drv(214000,214109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QImage{}
	_rp.SetDriver(__rv,53000,true)
	return _rp
}	
//QClipboard::image(QClipboard::Mode)
func (q *QClipboard) ImageWithMode(mode QClipboard_Mode) *QImage {
	var __rv uintptr
	q.Drv(214000,214110,unsafe.Pointer(&mode),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QImage{}
	_rp.SetDriver(__rv,53000,true)
	return _rp
}	
//QClipboard::mimeData()
func (q *QClipboard) MimeData() *QMimeData {
	var __rv uintptr
	q.Drv(214000,214111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QMimeData{}
	_rp.SetDriver(__rv,311000,false)
	return _rp
}	
//QClipboard::mimeData(QClipboard::Mode)
func (q *QClipboard) MimeDataWithMode(mode QClipboard_Mode) *QMimeData {
	var __rv uintptr
	q.Drv(214000,214112,unsafe.Pointer(&mode),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QMimeData{}
	_rp.SetDriver(__rv,311000,false)
	return _rp
}	
//QClipboard::ownsClipboard()
func (q *QClipboard) OwnsClipboard() bool {
	var __rv bool
	q.Drv(214000,214113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QClipboard::ownsFindBuffer()
func (q *QClipboard) OwnsFindBuffer() bool {
	var __rv bool
	q.Drv(214000,214114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QClipboard::ownsSelection()
func (q *QClipboard) OwnsSelection() bool {
	var __rv bool
	q.Drv(214000,214115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QClipboard::pixmap()
func (q *QClipboard) Pixmap() *QPixmap {
	var __rv uintptr
	q.Drv(214000,214116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPixmap{}
	_rp.SetDriver(__rv,96000,true)
	return _rp
}	
//QClipboard::pixmap(QClipboard::Mode)
func (q *QClipboard) PixmapWithMode(mode QClipboard_Mode) *QPixmap {
	var __rv uintptr
	q.Drv(214000,214117,unsafe.Pointer(&mode),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPixmap{}
	_rp.SetDriver(__rv,96000,true)
	return _rp
}	
//QClipboard::setImage(QImage const&)
func (q *QClipboard) SetImage(value *QImage)  {
	q.Drv(214000,214118,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QClipboard::setImage(QImage const&,QClipboard::Mode)
func (q *QClipboard) SetImageWithImageMode(value2 *QImage,mode QClipboard_Mode)  {
	q.Drv(214000,214119,Native(value2),unsafe.Pointer(&mode),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QClipboard::setMimeData(QMimeData*)
func (q *QClipboard) SetMimeData(data *QMimeData)  {
	q.Drv(214000,214120,Native(data),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QClipboard::setMimeData(QMimeData*,QClipboard::Mode)
func (q *QClipboard) SetMimeDataWithDataMode(data *QMimeData,mode QClipboard_Mode)  {
	q.Drv(214000,214121,Native(data),unsafe.Pointer(&mode),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QClipboard::setPixmap(QPixmap const&)
func (q *QClipboard) SetPixmap(value *QPixmap)  {
	q.Drv(214000,214122,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QClipboard::setPixmap(QPixmap const&,QClipboard::Mode)
func (q *QClipboard) SetPixmapWithPixmapMode(value2 *QPixmap,mode QClipboard_Mode)  {
	q.Drv(214000,214123,Native(value2),unsafe.Pointer(&mode),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QClipboard::setText(QString const&)
func (q *QClipboard) SetText(value string)  {
	q.Drv(214000,214124,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QClipboard::setText(QString const&,QClipboard::Mode)
func (q *QClipboard) SetTextWithStringMode(value2 string,mode QClipboard_Mode)  {
	q.Drv(214000,214125,unsafe.Pointer(&value2),unsafe.Pointer(&mode),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QClipboard::supportsFindBuffer()
func (q *QClipboard) SupportsFindBuffer() bool {
	var __rv bool
	q.Drv(214000,214126,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QClipboard::supportsSelection()
func (q *QClipboard) SupportsSelection() bool {
	var __rv bool
	q.Drv(214000,214127,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QClipboard::text()
func (q *QClipboard) Text() string {
	var __rv string
	q.Drv(214000,214128,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QClipboard::text(QClipboard::Mode)
func (q *QClipboard) TextWithMode(mode QClipboard_Mode) string {
	var __rv string
	q.Drv(214000,214129,unsafe.Pointer(&mode),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QClipboard::text(QString&)
func (q *QClipboard) TextWithSubtype(subtype *string) string {
	var __rv string
	q.Drv(214000,214130,unsafe.Pointer(&subtype),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QClipboard::text(QString&,QClipboard::Mode)
func (q *QClipboard) TextWithSubtypeMode(subtype *string,mode QClipboard_Mode) string {
	var __rv string
	q.Drv(214000,214131,unsafe.Pointer(&subtype),unsafe.Pointer(&mode),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
