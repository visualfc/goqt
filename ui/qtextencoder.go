// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QTextEncoder : QTextEncoder
type QTextEncoder struct {
	BaseDrv
}
//QTextEncoder::QTextEncoder(QTextCodec const*)	
func NewTextEncoder(codec *QTextCodec) *QTextEncoder {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),150000,150102,Native(codec),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTextEncoder{}
	_p.SetDriver(__rv,150000,true)
	return _p
} 
//QTextEncoder::QTextEncoder(QTextCodec const*,QFlags<QTextCodec::ConversionFlag>)	
func NewTextEncoderWithCodecFlags(codec *QTextCodec,flags QTextCodec_ConversionFlag) *QTextEncoder {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),150000,150103,Native(codec),unsafe.Pointer(&flags),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTextEncoder{}
	_p.SetDriver(__rv,150000,true)
	return _p
} 
//QTextEncoder::fromUnicode(QString const&)
func (q *QTextEncoder) FromUnicode(str string) []byte {
	var __rv []byte
	q.Drv(150000,150104,unsafe.Pointer(&str),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextEncoder::fromUnicode(QChar const*,int)
func (q *QTextEncoder) FromUnicodeWithUcLen(uc *rune,len int) []byte {
	var __rv []byte
	q.Drv(150000,150105,unsafe.Pointer(&uc),unsafe.Pointer(&len),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextEncoder::hasFailure()
func (q *QTextEncoder) HasFailure() bool {
	var __rv bool
	q.Drv(150000,150106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
