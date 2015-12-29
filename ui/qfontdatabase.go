// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QFontDatabase_WritingSystem - QFontDatabase::WritingSystem
type QFontDatabase_WritingSystem uint32
const (
	QFontDatabase_Any QFontDatabase_WritingSystem = 0
	QFontDatabase_Latin QFontDatabase_WritingSystem = 1
	QFontDatabase_Greek QFontDatabase_WritingSystem = 2
	QFontDatabase_Cyrillic QFontDatabase_WritingSystem = 3
	QFontDatabase_Armenian QFontDatabase_WritingSystem = 4
	QFontDatabase_Hebrew QFontDatabase_WritingSystem = 5
	QFontDatabase_Arabic QFontDatabase_WritingSystem = 6
	QFontDatabase_Syriac QFontDatabase_WritingSystem = 7
	QFontDatabase_Thaana QFontDatabase_WritingSystem = 8
	QFontDatabase_Devanagari QFontDatabase_WritingSystem = 9
	QFontDatabase_Bengali QFontDatabase_WritingSystem = 10
	QFontDatabase_Gurmukhi QFontDatabase_WritingSystem = 11
	QFontDatabase_Gujarati QFontDatabase_WritingSystem = 12
	QFontDatabase_Oriya QFontDatabase_WritingSystem = 13
	QFontDatabase_Tamil QFontDatabase_WritingSystem = 14
	QFontDatabase_Telugu QFontDatabase_WritingSystem = 15
	QFontDatabase_Kannada QFontDatabase_WritingSystem = 16
	QFontDatabase_Malayalam QFontDatabase_WritingSystem = 17
	QFontDatabase_Sinhala QFontDatabase_WritingSystem = 18
	QFontDatabase_Thai QFontDatabase_WritingSystem = 19
	QFontDatabase_Lao QFontDatabase_WritingSystem = 20
	QFontDatabase_Tibetan QFontDatabase_WritingSystem = 21
	QFontDatabase_Myanmar QFontDatabase_WritingSystem = 22
	QFontDatabase_Georgian QFontDatabase_WritingSystem = 23
	QFontDatabase_Khmer QFontDatabase_WritingSystem = 24
	QFontDatabase_SimplifiedChinese QFontDatabase_WritingSystem = 25
	QFontDatabase_TraditionalChinese QFontDatabase_WritingSystem = 26
	QFontDatabase_Japanese QFontDatabase_WritingSystem = 27
	QFontDatabase_Korean QFontDatabase_WritingSystem = 28
	QFontDatabase_Vietnamese QFontDatabase_WritingSystem = 29
	QFontDatabase_Symbol QFontDatabase_WritingSystem = 30
	QFontDatabase_Other QFontDatabase_WritingSystem = QFontDatabase_Symbol
	QFontDatabase_Ogham QFontDatabase_WritingSystem = QFontDatabase_Symbol+1
	QFontDatabase_Runic QFontDatabase_WritingSystem = QFontDatabase_Symbol+1+1
	QFontDatabase_Nko QFontDatabase_WritingSystem = QFontDatabase_Symbol+1+1+1
	QFontDatabase_WritingSystemsCount QFontDatabase_WritingSystem = QFontDatabase_Symbol+1+1+1+1
)
//struct QFontDatabase : QFontDatabase
type QFontDatabase struct {
	BaseDrv
}
//QFontDatabase::QFontDatabase()	
func NewFontDatabase() *QFontDatabase {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),38000,38102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QFontDatabase{}
	_p.SetDriver(__rv,38000,true)
	return _p
} 
//QFontDatabase::addApplicationFont(QString const&)	
func QFontDatabaseAddApplicationFont(fileName string) int {
	var __rv int
	DirectQtDrv(nil,38000,38103,unsafe.Pointer(&fileName),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontDatabase::addApplicationFont(QString const&)
func (q *QFontDatabase) AddApplicationFont(fileName string) int {
	var __rv int
	q.Drv(38000,38103,unsafe.Pointer(&fileName),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontDatabase::addApplicationFontFromData(QByteArray const&)	
func QFontDatabaseAddApplicationFontFromData(fontData []byte) int {
	var __rv int
	DirectQtDrv(nil,38000,38104,unsafe.Pointer(&fontData),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontDatabase::addApplicationFontFromData(QByteArray const&)
func (q *QFontDatabase) AddApplicationFontFromData(fontData []byte) int {
	var __rv int
	q.Drv(38000,38104,unsafe.Pointer(&fontData),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontDatabase::applicationFontFamilies(int)	
func QFontDatabaseApplicationFontFamilies(id int) []string {
	var __rv []string
	DirectQtDrv(nil,38000,38105,unsafe.Pointer(&id),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontDatabase::applicationFontFamilies(int)
func (q *QFontDatabase) ApplicationFontFamilies(id int) []string {
	var __rv []string
	q.Drv(38000,38105,unsafe.Pointer(&id),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontDatabase::bold(QString const&,QString const&)
func (q *QFontDatabase) Bold(family string,style string) bool {
	var __rv bool
	q.Drv(38000,38106,unsafe.Pointer(&family),unsafe.Pointer(&style),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontDatabase::families()
func (q *QFontDatabase) Families() []string {
	var __rv []string
	q.Drv(38000,38107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontDatabase::families(QFontDatabase::WritingSystem)
func (q *QFontDatabase) FamiliesWithWritingsystem(writingSystem QFontDatabase_WritingSystem) []string {
	var __rv []string
	q.Drv(38000,38108,unsafe.Pointer(&writingSystem),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontDatabase::font(QString const&,QString const&,int)
func (q *QFontDatabase) Font(family string,style string,pointSize int) *QFont {
	var __rv uintptr
	q.Drv(38000,38109,unsafe.Pointer(&family),unsafe.Pointer(&style),unsafe.Pointer(&pointSize),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QFont{}
	_rp.SetDriver(__rv,37000,true)
	return _rp
}	
//QFontDatabase::isBitmapScalable(QString const&)
func (q *QFontDatabase) IsBitmapScalable(family string) bool {
	var __rv bool
	q.Drv(38000,38110,unsafe.Pointer(&family),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontDatabase::isBitmapScalable(QString const&,QString const&)
func (q *QFontDatabase) IsBitmapScalableWithFamilyStyle(family string,style string) bool {
	var __rv bool
	q.Drv(38000,38111,unsafe.Pointer(&family),unsafe.Pointer(&style),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontDatabase::isFixedPitch(QString const&)
func (q *QFontDatabase) IsFixedPitch(family string) bool {
	var __rv bool
	q.Drv(38000,38112,unsafe.Pointer(&family),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontDatabase::isFixedPitch(QString const&,QString const&)
func (q *QFontDatabase) IsFixedPitchWithFamilyStyle(family string,style string) bool {
	var __rv bool
	q.Drv(38000,38113,unsafe.Pointer(&family),unsafe.Pointer(&style),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontDatabase::isScalable(QString const&)
func (q *QFontDatabase) IsScalable(family string) bool {
	var __rv bool
	q.Drv(38000,38114,unsafe.Pointer(&family),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontDatabase::isScalable(QString const&,QString const&)
func (q *QFontDatabase) IsScalableWithFamilyStyle(family string,style string) bool {
	var __rv bool
	q.Drv(38000,38115,unsafe.Pointer(&family),unsafe.Pointer(&style),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontDatabase::isSmoothlyScalable(QString const&)
func (q *QFontDatabase) IsSmoothlyScalable(family string) bool {
	var __rv bool
	q.Drv(38000,38116,unsafe.Pointer(&family),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontDatabase::isSmoothlyScalable(QString const&,QString const&)
func (q *QFontDatabase) IsSmoothlyScalableWithFamilyStyle(family string,style string) bool {
	var __rv bool
	q.Drv(38000,38117,unsafe.Pointer(&family),unsafe.Pointer(&style),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontDatabase::italic(QString const&,QString const&)
func (q *QFontDatabase) Italic(family string,style string) bool {
	var __rv bool
	q.Drv(38000,38118,unsafe.Pointer(&family),unsafe.Pointer(&style),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontDatabase::pointSizes(QString const&)
func (q *QFontDatabase) PointSizes(family string) []int {
	var __rv []int
	q.Drv(38000,38119,unsafe.Pointer(&family),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontDatabase::pointSizes(QString const&,QString const&)
func (q *QFontDatabase) PointSizesWithFamilyStyle(family string,style string) []int {
	var __rv []int
	q.Drv(38000,38120,unsafe.Pointer(&family),unsafe.Pointer(&style),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontDatabase::removeAllApplicationFonts()	
func QFontDatabaseRemoveAllApplicationFonts() bool {
	var __rv bool
	DirectQtDrv(nil,38000,38121,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontDatabase::removeAllApplicationFonts()
func (q *QFontDatabase) RemoveAllApplicationFonts() bool {
	var __rv bool
	q.Drv(38000,38121,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontDatabase::removeApplicationFont(int)	
func QFontDatabaseRemoveApplicationFont(id int) bool {
	var __rv bool
	DirectQtDrv(nil,38000,38122,unsafe.Pointer(&id),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontDatabase::removeApplicationFont(int)
func (q *QFontDatabase) RemoveApplicationFont(id int) bool {
	var __rv bool
	q.Drv(38000,38122,unsafe.Pointer(&id),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontDatabase::smoothSizes(QString const&,QString const&)
func (q *QFontDatabase) SmoothSizes(family string,style string) []int {
	var __rv []int
	q.Drv(38000,38123,unsafe.Pointer(&family),unsafe.Pointer(&style),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontDatabase::standardSizes()	
func QFontDatabaseStandardSizes() []int {
	var __rv []int
	DirectQtDrv(nil,38000,38124,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontDatabase::standardSizes()
func (q *QFontDatabase) StandardSizes() []int {
	var __rv []int
	q.Drv(38000,38124,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontDatabase::styleString(QFont const&)
func (q *QFontDatabase) StyleString(font *QFont) string {
	var __rv string
	q.Drv(38000,38125,Native(font),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontDatabase::styleString(QFontInfo const&)
func (q *QFontDatabase) StyleStringWithFontinfo(fontInfo *QFontInfo) string {
	var __rv string
	q.Drv(38000,38126,Native(fontInfo),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontDatabase::styles(QString const&)
func (q *QFontDatabase) Styles(family string) []string {
	var __rv []string
	q.Drv(38000,38127,unsafe.Pointer(&family),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontDatabase::supportsThreadedFontRendering()	
func QFontDatabaseSupportsThreadedFontRendering() bool {
	var __rv bool
	DirectQtDrv(nil,38000,38128,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontDatabase::supportsThreadedFontRendering()
func (q *QFontDatabase) SupportsThreadedFontRendering() bool {
	var __rv bool
	q.Drv(38000,38128,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontDatabase::weight(QString const&,QString const&)
func (q *QFontDatabase) Weight(family string,style string) int {
	var __rv int
	q.Drv(38000,38129,unsafe.Pointer(&family),unsafe.Pointer(&style),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontDatabase::writingSystemName(QFontDatabase::WritingSystem)	
func QFontDatabaseWritingSystemName(writingSystem QFontDatabase_WritingSystem) string {
	var __rv string
	DirectQtDrv(nil,38000,38130,unsafe.Pointer(&writingSystem),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontDatabase::writingSystemName(QFontDatabase::WritingSystem)
func (q *QFontDatabase) WritingSystemName(writingSystem QFontDatabase_WritingSystem) string {
	var __rv string
	q.Drv(38000,38130,unsafe.Pointer(&writingSystem),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontDatabase::writingSystemSample(QFontDatabase::WritingSystem)	
func QFontDatabaseWritingSystemSample(writingSystem QFontDatabase_WritingSystem) string {
	var __rv string
	DirectQtDrv(nil,38000,38131,unsafe.Pointer(&writingSystem),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontDatabase::writingSystemSample(QFontDatabase::WritingSystem)
func (q *QFontDatabase) WritingSystemSample(writingSystem QFontDatabase_WritingSystem) string {
	var __rv string
	q.Drv(38000,38131,unsafe.Pointer(&writingSystem),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontDatabase::writingSystems()
func (q *QFontDatabase) WritingSystems() []QFontDatabase_WritingSystem {
	var __rv []QFontDatabase_WritingSystem
	q.Drv(38000,38132,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontDatabase::writingSystems(QString const&)
func (q *QFontDatabase) WritingSystemsWithFamily(family string) []QFontDatabase_WritingSystem {
	var __rv []QFontDatabase_WritingSystem
	q.Drv(38000,38133,unsafe.Pointer(&family),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
