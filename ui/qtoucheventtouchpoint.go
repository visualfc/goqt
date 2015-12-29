// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QTouchEventTouchPoint : QTouchEvent::TouchPoint
type QTouchEventTouchPoint struct {
	BaseDrv
}
//QTouchEvent::TouchPoint::TouchPoint()	
func NewTouchEventTouchPoint() *QTouchEventTouchPoint {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),175000,175102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTouchEventTouchPoint{}
	_p.SetDriver(__rv,175000,true)
	return _p
} 
//QTouchEvent::TouchPoint::TouchPoint(QTouchEvent::TouchPoint const&)	
func NewTouchEventTouchPointCopy(other *QTouchEventTouchPoint) *QTouchEventTouchPoint {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),175000,175103,Native(other),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTouchEventTouchPoint{}
	_p.SetDriver(__rv,175000,true)
	return _p
} 
//QTouchEvent::TouchPoint::TouchPoint(int)	
func NewTouchEventTouchPointWithId(id int) *QTouchEventTouchPoint {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),175000,175104,unsafe.Pointer(&id),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTouchEventTouchPoint{}
	_p.SetDriver(__rv,175000,true)
	return _p
} 
//QTouchEvent::TouchPoint::id()
func (q *QTouchEventTouchPoint) Id() int {
	var __rv int
	q.Drv(175000,175105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTouchEvent::TouchPoint::lastNormalizedPos()
func (q *QTouchEventTouchPoint) LastNormalizedPos() *QPointF {
	var __rv uintptr
	q.Drv(175000,175106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QTouchEvent::TouchPoint::lastPos()
func (q *QTouchEventTouchPoint) LastPos() *QPointF {
	var __rv uintptr
	q.Drv(175000,175107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QTouchEvent::TouchPoint::lastScenePos()
func (q *QTouchEventTouchPoint) LastScenePos() *QPointF {
	var __rv uintptr
	q.Drv(175000,175108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QTouchEvent::TouchPoint::lastScreenPos()
func (q *QTouchEventTouchPoint) LastScreenPos() *QPointF {
	var __rv uintptr
	q.Drv(175000,175109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QTouchEvent::TouchPoint::normalizedPos()
func (q *QTouchEventTouchPoint) NormalizedPos() *QPointF {
	var __rv uintptr
	q.Drv(175000,175110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QTouchEvent::TouchPoint::pos()
func (q *QTouchEventTouchPoint) Pos() *QPointF {
	var __rv uintptr
	q.Drv(175000,175111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QTouchEvent::TouchPoint::pressure()
func (q *QTouchEventTouchPoint) Pressure() float64 {
	var __rv float64
	q.Drv(175000,175112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTouchEvent::TouchPoint::rect()
func (q *QTouchEventTouchPoint) Rect() *QRectF {
	var __rv uintptr
	q.Drv(175000,175113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRectF{}
	_rp.SetDriver(__rv,111000,true)
	return _rp
}	
//QTouchEvent::TouchPoint::scenePos()
func (q *QTouchEventTouchPoint) ScenePos() *QPointF {
	var __rv uintptr
	q.Drv(175000,175114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QTouchEvent::TouchPoint::sceneRect()
func (q *QTouchEventTouchPoint) SceneRect() *QRectF {
	var __rv uintptr
	q.Drv(175000,175115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRectF{}
	_rp.SetDriver(__rv,111000,true)
	return _rp
}	
//QTouchEvent::TouchPoint::screenPos()
func (q *QTouchEventTouchPoint) ScreenPos() *QPointF {
	var __rv uintptr
	q.Drv(175000,175116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QTouchEvent::TouchPoint::screenRect()
func (q *QTouchEventTouchPoint) ScreenRect() *QRectF {
	var __rv uintptr
	q.Drv(175000,175117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRectF{}
	_rp.SetDriver(__rv,111000,true)
	return _rp
}	
//QTouchEvent::TouchPoint::setId(int)
func (q *QTouchEventTouchPoint) SetId(id int)  {
	q.Drv(175000,175118,unsafe.Pointer(&id),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTouchEvent::TouchPoint::setLastNormalizedPos(QPointF const&)
func (q *QTouchEventTouchPoint) SetLastNormalizedPos(lastNormalizedPos *QPointF)  {
	q.Drv(175000,175119,Native(lastNormalizedPos),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTouchEvent::TouchPoint::setLastPos(QPointF const&)
func (q *QTouchEventTouchPoint) SetLastPos(lastPos *QPointF)  {
	q.Drv(175000,175120,Native(lastPos),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTouchEvent::TouchPoint::setLastScenePos(QPointF const&)
func (q *QTouchEventTouchPoint) SetLastScenePos(lastScenePos *QPointF)  {
	q.Drv(175000,175121,Native(lastScenePos),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTouchEvent::TouchPoint::setLastScreenPos(QPointF const&)
func (q *QTouchEventTouchPoint) SetLastScreenPos(lastScreenPos *QPointF)  {
	q.Drv(175000,175122,Native(lastScreenPos),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTouchEvent::TouchPoint::setNormalizedPos(QPointF const&)
func (q *QTouchEventTouchPoint) SetNormalizedPos(normalizedPos *QPointF)  {
	q.Drv(175000,175123,Native(normalizedPos),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTouchEvent::TouchPoint::setPos(QPointF const&)
func (q *QTouchEventTouchPoint) SetPos(pos *QPointF)  {
	q.Drv(175000,175124,Native(pos),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTouchEvent::TouchPoint::setPressure(double)
func (q *QTouchEventTouchPoint) SetPressure(pressure float64)  {
	q.Drv(175000,175125,unsafe.Pointer(&pressure),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTouchEvent::TouchPoint::setRect(QRectF const&)
func (q *QTouchEventTouchPoint) SetRect(rect *QRectF)  {
	q.Drv(175000,175126,Native(rect),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTouchEvent::TouchPoint::setScenePos(QPointF const&)
func (q *QTouchEventTouchPoint) SetScenePos(scenePos *QPointF)  {
	q.Drv(175000,175127,Native(scenePos),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTouchEvent::TouchPoint::setSceneRect(QRectF const&)
func (q *QTouchEventTouchPoint) SetSceneRect(sceneRect *QRectF)  {
	q.Drv(175000,175128,Native(sceneRect),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTouchEvent::TouchPoint::setScreenPos(QPointF const&)
func (q *QTouchEventTouchPoint) SetScreenPos(screenPos *QPointF)  {
	q.Drv(175000,175129,Native(screenPos),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTouchEvent::TouchPoint::setScreenRect(QRectF const&)
func (q *QTouchEventTouchPoint) SetScreenRect(screenRect *QRectF)  {
	q.Drv(175000,175130,Native(screenRect),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTouchEvent::TouchPoint::setStartNormalizedPos(QPointF const&)
func (q *QTouchEventTouchPoint) SetStartNormalizedPos(startNormalizedPos *QPointF)  {
	q.Drv(175000,175131,Native(startNormalizedPos),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTouchEvent::TouchPoint::setStartPos(QPointF const&)
func (q *QTouchEventTouchPoint) SetStartPos(startPos *QPointF)  {
	q.Drv(175000,175132,Native(startPos),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTouchEvent::TouchPoint::setStartScenePos(QPointF const&)
func (q *QTouchEventTouchPoint) SetStartScenePos(startScenePos *QPointF)  {
	q.Drv(175000,175133,Native(startScenePos),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTouchEvent::TouchPoint::setStartScreenPos(QPointF const&)
func (q *QTouchEventTouchPoint) SetStartScreenPos(startScreenPos *QPointF)  {
	q.Drv(175000,175134,Native(startScreenPos),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTouchEvent::TouchPoint::setState(QFlags<Qt::TouchPointState>)
func (q *QTouchEventTouchPoint) SetState(state Qt_TouchPointState)  {
	q.Drv(175000,175135,unsafe.Pointer(&state),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTouchEvent::TouchPoint::startNormalizedPos()
func (q *QTouchEventTouchPoint) StartNormalizedPos() *QPointF {
	var __rv uintptr
	q.Drv(175000,175136,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QTouchEvent::TouchPoint::startPos()
func (q *QTouchEventTouchPoint) StartPos() *QPointF {
	var __rv uintptr
	q.Drv(175000,175137,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QTouchEvent::TouchPoint::startScenePos()
func (q *QTouchEventTouchPoint) StartScenePos() *QPointF {
	var __rv uintptr
	q.Drv(175000,175138,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QTouchEvent::TouchPoint::startScreenPos()
func (q *QTouchEventTouchPoint) StartScreenPos() *QPointF {
	var __rv uintptr
	q.Drv(175000,175139,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QTouchEvent::TouchPoint::state()
func (q *QTouchEventTouchPoint) State() Qt_TouchPointState {
	var __rv Qt_TouchPointState
	q.Drv(175000,175140,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
