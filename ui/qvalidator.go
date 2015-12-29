// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QValidator_State - QValidator::State
type QValidator_State uint32
const (
	QValidator_Invalid QValidator_State = 0
	QValidator_Intermediate QValidator_State = 1
	QValidator_Acceptable QValidator_State = 2
)
//struct QValidator : QValidator
type QValidator struct {
	QObject
}
//QValidator::fixup(QString&)
func (q *QValidator) Fixup(value *string)  {
	q.Drv(393000,393102,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QValidator::locale()
func (q *QValidator) Locale() *QLocale {
	var __rv uintptr
	q.Drv(393000,393103,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QLocale{}
	_rp.SetDriver(__rv,72000,true)
	return _rp
}	
//QValidator::setLocale(QLocale const&)
func (q *QValidator) SetLocale(locale *QLocale)  {
	q.Drv(393000,393104,Native(locale),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QValidator::validate(QString&,int&)
func (q *QValidator) Validate(value2 *string,value3 *int) QValidator_State {
	var __rv QValidator_State
	q.Drv(393000,393105,unsafe.Pointer(&value2),unsafe.Pointer(&value3),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
