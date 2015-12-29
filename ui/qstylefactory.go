// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QStyleFactory : QStyleFactory
type QStyleFactory struct {
	BaseDrv
}
//QStyleFactory::QStyleFactory()	
func NewStyleFactory() *QStyleFactory {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),129000,129102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QStyleFactory{}
	_p.SetDriver(__rv,129000,true)
	return _p
} 
//QStyleFactory::create(QString const&)	
func QStyleFactoryCreate(value string) *QStyle {
	var __rv uintptr
	DirectQtDrv(nil,129000,129103,unsafe.Pointer(&value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QStyle{}
	_rp.SetDriver(__rv,357000,false)
	return _rp
}	
//QStyleFactory::create(QString const&)
func (q *QStyleFactory) Create(value string) *QStyle {
	var __rv uintptr
	q.Drv(129000,129103,unsafe.Pointer(&value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QStyle{}
	_rp.SetDriver(__rv,357000,false)
	return _rp
}	
//QStyleFactory::keys()	
func QStyleFactoryKeys() []string {
	var __rv []string
	DirectQtDrv(nil,129000,129104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStyleFactory::keys()
func (q *QStyleFactory) Keys() []string {
	var __rv []string
	q.Drv(129000,129104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
