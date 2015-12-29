// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QTextDecoder : QTextDecoder
type QTextDecoder struct {
	BaseDrv
}
//QTextDecoder::QTextDecoder(QTextCodec const*)	
func NewTextDecoder(codec *QTextCodec) *QTextDecoder {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),146000,146102,Native(codec),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTextDecoder{}
	_p.SetDriver(__rv,146000,true)
	return _p
} 
//QTextDecoder::QTextDecoder(QTextCodec const*,QFlags<QTextCodec::ConversionFlag>)	
func NewTextDecoderWithCodecFlags(codec *QTextCodec,flags QTextCodec_ConversionFlag) *QTextDecoder {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),146000,146103,Native(codec),unsafe.Pointer(&flags),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTextDecoder{}
	_p.SetDriver(__rv,146000,true)
	return _p
} 
//QTextDecoder::hasFailure()
func (q *QTextDecoder) HasFailure() bool {
	var __rv bool
	q.Drv(146000,146104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextDecoder::toUnicode(QByteArray const&)
func (q *QTextDecoder) ToUnicode(ba []byte) string {
	var __rv string
	q.Drv(146000,146105,unsafe.Pointer(&ba),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
