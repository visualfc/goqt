// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QMetaProperty : QMetaProperty
type QMetaProperty struct {
	BaseDrv
}
//QMetaProperty::QMetaProperty()	
func NewMetaProperty() *QMetaProperty {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),78000,78102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QMetaProperty{}
	_p.SetDriver(__rv,78000,true)
	return _p
} 
//QMetaProperty::enclosingMetaObject()
func (q *QMetaProperty) EnclosingMetaObject() *QMetaObject {
	var __rv uintptr
	q.Drv(78000,78103,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QMetaObject{}
	_rp.SetDriver(__rv,77000,true)
	return _rp
}	
//QMetaProperty::enumerator()
func (q *QMetaProperty) Enumerator() *QMetaEnum {
	var __rv uintptr
	q.Drv(78000,78104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QMetaEnum{}
	_rp.SetDriver(__rv,75000,true)
	return _rp
}	
//QMetaProperty::hasNotifySignal()
func (q *QMetaProperty) HasNotifySignal() bool {
	var __rv bool
	q.Drv(78000,78105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaProperty::hasStdCppSet()
func (q *QMetaProperty) HasStdCppSet() bool {
	var __rv bool
	q.Drv(78000,78106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaProperty::isConstant()
func (q *QMetaProperty) IsConstant() bool {
	var __rv bool
	q.Drv(78000,78107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaProperty::isDesignable()
func (q *QMetaProperty) IsDesignable() bool {
	var __rv bool
	q.Drv(78000,78108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaProperty::isDesignable(QObject const*)
func (q *QMetaProperty) IsDesignableWithObject(obj QObjectInterface) bool {
	var __rv bool
	q.Drv(78000,78109,Native(obj),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaProperty::isEditable()
func (q *QMetaProperty) IsEditable() bool {
	var __rv bool
	q.Drv(78000,78110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaProperty::isEditable(QObject const*)
func (q *QMetaProperty) IsEditableWithObject(obj QObjectInterface) bool {
	var __rv bool
	q.Drv(78000,78111,Native(obj),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaProperty::isEnumType()
func (q *QMetaProperty) IsEnumType() bool {
	var __rv bool
	q.Drv(78000,78112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaProperty::isFinal()
func (q *QMetaProperty) IsFinal() bool {
	var __rv bool
	q.Drv(78000,78113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaProperty::isFlagType()
func (q *QMetaProperty) IsFlagType() bool {
	var __rv bool
	q.Drv(78000,78114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaProperty::isReadable()
func (q *QMetaProperty) IsReadable() bool {
	var __rv bool
	q.Drv(78000,78115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaProperty::isResettable()
func (q *QMetaProperty) IsResettable() bool {
	var __rv bool
	q.Drv(78000,78116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaProperty::isScriptable()
func (q *QMetaProperty) IsScriptable() bool {
	var __rv bool
	q.Drv(78000,78117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaProperty::isScriptable(QObject const*)
func (q *QMetaProperty) IsScriptableWithObject(obj QObjectInterface) bool {
	var __rv bool
	q.Drv(78000,78118,Native(obj),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaProperty::isStored()
func (q *QMetaProperty) IsStored() bool {
	var __rv bool
	q.Drv(78000,78119,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaProperty::isStored(QObject const*)
func (q *QMetaProperty) IsStoredWithObject(obj QObjectInterface) bool {
	var __rv bool
	q.Drv(78000,78120,Native(obj),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaProperty::isUser()
func (q *QMetaProperty) IsUser() bool {
	var __rv bool
	q.Drv(78000,78121,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaProperty::isUser(QObject const*)
func (q *QMetaProperty) IsUserWithObject(obj QObjectInterface) bool {
	var __rv bool
	q.Drv(78000,78122,Native(obj),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaProperty::isValid()
func (q *QMetaProperty) IsValid() bool {
	var __rv bool
	q.Drv(78000,78123,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaProperty::isWritable()
func (q *QMetaProperty) IsWritable() bool {
	var __rv bool
	q.Drv(78000,78124,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaProperty::name()
func (q *QMetaProperty) Name() string {
	var __rv string
	q.Drv(78000,78125,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaProperty::notifySignal()
func (q *QMetaProperty) NotifySignal() *QMetaMethod {
	var __rv uintptr
	q.Drv(78000,78126,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QMetaMethod{}
	_rp.SetDriver(__rv,76000,true)
	return _rp
}	
//QMetaProperty::notifySignalIndex()
func (q *QMetaProperty) NotifySignalIndex() int {
	var __rv int
	q.Drv(78000,78127,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaProperty::propertyIndex()
func (q *QMetaProperty) PropertyIndex() int {
	var __rv int
	q.Drv(78000,78128,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaProperty::read(QObject const*)
func (q *QMetaProperty) Read(obj QObjectInterface) *QVariant {
	var __rv uintptr
	q.Drv(78000,78129,Native(obj),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVariant{}
	_rp.SetDriver(__rv,182000,true)
	return _rp
}	
//QMetaProperty::reset(QObject*)
func (q *QMetaProperty) Reset(obj QObjectInterface) bool {
	var __rv bool
	q.Drv(78000,78130,Native(obj),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaProperty::type()
func (q *QMetaProperty) Type() QVariant_Type {
	var __rv QVariant_Type
	q.Drv(78000,78131,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaProperty::typeName()
func (q *QMetaProperty) TypeName() string {
	var __rv string
	q.Drv(78000,78132,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaProperty::userType()
func (q *QMetaProperty) UserType() int {
	var __rv int
	q.Drv(78000,78133,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaProperty::write(QObject*,QVariant const&)
func (q *QMetaProperty) Write(obj QObjectInterface,value *QVariant) bool {
	var __rv bool
	q.Drv(78000,78134,Native(obj),Native(value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
