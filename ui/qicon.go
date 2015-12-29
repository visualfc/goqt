// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QIcon_Mode - QIcon::Mode
type QIcon_Mode uint32
const (
	QIcon_Normal QIcon_Mode = 0
	QIcon_Disabled QIcon_Mode = 1
	QIcon_Active QIcon_Mode = 2
	QIcon_Selected QIcon_Mode = 3
)
//enum QIcon_State - QIcon::State
type QIcon_State uint32
const (
	QIcon_On QIcon_State = 0
	QIcon_Off QIcon_State = 1
)
//struct QIcon : QIcon
type QIcon struct {
	BaseDrv
}
//QIcon::QIcon()	
func NewIcon() *QIcon {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),51000,51102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QIcon{}
	_p.SetDriver(__rv,51000,true)
	return _p
} 
//QIcon::QIcon(QIcon const&)	
func NewIconCopy(other *QIcon) *QIcon {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),51000,51103,Native(other),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QIcon{}
	_p.SetDriver(__rv,51000,true)
	return _p
} 
//QIcon::QIcon(QPixmap const&)	
func NewIconWithPixmap(pixmap *QPixmap) *QIcon {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),51000,51104,Native(pixmap),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QIcon{}
	_p.SetDriver(__rv,51000,true)
	return _p
} 
//QIcon::QIcon(QString const&)	
func NewIconWithFilename(fileName string) *QIcon {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),51000,51105,unsafe.Pointer(&fileName),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QIcon{}
	_p.SetDriver(__rv,51000,true)
	return _p
} 
//QIcon::actualSize(QSize const&)
func (q *QIcon) ActualSize(size *QSize) *QSize {
	var __rv uintptr
	q.Drv(51000,51106,Native(size),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QIcon::actualSize(QSize const&,QIcon::Mode,QIcon::State)
func (q *QIcon) ActualSizeWithSizeModeState(size *QSize,mode QIcon_Mode,state QIcon_State) *QSize {
	var __rv uintptr
	q.Drv(51000,51107,Native(size),unsafe.Pointer(&mode),unsafe.Pointer(&state),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QIcon::addFile(QString const&)
func (q *QIcon) AddFile(fileName string)  {
	q.Drv(51000,51108,unsafe.Pointer(&fileName),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QIcon::addFile(QString const&,QSize const&,QIcon::Mode,QIcon::State)
func (q *QIcon) AddFileWithFilenameSizeModeState(fileName string,size *QSize,mode QIcon_Mode,state QIcon_State)  {
	q.Drv(51000,51109,unsafe.Pointer(&fileName),Native(size),unsafe.Pointer(&mode),unsafe.Pointer(&state),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QIcon::addPixmap(QPixmap const&)
func (q *QIcon) AddPixmap(pixmap *QPixmap)  {
	q.Drv(51000,51110,Native(pixmap),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QIcon::addPixmap(QPixmap const&,QIcon::Mode,QIcon::State)
func (q *QIcon) AddPixmapWithPixmapModeState(pixmap *QPixmap,mode QIcon_Mode,state QIcon_State)  {
	q.Drv(51000,51111,Native(pixmap),unsafe.Pointer(&mode),unsafe.Pointer(&state),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QIcon::availableSizes()
func (q *QIcon) AvailableSizes() []*QSize {
	var __rv []*QSize
	q.Drv(51000,51112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QIcon::availableSizes(QIcon::Mode,QIcon::State)
func (q *QIcon) AvailableSizesWithModeState(mode QIcon_Mode,state QIcon_State) []*QSize {
	var __rv []*QSize
	q.Drv(51000,51113,unsafe.Pointer(&mode),unsafe.Pointer(&state),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QIcon::cacheKey()
func (q *QIcon) CacheKey() int64 {
	var __rv int64
	q.Drv(51000,51114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QIcon::detach()
func (q *QIcon) Detach()  {
	q.Drv(51000,51115,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QIcon::fromTheme(QString const&)	
func QIconFromTheme(name string) *QIcon {
	var __rv uintptr
	DirectQtDrv(nil,51000,51116,unsafe.Pointer(&name),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QIcon{}
	_rp.SetDriver(__rv,51000,true)
	return _rp
}	
//QIcon::fromTheme(QString const&)
func (q *QIcon) FromTheme(name string) *QIcon {
	var __rv uintptr
	q.Drv(51000,51116,unsafe.Pointer(&name),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QIcon{}
	_rp.SetDriver(__rv,51000,true)
	return _rp
}	
//QIcon::fromTheme(QString const&,QIcon const&)	
func QIconFromThemeWithNameFallback(name string,fallback *QIcon) *QIcon {
	var __rv uintptr
	DirectQtDrv(nil,51000,51117,unsafe.Pointer(&name),Native(fallback),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QIcon{}
	_rp.SetDriver(__rv,51000,true)
	return _rp
}	
//QIcon::fromTheme(QString const&,QIcon const&)
func (q *QIcon) FromThemeWithNameFallback(name string,fallback *QIcon) *QIcon {
	var __rv uintptr
	q.Drv(51000,51117,unsafe.Pointer(&name),Native(fallback),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QIcon{}
	_rp.SetDriver(__rv,51000,true)
	return _rp
}	
//QIcon::hasThemeIcon(QString const&)	
func QIconHasThemeIcon(name string) bool {
	var __rv bool
	DirectQtDrv(nil,51000,51118,unsafe.Pointer(&name),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QIcon::hasThemeIcon(QString const&)
func (q *QIcon) HasThemeIcon(name string) bool {
	var __rv bool
	q.Drv(51000,51118,unsafe.Pointer(&name),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QIcon::isDetached()
func (q *QIcon) IsDetached() bool {
	var __rv bool
	q.Drv(51000,51119,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QIcon::isNull()
func (q *QIcon) IsNull() bool {
	var __rv bool
	q.Drv(51000,51120,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QIcon::name()
func (q *QIcon) Name() string {
	var __rv string
	q.Drv(51000,51121,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QIcon::paint(QPainter*,QRect const&,QFlags<Qt::AlignmentFlag>,QIcon::Mode,QIcon::State)
func (q *QIcon) PaintWithPainterRectAlignmentModeState(painter *QPainter,rect *QRect,alignment Qt_AlignmentFlag,mode QIcon_Mode,state QIcon_State)  {
	q.Drv(51000,51122,Native(painter),Native(rect),unsafe.Pointer(&alignment),unsafe.Pointer(&mode),unsafe.Pointer(&state),nil,nil,nil,nil,nil,nil,nil)
}	
//QIcon::paint(QPainter*,int,int,int,int,QFlags<Qt::AlignmentFlag>,QIcon::Mode,QIcon::State)
func (q *QIcon) PaintWithPainterXYWidthHeightAlignmentModeState(painter *QPainter,x int,y int,w int,h int,alignment Qt_AlignmentFlag,mode QIcon_Mode,state QIcon_State)  {
	q.Drv(51000,51123,Native(painter),unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&w),unsafe.Pointer(&h),unsafe.Pointer(&alignment),unsafe.Pointer(&mode),unsafe.Pointer(&state),nil,nil,nil,nil)
}	
//QIcon::pixmap(QSize const&)
func (q *QIcon) Pixmap(size *QSize) *QPixmap {
	var __rv uintptr
	q.Drv(51000,51124,Native(size),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPixmap{}
	_rp.SetDriver(__rv,96000,true)
	return _rp
}	
//QIcon::pixmap(int)
func (q *QIcon) PixmapWithExtent(extent int) *QPixmap {
	var __rv uintptr
	q.Drv(51000,51125,unsafe.Pointer(&extent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPixmap{}
	_rp.SetDriver(__rv,96000,true)
	return _rp
}	
//QIcon::pixmap(QSize const&,QIcon::Mode,QIcon::State)
func (q *QIcon) PixmapWithSizeModeState(size *QSize,mode QIcon_Mode,state QIcon_State) *QPixmap {
	var __rv uintptr
	q.Drv(51000,51126,Native(size),unsafe.Pointer(&mode),unsafe.Pointer(&state),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPixmap{}
	_rp.SetDriver(__rv,96000,true)
	return _rp
}	
//QIcon::pixmap(int,QIcon::Mode,QIcon::State)
func (q *QIcon) PixmapWithExtentModeState(extent int,mode QIcon_Mode,state QIcon_State) *QPixmap {
	var __rv uintptr
	q.Drv(51000,51127,unsafe.Pointer(&extent),unsafe.Pointer(&mode),unsafe.Pointer(&state),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPixmap{}
	_rp.SetDriver(__rv,96000,true)
	return _rp
}	
//QIcon::pixmap(int,int,QIcon::Mode,QIcon::State)
func (q *QIcon) PixmapWithWidthHeightModeState(w int,h int,mode QIcon_Mode,state QIcon_State) *QPixmap {
	var __rv uintptr
	q.Drv(51000,51128,unsafe.Pointer(&w),unsafe.Pointer(&h),unsafe.Pointer(&mode),unsafe.Pointer(&state),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPixmap{}
	_rp.SetDriver(__rv,96000,true)
	return _rp
}	
//QIcon::setThemeName(QString const&)	
func QIconSetThemeName(path string)  {
	DirectQtDrv(nil,51000,51129,unsafe.Pointer(&path),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QIcon::setThemeName(QString const&)
func (q *QIcon) SetThemeName(path string)  {
	q.Drv(51000,51129,unsafe.Pointer(&path),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QIcon::setThemeSearchPaths(QStringList const&)	
func QIconSetThemeSearchPaths(searchpath []string)  {
	DirectQtDrv(nil,51000,51130,unsafe.Pointer(&searchpath),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QIcon::setThemeSearchPaths(QStringList const&)
func (q *QIcon) SetThemeSearchPaths(searchpath []string)  {
	q.Drv(51000,51130,unsafe.Pointer(&searchpath),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QIcon::themeName()	
func QIconThemeName() string {
	var __rv string
	DirectQtDrv(nil,51000,51131,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QIcon::themeName()
func (q *QIcon) ThemeName() string {
	var __rv string
	q.Drv(51000,51131,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QIcon::themeSearchPaths()	
func QIconThemeSearchPaths() []string {
	var __rv []string
	DirectQtDrv(nil,51000,51132,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QIcon::themeSearchPaths()
func (q *QIcon) ThemeSearchPaths() []string {
	var __rv []string
	q.Drv(51000,51132,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
