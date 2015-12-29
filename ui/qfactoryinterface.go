// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QFactoryInterface : QFactoryInterface
type QFactoryInterface struct {
	BaseDrv
}
//QFactoryInterface::keys()
func (q *QFactoryInterface) Keys() []string {
	var __rv []string
	q.Drv(32000,32102,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
