// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QGraphicsAnchor : QGraphicsAnchor
type QGraphicsAnchor struct {
	QObject
}
//QGraphicsAnchor::setSizePolicy(QSizePolicy::Policy)
func (q *QGraphicsAnchor) SetSizePolicy(policy QSizePolicy_Policy)  {
	q.Drv(248000,248102,unsafe.Pointer(&policy),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsAnchor::setSpacing(double)
func (q *QGraphicsAnchor) SetSpacing(spacing float64)  {
	q.Drv(248000,248103,unsafe.Pointer(&spacing),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsAnchor::sizePolicy()
func (q *QGraphicsAnchor) SizePolicy() QSizePolicy_Policy {
	var __rv QSizePolicy_Policy
	q.Drv(248000,248104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsAnchor::spacing()
func (q *QGraphicsAnchor) Spacing() float64 {
	var __rv float64
	q.Drv(248000,248105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsAnchor::unsetSpacing()
func (q *QGraphicsAnchor) UnsetSpacing()  {
	q.Drv(248000,248106,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
