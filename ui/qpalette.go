// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QPalette_ColorGroup - QPalette::ColorGroup
type QPalette_ColorGroup uint32
const (
	QPalette_Active QPalette_ColorGroup = 0
	QPalette_Disabled QPalette_ColorGroup = 1
	QPalette_Inactive QPalette_ColorGroup = 2
	QPalette_NColorGroups QPalette_ColorGroup = 3
	QPalette_Current QPalette_ColorGroup = 4
	QPalette_All QPalette_ColorGroup = 5
	QPalette_Normal QPalette_ColorGroup = QPalette_Active
)
//enum QPalette_ColorRole - QPalette::ColorRole
type QPalette_ColorRole uint32
const (
	QPalette_WindowText QPalette_ColorRole = 0
	QPalette_Button QPalette_ColorRole = 1
	QPalette_Light QPalette_ColorRole = 2
	QPalette_Midlight QPalette_ColorRole = 3
	QPalette_Dark QPalette_ColorRole = 4
	QPalette_Mid QPalette_ColorRole = 5
	QPalette_Text QPalette_ColorRole = 6
	QPalette_BrightText QPalette_ColorRole = 7
	QPalette_ButtonText QPalette_ColorRole = 8
	QPalette_Base QPalette_ColorRole = 9
	QPalette_Window QPalette_ColorRole = 10
	QPalette_Shadow QPalette_ColorRole = 11
	QPalette_Highlight QPalette_ColorRole = 12
	QPalette_HighlightedText QPalette_ColorRole = 13
	QPalette_Link QPalette_ColorRole = 14
	QPalette_LinkVisited QPalette_ColorRole = 15
	QPalette_AlternateBase QPalette_ColorRole = 16
	QPalette_NoRole QPalette_ColorRole = 17
	QPalette_ToolTipBase QPalette_ColorRole = 18
	QPalette_ToolTipText QPalette_ColorRole = 19
	QPalette_NColorRoles QPalette_ColorRole = QPalette_ToolTipText+1
	QPalette_Foreground QPalette_ColorRole = QPalette_WindowText
	QPalette_Background QPalette_ColorRole = QPalette_Window
)
//struct QPalette : QPalette
type QPalette struct {
	BaseDrv
}
//QPalette::QPalette()	
func NewPalette() *QPalette {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),91000,91102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPalette{}
	_p.SetDriver(__rv,91000,true)
	return _p
} 
//QPalette::QPalette(QColor const&)	
func NewPaletteWithColor(button *QColor) *QPalette {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),91000,91103,Native(button),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPalette{}
	_p.SetDriver(__rv,91000,true)
	return _p
} 
//QPalette::QPalette(QPalette const&)	
func NewPaletteCopy(palette *QPalette) *QPalette {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),91000,91104,Native(palette),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPalette{}
	_p.SetDriver(__rv,91000,true)
	return _p
} 
//QPalette::QPalette(Qt::GlobalColor)	
func NewPaletteWithGlobalcolor(button Qt_GlobalColor) *QPalette {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),91000,91105,unsafe.Pointer(&button),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPalette{}
	_p.SetDriver(__rv,91000,true)
	return _p
} 
//QPalette::QPalette(QColor const&,QColor const&)	
func NewPaletteWithColorColor(button *QColor,window *QColor) *QPalette {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),91000,91106,Native(button),Native(window),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPalette{}
	_p.SetDriver(__rv,91000,true)
	return _p
} 
//QPalette::QPalette(QColor const&,QColor const&,QColor const&,QColor const&,QColor const&,QColor const&,QColor const&)	
func NewPaletteWithColorColorColorColorColorColorColor(windowText *QColor,window *QColor,light *QColor,dark *QColor,mid *QColor,text *QColor,base *QColor) *QPalette {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),91000,91107,Native(windowText),Native(window),Native(light),Native(dark),Native(mid),Native(text),Native(base),nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPalette{}
	_p.SetDriver(__rv,91000,true)
	return _p
} 
//QPalette::QPalette(QBrush const&,QBrush const&,QBrush const&,QBrush const&,QBrush const&,QBrush const&,QBrush const&,QBrush const&,QBrush const&)	
func NewPaletteWithWindowtextButtonLightDarkMidTextBright_textBaseWindow(windowText *QBrush,button *QBrush,light *QBrush,dark *QBrush,mid *QBrush,text *QBrush,bright_text *QBrush,base *QBrush,window *QBrush) *QPalette {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),91000,91108,Native(windowText),Native(button),Native(light),Native(dark),Native(mid),Native(text),Native(bright_text),Native(base),Native(window),nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPalette{}
	_p.SetDriver(__rv,91000,true)
	return _p
} 
//QPalette::alternateBase()
func (q *QPalette) AlternateBase() *QBrush {
	var __rv uintptr
	q.Drv(91000,91109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QBrush{}
	_rp.SetDriver(__rv,9000,true)
	return _rp
}	
//QPalette::background()
func (q *QPalette) Background() *QBrush {
	var __rv uintptr
	q.Drv(91000,91110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QBrush{}
	_rp.SetDriver(__rv,9000,true)
	return _rp
}	
//QPalette::base()
func (q *QPalette) Base() *QBrush {
	var __rv uintptr
	q.Drv(91000,91111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QBrush{}
	_rp.SetDriver(__rv,9000,true)
	return _rp
}	
//QPalette::brightText()
func (q *QPalette) BrightText() *QBrush {
	var __rv uintptr
	q.Drv(91000,91112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QBrush{}
	_rp.SetDriver(__rv,9000,true)
	return _rp
}	
//QPalette::brush(QPalette::ColorRole)
func (q *QPalette) Brush(cr QPalette_ColorRole) *QBrush {
	var __rv uintptr
	q.Drv(91000,91113,unsafe.Pointer(&cr),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QBrush{}
	_rp.SetDriver(__rv,9000,true)
	return _rp
}	
//QPalette::brush(QPalette::ColorGroup,QPalette::ColorRole)
func (q *QPalette) BrushWithCgCr(cg QPalette_ColorGroup,cr QPalette_ColorRole) *QBrush {
	var __rv uintptr
	q.Drv(91000,91114,unsafe.Pointer(&cg),unsafe.Pointer(&cr),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QBrush{}
	_rp.SetDriver(__rv,9000,true)
	return _rp
}	
//QPalette::button()
func (q *QPalette) Button() *QBrush {
	var __rv uintptr
	q.Drv(91000,91115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QBrush{}
	_rp.SetDriver(__rv,9000,true)
	return _rp
}	
//QPalette::buttonText()
func (q *QPalette) ButtonText() *QBrush {
	var __rv uintptr
	q.Drv(91000,91116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QBrush{}
	_rp.SetDriver(__rv,9000,true)
	return _rp
}	
//QPalette::cacheKey()
func (q *QPalette) CacheKey() int64 {
	var __rv int64
	q.Drv(91000,91117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPalette::color(QPalette::ColorRole)
func (q *QPalette) Color(cr QPalette_ColorRole) *QColor {
	var __rv uintptr
	q.Drv(91000,91118,unsafe.Pointer(&cr),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QColor{}
	_rp.SetDriver(__rv,13000,true)
	return _rp
}	
//QPalette::color(QPalette::ColorGroup,QPalette::ColorRole)
func (q *QPalette) ColorWithCgCr(cg QPalette_ColorGroup,cr QPalette_ColorRole) *QColor {
	var __rv uintptr
	q.Drv(91000,91119,unsafe.Pointer(&cg),unsafe.Pointer(&cr),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QColor{}
	_rp.SetDriver(__rv,13000,true)
	return _rp
}	
//QPalette::currentColorGroup()
func (q *QPalette) CurrentColorGroup() QPalette_ColorGroup {
	var __rv QPalette_ColorGroup
	q.Drv(91000,91120,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPalette::dark()
func (q *QPalette) Dark() *QBrush {
	var __rv uintptr
	q.Drv(91000,91121,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QBrush{}
	_rp.SetDriver(__rv,9000,true)
	return _rp
}	
//QPalette::foreground()
func (q *QPalette) Foreground() *QBrush {
	var __rv uintptr
	q.Drv(91000,91122,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QBrush{}
	_rp.SetDriver(__rv,9000,true)
	return _rp
}	
//QPalette::highlight()
func (q *QPalette) Highlight() *QBrush {
	var __rv uintptr
	q.Drv(91000,91123,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QBrush{}
	_rp.SetDriver(__rv,9000,true)
	return _rp
}	
//QPalette::highlightedText()
func (q *QPalette) HighlightedText() *QBrush {
	var __rv uintptr
	q.Drv(91000,91124,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QBrush{}
	_rp.SetDriver(__rv,9000,true)
	return _rp
}	
//QPalette::isBrushSet(QPalette::ColorGroup,QPalette::ColorRole)
func (q *QPalette) IsBrushSet(cg QPalette_ColorGroup,cr QPalette_ColorRole) bool {
	var __rv bool
	q.Drv(91000,91125,unsafe.Pointer(&cg),unsafe.Pointer(&cr),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPalette::isCopyOf(QPalette const&)
func (q *QPalette) IsCopyOf(p *QPalette) bool {
	var __rv bool
	q.Drv(91000,91126,Native(p),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPalette::isEqual(QPalette::ColorGroup,QPalette::ColorGroup)
func (q *QPalette) IsEqual(cr1 QPalette_ColorGroup,cr2 QPalette_ColorGroup) bool {
	var __rv bool
	q.Drv(91000,91127,unsafe.Pointer(&cr1),unsafe.Pointer(&cr2),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPalette::light()
func (q *QPalette) Light() *QBrush {
	var __rv uintptr
	q.Drv(91000,91128,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QBrush{}
	_rp.SetDriver(__rv,9000,true)
	return _rp
}	
//QPalette::link()
func (q *QPalette) Link() *QBrush {
	var __rv uintptr
	q.Drv(91000,91129,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QBrush{}
	_rp.SetDriver(__rv,9000,true)
	return _rp
}	
//QPalette::linkVisited()
func (q *QPalette) LinkVisited() *QBrush {
	var __rv uintptr
	q.Drv(91000,91130,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QBrush{}
	_rp.SetDriver(__rv,9000,true)
	return _rp
}	
//QPalette::mid()
func (q *QPalette) Mid() *QBrush {
	var __rv uintptr
	q.Drv(91000,91131,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QBrush{}
	_rp.SetDriver(__rv,9000,true)
	return _rp
}	
//QPalette::midlight()
func (q *QPalette) Midlight() *QBrush {
	var __rv uintptr
	q.Drv(91000,91132,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QBrush{}
	_rp.SetDriver(__rv,9000,true)
	return _rp
}	
//QPalette::resolve()
func (q *QPalette) Resolve() uint {
	var __rv uint
	q.Drv(91000,91133,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPalette::resolve(QPalette const&)
func (q *QPalette) ResolveWithPalette(value *QPalette) *QPalette {
	var __rv uintptr
	q.Drv(91000,91134,Native(value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPalette{}
	_rp.SetDriver(__rv,91000,true)
	return _rp
}	
//QPalette::resolve(unsigned int)
func (q *QPalette) ResolveWithMask(mask uint)  {
	q.Drv(91000,91135,unsafe.Pointer(&mask),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPalette::setBrush(QPalette::ColorRole,QBrush const&)
func (q *QPalette) SetBrushWithCrBrush(cr QPalette_ColorRole,brush *QBrush)  {
	q.Drv(91000,91136,unsafe.Pointer(&cr),Native(brush),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPalette::setBrush(QPalette::ColorGroup,QPalette::ColorRole,QBrush const&)
func (q *QPalette) SetBrushWithCgCrBrush(cg QPalette_ColorGroup,cr QPalette_ColorRole,brush *QBrush)  {
	q.Drv(91000,91137,unsafe.Pointer(&cg),unsafe.Pointer(&cr),Native(brush),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPalette::setColor(QPalette::ColorRole,QColor const&)
func (q *QPalette) SetColorWithCrColor(cr QPalette_ColorRole,color *QColor)  {
	q.Drv(91000,91138,unsafe.Pointer(&cr),Native(color),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPalette::setColor(QPalette::ColorGroup,QPalette::ColorRole,QColor const&)
func (q *QPalette) SetColorWithCgCrColor(cg QPalette_ColorGroup,cr QPalette_ColorRole,color *QColor)  {
	q.Drv(91000,91139,unsafe.Pointer(&cg),unsafe.Pointer(&cr),Native(color),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPalette::setColorGroup(QPalette::ColorGroup,QBrush const&,QBrush const&,QBrush const&,QBrush const&,QBrush const&,QBrush const&,QBrush const&,QBrush const&,QBrush const&)
func (q *QPalette) SetColorGroup(cr QPalette_ColorGroup,windowText *QBrush,button *QBrush,light *QBrush,dark *QBrush,mid *QBrush,text *QBrush,bright_text *QBrush,base *QBrush,window *QBrush)  {
	q.Drv(91000,91140,unsafe.Pointer(&cr),Native(windowText),Native(button),Native(light),Native(dark),Native(mid),Native(text),Native(bright_text),Native(base),Native(window),nil,nil)
}	
//QPalette::setCurrentColorGroup(QPalette::ColorGroup)
func (q *QPalette) SetCurrentColorGroup(cg QPalette_ColorGroup)  {
	q.Drv(91000,91141,unsafe.Pointer(&cg),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPalette::shadow()
func (q *QPalette) Shadow() *QBrush {
	var __rv uintptr
	q.Drv(91000,91142,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QBrush{}
	_rp.SetDriver(__rv,9000,true)
	return _rp
}	
//QPalette::text()
func (q *QPalette) Text() *QBrush {
	var __rv uintptr
	q.Drv(91000,91143,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QBrush{}
	_rp.SetDriver(__rv,9000,true)
	return _rp
}	
//QPalette::toolTipBase()
func (q *QPalette) ToolTipBase() *QBrush {
	var __rv uintptr
	q.Drv(91000,91144,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QBrush{}
	_rp.SetDriver(__rv,9000,true)
	return _rp
}	
//QPalette::toolTipText()
func (q *QPalette) ToolTipText() *QBrush {
	var __rv uintptr
	q.Drv(91000,91145,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QBrush{}
	_rp.SetDriver(__rv,9000,true)
	return _rp
}	
//QPalette::window()
func (q *QPalette) Window() *QBrush {
	var __rv uintptr
	q.Drv(91000,91146,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QBrush{}
	_rp.SetDriver(__rv,9000,true)
	return _rp
}	
//QPalette::windowText()
func (q *QPalette) WindowText() *QBrush {
	var __rv uintptr
	q.Drv(91000,91147,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QBrush{}
	_rp.SetDriver(__rv,9000,true)
	return _rp
}	
