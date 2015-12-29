// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QFont_StyleStrategy - QFont::StyleStrategy
type QFont_StyleStrategy uint32
const (
	QFont_PreferDefault QFont_StyleStrategy = 0x0001
	QFont_PreferBitmap QFont_StyleStrategy = 0x0002
	QFont_PreferDevice QFont_StyleStrategy = 0x0004
	QFont_PreferOutline QFont_StyleStrategy = 0x0008
	QFont_ForceOutline QFont_StyleStrategy = 0x0010
	QFont_PreferMatch QFont_StyleStrategy = 0x0020
	QFont_PreferQuality QFont_StyleStrategy = 0x0040
	QFont_PreferAntialias QFont_StyleStrategy = 0x0080
	QFont_NoAntialias QFont_StyleStrategy = 0x0100
	QFont_OpenGLCompatible QFont_StyleStrategy = 0x0200
	QFont_ForceIntegerMetrics QFont_StyleStrategy = 0x0400
	QFont_NoFontMerging QFont_StyleStrategy = 0x8000
)
//enum QFont_SpacingType - QFont::SpacingType
type QFont_SpacingType uint32
const (
	QFont_PercentageSpacing QFont_SpacingType = 0
	QFont_AbsoluteSpacing QFont_SpacingType = 1
)
//enum QFont_ResolveProperties - QFont::ResolveProperties
type QFont_ResolveProperties uint32
const (
	QFont_FamilyResolved QFont_ResolveProperties = 0x0001
	QFont_SizeResolved QFont_ResolveProperties = 0x0002
	QFont_StyleHintResolved QFont_ResolveProperties = 0x0004
	QFont_StyleStrategyResolved QFont_ResolveProperties = 0x0008
	QFont_WeightResolved QFont_ResolveProperties = 0x0010
	QFont_StyleResolved QFont_ResolveProperties = 0x0020
	QFont_UnderlineResolved QFont_ResolveProperties = 0x0040
	QFont_OverlineResolved QFont_ResolveProperties = 0x0080
	QFont_StrikeOutResolved QFont_ResolveProperties = 0x0100
	QFont_FixedPitchResolved QFont_ResolveProperties = 0x0200
	QFont_StretchResolved QFont_ResolveProperties = 0x0400
	QFont_KerningResolved QFont_ResolveProperties = 0x0800
	QFont_CapitalizationResolved QFont_ResolveProperties = 0x1000
	QFont_LetterSpacingResolved QFont_ResolveProperties = 0x2000
	QFont_WordSpacingResolved QFont_ResolveProperties = 0x4000
	QFont_AllPropertiesResolved QFont_ResolveProperties = 0x7fff
)
//enum QFont_StyleHint - QFont::StyleHint
type QFont_StyleHint uint32
const (
	QFont_Helvetica QFont_StyleHint = 0
	QFont_SansSerif QFont_StyleHint = QFont_Helvetica
	QFont_Times QFont_StyleHint = QFont_Helvetica+1
	QFont_Serif QFont_StyleHint = QFont_Times
	QFont_Courier QFont_StyleHint = QFont_Times+1
	QFont_TypeWriter QFont_StyleHint = QFont_Courier
	QFont_OldEnglish QFont_StyleHint = QFont_Courier+1
	QFont_Decorative QFont_StyleHint = QFont_OldEnglish
	QFont_System QFont_StyleHint = QFont_OldEnglish+1
	QFont_AnyStyle QFont_StyleHint = QFont_OldEnglish+1+1
	QFont_Cursive QFont_StyleHint = QFont_OldEnglish+1+1+1
	QFont_Monospace QFont_StyleHint = QFont_OldEnglish+1+1+1+1
	QFont_Fantasy QFont_StyleHint = QFont_OldEnglish+1+1+1+1+1
)
//enum QFont_Weight - QFont::Weight
type QFont_Weight uint32
const (
	QFont_Light QFont_Weight = 25
	QFont_Normal QFont_Weight = 50
	QFont_DemiBold QFont_Weight = 63
	QFont_Bold QFont_Weight = 75
	QFont_Black QFont_Weight = 87
)
//enum QFont_Capitalization - QFont::Capitalization
type QFont_Capitalization uint32
const (
	QFont_MixedCase QFont_Capitalization = 0
	QFont_AllUppercase QFont_Capitalization = 1
	QFont_AllLowercase QFont_Capitalization = 2
	QFont_SmallCaps QFont_Capitalization = 3
	QFont_Capitalize QFont_Capitalization = 4
)
//enum QFont_Stretch - QFont::Stretch
type QFont_Stretch uint32
const (
	QFont_UltraCondensed QFont_Stretch = 50
	QFont_ExtraCondensed QFont_Stretch = 62
	QFont_Condensed QFont_Stretch = 75
	QFont_SemiCondensed QFont_Stretch = 87
	QFont_Unstretched QFont_Stretch = 100
	QFont_SemiExpanded QFont_Stretch = 112
	QFont_Expanded QFont_Stretch = 125
	QFont_ExtraExpanded QFont_Stretch = 150
	QFont_UltraExpanded QFont_Stretch = 200
)
//enum QFont_Style - QFont::Style
type QFont_Style uint32
const (
	QFont_StyleNormal QFont_Style = 0
	QFont_StyleItalic QFont_Style = 1
	QFont_StyleOblique QFont_Style = 2
)
//struct QFont : QFont
type QFont struct {
	BaseDrv
}
//QFont::QFont()	
func NewFont() *QFont {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),37000,37102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QFont{}
	_p.SetDriver(__rv,37000,true)
	return _p
} 
//QFont::QFont(QFont const&)	
func NewFontCopy(value *QFont) *QFont {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),37000,37103,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QFont{}
	_p.SetDriver(__rv,37000,true)
	return _p
} 
//QFont::QFont(QFont const&,QPaintDevice*)	
func NewFontWithFontPaintDevice(value2 *QFont,pd QPaintDeviceInterface) *QFont {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),37000,37104,Native(value2),unsafe.Pointer(new_pd_head(pd)),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QFont{}
	_p.SetDriver(__rv,37000,true)
	return _p
} 
//QFont::QFont(QString const&,int,int,bool)	
func NewFontWithFamilyPointsizeWeightItalic(family string,pointSize int,weight int,italic bool) *QFont {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),37000,37105,unsafe.Pointer(&family),unsafe.Pointer(&pointSize),unsafe.Pointer(&weight),unsafe.Pointer(&italic),nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QFont{}
	_p.SetDriver(__rv,37000,true)
	return _p
} 
//QFont::bold()
func (q *QFont) Bold() bool {
	var __rv bool
	q.Drv(37000,37106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFont::cacheStatistics()	
func QFontCacheStatistics()  {
	DirectQtDrv(nil,37000,37107,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFont::cacheStatistics()
func (q *QFont) CacheStatistics()  {
	q.Drv(37000,37107,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFont::capitalization()
func (q *QFont) Capitalization() QFont_Capitalization {
	var __rv QFont_Capitalization
	q.Drv(37000,37108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFont::cleanup()	
func QFontCleanup()  {
	DirectQtDrv(nil,37000,37109,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFont::cleanup()
func (q *QFont) Cleanup()  {
	q.Drv(37000,37109,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFont::defaultFamily()
func (q *QFont) DefaultFamily() string {
	var __rv string
	q.Drv(37000,37110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFont::exactMatch()
func (q *QFont) ExactMatch() bool {
	var __rv bool
	q.Drv(37000,37111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFont::family()
func (q *QFont) Family() string {
	var __rv string
	q.Drv(37000,37112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFont::fixedPitch()
func (q *QFont) FixedPitch() bool {
	var __rv bool
	q.Drv(37000,37113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFont::fromString(QString const&)
func (q *QFont) FromString(value string) bool {
	var __rv bool
	q.Drv(37000,37114,unsafe.Pointer(&value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFont::initialize()	
func QFontInitialize()  {
	DirectQtDrv(nil,37000,37115,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFont::initialize()
func (q *QFont) Initialize()  {
	q.Drv(37000,37115,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFont::insertSubstitution(QString const&,QString const&)	
func QFontInsertSubstitution(value2 string,value3 string)  {
	DirectQtDrv(nil,37000,37116,unsafe.Pointer(&value2),unsafe.Pointer(&value3),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFont::insertSubstitution(QString const&,QString const&)
func (q *QFont) InsertSubstitution(value2 string,value3 string)  {
	q.Drv(37000,37116,unsafe.Pointer(&value2),unsafe.Pointer(&value3),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFont::insertSubstitutions(QString const&,QStringList const&)	
func QFontInsertSubstitutions(value2 string,value3 []string)  {
	DirectQtDrv(nil,37000,37117,unsafe.Pointer(&value2),unsafe.Pointer(&value3),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFont::insertSubstitutions(QString const&,QStringList const&)
func (q *QFont) InsertSubstitutions(value2 string,value3 []string)  {
	q.Drv(37000,37117,unsafe.Pointer(&value2),unsafe.Pointer(&value3),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFont::isCopyOf(QFont const&)
func (q *QFont) IsCopyOf(value *QFont) bool {
	var __rv bool
	q.Drv(37000,37118,Native(value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFont::italic()
func (q *QFont) Italic() bool {
	var __rv bool
	q.Drv(37000,37119,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFont::kerning()
func (q *QFont) Kerning() bool {
	var __rv bool
	q.Drv(37000,37120,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFont::key()
func (q *QFont) Key() string {
	var __rv string
	q.Drv(37000,37121,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFont::lastResortFamily()
func (q *QFont) LastResortFamily() string {
	var __rv string
	q.Drv(37000,37122,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFont::lastResortFont()
func (q *QFont) LastResortFont() string {
	var __rv string
	q.Drv(37000,37123,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFont::letterSpacing()
func (q *QFont) LetterSpacing() float64 {
	var __rv float64
	q.Drv(37000,37124,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFont::letterSpacingType()
func (q *QFont) LetterSpacingType() QFont_SpacingType {
	var __rv QFont_SpacingType
	q.Drv(37000,37125,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFont::overline()
func (q *QFont) Overline() bool {
	var __rv bool
	q.Drv(37000,37126,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFont::pixelSize()
func (q *QFont) PixelSize() int {
	var __rv int
	q.Drv(37000,37127,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFont::pointSize()
func (q *QFont) PointSize() int {
	var __rv int
	q.Drv(37000,37128,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFont::pointSizeF()
func (q *QFont) PointSizeF() float64 {
	var __rv float64
	q.Drv(37000,37129,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFont::rawMode()
func (q *QFont) RawMode() bool {
	var __rv bool
	q.Drv(37000,37130,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFont::rawName()
func (q *QFont) RawName() string {
	var __rv string
	q.Drv(37000,37131,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFont::removeSubstitution(QString const&)	
func QFontRemoveSubstitutions(value string)  {
	DirectQtDrv(nil,37000,37132,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFont::removeSubstitution(QString const&)
func (q *QFont) RemoveSubstitutions(value string)  {
	q.Drv(37000,37132,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFont::resolve()
func (q *QFont) Resolve() uint {
	var __rv uint
	q.Drv(37000,37133,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFont::resolve(QFont const&)
func (q *QFont) ResolveWithFont(value *QFont) *QFont {
	var __rv uintptr
	q.Drv(37000,37134,Native(value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QFont{}
	_rp.SetDriver(__rv,37000,true)
	return _rp
}	
//QFont::resolve(unsigned int)
func (q *QFont) ResolveWithMask(mask uint)  {
	q.Drv(37000,37135,unsafe.Pointer(&mask),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFont::setBold(bool)
func (q *QFont) SetBold(value bool)  {
	q.Drv(37000,37136,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFont::setCapitalization(QFont::Capitalization)
func (q *QFont) SetCapitalization(value QFont_Capitalization)  {
	q.Drv(37000,37137,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFont::setFamily(QString const&)
func (q *QFont) SetFamily(value string)  {
	q.Drv(37000,37138,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFont::setFixedPitch(bool)
func (q *QFont) SetFixedPitch(value bool)  {
	q.Drv(37000,37139,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFont::setItalic(bool)
func (q *QFont) SetItalic(b bool)  {
	q.Drv(37000,37140,unsafe.Pointer(&b),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFont::setKerning(bool)
func (q *QFont) SetKerning(value bool)  {
	q.Drv(37000,37141,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFont::setLetterSpacing(QFont::SpacingType,double)
func (q *QFont) SetLetterSpacing(_type QFont_SpacingType,spacing float64)  {
	q.Drv(37000,37142,unsafe.Pointer(&_type),unsafe.Pointer(&spacing),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFont::setOverline(bool)
func (q *QFont) SetOverline(value bool)  {
	q.Drv(37000,37143,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFont::setPixelSize(int)
func (q *QFont) SetPixelSize(value int)  {
	q.Drv(37000,37144,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFont::setPointSize(int)
func (q *QFont) SetPointSize(value int)  {
	q.Drv(37000,37145,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFont::setPointSizeF(double)
func (q *QFont) SetPointSizeF(value float64)  {
	q.Drv(37000,37146,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFont::setRawMode(bool)
func (q *QFont) SetRawMode(value bool)  {
	q.Drv(37000,37147,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFont::setRawName(QString const&)
func (q *QFont) SetRawName(value string)  {
	q.Drv(37000,37148,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFont::setStretch(int)
func (q *QFont) SetStretch(value int)  {
	q.Drv(37000,37149,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFont::setStrikeOut(bool)
func (q *QFont) SetStrikeOut(value bool)  {
	q.Drv(37000,37150,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFont::setStyle(QFont::Style)
func (q *QFont) SetStyle(style QFont_Style)  {
	q.Drv(37000,37151,unsafe.Pointer(&style),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFont::setStyleHint(QFont::StyleHint)
func (q *QFont) SetStyleHint(value QFont_StyleHint)  {
	q.Drv(37000,37152,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFont::setStyleHint(QFont::StyleHint,QFont::StyleStrategy)
func (q *QFont) SetStyleHintWithStylehintStylestrategy(value2 QFont_StyleHint,value3 QFont_StyleStrategy)  {
	q.Drv(37000,37153,unsafe.Pointer(&value2),unsafe.Pointer(&value3),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFont::setStyleStrategy(QFont::StyleStrategy)
func (q *QFont) SetStyleStrategy(s QFont_StyleStrategy)  {
	q.Drv(37000,37154,unsafe.Pointer(&s),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFont::setUnderline(bool)
func (q *QFont) SetUnderline(value bool)  {
	q.Drv(37000,37155,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFont::setWeight(int)
func (q *QFont) SetWeight(value int)  {
	q.Drv(37000,37156,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFont::setWordSpacing(double)
func (q *QFont) SetWordSpacing(spacing float64)  {
	q.Drv(37000,37157,unsafe.Pointer(&spacing),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFont::stretch()
func (q *QFont) Stretch() int {
	var __rv int
	q.Drv(37000,37158,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFont::strikeOut()
func (q *QFont) StrikeOut() bool {
	var __rv bool
	q.Drv(37000,37159,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFont::style()
func (q *QFont) Style() QFont_Style {
	var __rv QFont_Style
	q.Drv(37000,37160,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFont::styleHint()
func (q *QFont) StyleHint() QFont_StyleHint {
	var __rv QFont_StyleHint
	q.Drv(37000,37161,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFont::styleStrategy()
func (q *QFont) StyleStrategy() QFont_StyleStrategy {
	var __rv QFont_StyleStrategy
	q.Drv(37000,37162,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFont::substitute(QString const&)	
func QFontSubstitute(value string) string {
	var __rv string
	DirectQtDrv(nil,37000,37163,unsafe.Pointer(&value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFont::substitute(QString const&)
func (q *QFont) Substitute(value string) string {
	var __rv string
	q.Drv(37000,37163,unsafe.Pointer(&value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFont::substitutes(QString const&)	
func QFontSubstitutes(value string) []string {
	var __rv []string
	DirectQtDrv(nil,37000,37164,unsafe.Pointer(&value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFont::substitutes(QString const&)
func (q *QFont) Substitutes(value string) []string {
	var __rv []string
	q.Drv(37000,37164,unsafe.Pointer(&value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFont::substitutions()	
func QFontSubstitutions() []string {
	var __rv []string
	DirectQtDrv(nil,37000,37165,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFont::substitutions()
func (q *QFont) Substitutions() []string {
	var __rv []string
	q.Drv(37000,37165,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFont::toString()
func (q *QFont) ToString() string {
	var __rv string
	q.Drv(37000,37166,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFont::underline()
func (q *QFont) Underline() bool {
	var __rv bool
	q.Drv(37000,37167,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFont::weight()
func (q *QFont) Weight() int {
	var __rv int
	q.Drv(37000,37168,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFont::wordSpacing()
func (q *QFont) WordSpacing() float64 {
	var __rv float64
	q.Drv(37000,37169,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
