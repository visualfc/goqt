// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QTextFormat_Property - QTextFormat::Property
type QTextFormat_Property uint32
const (
	QTextFormat_ObjectIndex QTextFormat_Property = 0x0
	QTextFormat_CssFloat QTextFormat_Property = 0x0800
	QTextFormat_LayoutDirection QTextFormat_Property = 0x0801
	QTextFormat_OutlinePen QTextFormat_Property = 0x810
	QTextFormat_BackgroundBrush QTextFormat_Property = 0x820
	QTextFormat_ForegroundBrush QTextFormat_Property = 0x821
	QTextFormat_BackgroundImageUrl QTextFormat_Property = 0x823
	QTextFormat_BlockAlignment QTextFormat_Property = 0x1010
	QTextFormat_BlockTopMargin QTextFormat_Property = 0x1030
	QTextFormat_BlockBottomMargin QTextFormat_Property = 0x1031
	QTextFormat_BlockLeftMargin QTextFormat_Property = 0x1032
	QTextFormat_BlockRightMargin QTextFormat_Property = 0x1033
	QTextFormat_TextIndent QTextFormat_Property = 0x1034
	QTextFormat_TabPositions QTextFormat_Property = 0x1035
	QTextFormat_BlockIndent QTextFormat_Property = 0x1040
	QTextFormat_BlockNonBreakableLines QTextFormat_Property = 0x1050
	QTextFormat_BlockTrailingHorizontalRulerWidth QTextFormat_Property = 0x1060
	QTextFormat_FirstFontProperty QTextFormat_Property = 0x1FE0
	QTextFormat_FontCapitalization QTextFormat_Property = QTextFormat_FirstFontProperty
	QTextFormat_FontLetterSpacing QTextFormat_Property = 0x1FE1
	QTextFormat_FontWordSpacing QTextFormat_Property = 0x1FE2
	QTextFormat_FontStyleHint QTextFormat_Property = 0x1FE3
	QTextFormat_FontStyleStrategy QTextFormat_Property = 0x1FE4
	QTextFormat_FontKerning QTextFormat_Property = 0x1FE5
	QTextFormat_FontFamily QTextFormat_Property = 0x2000
	QTextFormat_FontPointSize QTextFormat_Property = 0x2001
	QTextFormat_FontSizeAdjustment QTextFormat_Property = 0x2002
	QTextFormat_FontSizeIncrement QTextFormat_Property = QTextFormat_FontSizeAdjustment
	QTextFormat_FontWeight QTextFormat_Property = 0x2003
	QTextFormat_FontItalic QTextFormat_Property = 0x2004
	QTextFormat_FontUnderline QTextFormat_Property = 0x2005
	QTextFormat_FontOverline QTextFormat_Property = 0x2006
	QTextFormat_FontStrikeOut QTextFormat_Property = 0x2007
	QTextFormat_FontFixedPitch QTextFormat_Property = 0x2008
	QTextFormat_FontPixelSize QTextFormat_Property = 0x2009
	QTextFormat_LastFontProperty QTextFormat_Property = QTextFormat_FontPixelSize
	QTextFormat_TextUnderlineColor QTextFormat_Property = 0x2010
	QTextFormat_TextVerticalAlignment QTextFormat_Property = 0x2021
	QTextFormat_TextOutline QTextFormat_Property = 0x2022
	QTextFormat_TextUnderlineStyle QTextFormat_Property = 0x2023
	QTextFormat_TextToolTip QTextFormat_Property = 0x2024
	QTextFormat_IsAnchor QTextFormat_Property = 0x2030
	QTextFormat_AnchorHref QTextFormat_Property = 0x2031
	QTextFormat_AnchorName QTextFormat_Property = 0x2032
	QTextFormat_ObjectType QTextFormat_Property = 0x2f00
	QTextFormat_ListStyle QTextFormat_Property = 0x3000
	QTextFormat_ListIndent QTextFormat_Property = 0x3001
	QTextFormat_FrameBorder QTextFormat_Property = 0x4000
	QTextFormat_FrameMargin QTextFormat_Property = 0x4001
	QTextFormat_FramePadding QTextFormat_Property = 0x4002
	QTextFormat_FrameWidth QTextFormat_Property = 0x4003
	QTextFormat_FrameHeight QTextFormat_Property = 0x4004
	QTextFormat_FrameTopMargin QTextFormat_Property = 0x4005
	QTextFormat_FrameBottomMargin QTextFormat_Property = 0x4006
	QTextFormat_FrameLeftMargin QTextFormat_Property = 0x4007
	QTextFormat_FrameRightMargin QTextFormat_Property = 0x4008
	QTextFormat_FrameBorderBrush QTextFormat_Property = 0x4009
	QTextFormat_FrameBorderStyle QTextFormat_Property = 0x4010
	QTextFormat_TableColumns QTextFormat_Property = 0x4100
	QTextFormat_TableColumnWidthConstraints QTextFormat_Property = 0x4101
	QTextFormat_TableCellSpacing QTextFormat_Property = 0x4102
	QTextFormat_TableCellPadding QTextFormat_Property = 0x4103
	QTextFormat_TableHeaderRowCount QTextFormat_Property = 0x4104
	QTextFormat_TableCellRowSpan QTextFormat_Property = 0x4810
	QTextFormat_TableCellColumnSpan QTextFormat_Property = 0x4811
	QTextFormat_TableCellTopPadding QTextFormat_Property = 0x4812
	QTextFormat_TableCellBottomPadding QTextFormat_Property = 0x4813
	QTextFormat_TableCellLeftPadding QTextFormat_Property = 0x4814
	QTextFormat_TableCellRightPadding QTextFormat_Property = 0x4815
	QTextFormat_ImageName QTextFormat_Property = 0x5000
	QTextFormat_ImageWidth QTextFormat_Property = 0x5010
	QTextFormat_ImageHeight QTextFormat_Property = 0x5011
	QTextFormat_FullWidthSelection QTextFormat_Property = 0x06000
	QTextFormat_PageBreakPolicy QTextFormat_Property = 0x7000
	QTextFormat_UserProperty QTextFormat_Property = 0x100000
)
//enum QTextFormat_FormatType - QTextFormat::FormatType
type QTextFormat_FormatType int32
const (
	QTextFormat_InvalidFormat QTextFormat_FormatType = -1
	QTextFormat_BlockFormat QTextFormat_FormatType = 1
	QTextFormat_CharFormat QTextFormat_FormatType = 2
	QTextFormat_ListFormat QTextFormat_FormatType = 3
	QTextFormat_TableFormat QTextFormat_FormatType = 4
	QTextFormat_FrameFormat QTextFormat_FormatType = 5
	QTextFormat_UserFormat QTextFormat_FormatType = 100
)
//enum QTextFormat_ObjectTypes - QTextFormat::ObjectTypes
type QTextFormat_ObjectTypes uint32
const (
	QTextFormat_NoObject QTextFormat_ObjectTypes = 0
	QTextFormat_ImageObject QTextFormat_ObjectTypes = 1
	QTextFormat_TableObject QTextFormat_ObjectTypes = 2
	QTextFormat_TableCellObject QTextFormat_ObjectTypes = 3
	QTextFormat_UserObject QTextFormat_ObjectTypes = 0x1000
)
//enum QTextFormat_PageBreakFlag - QTextFormat::PageBreakFlag
type QTextFormat_PageBreakFlag uint32
const (
	QTextFormat_PageBreak_Auto QTextFormat_PageBreakFlag = 0
	QTextFormat_PageBreak_AlwaysBefore QTextFormat_PageBreakFlag = 0x001
	QTextFormat_PageBreak_AlwaysAfter QTextFormat_PageBreakFlag = 0x010
)
//struct QTextFormat : QTextFormat
type QTextFormat struct {
	BaseDrv
}
//QTextFormat::QTextFormat()	
func NewTextFormat() *QTextFormat {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),151000,151102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTextFormat{}
	_p.SetDriver(__rv,151000,true)
	return _p
} 
//QTextFormat::QTextFormat(QTextFormat const&)	
func NewTextFormatCopy(rhs *QTextFormat) *QTextFormat {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),151000,151103,Native(rhs),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTextFormat{}
	_p.SetDriver(__rv,151000,true)
	return _p
} 
//QTextFormat::QTextFormat(int)	
func NewTextFormatWithType(_type int) *QTextFormat {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),151000,151104,unsafe.Pointer(&_type),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTextFormat{}
	_p.SetDriver(__rv,151000,true)
	return _p
} 
//QTextFormat::background()
func (q *QTextFormat) Background() *QBrush {
	var __rv uintptr
	q.Drv(151000,151105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QBrush{}
	_rp.SetDriver(__rv,9000,true)
	return _rp
}	
//QTextFormat::boolProperty(int)
func (q *QTextFormat) BoolProperty(propertyId int) bool {
	var __rv bool
	q.Drv(151000,151106,unsafe.Pointer(&propertyId),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextFormat::brushProperty(int)
func (q *QTextFormat) BrushProperty(propertyId int) *QBrush {
	var __rv uintptr
	q.Drv(151000,151107,unsafe.Pointer(&propertyId),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QBrush{}
	_rp.SetDriver(__rv,9000,true)
	return _rp
}	
//QTextFormat::clearBackground()
func (q *QTextFormat) ClearBackground()  {
	q.Drv(151000,151108,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextFormat::clearForeground()
func (q *QTextFormat) ClearForeground()  {
	q.Drv(151000,151109,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextFormat::clearProperty(int)
func (q *QTextFormat) ClearProperty(propertyId int)  {
	q.Drv(151000,151110,unsafe.Pointer(&propertyId),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextFormat::colorProperty(int)
func (q *QTextFormat) ColorProperty(propertyId int) *QColor {
	var __rv uintptr
	q.Drv(151000,151111,unsafe.Pointer(&propertyId),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QColor{}
	_rp.SetDriver(__rv,13000,true)
	return _rp
}	
//QTextFormat::doubleProperty(int)
func (q *QTextFormat) DoubleProperty(propertyId int) float64 {
	var __rv float64
	q.Drv(151000,151112,unsafe.Pointer(&propertyId),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextFormat::foreground()
func (q *QTextFormat) Foreground() *QBrush {
	var __rv uintptr
	q.Drv(151000,151113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QBrush{}
	_rp.SetDriver(__rv,9000,true)
	return _rp
}	
//QTextFormat::hasProperty(int)
func (q *QTextFormat) HasProperty(propertyId int) bool {
	var __rv bool
	q.Drv(151000,151114,unsafe.Pointer(&propertyId),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextFormat::intProperty(int)
func (q *QTextFormat) IntProperty(propertyId int) int {
	var __rv int
	q.Drv(151000,151115,unsafe.Pointer(&propertyId),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextFormat::isBlockFormat()
func (q *QTextFormat) IsBlockFormat() bool {
	var __rv bool
	q.Drv(151000,151116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextFormat::isCharFormat()
func (q *QTextFormat) IsCharFormat() bool {
	var __rv bool
	q.Drv(151000,151117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextFormat::isFrameFormat()
func (q *QTextFormat) IsFrameFormat() bool {
	var __rv bool
	q.Drv(151000,151118,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextFormat::isImageFormat()
func (q *QTextFormat) IsImageFormat() bool {
	var __rv bool
	q.Drv(151000,151119,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextFormat::isListFormat()
func (q *QTextFormat) IsListFormat() bool {
	var __rv bool
	q.Drv(151000,151120,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextFormat::isTableCellFormat()
func (q *QTextFormat) IsTableCellFormat() bool {
	var __rv bool
	q.Drv(151000,151121,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextFormat::isTableFormat()
func (q *QTextFormat) IsTableFormat() bool {
	var __rv bool
	q.Drv(151000,151122,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextFormat::isValid()
func (q *QTextFormat) IsValid() bool {
	var __rv bool
	q.Drv(151000,151123,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextFormat::layoutDirection()
func (q *QTextFormat) LayoutDirection() Qt_LayoutDirection {
	var __rv Qt_LayoutDirection
	q.Drv(151000,151124,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextFormat::lengthProperty(int)
func (q *QTextFormat) LengthProperty(propertyId int) *QTextLength {
	var __rv uintptr
	q.Drv(151000,151125,unsafe.Pointer(&propertyId),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextLength{}
	_rp.SetDriver(__rv,160000,true)
	return _rp
}	
//QTextFormat::lengthVectorProperty(int)
func (q *QTextFormat) LengthVectorProperty(propertyId int) []*QTextLength {
	var __rv []*QTextLength
	q.Drv(151000,151126,unsafe.Pointer(&propertyId),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextFormat::merge(QTextFormat const&)
func (q *QTextFormat) Merge(other *QTextFormat)  {
	q.Drv(151000,151127,Native(other),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextFormat::objectIndex()
func (q *QTextFormat) ObjectIndex() int {
	var __rv int
	q.Drv(151000,151128,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextFormat::objectType()
func (q *QTextFormat) ObjectType() int {
	var __rv int
	q.Drv(151000,151129,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextFormat::penProperty(int)
func (q *QTextFormat) PenProperty(propertyId int) *QPen {
	var __rv uintptr
	q.Drv(151000,151130,unsafe.Pointer(&propertyId),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPen{}
	_rp.SetDriver(__rv,92000,true)
	return _rp
}	
//QTextFormat::properties()
func (q *QTextFormat) Properties() map[int]*QVariant {
	__rv := make(map[int]*QVariant)
	q.Drv(151000,151131,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextFormat::property(int)
func (q *QTextFormat) Property(propertyId int) *QVariant {
	var __rv uintptr
	q.Drv(151000,151132,unsafe.Pointer(&propertyId),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVariant{}
	_rp.SetDriver(__rv,182000,true)
	return _rp
}	
//QTextFormat::propertyCount()
func (q *QTextFormat) PropertyCount() int {
	var __rv int
	q.Drv(151000,151133,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextFormat::setBackground(QBrush const&)
func (q *QTextFormat) SetBackground(brush *QBrush)  {
	q.Drv(151000,151134,Native(brush),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextFormat::setForeground(QBrush const&)
func (q *QTextFormat) SetForeground(brush *QBrush)  {
	q.Drv(151000,151135,Native(brush),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextFormat::setLayoutDirection(Qt::LayoutDirection)
func (q *QTextFormat) SetLayoutDirection(direction Qt_LayoutDirection)  {
	q.Drv(151000,151136,unsafe.Pointer(&direction),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextFormat::setObjectIndex(int)
func (q *QTextFormat) SetObjectIndex(object int)  {
	q.Drv(151000,151137,unsafe.Pointer(&object),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextFormat::setObjectType(int)
func (q *QTextFormat) SetObjectType(_type int)  {
	q.Drv(151000,151138,unsafe.Pointer(&_type),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextFormat::setProperty(int,QVariant const&)
func (q *QTextFormat) SetPropertyWithPropertyidValue(propertyId int,value *QVariant)  {
	q.Drv(151000,151139,unsafe.Pointer(&propertyId),Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextFormat::setProperty(int,QVector<QTextLength> const&)
func (q *QTextFormat) SetPropertyWithPropertyidLengths(propertyId int,lengths []*QTextLength)  {
	q.Drv(151000,151140,unsafe.Pointer(&propertyId),unsafe.Pointer(&lengths),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextFormat::stringProperty(int)
func (q *QTextFormat) StringProperty(propertyId int) string {
	var __rv string
	q.Drv(151000,151141,unsafe.Pointer(&propertyId),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextFormat::toBlockFormat()
func (q *QTextFormat) ToBlockFormat() *QTextBlockFormat {
	var __rv uintptr
	q.Drv(151000,151142,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextBlockFormat{}
	_rp.SetDriver(__rv,139000,true)
	return _rp
}	
//QTextFormat::toCharFormat()
func (q *QTextFormat) ToCharFormat() *QTextCharFormat {
	var __rv uintptr
	q.Drv(151000,151143,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextCharFormat{}
	_rp.SetDriver(__rv,142000,true)
	return _rp
}	
//QTextFormat::toFrameFormat()
func (q *QTextFormat) ToFrameFormat() *QTextFrameFormat {
	var __rv uintptr
	q.Drv(151000,151144,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextFrameFormat{}
	_rp.SetDriver(__rv,154000,true)
	return _rp
}	
//QTextFormat::toImageFormat()
func (q *QTextFormat) ToImageFormat() *QTextImageFormat {
	var __rv uintptr
	q.Drv(151000,151145,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextImageFormat{}
	_rp.SetDriver(__rv,155000,true)
	return _rp
}	
//QTextFormat::toListFormat()
func (q *QTextFormat) ToListFormat() *QTextListFormat {
	var __rv uintptr
	q.Drv(151000,151146,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextListFormat{}
	_rp.SetDriver(__rv,162000,true)
	return _rp
}	
//QTextFormat::toTableCellFormat()
func (q *QTextFormat) ToTableCellFormat() *QTextTableCellFormat {
	var __rv uintptr
	q.Drv(151000,151147,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextTableCellFormat{}
	_rp.SetDriver(__rv,167000,true)
	return _rp
}	
//QTextFormat::toTableFormat()
func (q *QTextFormat) ToTableFormat() *QTextTableFormat {
	var __rv uintptr
	q.Drv(151000,151148,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextTableFormat{}
	_rp.SetDriver(__rv,168000,true)
	return _rp
}	
//QTextFormat::type()
func (q *QTextFormat) Type() int {
	var __rv int
	q.Drv(151000,151149,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
