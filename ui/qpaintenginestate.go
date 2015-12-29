// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QPaintEngineState : QPaintEngineState
type QPaintEngineState struct {
	BaseDrv
}
//QPaintEngineState::QPaintEngineState()	
func NewPaintEngineState() *QPaintEngineState {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),84000,84102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPaintEngineState{}
	_p.SetDriver(__rv,84000,true)
	return _p
} 
//QPaintEngineState::backgroundBrush()
func (q *QPaintEngineState) BackgroundBrush() *QBrush {
	var __rv uintptr
	q.Drv(84000,84103,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QBrush{}
	_rp.SetDriver(__rv,9000,true)
	return _rp
}	
//QPaintEngineState::backgroundMode()
func (q *QPaintEngineState) BackgroundMode() Qt_BGMode {
	var __rv Qt_BGMode
	q.Drv(84000,84104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPaintEngineState::brush()
func (q *QPaintEngineState) Brush() *QBrush {
	var __rv uintptr
	q.Drv(84000,84105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QBrush{}
	_rp.SetDriver(__rv,9000,true)
	return _rp
}	
//QPaintEngineState::brushNeedsResolving()
func (q *QPaintEngineState) BrushNeedsResolving() bool {
	var __rv bool
	q.Drv(84000,84106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPaintEngineState::brushOrigin()
func (q *QPaintEngineState) BrushOrigin() *QPointF {
	var __rv uintptr
	q.Drv(84000,84107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QPaintEngineState::clipOperation()
func (q *QPaintEngineState) ClipOperation() Qt_ClipOperation {
	var __rv Qt_ClipOperation
	q.Drv(84000,84108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPaintEngineState::clipPath()
func (q *QPaintEngineState) ClipPath() *QPainterPath {
	var __rv uintptr
	q.Drv(84000,84109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPainterPath{}
	_rp.SetDriver(__rv,88000,true)
	return _rp
}	
//QPaintEngineState::clipRegion()
func (q *QPaintEngineState) ClipRegion() *QRegion {
	var __rv uintptr
	q.Drv(84000,84110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRegion{}
	_rp.SetDriver(__rv,113000,true)
	return _rp
}	
//QPaintEngineState::compositionMode()
func (q *QPaintEngineState) CompositionMode() QPainter_CompositionMode {
	var __rv QPainter_CompositionMode
	q.Drv(84000,84111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPaintEngineState::font()
func (q *QPaintEngineState) Font() *QFont {
	var __rv uintptr
	q.Drv(84000,84112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QFont{}
	_rp.SetDriver(__rv,37000,true)
	return _rp
}	
//QPaintEngineState::isClipEnabled()
func (q *QPaintEngineState) IsClipEnabled() bool {
	var __rv bool
	q.Drv(84000,84113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPaintEngineState::matrix()
func (q *QPaintEngineState) Matrix() *QMatrix {
	var __rv uintptr
	q.Drv(84000,84114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QMatrix{}
	_rp.SetDriver(__rv,74000,true)
	return _rp
}	
//QPaintEngineState::opacity()
func (q *QPaintEngineState) Opacity() float64 {
	var __rv float64
	q.Drv(84000,84115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPaintEngineState::painter()
func (q *QPaintEngineState) Painter() *QPainter {
	var __rv uintptr
	q.Drv(84000,84116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPainter{}
	_rp.SetDriver(__rv,86000,true)
	return _rp
}	
//QPaintEngineState::pen()
func (q *QPaintEngineState) Pen() *QPen {
	var __rv uintptr
	q.Drv(84000,84117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPen{}
	_rp.SetDriver(__rv,92000,true)
	return _rp
}	
//QPaintEngineState::penNeedsResolving()
func (q *QPaintEngineState) PenNeedsResolving() bool {
	var __rv bool
	q.Drv(84000,84118,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPaintEngineState::renderHints()
func (q *QPaintEngineState) RenderHints() QPainter_RenderHint {
	var __rv QPainter_RenderHint
	q.Drv(84000,84119,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPaintEngineState::state()
func (q *QPaintEngineState) State() QPaintEngine_DirtyFlag {
	var __rv QPaintEngine_DirtyFlag
	q.Drv(84000,84120,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPaintEngineState::transform()
func (q *QPaintEngineState) Transform() *QTransform {
	var __rv uintptr
	q.Drv(84000,84121,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTransform{}
	_rp.SetDriver(__rv,176000,true)
	return _rp
}	
