// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QTileRules : QTileRules
type QTileRules struct {
	BaseDrv
}
//QTileRules::QTileRules()	
func NewTileRules() *QTileRules {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),169000,169102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTileRules{}
	_p.SetDriver(__rv,169000,true)
	return _p
} 
//QTileRules::QTileRules(Qt::TileRule)	
func NewTileRulesWithRule(rule Qt_TileRule) *QTileRules {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),169000,169103,unsafe.Pointer(&rule),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTileRules{}
	_p.SetDriver(__rv,169000,true)
	return _p
} 
//QTileRules::QTileRules(Qt::TileRule,Qt::TileRule)	
func NewTileRulesWithHorizontalruleVerticalrule(horizontalRule Qt_TileRule,verticalRule Qt_TileRule) *QTileRules {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),169000,169104,unsafe.Pointer(&horizontalRule),unsafe.Pointer(&verticalRule),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTileRules{}
	_p.SetDriver(__rv,169000,true)
	return _p
} 
