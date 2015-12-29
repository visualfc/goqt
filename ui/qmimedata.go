// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QMimeData : QMimeData
type QMimeData struct {
	QObject
}
//QMimeData::QMimeData()	
func NewMimeData() *QMimeData {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),311000,311102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QMimeData{}
	_p.SetDriver(__rv,311000,false)
	return _p
} 
//QMimeData::clear()
func (q *QMimeData) Clear()  {
	q.Drv(311000,311103,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMimeData::colorData()
func (q *QMimeData) ColorData() *QVariant {
	var __rv uintptr
	q.Drv(311000,311104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVariant{}
	_rp.SetDriver(__rv,182000,true)
	return _rp
}	
//QMimeData::data(QString const&)
func (q *QMimeData) Data(mimetype string) []byte {
	var __rv []byte
	q.Drv(311000,311105,unsafe.Pointer(&mimetype),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMimeData::formats()
func (q *QMimeData) Formats() []string {
	var __rv []string
	q.Drv(311000,311106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMimeData::hasColor()
func (q *QMimeData) HasColor() bool {
	var __rv bool
	q.Drv(311000,311107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMimeData::hasFormat(QString const&)
func (q *QMimeData) HasFormat(mimetype string) bool {
	var __rv bool
	q.Drv(311000,311108,unsafe.Pointer(&mimetype),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMimeData::hasHtml()
func (q *QMimeData) HasHtml() bool {
	var __rv bool
	q.Drv(311000,311109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMimeData::hasImage()
func (q *QMimeData) HasImage() bool {
	var __rv bool
	q.Drv(311000,311110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMimeData::hasText()
func (q *QMimeData) HasText() bool {
	var __rv bool
	q.Drv(311000,311111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMimeData::hasUrls()
func (q *QMimeData) HasUrls() bool {
	var __rv bool
	q.Drv(311000,311112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMimeData::html()
func (q *QMimeData) Html() string {
	var __rv string
	q.Drv(311000,311113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMimeData::imageData()
func (q *QMimeData) ImageData() *QVariant {
	var __rv uintptr
	q.Drv(311000,311114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVariant{}
	_rp.SetDriver(__rv,182000,true)
	return _rp
}	
//QMimeData::removeFormat(QString const&)
func (q *QMimeData) RemoveFormat(mimetype string)  {
	q.Drv(311000,311115,unsafe.Pointer(&mimetype),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMimeData::retrieveData(QString const&,QVariant::Type)
func (q *QMimeData) RetrieveData(mimetype string,preferredType QVariant_Type) *QVariant {
	var __rv uintptr
	q.Drv(311000,311116,unsafe.Pointer(&mimetype),unsafe.Pointer(&preferredType),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVariant{}
	_rp.SetDriver(__rv,182000,true)
	return _rp
}	
//QMimeData::setColorData(QVariant const&)
func (q *QMimeData) SetColorData(color *QVariant)  {
	q.Drv(311000,311117,Native(color),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMimeData::setData(QString const&,QByteArray const&)
func (q *QMimeData) SetData(mimetype string,data []byte)  {
	q.Drv(311000,311118,unsafe.Pointer(&mimetype),unsafe.Pointer(&data),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMimeData::setHtml(QString const&)
func (q *QMimeData) SetHtml(html string)  {
	q.Drv(311000,311119,unsafe.Pointer(&html),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMimeData::setImageData(QVariant const&)
func (q *QMimeData) SetImageData(image *QVariant)  {
	q.Drv(311000,311120,Native(image),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMimeData::setText(QString const&)
func (q *QMimeData) SetText(text string)  {
	q.Drv(311000,311121,unsafe.Pointer(&text),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMimeData::setUrls(QList<QUrl> const&)
func (q *QMimeData) SetUrls(urls []*QUrl)  {
	q.Drv(311000,311122,unsafe.Pointer(&urls),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMimeData::text()
func (q *QMimeData) Text() string {
	var __rv string
	q.Drv(311000,311123,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMimeData::urls()
func (q *QMimeData) Urls() []*QUrl {
	var __rv []*QUrl
	q.Drv(311000,311124,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
