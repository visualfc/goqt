// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QSignalMapper : QSignalMapper
type QSignalMapper struct {
	QObject
}
func (q *QSignalMapper) OnMapped(fn func(*QObject)) uintptr {
	var __rv uintptr
	q.Drv(342000,342102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QSignalMapper) OnMappedWithWidget(fn func(*QWidget)) uintptr {
	var __rv uintptr
	q.Drv(342000,342103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QSignalMapper) OnMappedWithString(fn func(string)) uintptr {
	var __rv uintptr
	q.Drv(342000,342104,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QSignalMapper) OnMappedWithInt(fn func(int)) uintptr {
	var __rv uintptr
	q.Drv(342000,342105,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QSignalMapper::QSignalMapper()	
func NewSignalMapper() *QSignalMapper {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),342000,342106,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QSignalMapper{}
	_p.SetDriver(__rv,342000,false)
	return _p
} 
//QSignalMapper::QSignalMapper(QObject*)	
func NewSignalMapperWithParent(parent QObjectInterface) *QSignalMapper {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),342000,342107,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QSignalMapper{}
	_p.SetDriver(__rv,342000,false)
	return _p
} 
//QSignalMapper::map()
func (q *QSignalMapper) Map()  {
	q.Drv(342000,342108,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSignalMapper::map(QObject*)
func (q *QSignalMapper) MapWithObject(sender QObjectInterface)  {
	q.Drv(342000,342109,Native(sender),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSignalMapper::mapping(QObject*)
func (q *QSignalMapper) Mapping(object QObjectInterface) *QObject {
	var __rv uintptr
	q.Drv(342000,342110,Native(object),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QObject{}
	_rp.SetDriver(__rv,314000,false)
	return _rp
}	
//QSignalMapper::mapping(QString const&)
func (q *QSignalMapper) MappingWithText(text string) *QObject {
	var __rv uintptr
	q.Drv(342000,342111,unsafe.Pointer(&text),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QObject{}
	_rp.SetDriver(__rv,314000,false)
	return _rp
}	
//QSignalMapper::mapping(QWidget*)
func (q *QSignalMapper) MappingWithWidget(widget QWidgetInterface) *QObject {
	var __rv uintptr
	q.Drv(342000,342112,Native(widget),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QObject{}
	_rp.SetDriver(__rv,314000,false)
	return _rp
}	
//QSignalMapper::mapping(int)
func (q *QSignalMapper) MappingWithId(id int) *QObject {
	var __rv uintptr
	q.Drv(342000,342113,unsafe.Pointer(&id),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QObject{}
	_rp.SetDriver(__rv,314000,false)
	return _rp
}	
//QSignalMapper::removeMappings(QObject*)
func (q *QSignalMapper) RemoveMappings(sender QObjectInterface)  {
	q.Drv(342000,342114,Native(sender),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSignalMapper::setMapping(QObject*,QObject*)
func (q *QSignalMapper) SetMappingWithObjectObject(sender QObjectInterface,object QObjectInterface)  {
	q.Drv(342000,342115,Native(sender),Native(object),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSignalMapper::setMapping(QObject*,QString const&)
func (q *QSignalMapper) SetMappingWithObjectText(sender QObjectInterface,text string)  {
	q.Drv(342000,342116,Native(sender),unsafe.Pointer(&text),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSignalMapper::setMapping(QObject*,QWidget*)
func (q *QSignalMapper) SetMappingWithObjectWidget(sender QObjectInterface,widget QWidgetInterface)  {
	q.Drv(342000,342117,Native(sender),Native(widget),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSignalMapper::setMapping(QObject*,int)
func (q *QSignalMapper) SetMappingWithObjectId(sender QObjectInterface,id int)  {
	q.Drv(342000,342118,Native(sender),unsafe.Pointer(&id),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
