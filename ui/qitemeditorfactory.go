// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QItemEditorFactory : QItemEditorFactory
type QItemEditorFactory struct {
	BaseDrv
}
//QItemEditorFactory::QItemEditorFactory()	
func NewItemEditorFactory() *QItemEditorFactory {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),61000,61102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QItemEditorFactory{}
	_p.SetDriver(__rv,61000,true)
	return _p
} 
//QItemEditorFactory::createEditor(QVariant::Type,QWidget*)
func (q *QItemEditorFactory) CreateEditor(_type QVariant_Type,parent QWidgetInterface) *QWidget {
	var __rv uintptr
	q.Drv(61000,61103,unsafe.Pointer(&_type),Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QItemEditorFactory::defaultFactory()	
func QItemEditorFactoryDefaultFactory() *QItemEditorFactory {
	var __rv uintptr
	DirectQtDrv(nil,61000,61104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QItemEditorFactory{}
	_rp.SetDriver(__rv,61000,true)
	return _rp
}	
//QItemEditorFactory::defaultFactory()
func (q *QItemEditorFactory) DefaultFactory() *QItemEditorFactory {
	var __rv uintptr
	q.Drv(61000,61104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QItemEditorFactory{}
	_rp.SetDriver(__rv,61000,true)
	return _rp
}	
//QItemEditorFactory::registerEditor(QVariant::Type,QItemEditorCreatorBase*)
func (q *QItemEditorFactory) RegisterEditor(_type QVariant_Type,creator *QItemEditorCreatorBase)  {
	q.Drv(61000,61105,unsafe.Pointer(&_type),Native(creator),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QItemEditorFactory::setDefaultFactory(QItemEditorFactory*)	
func QItemEditorFactorySetDefaultFactory(factory *QItemEditorFactory)  {
	DirectQtDrv(nil,61000,61106,Native(factory),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QItemEditorFactory::setDefaultFactory(QItemEditorFactory*)
func (q *QItemEditorFactory) SetDefaultFactory(factory *QItemEditorFactory)  {
	q.Drv(61000,61106,Native(factory),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QItemEditorFactory::valuePropertyName(QVariant::Type)
func (q *QItemEditorFactory) ValuePropertyName(_type QVariant_Type) []byte {
	var __rv []byte
	q.Drv(61000,61107,unsafe.Pointer(&_type),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
