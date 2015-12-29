// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QColor_Spec - QColor::Spec
type QColor_Spec uint32
const (
	QColor_Invalid QColor_Spec = 0
	QColor_Rgb QColor_Spec = 1
	QColor_Hsv QColor_Spec = 2
	QColor_Cmyk QColor_Spec = 3
	QColor_Hsl QColor_Spec = 4
)
//struct QColor : QColor
type QColor struct {
	BaseDrv
}
//QColor::QColor()	
func NewColor() *QColor {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),13000,13102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QColor{}
	_p.SetDriver(__rv,13000,true)
	return _p
} 
//QColor::QColor(QColor const&)	
func NewColorCopy(color *QColor) *QColor {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),13000,13103,Native(color),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QColor{}
	_p.SetDriver(__rv,13000,true)
	return _p
} 
//QColor::QColor(QColor::Spec)	
func NewColorWithSpec(spec QColor_Spec) *QColor {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),13000,13104,unsafe.Pointer(&spec),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QColor{}
	_p.SetDriver(__rv,13000,true)
	return _p
} 
//QColor::QColor(QString const&)	
func NewColorWithName(name string) *QColor {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),13000,13105,unsafe.Pointer(&name),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QColor{}
	_p.SetDriver(__rv,13000,true)
	return _p
} 
//QColor::QColor(Qt::GlobalColor)	
func NewColorWithGlobalcolor(color Qt_GlobalColor) *QColor {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),13000,13106,unsafe.Pointer(&color),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QColor{}
	_p.SetDriver(__rv,13000,true)
	return _p
} 
//QColor::QColor(unsigned int)	
func NewColorWithRgb(rgb uint) *QColor {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),13000,13107,unsafe.Pointer(&rgb),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QColor{}
	_p.SetDriver(__rv,13000,true)
	return _p
} 
//QColor::QColor(int,int,int,int)	
func NewColorWithRGBA(r int,g int,b int,a int) *QColor {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),13000,13108,unsafe.Pointer(&r),unsafe.Pointer(&g),unsafe.Pointer(&b),unsafe.Pointer(&a),nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QColor{}
	_p.SetDriver(__rv,13000,true)
	return _p
} 
//QColor::alpha()
func (q *QColor) Alpha() int {
	var __rv int
	q.Drv(13000,13109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QColor::alphaF()
func (q *QColor) AlphaF() float64 {
	var __rv float64
	q.Drv(13000,13110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QColor::black()
func (q *QColor) Black() int {
	var __rv int
	q.Drv(13000,13111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QColor::blackF()
func (q *QColor) BlackF() float64 {
	var __rv float64
	q.Drv(13000,13112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QColor::blue()
func (q *QColor) Blue() int {
	var __rv int
	q.Drv(13000,13113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QColor::blueF()
func (q *QColor) BlueF() float64 {
	var __rv float64
	q.Drv(13000,13114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QColor::colorNames()	
func QColorColorNames() []string {
	var __rv []string
	DirectQtDrv(nil,13000,13115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QColor::colorNames()
func (q *QColor) ColorNames() []string {
	var __rv []string
	q.Drv(13000,13115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QColor::convertTo(QColor::Spec)
func (q *QColor) ConvertTo(colorSpec QColor_Spec) *QColor {
	var __rv uintptr
	q.Drv(13000,13116,unsafe.Pointer(&colorSpec),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QColor{}
	_rp.SetDriver(__rv,13000,true)
	return _rp
}	
//QColor::cyan()
func (q *QColor) Cyan() int {
	var __rv int
	q.Drv(13000,13117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QColor::cyanF()
func (q *QColor) CyanF() float64 {
	var __rv float64
	q.Drv(13000,13118,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QColor::dark()
func (q *QColor) Dark() *QColor {
	var __rv uintptr
	q.Drv(13000,13119,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QColor{}
	_rp.SetDriver(__rv,13000,true)
	return _rp
}	
//QColor::dark(int)
func (q *QColor) DarkWithInt(f int) *QColor {
	var __rv uintptr
	q.Drv(13000,13120,unsafe.Pointer(&f),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QColor{}
	_rp.SetDriver(__rv,13000,true)
	return _rp
}	
//QColor::darker()
func (q *QColor) Darker() *QColor {
	var __rv uintptr
	q.Drv(13000,13121,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QColor{}
	_rp.SetDriver(__rv,13000,true)
	return _rp
}	
//QColor::darker(int)
func (q *QColor) DarkerWithInt(f int) *QColor {
	var __rv uintptr
	q.Drv(13000,13122,unsafe.Pointer(&f),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QColor{}
	_rp.SetDriver(__rv,13000,true)
	return _rp
}	
//QColor::fromCmyk(int,int,int,int,int)	
func QColorFromCmyk(c int,m int,y int,k int,a int) *QColor {
	var __rv uintptr
	DirectQtDrv(nil,13000,13123,unsafe.Pointer(&c),unsafe.Pointer(&m),unsafe.Pointer(&y),unsafe.Pointer(&k),unsafe.Pointer(&a),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QColor{}
	_rp.SetDriver(__rv,13000,true)
	return _rp
}	
//QColor::fromCmyk(int,int,int,int,int)
func (q *QColor) FromCmyk(c int,m int,y int,k int,a int) *QColor {
	var __rv uintptr
	q.Drv(13000,13123,unsafe.Pointer(&c),unsafe.Pointer(&m),unsafe.Pointer(&y),unsafe.Pointer(&k),unsafe.Pointer(&a),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QColor{}
	_rp.SetDriver(__rv,13000,true)
	return _rp
}	
//QColor::fromCmykF(double,double,double,double,double)	
func QColorFromCmykF(c float64,m float64,y float64,k float64,a float64) *QColor {
	var __rv uintptr
	DirectQtDrv(nil,13000,13124,unsafe.Pointer(&c),unsafe.Pointer(&m),unsafe.Pointer(&y),unsafe.Pointer(&k),unsafe.Pointer(&a),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QColor{}
	_rp.SetDriver(__rv,13000,true)
	return _rp
}	
//QColor::fromCmykF(double,double,double,double,double)
func (q *QColor) FromCmykF(c float64,m float64,y float64,k float64,a float64) *QColor {
	var __rv uintptr
	q.Drv(13000,13124,unsafe.Pointer(&c),unsafe.Pointer(&m),unsafe.Pointer(&y),unsafe.Pointer(&k),unsafe.Pointer(&a),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QColor{}
	_rp.SetDriver(__rv,13000,true)
	return _rp
}	
//QColor::fromHsl(int,int,int,int)	
func QColorFromHsl(h int,s int,l int,a int) *QColor {
	var __rv uintptr
	DirectQtDrv(nil,13000,13125,unsafe.Pointer(&h),unsafe.Pointer(&s),unsafe.Pointer(&l),unsafe.Pointer(&a),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QColor{}
	_rp.SetDriver(__rv,13000,true)
	return _rp
}	
//QColor::fromHsl(int,int,int,int)
func (q *QColor) FromHsl(h int,s int,l int,a int) *QColor {
	var __rv uintptr
	q.Drv(13000,13125,unsafe.Pointer(&h),unsafe.Pointer(&s),unsafe.Pointer(&l),unsafe.Pointer(&a),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QColor{}
	_rp.SetDriver(__rv,13000,true)
	return _rp
}	
//QColor::fromHslF(double,double,double,double)	
func QColorFromHslF(h float64,s float64,l float64,a float64) *QColor {
	var __rv uintptr
	DirectQtDrv(nil,13000,13126,unsafe.Pointer(&h),unsafe.Pointer(&s),unsafe.Pointer(&l),unsafe.Pointer(&a),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QColor{}
	_rp.SetDriver(__rv,13000,true)
	return _rp
}	
//QColor::fromHslF(double,double,double,double)
func (q *QColor) FromHslF(h float64,s float64,l float64,a float64) *QColor {
	var __rv uintptr
	q.Drv(13000,13126,unsafe.Pointer(&h),unsafe.Pointer(&s),unsafe.Pointer(&l),unsafe.Pointer(&a),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QColor{}
	_rp.SetDriver(__rv,13000,true)
	return _rp
}	
//QColor::fromHsv(int,int,int,int)	
func QColorFromHsv(h int,s int,v int,a int) *QColor {
	var __rv uintptr
	DirectQtDrv(nil,13000,13127,unsafe.Pointer(&h),unsafe.Pointer(&s),unsafe.Pointer(&v),unsafe.Pointer(&a),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QColor{}
	_rp.SetDriver(__rv,13000,true)
	return _rp
}	
//QColor::fromHsv(int,int,int,int)
func (q *QColor) FromHsv(h int,s int,v int,a int) *QColor {
	var __rv uintptr
	q.Drv(13000,13127,unsafe.Pointer(&h),unsafe.Pointer(&s),unsafe.Pointer(&v),unsafe.Pointer(&a),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QColor{}
	_rp.SetDriver(__rv,13000,true)
	return _rp
}	
//QColor::fromHsvF(double,double,double,double)	
func QColorFromHsvF(h float64,s float64,v float64,a float64) *QColor {
	var __rv uintptr
	DirectQtDrv(nil,13000,13128,unsafe.Pointer(&h),unsafe.Pointer(&s),unsafe.Pointer(&v),unsafe.Pointer(&a),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QColor{}
	_rp.SetDriver(__rv,13000,true)
	return _rp
}	
//QColor::fromHsvF(double,double,double,double)
func (q *QColor) FromHsvF(h float64,s float64,v float64,a float64) *QColor {
	var __rv uintptr
	q.Drv(13000,13128,unsafe.Pointer(&h),unsafe.Pointer(&s),unsafe.Pointer(&v),unsafe.Pointer(&a),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QColor{}
	_rp.SetDriver(__rv,13000,true)
	return _rp
}	
//QColor::fromRgb(unsigned int)	
func QColorFromRgb(rgb uint) *QColor {
	var __rv uintptr
	DirectQtDrv(nil,13000,13129,unsafe.Pointer(&rgb),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QColor{}
	_rp.SetDriver(__rv,13000,true)
	return _rp
}	
//QColor::fromRgb(unsigned int)
func (q *QColor) FromRgb(rgb uint) *QColor {
	var __rv uintptr
	q.Drv(13000,13129,unsafe.Pointer(&rgb),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QColor{}
	_rp.SetDriver(__rv,13000,true)
	return _rp
}	
//QColor::fromRgb(int,int,int,int)	
func QColorFromRgbWithRGBA(r int,g int,b int,a int) *QColor {
	var __rv uintptr
	DirectQtDrv(nil,13000,13130,unsafe.Pointer(&r),unsafe.Pointer(&g),unsafe.Pointer(&b),unsafe.Pointer(&a),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QColor{}
	_rp.SetDriver(__rv,13000,true)
	return _rp
}	
//QColor::fromRgb(int,int,int,int)
func (q *QColor) FromRgbWithRGBA(r int,g int,b int,a int) *QColor {
	var __rv uintptr
	q.Drv(13000,13130,unsafe.Pointer(&r),unsafe.Pointer(&g),unsafe.Pointer(&b),unsafe.Pointer(&a),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QColor{}
	_rp.SetDriver(__rv,13000,true)
	return _rp
}	
//QColor::fromRgbF(double,double,double,double)	
func QColorFromRgbF(r float64,g float64,b float64,a float64) *QColor {
	var __rv uintptr
	DirectQtDrv(nil,13000,13131,unsafe.Pointer(&r),unsafe.Pointer(&g),unsafe.Pointer(&b),unsafe.Pointer(&a),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QColor{}
	_rp.SetDriver(__rv,13000,true)
	return _rp
}	
//QColor::fromRgbF(double,double,double,double)
func (q *QColor) FromRgbF(r float64,g float64,b float64,a float64) *QColor {
	var __rv uintptr
	q.Drv(13000,13131,unsafe.Pointer(&r),unsafe.Pointer(&g),unsafe.Pointer(&b),unsafe.Pointer(&a),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QColor{}
	_rp.SetDriver(__rv,13000,true)
	return _rp
}	
//QColor::fromRgba(unsigned int)	
func QColorFromRgba(rgba uint) *QColor {
	var __rv uintptr
	DirectQtDrv(nil,13000,13132,unsafe.Pointer(&rgba),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QColor{}
	_rp.SetDriver(__rv,13000,true)
	return _rp
}	
//QColor::fromRgba(unsigned int)
func (q *QColor) FromRgba(rgba uint) *QColor {
	var __rv uintptr
	q.Drv(13000,13132,unsafe.Pointer(&rgba),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QColor{}
	_rp.SetDriver(__rv,13000,true)
	return _rp
}	
//QColor::getCmyk(int*,int*,int*,int*,int*)
func (q *QColor) GetCmyk(c *int,m *int,y *int,k *int,a *int)  {
	q.Drv(13000,13133,unsafe.Pointer(&c),unsafe.Pointer(&m),unsafe.Pointer(&y),unsafe.Pointer(&k),unsafe.Pointer(&a),nil,nil,nil,nil,nil,nil,nil)
}	
//QColor::getCmykF(double*,double*,double*,double*,double*)
func (q *QColor) GetCmykF(c *float64,m *float64,y *float64,k *float64,a *float64)  {
	q.Drv(13000,13134,unsafe.Pointer(&c),unsafe.Pointer(&m),unsafe.Pointer(&y),unsafe.Pointer(&k),unsafe.Pointer(&a),nil,nil,nil,nil,nil,nil,nil)
}	
//QColor::getHsl(int*,int*,int*,int*)
func (q *QColor) GetHsl(h *int,s *int,l *int,a *int)  {
	q.Drv(13000,13135,unsafe.Pointer(&h),unsafe.Pointer(&s),unsafe.Pointer(&l),unsafe.Pointer(&a),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QColor::getHslF(double*,double*,double*,double*)
func (q *QColor) GetHslF(h *float64,s *float64,l *float64,a *float64)  {
	q.Drv(13000,13136,unsafe.Pointer(&h),unsafe.Pointer(&s),unsafe.Pointer(&l),unsafe.Pointer(&a),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QColor::getHsv(int*,int*,int*,int*)
func (q *QColor) GetHsv(h *int,s *int,v *int,a *int)  {
	q.Drv(13000,13137,unsafe.Pointer(&h),unsafe.Pointer(&s),unsafe.Pointer(&v),unsafe.Pointer(&a),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QColor::getHsvF(double*,double*,double*,double*)
func (q *QColor) GetHsvF(h *float64,s *float64,v *float64,a *float64)  {
	q.Drv(13000,13138,unsafe.Pointer(&h),unsafe.Pointer(&s),unsafe.Pointer(&v),unsafe.Pointer(&a),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QColor::getRgb(int*,int*,int*,int*)
func (q *QColor) GetRgb(r *int,g *int,b *int,a *int)  {
	q.Drv(13000,13139,unsafe.Pointer(&r),unsafe.Pointer(&g),unsafe.Pointer(&b),unsafe.Pointer(&a),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QColor::getRgbF(double*,double*,double*,double*)
func (q *QColor) GetRgbF(r *float64,g *float64,b *float64,a *float64)  {
	q.Drv(13000,13140,unsafe.Pointer(&r),unsafe.Pointer(&g),unsafe.Pointer(&b),unsafe.Pointer(&a),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QColor::green()
func (q *QColor) Green() int {
	var __rv int
	q.Drv(13000,13141,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QColor::greenF()
func (q *QColor) GreenF() float64 {
	var __rv float64
	q.Drv(13000,13142,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QColor::hslHue()
func (q *QColor) HslHue() int {
	var __rv int
	q.Drv(13000,13143,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QColor::hslHueF()
func (q *QColor) HslHueF() float64 {
	var __rv float64
	q.Drv(13000,13144,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QColor::hslSaturation()
func (q *QColor) HslSaturation() int {
	var __rv int
	q.Drv(13000,13145,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QColor::hslSaturationF()
func (q *QColor) HslSaturationF() float64 {
	var __rv float64
	q.Drv(13000,13146,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QColor::hsvHue()
func (q *QColor) HsvHue() int {
	var __rv int
	q.Drv(13000,13147,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QColor::hsvHueF()
func (q *QColor) HsvHueF() float64 {
	var __rv float64
	q.Drv(13000,13148,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QColor::hsvSaturation()
func (q *QColor) HsvSaturation() int {
	var __rv int
	q.Drv(13000,13149,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QColor::hsvSaturationF()
func (q *QColor) HsvSaturationF() float64 {
	var __rv float64
	q.Drv(13000,13150,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QColor::hue()
func (q *QColor) Hue() int {
	var __rv int
	q.Drv(13000,13151,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QColor::hueF()
func (q *QColor) HueF() float64 {
	var __rv float64
	q.Drv(13000,13152,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QColor::isValid()
func (q *QColor) IsValid() bool {
	var __rv bool
	q.Drv(13000,13153,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QColor::isValidColor(QString const&)	
func QColorIsValidColor(name string) bool {
	var __rv bool
	DirectQtDrv(nil,13000,13154,unsafe.Pointer(&name),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QColor::isValidColor(QString const&)
func (q *QColor) IsValidColor(name string) bool {
	var __rv bool
	q.Drv(13000,13154,unsafe.Pointer(&name),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QColor::light()
func (q *QColor) Light() *QColor {
	var __rv uintptr
	q.Drv(13000,13155,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QColor{}
	_rp.SetDriver(__rv,13000,true)
	return _rp
}	
//QColor::light(int)
func (q *QColor) LightWithInt(f int) *QColor {
	var __rv uintptr
	q.Drv(13000,13156,unsafe.Pointer(&f),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QColor{}
	_rp.SetDriver(__rv,13000,true)
	return _rp
}	
//QColor::lighter()
func (q *QColor) Lighter() *QColor {
	var __rv uintptr
	q.Drv(13000,13157,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QColor{}
	_rp.SetDriver(__rv,13000,true)
	return _rp
}	
//QColor::lighter(int)
func (q *QColor) LighterWithInt(f int) *QColor {
	var __rv uintptr
	q.Drv(13000,13158,unsafe.Pointer(&f),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QColor{}
	_rp.SetDriver(__rv,13000,true)
	return _rp
}	
//QColor::lightness()
func (q *QColor) Lightness() int {
	var __rv int
	q.Drv(13000,13159,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QColor::lightnessF()
func (q *QColor) LightnessF() float64 {
	var __rv float64
	q.Drv(13000,13160,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QColor::magenta()
func (q *QColor) Magenta() int {
	var __rv int
	q.Drv(13000,13161,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QColor::magentaF()
func (q *QColor) MagentaF() float64 {
	var __rv float64
	q.Drv(13000,13162,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QColor::name()
func (q *QColor) Name() string {
	var __rv string
	q.Drv(13000,13163,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QColor::red()
func (q *QColor) Red() int {
	var __rv int
	q.Drv(13000,13164,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QColor::redF()
func (q *QColor) RedF() float64 {
	var __rv float64
	q.Drv(13000,13165,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QColor::rgb()
func (q *QColor) Rgb() uint {
	var __rv uint
	q.Drv(13000,13166,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QColor::rgba()
func (q *QColor) Rgba() uint {
	var __rv uint
	q.Drv(13000,13167,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QColor::saturation()
func (q *QColor) Saturation() int {
	var __rv int
	q.Drv(13000,13168,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QColor::saturationF()
func (q *QColor) SaturationF() float64 {
	var __rv float64
	q.Drv(13000,13169,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QColor::setAlpha(int)
func (q *QColor) SetAlpha(alpha int)  {
	q.Drv(13000,13170,unsafe.Pointer(&alpha),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QColor::setAlphaF(double)
func (q *QColor) SetAlphaF(alpha float64)  {
	q.Drv(13000,13171,unsafe.Pointer(&alpha),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QColor::setBlue(int)
func (q *QColor) SetBlue(blue int)  {
	q.Drv(13000,13172,unsafe.Pointer(&blue),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QColor::setBlueF(double)
func (q *QColor) SetBlueF(blue float64)  {
	q.Drv(13000,13173,unsafe.Pointer(&blue),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QColor::setCmyk(int,int,int,int,int)
func (q *QColor) SetCmyk(c int,m int,y int,k int,a int)  {
	q.Drv(13000,13174,unsafe.Pointer(&c),unsafe.Pointer(&m),unsafe.Pointer(&y),unsafe.Pointer(&k),unsafe.Pointer(&a),nil,nil,nil,nil,nil,nil,nil)
}	
//QColor::setCmykF(double,double,double,double,double)
func (q *QColor) SetCmykF(c float64,m float64,y float64,k float64,a float64)  {
	q.Drv(13000,13175,unsafe.Pointer(&c),unsafe.Pointer(&m),unsafe.Pointer(&y),unsafe.Pointer(&k),unsafe.Pointer(&a),nil,nil,nil,nil,nil,nil,nil)
}	
//QColor::setGreen(int)
func (q *QColor) SetGreen(green int)  {
	q.Drv(13000,13176,unsafe.Pointer(&green),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QColor::setGreenF(double)
func (q *QColor) SetGreenF(green float64)  {
	q.Drv(13000,13177,unsafe.Pointer(&green),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QColor::setHsl(int,int,int,int)
func (q *QColor) SetHsl(h int,s int,l int,a int)  {
	q.Drv(13000,13178,unsafe.Pointer(&h),unsafe.Pointer(&s),unsafe.Pointer(&l),unsafe.Pointer(&a),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QColor::setHslF(double,double,double,double)
func (q *QColor) SetHslF(h float64,s float64,l float64,a float64)  {
	q.Drv(13000,13179,unsafe.Pointer(&h),unsafe.Pointer(&s),unsafe.Pointer(&l),unsafe.Pointer(&a),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QColor::setHsv(int,int,int,int)
func (q *QColor) SetHsv(h int,s int,v int,a int)  {
	q.Drv(13000,13180,unsafe.Pointer(&h),unsafe.Pointer(&s),unsafe.Pointer(&v),unsafe.Pointer(&a),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QColor::setHsvF(double,double,double,double)
func (q *QColor) SetHsvF(h float64,s float64,v float64,a float64)  {
	q.Drv(13000,13181,unsafe.Pointer(&h),unsafe.Pointer(&s),unsafe.Pointer(&v),unsafe.Pointer(&a),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QColor::setNamedColor(QString const&)
func (q *QColor) SetNamedColor(name string)  {
	q.Drv(13000,13182,unsafe.Pointer(&name),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QColor::setRed(int)
func (q *QColor) SetRed(red int)  {
	q.Drv(13000,13183,unsafe.Pointer(&red),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QColor::setRedF(double)
func (q *QColor) SetRedF(red float64)  {
	q.Drv(13000,13184,unsafe.Pointer(&red),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QColor::setRgb(unsigned int)
func (q *QColor) SetRgb(rgb uint)  {
	q.Drv(13000,13185,unsafe.Pointer(&rgb),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QColor::setRgb(int,int,int,int)
func (q *QColor) SetRgbWithRGBA(r int,g int,b int,a int)  {
	q.Drv(13000,13186,unsafe.Pointer(&r),unsafe.Pointer(&g),unsafe.Pointer(&b),unsafe.Pointer(&a),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QColor::setRgbF(double,double,double,double)
func (q *QColor) SetRgbF(r float64,g float64,b float64,a float64)  {
	q.Drv(13000,13187,unsafe.Pointer(&r),unsafe.Pointer(&g),unsafe.Pointer(&b),unsafe.Pointer(&a),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QColor::setRgba(unsigned int)
func (q *QColor) SetRgba(rgba uint)  {
	q.Drv(13000,13188,unsafe.Pointer(&rgba),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QColor::spec()
func (q *QColor) Spec() QColor_Spec {
	var __rv QColor_Spec
	q.Drv(13000,13189,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QColor::toCmyk()
func (q *QColor) ToCmyk() *QColor {
	var __rv uintptr
	q.Drv(13000,13190,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QColor{}
	_rp.SetDriver(__rv,13000,true)
	return _rp
}	
//QColor::toHsl()
func (q *QColor) ToHsl() *QColor {
	var __rv uintptr
	q.Drv(13000,13191,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QColor{}
	_rp.SetDriver(__rv,13000,true)
	return _rp
}	
//QColor::toHsv()
func (q *QColor) ToHsv() *QColor {
	var __rv uintptr
	q.Drv(13000,13192,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QColor{}
	_rp.SetDriver(__rv,13000,true)
	return _rp
}	
//QColor::toRgb()
func (q *QColor) ToRgb() *QColor {
	var __rv uintptr
	q.Drv(13000,13193,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QColor{}
	_rp.SetDriver(__rv,13000,true)
	return _rp
}	
//QColor::value()
func (q *QColor) Value() int {
	var __rv int
	q.Drv(13000,13194,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QColor::valueF()
func (q *QColor) ValueF() float64 {
	var __rv float64
	q.Drv(13000,13195,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QColor::yellow()
func (q *QColor) Yellow() int {
	var __rv int
	q.Drv(13000,13196,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QColor::yellowF()
func (q *QColor) YellowF() float64 {
	var __rv float64
	q.Drv(13000,13197,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
