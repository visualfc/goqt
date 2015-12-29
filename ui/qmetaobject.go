// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QMetaObject_Call - QMetaObject::Call
type QMetaObject_Call uint32
const (
	QMetaObject_InvokeMetaMethod QMetaObject_Call = 0
	QMetaObject_ReadProperty QMetaObject_Call = 1
	QMetaObject_WriteProperty QMetaObject_Call = 2
	QMetaObject_ResetProperty QMetaObject_Call = 3
	QMetaObject_QueryPropertyDesignable QMetaObject_Call = 4
	QMetaObject_QueryPropertyScriptable QMetaObject_Call = 5
	QMetaObject_QueryPropertyStored QMetaObject_Call = 6
	QMetaObject_QueryPropertyEditable QMetaObject_Call = 7
	QMetaObject_QueryPropertyUser QMetaObject_Call = 8
	QMetaObject_CreateInstance QMetaObject_Call = 9
)
//struct QMetaObject : QMetaObject
type QMetaObject struct {
	BaseDrv
}
//QMetaObject::QMetaObject()	
func NewMetaObject() *QMetaObject {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),77000,77102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QMetaObject{}
	_p.SetDriver(__rv,77000,true)
	return _p
} 
//QMetaObject::cast(QObject*)
func (q *QMetaObject) Cast(obj QObjectInterface) *QObject {
	var __rv uintptr
	q.Drv(77000,77103,Native(obj),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QObject{}
	_rp.SetDriver(__rv,314000,false)
	return _rp
}	
//QMetaObject::checkConnectArgs(char const*,char const*)	
func QMetaObjectCheckConnectArgs(signal string,method string) bool {
	var __rv bool
	DirectQtDrv(nil,77000,77104,unsafe.Pointer(&signal),unsafe.Pointer(&method),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaObject::checkConnectArgs(char const*,char const*)
func (q *QMetaObject) CheckConnectArgs(signal string,method string) bool {
	var __rv bool
	q.Drv(77000,77104,unsafe.Pointer(&signal),unsafe.Pointer(&method),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaObject::classInfoCount()
func (q *QMetaObject) ClassInfoCount() int {
	var __rv int
	q.Drv(77000,77105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaObject::classInfoOffset()
func (q *QMetaObject) ClassInfoOffset() int {
	var __rv int
	q.Drv(77000,77106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaObject::className()
func (q *QMetaObject) ClassName() string {
	var __rv string
	q.Drv(77000,77107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaObject::connect(QObject const*,int,QObject const*,int,int,int*)	
func QMetaObjectConnect(sender QObjectInterface,signal_index int,receiver QObjectInterface,method_index int,_type int,types *int) bool {
	var __rv bool
	DirectQtDrv(nil,77000,77108,Native(sender),unsafe.Pointer(&signal_index),Native(receiver),unsafe.Pointer(&method_index),unsafe.Pointer(&_type),unsafe.Pointer(&types),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaObject::connect(QObject const*,int,QObject const*,int,int,int*)
func (q *QMetaObject) Connect(sender QObjectInterface,signal_index int,receiver QObjectInterface,method_index int,_type int,types *int) bool {
	var __rv bool
	q.Drv(77000,77108,Native(sender),unsafe.Pointer(&signal_index),Native(receiver),unsafe.Pointer(&method_index),unsafe.Pointer(&_type),unsafe.Pointer(&types),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaObject::connectSlotsByName(QObject*)	
func QMetaObjectConnectSlotsByName(o QObjectInterface)  {
	DirectQtDrv(nil,77000,77109,Native(o),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMetaObject::connectSlotsByName(QObject*)
func (q *QMetaObject) ConnectSlotsByName(o QObjectInterface)  {
	q.Drv(77000,77109,Native(o),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMetaObject::constructor(int)
func (q *QMetaObject) Constructor(index int) *QMetaMethod {
	var __rv uintptr
	q.Drv(77000,77110,unsafe.Pointer(&index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QMetaMethod{}
	_rp.SetDriver(__rv,76000,true)
	return _rp
}	
//QMetaObject::constructorCount()
func (q *QMetaObject) ConstructorCount() int {
	var __rv int
	q.Drv(77000,77111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaObject::disconnect(QObject const*,int,QObject const*,int)	
func QMetaObjectDisconnect(sender QObjectInterface,signal_index int,receiver QObjectInterface,method_index int) bool {
	var __rv bool
	DirectQtDrv(nil,77000,77112,Native(sender),unsafe.Pointer(&signal_index),Native(receiver),unsafe.Pointer(&method_index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaObject::disconnect(QObject const*,int,QObject const*,int)
func (q *QMetaObject) Disconnect(sender QObjectInterface,signal_index int,receiver QObjectInterface,method_index int) bool {
	var __rv bool
	q.Drv(77000,77112,Native(sender),unsafe.Pointer(&signal_index),Native(receiver),unsafe.Pointer(&method_index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaObject::disconnectOne(QObject const*,int,QObject const*,int)	
func QMetaObjectDisconnectOne(sender QObjectInterface,signal_index int,receiver QObjectInterface,method_index int) bool {
	var __rv bool
	DirectQtDrv(nil,77000,77113,Native(sender),unsafe.Pointer(&signal_index),Native(receiver),unsafe.Pointer(&method_index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaObject::disconnectOne(QObject const*,int,QObject const*,int)
func (q *QMetaObject) DisconnectOne(sender QObjectInterface,signal_index int,receiver QObjectInterface,method_index int) bool {
	var __rv bool
	q.Drv(77000,77113,Native(sender),unsafe.Pointer(&signal_index),Native(receiver),unsafe.Pointer(&method_index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaObject::enumerator(int)
func (q *QMetaObject) Enumerator(index int) *QMetaEnum {
	var __rv uintptr
	q.Drv(77000,77114,unsafe.Pointer(&index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QMetaEnum{}
	_rp.SetDriver(__rv,75000,true)
	return _rp
}	
//QMetaObject::enumeratorCount()
func (q *QMetaObject) EnumeratorCount() int {
	var __rv int
	q.Drv(77000,77115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaObject::enumeratorOffset()
func (q *QMetaObject) EnumeratorOffset() int {
	var __rv int
	q.Drv(77000,77116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaObject::indexOfClassInfo(char const*)
func (q *QMetaObject) IndexOfClassInfo(name string) int {
	var __rv int
	q.Drv(77000,77117,unsafe.Pointer(&name),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaObject::indexOfConstructor(char const*)
func (q *QMetaObject) IndexOfConstructor(constructor string) int {
	var __rv int
	q.Drv(77000,77118,unsafe.Pointer(&constructor),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaObject::indexOfEnumerator(char const*)
func (q *QMetaObject) IndexOfEnumerator(name string) int {
	var __rv int
	q.Drv(77000,77119,unsafe.Pointer(&name),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaObject::indexOfMethod(char const*)
func (q *QMetaObject) IndexOfMethod(method string) int {
	var __rv int
	q.Drv(77000,77120,unsafe.Pointer(&method),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaObject::indexOfProperty(char const*)
func (q *QMetaObject) IndexOfProperty(name string) int {
	var __rv int
	q.Drv(77000,77121,unsafe.Pointer(&name),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaObject::indexOfSignal(char const*)
func (q *QMetaObject) IndexOfSignal(signal string) int {
	var __rv int
	q.Drv(77000,77122,unsafe.Pointer(&signal),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaObject::indexOfSlot(char const*)
func (q *QMetaObject) IndexOfSlot(slot string) int {
	var __rv int
	q.Drv(77000,77123,unsafe.Pointer(&slot),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaObject::method(int)
func (q *QMetaObject) Method(index int) *QMetaMethod {
	var __rv uintptr
	q.Drv(77000,77124,unsafe.Pointer(&index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QMetaMethod{}
	_rp.SetDriver(__rv,76000,true)
	return _rp
}	
//QMetaObject::methodCount()
func (q *QMetaObject) MethodCount() int {
	var __rv int
	q.Drv(77000,77125,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaObject::methodOffset()
func (q *QMetaObject) MethodOffset() int {
	var __rv int
	q.Drv(77000,77126,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaObject::newInstance()
func (q *QMetaObject) NewInstance() *QObject {
	var __rv uintptr
	q.Drv(77000,77127,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QObject{}
	_rp.SetDriver(__rv,314000,false)
	return _rp
}	
//QMetaObject::newInstance(QGenericArgument,QGenericArgument,QGenericArgument,QGenericArgument,QGenericArgument,QGenericArgument,QGenericArgument,QGenericArgument,QGenericArgument,QGenericArgument)
func (q *QMetaObject) NewInstanceWithVal0Val1Val2Val3Val4Val5Val6Val7Val8Val9(val0 *QGenericArgument,val1 *QGenericArgument,val2 *QGenericArgument,val3 *QGenericArgument,val4 *QGenericArgument,val5 *QGenericArgument,val6 *QGenericArgument,val7 *QGenericArgument,val8 *QGenericArgument,val9 *QGenericArgument) *QObject {
	var __rv uintptr
	q.Drv(77000,77128,Native(val0),Native(val1),Native(val2),Native(val3),Native(val4),Native(val5),Native(val6),Native(val7),Native(val8),Native(val9),unsafe.Pointer(&__rv),nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QObject{}
	_rp.SetDriver(__rv,314000,false)
	return _rp
}	
//QMetaObject::normalizedSignature(char const*)	
func QMetaObjectNormalizedSignature(method string) []byte {
	var __rv []byte
	DirectQtDrv(nil,77000,77129,unsafe.Pointer(&method),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaObject::normalizedSignature(char const*)
func (q *QMetaObject) NormalizedSignature(method string) []byte {
	var __rv []byte
	q.Drv(77000,77129,unsafe.Pointer(&method),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaObject::normalizedType(char const*)	
func QMetaObjectNormalizedType(_type string) []byte {
	var __rv []byte
	DirectQtDrv(nil,77000,77130,unsafe.Pointer(&_type),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaObject::normalizedType(char const*)
func (q *QMetaObject) NormalizedType(_type string) []byte {
	var __rv []byte
	q.Drv(77000,77130,unsafe.Pointer(&_type),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaObject::property(int)
func (q *QMetaObject) Property(index int) *QMetaProperty {
	var __rv uintptr
	q.Drv(77000,77131,unsafe.Pointer(&index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QMetaProperty{}
	_rp.SetDriver(__rv,78000,true)
	return _rp
}	
//QMetaObject::propertyCount()
func (q *QMetaObject) PropertyCount() int {
	var __rv int
	q.Drv(77000,77132,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaObject::propertyOffset()
func (q *QMetaObject) PropertyOffset() int {
	var __rv int
	q.Drv(77000,77133,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaObject::superClass()
func (q *QMetaObject) SuperClass() *QMetaObject {
	var __rv uintptr
	q.Drv(77000,77134,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QMetaObject{}
	_rp.SetDriver(__rv,77000,true)
	return _rp
}	
//QMetaObject::tr(char const*,char const*)
func (q *QMetaObject) TrWithSC(s string,c string) string {
	var __rv string
	q.Drv(77000,77135,unsafe.Pointer(&s),unsafe.Pointer(&c),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaObject::tr(char const*,char const*,int)
func (q *QMetaObject) TrWithSCN(s string,c string,n int) string {
	var __rv string
	q.Drv(77000,77136,unsafe.Pointer(&s),unsafe.Pointer(&c),unsafe.Pointer(&n),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaObject::userProperty()
func (q *QMetaObject) UserProperty() *QMetaProperty {
	var __rv uintptr
	q.Drv(77000,77137,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QMetaProperty{}
	_rp.SetDriver(__rv,78000,true)
	return _rp
}	
