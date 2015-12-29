// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QVariantComparisonHelper : QVariantComparisonHelper
type QVariantComparisonHelper struct {
	BaseDrv
}
//QVariantComparisonHelper::QVariantComparisonHelper(QVariant const&)	
func NewVariantComparisonHelper(_var *QVariant) *QVariantComparisonHelper {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),183000,183102,Native(_var),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QVariantComparisonHelper{}
	_p.SetDriver(__rv,183000,true)
	return _p
} 
