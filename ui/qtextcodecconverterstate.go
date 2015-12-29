// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QTextCodecConverterState : QTextCodec::ConverterState
type QTextCodecConverterState struct {
	BaseDrv
}
//QTextCodec::ConverterState::ConverterState()	
func NewTextCodecConverterState() *QTextCodecConverterState {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),144000,144102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTextCodecConverterState{}
	_p.SetDriver(__rv,144000,true)
	return _p
} 
//QTextCodec::ConverterState::ConverterState(QFlags<QTextCodec::ConversionFlag>)	
func NewTextCodecConverterStateWithConversionflag(f QTextCodec_ConversionFlag) *QTextCodecConverterState {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),144000,144103,unsafe.Pointer(&f),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTextCodecConverterState{}
	_p.SetDriver(__rv,144000,true)
	return _p
} 
