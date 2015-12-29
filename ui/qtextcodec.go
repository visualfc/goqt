// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QTextCodec_ConversionFlag - QTextCodec::ConversionFlag
type QTextCodec_ConversionFlag uint32
const (
	QTextCodec_DefaultConversion QTextCodec_ConversionFlag = 0
	QTextCodec_ConvertInvalidToNull QTextCodec_ConversionFlag = 0x80000000
	QTextCodec_IgnoreHeader QTextCodec_ConversionFlag = 0x1
	QTextCodec_FreeFunction QTextCodec_ConversionFlag = 0x2
)
//struct QTextCodec : QTextCodec
type QTextCodec struct {
	BaseDrv
}
//QTextCodec::aliases()
func (q *QTextCodec) Aliases() [][]byte {
	var __rv [][]byte
	q.Drv(143000,143102,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextCodec::availableCodecs()	
func QTextCodecAvailableCodecs() [][]byte {
	var __rv [][]byte
	DirectQtDrv(nil,143000,143103,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextCodec::availableCodecs()
func (q *QTextCodec) AvailableCodecs() [][]byte {
	var __rv [][]byte
	q.Drv(143000,143103,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextCodec::availableMibs()	
func QTextCodecAvailableMibs() []int {
	var __rv []int
	DirectQtDrv(nil,143000,143104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextCodec::availableMibs()
func (q *QTextCodec) AvailableMibs() []int {
	var __rv []int
	q.Drv(143000,143104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextCodec::canEncode(QChar)
func (q *QTextCodec) CanEncode(value rune) bool {
	var __rv bool
	q.Drv(143000,143105,unsafe.Pointer(&value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextCodec::canEncode(QString const&)
func (q *QTextCodec) CanEncodeWithString(value string) bool {
	var __rv bool
	q.Drv(143000,143106,unsafe.Pointer(&value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextCodec::codecForHtml(QByteArray const&)	
func QTextCodecCodecForHtml(ba []byte) *QTextCodec {
	var __rv uintptr
	DirectQtDrv(nil,143000,143107,unsafe.Pointer(&ba),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextCodec{}
	_rp.SetDriver(__rv,143000,true)
	return _rp
}	
//QTextCodec::codecForHtml(QByteArray const&)
func (q *QTextCodec) CodecForHtml(ba []byte) *QTextCodec {
	var __rv uintptr
	q.Drv(143000,143107,unsafe.Pointer(&ba),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextCodec{}
	_rp.SetDriver(__rv,143000,true)
	return _rp
}	
//QTextCodec::codecForHtml(QByteArray const&,QTextCodec*)	
func QTextCodecCodecForHtmlWithBaDefaultcodec(ba []byte,defaultCodec *QTextCodec) *QTextCodec {
	var __rv uintptr
	DirectQtDrv(nil,143000,143108,unsafe.Pointer(&ba),Native(defaultCodec),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextCodec{}
	_rp.SetDriver(__rv,143000,true)
	return _rp
}	
//QTextCodec::codecForHtml(QByteArray const&,QTextCodec*)
func (q *QTextCodec) CodecForHtmlWithBaDefaultcodec(ba []byte,defaultCodec *QTextCodec) *QTextCodec {
	var __rv uintptr
	q.Drv(143000,143108,unsafe.Pointer(&ba),Native(defaultCodec),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextCodec{}
	_rp.SetDriver(__rv,143000,true)
	return _rp
}	
//QTextCodec::codecForLocale()	
func QTextCodecCodecForLocale() *QTextCodec {
	var __rv uintptr
	DirectQtDrv(nil,143000,143109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextCodec{}
	_rp.SetDriver(__rv,143000,true)
	return _rp
}	
//QTextCodec::codecForLocale()
func (q *QTextCodec) CodecForLocale() *QTextCodec {
	var __rv uintptr
	q.Drv(143000,143109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextCodec{}
	_rp.SetDriver(__rv,143000,true)
	return _rp
}	
//QTextCodec::codecForMib(int)	
func QTextCodecCodecForMib(mib int) *QTextCodec {
	var __rv uintptr
	DirectQtDrv(nil,143000,143110,unsafe.Pointer(&mib),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextCodec{}
	_rp.SetDriver(__rv,143000,true)
	return _rp
}	
//QTextCodec::codecForMib(int)
func (q *QTextCodec) CodecForMib(mib int) *QTextCodec {
	var __rv uintptr
	q.Drv(143000,143110,unsafe.Pointer(&mib),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextCodec{}
	_rp.SetDriver(__rv,143000,true)
	return _rp
}	
//QTextCodec::codecForName(QByteArray const&)	
func QTextCodecCodecForName(name []byte) *QTextCodec {
	var __rv uintptr
	DirectQtDrv(nil,143000,143111,unsafe.Pointer(&name),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextCodec{}
	_rp.SetDriver(__rv,143000,true)
	return _rp
}	
//QTextCodec::codecForName(QByteArray const&)
func (q *QTextCodec) CodecForName(name []byte) *QTextCodec {
	var __rv uintptr
	q.Drv(143000,143111,unsafe.Pointer(&name),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextCodec{}
	_rp.SetDriver(__rv,143000,true)
	return _rp
}	
//QTextCodec::codecForName(char const*)	
func QTextCodecCodecForNameWithName(name string) *QTextCodec {
	var __rv uintptr
	DirectQtDrv(nil,143000,143112,unsafe.Pointer(&name),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextCodec{}
	_rp.SetDriver(__rv,143000,true)
	return _rp
}	
//QTextCodec::codecForName(char const*)
func (q *QTextCodec) CodecForNameWithName(name string) *QTextCodec {
	var __rv uintptr
	q.Drv(143000,143112,unsafe.Pointer(&name),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextCodec{}
	_rp.SetDriver(__rv,143000,true)
	return _rp
}	
//QTextCodec::codecForUtfText(QByteArray const&)	
func QTextCodecCodecForUtfText(ba []byte) *QTextCodec {
	var __rv uintptr
	DirectQtDrv(nil,143000,143113,unsafe.Pointer(&ba),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextCodec{}
	_rp.SetDriver(__rv,143000,true)
	return _rp
}	
//QTextCodec::codecForUtfText(QByteArray const&)
func (q *QTextCodec) CodecForUtfText(ba []byte) *QTextCodec {
	var __rv uintptr
	q.Drv(143000,143113,unsafe.Pointer(&ba),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextCodec{}
	_rp.SetDriver(__rv,143000,true)
	return _rp
}	
//QTextCodec::codecForUtfText(QByteArray const&,QTextCodec*)	
func QTextCodecCodecForUtfTextWithBaDefaultcodec(ba []byte,defaultCodec *QTextCodec) *QTextCodec {
	var __rv uintptr
	DirectQtDrv(nil,143000,143114,unsafe.Pointer(&ba),Native(defaultCodec),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextCodec{}
	_rp.SetDriver(__rv,143000,true)
	return _rp
}	
//QTextCodec::codecForUtfText(QByteArray const&,QTextCodec*)
func (q *QTextCodec) CodecForUtfTextWithBaDefaultcodec(ba []byte,defaultCodec *QTextCodec) *QTextCodec {
	var __rv uintptr
	q.Drv(143000,143114,unsafe.Pointer(&ba),Native(defaultCodec),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextCodec{}
	_rp.SetDriver(__rv,143000,true)
	return _rp
}	
//QTextCodec::fromUnicode(QString const&)
func (q *QTextCodec) FromUnicode(uc string) []byte {
	var __rv []byte
	q.Drv(143000,143115,unsafe.Pointer(&uc),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextCodec::fromUnicode(QChar const*,int,QTextCodec::ConverterState*)
func (q *QTextCodec) FromUnicodeWithInLengthState(in *rune,length int,state *QTextCodecConverterState) []byte {
	var __rv []byte
	q.Drv(143000,143116,unsafe.Pointer(&in),unsafe.Pointer(&length),Native(state),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextCodec::makeDecoder()
func (q *QTextCodec) MakeDecoder() *QTextDecoder {
	var __rv uintptr
	q.Drv(143000,143117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextDecoder{}
	_rp.SetDriver(__rv,146000,true)
	return _rp
}	
//QTextCodec::makeDecoder(QFlags<QTextCodec::ConversionFlag>)
func (q *QTextCodec) MakeDecoderWithFlags(flags QTextCodec_ConversionFlag) *QTextDecoder {
	var __rv uintptr
	q.Drv(143000,143118,unsafe.Pointer(&flags),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextDecoder{}
	_rp.SetDriver(__rv,146000,true)
	return _rp
}	
//QTextCodec::makeEncoder()
func (q *QTextCodec) MakeEncoder() *QTextEncoder {
	var __rv uintptr
	q.Drv(143000,143119,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextEncoder{}
	_rp.SetDriver(__rv,150000,true)
	return _rp
}	
//QTextCodec::makeEncoder(QFlags<QTextCodec::ConversionFlag>)
func (q *QTextCodec) MakeEncoderWithFlags(flags QTextCodec_ConversionFlag) *QTextEncoder {
	var __rv uintptr
	q.Drv(143000,143120,unsafe.Pointer(&flags),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextEncoder{}
	_rp.SetDriver(__rv,150000,true)
	return _rp
}	
//QTextCodec::mibEnum()
func (q *QTextCodec) MibEnum() int {
	var __rv int
	q.Drv(143000,143121,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextCodec::name()
func (q *QTextCodec) Name() []byte {
	var __rv []byte
	q.Drv(143000,143122,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextCodec::setCodecForLocale(QTextCodec*)	
func QTextCodecSetCodecForLocale(c *QTextCodec)  {
	DirectQtDrv(nil,143000,143123,Native(c),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextCodec::setCodecForLocale(QTextCodec*)
func (q *QTextCodec) SetCodecForLocale(c *QTextCodec)  {
	q.Drv(143000,143123,Native(c),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextCodec::toUnicode(QByteArray const&)
func (q *QTextCodec) ToUnicode(value []byte) string {
	var __rv string
	q.Drv(143000,143124,unsafe.Pointer(&value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextCodec::toUnicode(char const*)
func (q *QTextCodec) ToUnicodeWithChars(chars string) string {
	var __rv string
	q.Drv(143000,143125,unsafe.Pointer(&chars),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
