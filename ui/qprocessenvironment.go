// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QProcessEnvironment : QProcessEnvironment
type QProcessEnvironment struct {
	BaseDrv
}
//QProcessEnvironment::QProcessEnvironment()	
func NewProcessEnvironment() *QProcessEnvironment {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),106000,106102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QProcessEnvironment{}
	_p.SetDriver(__rv,106000,true)
	return _p
} 
//QProcessEnvironment::QProcessEnvironment(QProcessEnvironment const&)	
func NewProcessEnvironmentCopy(other *QProcessEnvironment) *QProcessEnvironment {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),106000,106103,Native(other),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QProcessEnvironment{}
	_p.SetDriver(__rv,106000,true)
	return _p
} 
//QProcessEnvironment::clear()
func (q *QProcessEnvironment) Clear()  {
	q.Drv(106000,106104,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QProcessEnvironment::contains(QString const&)
func (q *QProcessEnvironment) Contains(name string) bool {
	var __rv bool
	q.Drv(106000,106105,unsafe.Pointer(&name),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QProcessEnvironment::insert(QString const&,QString const&)
func (q *QProcessEnvironment) Insert(name string,value string)  {
	q.Drv(106000,106106,unsafe.Pointer(&name),unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QProcessEnvironment::isEmpty()
func (q *QProcessEnvironment) IsEmpty() bool {
	var __rv bool
	q.Drv(106000,106107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QProcessEnvironment::remove(QString const&)
func (q *QProcessEnvironment) Remove(name string)  {
	q.Drv(106000,106108,unsafe.Pointer(&name),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QProcessEnvironment::systemEnvironment()	
func QProcessEnvironmentSystemEnvironment() *QProcessEnvironment {
	var __rv uintptr
	DirectQtDrv(nil,106000,106109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QProcessEnvironment{}
	_rp.SetDriver(__rv,106000,true)
	return _rp
}	
//QProcessEnvironment::systemEnvironment()
func (q *QProcessEnvironment) SystemEnvironment() *QProcessEnvironment {
	var __rv uintptr
	q.Drv(106000,106109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QProcessEnvironment{}
	_rp.SetDriver(__rv,106000,true)
	return _rp
}	
//QProcessEnvironment::toStringList()
func (q *QProcessEnvironment) ToStringList() []string {
	var __rv []string
	q.Drv(106000,106110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QProcessEnvironment::value(QString const&)
func (q *QProcessEnvironment) Value(name string) string {
	var __rv string
	q.Drv(106000,106111,unsafe.Pointer(&name),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QProcessEnvironment::value(QString const&,QString const&)
func (q *QProcessEnvironment) ValueWithNameDefaultvalue(name string,defaultValue string) string {
	var __rv string
	q.Drv(106000,106112,unsafe.Pointer(&name),unsafe.Pointer(&defaultValue),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
