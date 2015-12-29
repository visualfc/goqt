// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QTextCharFormat_VerticalAlignment - QTextCharFormat::VerticalAlignment
type QTextCharFormat_VerticalAlignment uint32
const (
	QTextCharFormat_AlignNormal QTextCharFormat_VerticalAlignment = 0
	QTextCharFormat_AlignSuperScript QTextCharFormat_VerticalAlignment = 0
	QTextCharFormat_AlignSubScript QTextCharFormat_VerticalAlignment = 1
	QTextCharFormat_AlignMiddle QTextCharFormat_VerticalAlignment = 2
	QTextCharFormat_AlignTop QTextCharFormat_VerticalAlignment = 3
	QTextCharFormat_AlignBottom QTextCharFormat_VerticalAlignment = 4
)
//enum QTextCharFormat_UnderlineStyle - QTextCharFormat::UnderlineStyle
type QTextCharFormat_UnderlineStyle uint32
const (
	QTextCharFormat_NoUnderline QTextCharFormat_UnderlineStyle = 0
	QTextCharFormat_SingleUnderline QTextCharFormat_UnderlineStyle = 1
	QTextCharFormat_DashUnderline QTextCharFormat_UnderlineStyle = 2
	QTextCharFormat_DotLine QTextCharFormat_UnderlineStyle = 3
	QTextCharFormat_DashDotLine QTextCharFormat_UnderlineStyle = 4
	QTextCharFormat_DashDotDotLine QTextCharFormat_UnderlineStyle = 5
	QTextCharFormat_WaveUnderline QTextCharFormat_UnderlineStyle = 6
	QTextCharFormat_SpellCheckUnderline QTextCharFormat_UnderlineStyle = 7
)
//struct QTextCharFormat : QTextCharFormat
type QTextCharFormat struct {
	QTextFormat
}
//QTextCharFormat::QTextCharFormat()	
func NewTextCharFormat() *QTextCharFormat {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),142000,142102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTextCharFormat{}
	_p.SetDriver(__rv,142000,true)
	return _p
} 
//QTextCharFormat::anchorHref()
func (q *QTextCharFormat) AnchorHref() string {
	var __rv string
	q.Drv(142000,142103,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextCharFormat::anchorName()
func (q *QTextCharFormat) AnchorName() string {
	var __rv string
	q.Drv(142000,142104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextCharFormat::anchorNames()
func (q *QTextCharFormat) AnchorNames() []string {
	var __rv []string
	q.Drv(142000,142105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextCharFormat::font()
func (q *QTextCharFormat) Font() *QFont {
	var __rv uintptr
	q.Drv(142000,142106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QFont{}
	_rp.SetDriver(__rv,37000,true)
	return _rp
}	
//QTextCharFormat::fontCapitalization()
func (q *QTextCharFormat) FontCapitalization() QFont_Capitalization {
	var __rv QFont_Capitalization
	q.Drv(142000,142107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextCharFormat::fontFamily()
func (q *QTextCharFormat) FontFamily() string {
	var __rv string
	q.Drv(142000,142108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextCharFormat::fontFixedPitch()
func (q *QTextCharFormat) FontFixedPitch() bool {
	var __rv bool
	q.Drv(142000,142109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextCharFormat::fontItalic()
func (q *QTextCharFormat) FontItalic() bool {
	var __rv bool
	q.Drv(142000,142110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextCharFormat::fontKerning()
func (q *QTextCharFormat) FontKerning() bool {
	var __rv bool
	q.Drv(142000,142111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextCharFormat::fontLetterSpacing()
func (q *QTextCharFormat) FontLetterSpacing() float64 {
	var __rv float64
	q.Drv(142000,142112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextCharFormat::fontOverline()
func (q *QTextCharFormat) FontOverline() bool {
	var __rv bool
	q.Drv(142000,142113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextCharFormat::fontPointSize()
func (q *QTextCharFormat) FontPointSize() float64 {
	var __rv float64
	q.Drv(142000,142114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextCharFormat::fontStrikeOut()
func (q *QTextCharFormat) FontStrikeOut() bool {
	var __rv bool
	q.Drv(142000,142115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextCharFormat::fontStyleHint()
func (q *QTextCharFormat) FontStyleHint() QFont_StyleHint {
	var __rv QFont_StyleHint
	q.Drv(142000,142116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextCharFormat::fontStyleStrategy()
func (q *QTextCharFormat) FontStyleStrategy() QFont_StyleStrategy {
	var __rv QFont_StyleStrategy
	q.Drv(142000,142117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextCharFormat::fontUnderline()
func (q *QTextCharFormat) FontUnderline() bool {
	var __rv bool
	q.Drv(142000,142118,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextCharFormat::fontWeight()
func (q *QTextCharFormat) FontWeight() int {
	var __rv int
	q.Drv(142000,142119,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextCharFormat::fontWordSpacing()
func (q *QTextCharFormat) FontWordSpacing() float64 {
	var __rv float64
	q.Drv(142000,142120,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextCharFormat::isAnchor()
func (q *QTextCharFormat) IsAnchor() bool {
	var __rv bool
	q.Drv(142000,142121,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextCharFormat::isValid()
func (q *QTextCharFormat) IsValid() bool {
	var __rv bool
	q.Drv(142000,142122,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextCharFormat::setAnchor(bool)
func (q *QTextCharFormat) SetAnchor(anchor bool)  {
	q.Drv(142000,142123,unsafe.Pointer(&anchor),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextCharFormat::setAnchorHref(QString const&)
func (q *QTextCharFormat) SetAnchorHref(value string)  {
	q.Drv(142000,142124,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextCharFormat::setAnchorName(QString const&)
func (q *QTextCharFormat) SetAnchorName(name string)  {
	q.Drv(142000,142125,unsafe.Pointer(&name),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextCharFormat::setAnchorNames(QStringList const&)
func (q *QTextCharFormat) SetAnchorNames(names []string)  {
	q.Drv(142000,142126,unsafe.Pointer(&names),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextCharFormat::setFont(QFont const&)
func (q *QTextCharFormat) SetFont(font *QFont)  {
	q.Drv(142000,142127,Native(font),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextCharFormat::setFontCapitalization(QFont::Capitalization)
func (q *QTextCharFormat) SetFontCapitalization(capitalization QFont_Capitalization)  {
	q.Drv(142000,142128,unsafe.Pointer(&capitalization),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextCharFormat::setFontFamily(QString const&)
func (q *QTextCharFormat) SetFontFamily(family string)  {
	q.Drv(142000,142129,unsafe.Pointer(&family),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextCharFormat::setFontFixedPitch(bool)
func (q *QTextCharFormat) SetFontFixedPitch(fixedPitch bool)  {
	q.Drv(142000,142130,unsafe.Pointer(&fixedPitch),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextCharFormat::setFontItalic(bool)
func (q *QTextCharFormat) SetFontItalic(italic bool)  {
	q.Drv(142000,142131,unsafe.Pointer(&italic),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextCharFormat::setFontKerning(bool)
func (q *QTextCharFormat) SetFontKerning(enable bool)  {
	q.Drv(142000,142132,unsafe.Pointer(&enable),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextCharFormat::setFontLetterSpacing(double)
func (q *QTextCharFormat) SetFontLetterSpacing(spacing float64)  {
	q.Drv(142000,142133,unsafe.Pointer(&spacing),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextCharFormat::setFontOverline(bool)
func (q *QTextCharFormat) SetFontOverline(overline bool)  {
	q.Drv(142000,142134,unsafe.Pointer(&overline),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextCharFormat::setFontPointSize(double)
func (q *QTextCharFormat) SetFontPointSize(size float64)  {
	q.Drv(142000,142135,unsafe.Pointer(&size),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextCharFormat::setFontStrikeOut(bool)
func (q *QTextCharFormat) SetFontStrikeOut(strikeOut bool)  {
	q.Drv(142000,142136,unsafe.Pointer(&strikeOut),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextCharFormat::setFontStyleHint(QFont::StyleHint)
func (q *QTextCharFormat) SetFontStyleHint(hint QFont_StyleHint)  {
	q.Drv(142000,142137,unsafe.Pointer(&hint),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextCharFormat::setFontStyleHint(QFont::StyleHint,QFont::StyleStrategy)
func (q *QTextCharFormat) SetFontStyleHintWithHintStrategy(hint QFont_StyleHint,strategy QFont_StyleStrategy)  {
	q.Drv(142000,142138,unsafe.Pointer(&hint),unsafe.Pointer(&strategy),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextCharFormat::setFontStyleStrategy(QFont::StyleStrategy)
func (q *QTextCharFormat) SetFontStyleStrategy(strategy QFont_StyleStrategy)  {
	q.Drv(142000,142139,unsafe.Pointer(&strategy),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextCharFormat::setFontUnderline(bool)
func (q *QTextCharFormat) SetFontUnderline(underline bool)  {
	q.Drv(142000,142140,unsafe.Pointer(&underline),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextCharFormat::setFontWeight(int)
func (q *QTextCharFormat) SetFontWeight(weight int)  {
	q.Drv(142000,142141,unsafe.Pointer(&weight),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextCharFormat::setFontWordSpacing(double)
func (q *QTextCharFormat) SetFontWordSpacing(spacing float64)  {
	q.Drv(142000,142142,unsafe.Pointer(&spacing),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextCharFormat::setTableCellColumnSpan(int)
func (q *QTextCharFormat) SetTableCellColumnSpan(tableCellColumnSpan int)  {
	q.Drv(142000,142143,unsafe.Pointer(&tableCellColumnSpan),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextCharFormat::setTableCellRowSpan(int)
func (q *QTextCharFormat) SetTableCellRowSpan(tableCellRowSpan int)  {
	q.Drv(142000,142144,unsafe.Pointer(&tableCellRowSpan),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextCharFormat::setTextOutline(QPen const&)
func (q *QTextCharFormat) SetTextOutline(pen *QPen)  {
	q.Drv(142000,142145,Native(pen),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextCharFormat::setToolTip(QString const&)
func (q *QTextCharFormat) SetToolTip(tip string)  {
	q.Drv(142000,142146,unsafe.Pointer(&tip),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextCharFormat::setUnderlineColor(QColor const&)
func (q *QTextCharFormat) SetUnderlineColor(color *QColor)  {
	q.Drv(142000,142147,Native(color),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextCharFormat::setUnderlineStyle(QTextCharFormat::UnderlineStyle)
func (q *QTextCharFormat) SetUnderlineStyle(style QTextCharFormat_UnderlineStyle)  {
	q.Drv(142000,142148,unsafe.Pointer(&style),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextCharFormat::setVerticalAlignment(QTextCharFormat::VerticalAlignment)
func (q *QTextCharFormat) SetVerticalAlignment(alignment QTextCharFormat_VerticalAlignment)  {
	q.Drv(142000,142149,unsafe.Pointer(&alignment),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextCharFormat::tableCellColumnSpan()
func (q *QTextCharFormat) TableCellColumnSpan() int {
	var __rv int
	q.Drv(142000,142150,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextCharFormat::tableCellRowSpan()
func (q *QTextCharFormat) TableCellRowSpan() int {
	var __rv int
	q.Drv(142000,142151,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextCharFormat::textOutline()
func (q *QTextCharFormat) TextOutline() *QPen {
	var __rv uintptr
	q.Drv(142000,142152,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPen{}
	_rp.SetDriver(__rv,92000,true)
	return _rp
}	
//QTextCharFormat::toolTip()
func (q *QTextCharFormat) ToolTip() string {
	var __rv string
	q.Drv(142000,142153,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextCharFormat::underlineColor()
func (q *QTextCharFormat) UnderlineColor() *QColor {
	var __rv uintptr
	q.Drv(142000,142154,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QColor{}
	_rp.SetDriver(__rv,13000,true)
	return _rp
}	
//QTextCharFormat::underlineStyle()
func (q *QTextCharFormat) UnderlineStyle() QTextCharFormat_UnderlineStyle {
	var __rv QTextCharFormat_UnderlineStyle
	q.Drv(142000,142155,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextCharFormat::verticalAlignment()
func (q *QTextCharFormat) VerticalAlignment() QTextCharFormat_VerticalAlignment {
	var __rv QTextCharFormat_VerticalAlignment
	q.Drv(142000,142156,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
