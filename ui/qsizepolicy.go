// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QSizePolicy_ControlType - QSizePolicy::ControlType
type QSizePolicy_ControlType uint32
const (
	QSizePolicy_DefaultType QSizePolicy_ControlType = 0x00000001
	QSizePolicy_ButtonBox QSizePolicy_ControlType = 0x00000002
	QSizePolicy_CheckBox QSizePolicy_ControlType = 0x00000004
	QSizePolicy_ComboBox QSizePolicy_ControlType = 0x00000008
	QSizePolicy_Frame QSizePolicy_ControlType = 0x00000010
	QSizePolicy_GroupBox QSizePolicy_ControlType = 0x00000020
	QSizePolicy_Label QSizePolicy_ControlType = 0x00000040
	QSizePolicy_Line QSizePolicy_ControlType = 0x00000080
	QSizePolicy_LineEdit QSizePolicy_ControlType = 0x00000100
	QSizePolicy_PushButton QSizePolicy_ControlType = 0x00000200
	QSizePolicy_RadioButton QSizePolicy_ControlType = 0x00000400
	QSizePolicy_Slider QSizePolicy_ControlType = 0x00000800
	QSizePolicy_SpinBox QSizePolicy_ControlType = 0x00001000
	QSizePolicy_TabWidget QSizePolicy_ControlType = 0x00002000
	QSizePolicy_ToolButton QSizePolicy_ControlType = 0x00004000
)
//enum QSizePolicy_PolicyFlag - QSizePolicy::PolicyFlag
type QSizePolicy_PolicyFlag uint32
const (
	QSizePolicy_GrowFlag QSizePolicy_PolicyFlag = 1
	QSizePolicy_ExpandFlag QSizePolicy_PolicyFlag = 2
	QSizePolicy_ShrinkFlag QSizePolicy_PolicyFlag = 4
	QSizePolicy_IgnoreFlag QSizePolicy_PolicyFlag = 8
)
//enum QSizePolicy_SizePolicyMasks - QSizePolicy::SizePolicyMasks
type QSizePolicy_SizePolicyMasks uint32
const (
	QSizePolicy_HSize QSizePolicy_SizePolicyMasks = 4
	QSizePolicy_HMask QSizePolicy_SizePolicyMasks = 0x0f
	QSizePolicy_VMask QSizePolicy_SizePolicyMasks = QSizePolicy_HMask<<QSizePolicy_HSize
	QSizePolicy_CTShift QSizePolicy_SizePolicyMasks = 9
	QSizePolicy_CTSize QSizePolicy_SizePolicyMasks = 5
	QSizePolicy_WFHShift QSizePolicy_SizePolicyMasks = QSizePolicy_CTShift+QSizePolicy_CTSize
	QSizePolicy_CTMask QSizePolicy_SizePolicyMasks = ((0x1<<QSizePolicy_CTSize)-1)<<QSizePolicy_CTShift
	QSizePolicy_UnusedShift QSizePolicy_SizePolicyMasks = QSizePolicy_CTShift+QSizePolicy_CTSize
	QSizePolicy_UnusedSize QSizePolicy_SizePolicyMasks = 2
)
//enum QSizePolicy_Policy - QSizePolicy::Policy
type QSizePolicy_Policy uint32
const (
	QSizePolicy_Fixed QSizePolicy_Policy = 0
	QSizePolicy_Minimum QSizePolicy_Policy = QSizePolicy_Policy(QSizePolicy_GrowFlag)
	QSizePolicy_Maximum QSizePolicy_Policy = QSizePolicy_Policy(QSizePolicy_ShrinkFlag)
	QSizePolicy_Preferred QSizePolicy_Policy = QSizePolicy_Policy(QSizePolicy_GrowFlag)|QSizePolicy_Policy(QSizePolicy_ShrinkFlag)
	QSizePolicy_MinimumExpanding QSizePolicy_Policy = QSizePolicy_Policy(QSizePolicy_GrowFlag)|QSizePolicy_Policy(QSizePolicy_ExpandFlag)
	QSizePolicy_Expanding QSizePolicy_Policy = QSizePolicy_Policy(QSizePolicy_GrowFlag)|QSizePolicy_Policy(QSizePolicy_ShrinkFlag)|QSizePolicy_Policy(QSizePolicy_ExpandFlag)
	QSizePolicy_Ignored QSizePolicy_Policy = QSizePolicy_Policy(QSizePolicy_ShrinkFlag)|QSizePolicy_Policy(QSizePolicy_GrowFlag)|QSizePolicy_Policy(QSizePolicy_IgnoreFlag)
)
//struct QSizePolicy : QSizePolicy
type QSizePolicy struct {
	BaseDrv
}
//QSizePolicy::QSizePolicy()	
func NewSizePolicy() *QSizePolicy {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),121000,121102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QSizePolicy{}
	_p.SetDriver(__rv,121000,true)
	return _p
} 
//QSizePolicy::QSizePolicy(QSizePolicy const&)	
func NewSizePolicyCopy(other *QSizePolicy) *QSizePolicy {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),121000,121103,Native(other),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QSizePolicy{}
	_p.SetDriver(__rv,121000,true)
	return _p
} 
//QSizePolicy::QSizePolicy(QSizePolicy::Policy,QSizePolicy::Policy)	
func NewSizePolicyWithHorizontalVertical(horizontal QSizePolicy_Policy,vertical QSizePolicy_Policy) *QSizePolicy {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),121000,121104,unsafe.Pointer(&horizontal),unsafe.Pointer(&vertical),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QSizePolicy{}
	_p.SetDriver(__rv,121000,true)
	return _p
} 
//QSizePolicy::QSizePolicy(QSizePolicy::Policy,QSizePolicy::Policy,QSizePolicy::ControlType)	
func NewSizePolicyWithHorizontalVerticalType(horizontal QSizePolicy_Policy,vertical QSizePolicy_Policy,_type QSizePolicy_ControlType) *QSizePolicy {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),121000,121105,unsafe.Pointer(&horizontal),unsafe.Pointer(&vertical),unsafe.Pointer(&_type),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QSizePolicy{}
	_p.SetDriver(__rv,121000,true)
	return _p
} 
//QSizePolicy::controlType()
func (q *QSizePolicy) ControlType() QSizePolicy_ControlType {
	var __rv QSizePolicy_ControlType
	q.Drv(121000,121106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSizePolicy::expandingDirections()
func (q *QSizePolicy) ExpandingDirections() Qt_Orientation {
	var __rv Qt_Orientation
	q.Drv(121000,121107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSizePolicy::hasHeightForWidth()
func (q *QSizePolicy) HasHeightForWidth() bool {
	var __rv bool
	q.Drv(121000,121108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSizePolicy::horizontalPolicy()
func (q *QSizePolicy) HorizontalPolicy() QSizePolicy_Policy {
	var __rv QSizePolicy_Policy
	q.Drv(121000,121109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSizePolicy::horizontalStretch()
func (q *QSizePolicy) HorizontalStretch() int {
	var __rv int
	q.Drv(121000,121110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSizePolicy::setControlType(QSizePolicy::ControlType)
func (q *QSizePolicy) SetControlType(_type QSizePolicy_ControlType)  {
	q.Drv(121000,121111,unsafe.Pointer(&_type),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSizePolicy::setHeightForWidth(bool)
func (q *QSizePolicy) SetHeightForWidth(b bool)  {
	q.Drv(121000,121112,unsafe.Pointer(&b),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSizePolicy::setHorizontalPolicy(QSizePolicy::Policy)
func (q *QSizePolicy) SetHorizontalPolicy(d QSizePolicy_Policy)  {
	q.Drv(121000,121113,unsafe.Pointer(&d),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSizePolicy::setHorizontalStretch(unsigned char)
func (q *QSizePolicy) SetHorizontalStretch(stretchFactor byte)  {
	q.Drv(121000,121114,unsafe.Pointer(&stretchFactor),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSizePolicy::setVerticalPolicy(QSizePolicy::Policy)
func (q *QSizePolicy) SetVerticalPolicy(d QSizePolicy_Policy)  {
	q.Drv(121000,121115,unsafe.Pointer(&d),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSizePolicy::setVerticalStretch(unsigned char)
func (q *QSizePolicy) SetVerticalStretch(stretchFactor byte)  {
	q.Drv(121000,121116,unsafe.Pointer(&stretchFactor),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSizePolicy::transpose()
func (q *QSizePolicy) Transpose()  {
	q.Drv(121000,121117,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSizePolicy::verticalPolicy()
func (q *QSizePolicy) VerticalPolicy() QSizePolicy_Policy {
	var __rv QSizePolicy_Policy
	q.Drv(121000,121118,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSizePolicy::verticalStretch()
func (q *QSizePolicy) VerticalStretch() int {
	var __rv int
	q.Drv(121000,121119,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
