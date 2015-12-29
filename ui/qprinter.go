// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QPrinter_OutputFormat - QPrinter::OutputFormat
type QPrinter_OutputFormat uint32
const (
	QPrinter_NativeFormat QPrinter_OutputFormat = 0
	QPrinter_PdfFormat QPrinter_OutputFormat = 1
	QPrinter_PostScriptFormat QPrinter_OutputFormat = 2
)
//enum QPrinter_DuplexMode - QPrinter::DuplexMode
type QPrinter_DuplexMode uint32
const (
	QPrinter_DuplexNone QPrinter_DuplexMode = 0
	QPrinter_DuplexAuto QPrinter_DuplexMode = 0
	QPrinter_DuplexLongSide QPrinter_DuplexMode = 1
	QPrinter_DuplexShortSide QPrinter_DuplexMode = 2
)
//enum QPrinter_PrintRange - QPrinter::PrintRange
type QPrinter_PrintRange uint32
const (
	QPrinter_AllPages QPrinter_PrintRange = 0
	QPrinter_Selection QPrinter_PrintRange = 1
	QPrinter_PageRange QPrinter_PrintRange = 2
	QPrinter_CurrentPage QPrinter_PrintRange = 3
)
//enum QPrinter_Unit - QPrinter::Unit
type QPrinter_Unit uint32
const (
	QPrinter_Millimeter QPrinter_Unit = 0
	QPrinter_Point QPrinter_Unit = 1
	QPrinter_Inch QPrinter_Unit = 2
	QPrinter_Pica QPrinter_Unit = 3
	QPrinter_Didot QPrinter_Unit = 4
	QPrinter_Cicero QPrinter_Unit = 5
	QPrinter_DevicePixel QPrinter_Unit = 6
)
//enum QPrinter_PrinterMode - QPrinter::PrinterMode
type QPrinter_PrinterMode uint32
const (
	QPrinter_ScreenResolution QPrinter_PrinterMode = 0
	QPrinter_PrinterResolution QPrinter_PrinterMode = 1
	QPrinter_HighResolution QPrinter_PrinterMode = 2
)
//enum QPrinter_Orientation - QPrinter::Orientation
type QPrinter_Orientation uint32
const (
	QPrinter_Portrait QPrinter_Orientation = 0
	QPrinter_Landscape QPrinter_Orientation = 1
)
//enum QPrinter_PrinterState - QPrinter::PrinterState
type QPrinter_PrinterState uint32
const (
	QPrinter_Idle QPrinter_PrinterState = 0
	QPrinter_Active QPrinter_PrinterState = 1
	QPrinter_Aborted QPrinter_PrinterState = 2
	QPrinter_Error QPrinter_PrinterState = 3
)
//enum QPrinter_ColorMode - QPrinter::ColorMode
type QPrinter_ColorMode uint32
const (
	QPrinter_GrayScale QPrinter_ColorMode = 0
	QPrinter_Color QPrinter_ColorMode = 1
)
//enum QPrinter_PaperSource - QPrinter::PaperSource
type QPrinter_PaperSource uint32
const (
	QPrinter_OnlyOne QPrinter_PaperSource = 0
	QPrinter_Lower QPrinter_PaperSource = 1
	QPrinter_Middle QPrinter_PaperSource = 2
	QPrinter_Manual QPrinter_PaperSource = 3
	QPrinter_Envelope QPrinter_PaperSource = 4
	QPrinter_EnvelopeManual QPrinter_PaperSource = 5
	QPrinter_Auto QPrinter_PaperSource = 6
	QPrinter_Tractor QPrinter_PaperSource = 7
	QPrinter_SmallFormat QPrinter_PaperSource = 8
	QPrinter_LargeFormat QPrinter_PaperSource = 9
	QPrinter_LargeCapacity QPrinter_PaperSource = 10
	QPrinter_Cassette QPrinter_PaperSource = 11
	QPrinter_FormSource QPrinter_PaperSource = 12
	QPrinter_MaxPageSource QPrinter_PaperSource = 13
)
//enum QPrinter_PageOrder - QPrinter::PageOrder
type QPrinter_PageOrder uint32
const (
	QPrinter_FirstPageFirst QPrinter_PageOrder = 0
	QPrinter_LastPageFirst QPrinter_PageOrder = 1
)
//enum QPrinter_PageSize - QPrinter::PageSize
type QPrinter_PageSize uint32
const (
	QPrinter_A4 QPrinter_PageSize = 0
	QPrinter_B5 QPrinter_PageSize = 1
	QPrinter_Letter QPrinter_PageSize = 2
	QPrinter_Legal QPrinter_PageSize = 3
	QPrinter_Executive QPrinter_PageSize = 4
	QPrinter_A0 QPrinter_PageSize = 5
	QPrinter_A1 QPrinter_PageSize = 6
	QPrinter_A2 QPrinter_PageSize = 7
	QPrinter_A3 QPrinter_PageSize = 8
	QPrinter_A5 QPrinter_PageSize = 9
	QPrinter_A6 QPrinter_PageSize = 10
	QPrinter_A7 QPrinter_PageSize = 11
	QPrinter_A8 QPrinter_PageSize = 12
	QPrinter_A9 QPrinter_PageSize = 13
	QPrinter_B0 QPrinter_PageSize = 14
	QPrinter_B1 QPrinter_PageSize = 15
	QPrinter_B10 QPrinter_PageSize = 16
	QPrinter_B2 QPrinter_PageSize = 17
	QPrinter_B3 QPrinter_PageSize = 18
	QPrinter_B4 QPrinter_PageSize = 19
	QPrinter_B6 QPrinter_PageSize = 20
	QPrinter_B7 QPrinter_PageSize = 21
	QPrinter_B8 QPrinter_PageSize = 22
	QPrinter_B9 QPrinter_PageSize = 23
	QPrinter_C5E QPrinter_PageSize = 24
	QPrinter_Comm10E QPrinter_PageSize = 25
	QPrinter_DLE QPrinter_PageSize = 26
	QPrinter_Folio QPrinter_PageSize = 27
	QPrinter_Ledger QPrinter_PageSize = 28
	QPrinter_Tabloid QPrinter_PageSize = 29
	QPrinter_Custom QPrinter_PageSize = 30
	QPrinter_NPageSize QPrinter_PageSize = QPrinter_Custom
	QPrinter_NPaperSize QPrinter_PageSize = QPrinter_Custom
)
//enum QPrinter_PaperSize - QPrinter::PaperSize
type QPrinter_PaperSize uint32
const (
	QPrinter_PaperSize_A4 QPrinter_PaperSize = 0
	QPrinter_PaperSize_B5 QPrinter_PaperSize = 1
	QPrinter_PaperSize_Letter QPrinter_PaperSize = 2
	QPrinter_PaperSize_Legal QPrinter_PaperSize = 3
	QPrinter_PaperSize_Executive QPrinter_PaperSize = 4
	QPrinter_PaperSize_A0 QPrinter_PaperSize = 5
	QPrinter_PaperSize_A1 QPrinter_PaperSize = 6
	QPrinter_PaperSize_A2 QPrinter_PaperSize = 7
	QPrinter_PaperSize_A3 QPrinter_PaperSize = 8
	QPrinter_PaperSize_A5 QPrinter_PaperSize = 9
	QPrinter_PaperSize_A6 QPrinter_PaperSize = 10
	QPrinter_PaperSize_A7 QPrinter_PaperSize = 11
	QPrinter_PaperSize_A8 QPrinter_PaperSize = 12
	QPrinter_PaperSize_A9 QPrinter_PaperSize = 13
	QPrinter_PaperSize_B0 QPrinter_PaperSize = 14
	QPrinter_PaperSize_B1 QPrinter_PaperSize = 15
	QPrinter_PaperSize_B10 QPrinter_PaperSize = 16
	QPrinter_PaperSize_B2 QPrinter_PaperSize = 17
	QPrinter_PaperSize_B3 QPrinter_PaperSize = 18
	QPrinter_PaperSize_B4 QPrinter_PaperSize = 19
	QPrinter_PaperSize_B6 QPrinter_PaperSize = 20
	QPrinter_PaperSize_B7 QPrinter_PaperSize = 21
	QPrinter_PaperSize_B8 QPrinter_PaperSize = 22
	QPrinter_PaperSize_B9 QPrinter_PaperSize = 23
	QPrinter_PaperSize_C5E QPrinter_PaperSize = 24
	QPrinter_PaperSize_Comm10E QPrinter_PaperSize = 25
	QPrinter_PaperSize_DLE QPrinter_PaperSize = 26
	QPrinter_PaperSize_Folio QPrinter_PaperSize = 27
	QPrinter_PaperSize_Ledger QPrinter_PaperSize = 28
	QPrinter_PaperSize_Tabloid QPrinter_PaperSize = 29
	QPrinter_PaperSize_Custom QPrinter_PaperSize = 30
	QPrinter_PaperSize_NPageSize QPrinter_PaperSize = QPrinter_PaperSize(QPrinter_Custom)
	QPrinter_PaperSize_NPaperSize QPrinter_PaperSize = QPrinter_PaperSize(QPrinter_Custom)
)
//struct QPrinter : QPrinter
type QPrinter struct {
	QPaintDevice
}
//QPrinter::QPrinter()	
func NewPrinter() *QPrinter {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),104000,104102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPrinter{}
	_p.SetDriver(__rv,104000,true)
	return _p
} 
//QPrinter::QPrinter(QPrinter::PrinterMode)	
func NewPrinterWithMode(mode QPrinter_PrinterMode) *QPrinter {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),104000,104103,unsafe.Pointer(&mode),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPrinter{}
	_p.SetDriver(__rv,104000,true)
	return _p
} 
//QPrinter::QPrinter(QPrinterInfo const&,QPrinter::PrinterMode)	
func NewPrinterWithPrinterMode(printer *QPrinterInfo,mode QPrinter_PrinterMode) *QPrinter {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),104000,104104,Native(printer),unsafe.Pointer(&mode),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPrinter{}
	_p.SetDriver(__rv,104000,true)
	return _p
} 
//QPrinter::abort()
func (q *QPrinter) Abort() bool {
	var __rv bool
	q.Drv(104000,104105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPrinter::actualNumCopies()
func (q *QPrinter) ActualNumCopies() int {
	var __rv int
	q.Drv(104000,104106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPrinter::collateCopies()
func (q *QPrinter) CollateCopies() bool {
	var __rv bool
	q.Drv(104000,104107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPrinter::colorMode()
func (q *QPrinter) ColorMode() QPrinter_ColorMode {
	var __rv QPrinter_ColorMode
	q.Drv(104000,104108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPrinter::copyCount()
func (q *QPrinter) CopyCount() int {
	var __rv int
	q.Drv(104000,104109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPrinter::creator()
func (q *QPrinter) Creator() string {
	var __rv string
	q.Drv(104000,104110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPrinter::devType()
func (q *QPrinter) DevType() int {
	var __rv int
	q.Drv(104000,104111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPrinter::docName()
func (q *QPrinter) DocName() string {
	var __rv string
	q.Drv(104000,104112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPrinter::doubleSidedPrinting()
func (q *QPrinter) DoubleSidedPrinting() bool {
	var __rv bool
	q.Drv(104000,104113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPrinter::duplex()
func (q *QPrinter) Duplex() QPrinter_DuplexMode {
	var __rv QPrinter_DuplexMode
	q.Drv(104000,104114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPrinter::fontEmbeddingEnabled()
func (q *QPrinter) FontEmbeddingEnabled() bool {
	var __rv bool
	q.Drv(104000,104115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPrinter::fromPage()
func (q *QPrinter) FromPage() int {
	var __rv int
	q.Drv(104000,104116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPrinter::fullPage()
func (q *QPrinter) FullPage() bool {
	var __rv bool
	q.Drv(104000,104117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPrinter::getPageMargins(double*,double*,double*,double*,QPrinter::Unit)
func (q *QPrinter) GetPageMargins(left *float64,top *float64,right *float64,bottom *float64,unit QPrinter_Unit)  {
	q.Drv(104000,104118,unsafe.Pointer(&left),unsafe.Pointer(&top),unsafe.Pointer(&right),unsafe.Pointer(&bottom),unsafe.Pointer(&unit),nil,nil,nil,nil,nil,nil,nil)
}	
//QPrinter::isValid()
func (q *QPrinter) IsValid() bool {
	var __rv bool
	q.Drv(104000,104119,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPrinter::newPage()
func (q *QPrinter) NewPage() bool {
	var __rv bool
	q.Drv(104000,104120,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPrinter::numCopies()
func (q *QPrinter) NumCopies() int {
	var __rv int
	q.Drv(104000,104121,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPrinter::orientation()
func (q *QPrinter) Orientation() QPrinter_Orientation {
	var __rv QPrinter_Orientation
	q.Drv(104000,104122,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPrinter::outputFileName()
func (q *QPrinter) OutputFileName() string {
	var __rv string
	q.Drv(104000,104123,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPrinter::outputFormat()
func (q *QPrinter) OutputFormat() QPrinter_OutputFormat {
	var __rv QPrinter_OutputFormat
	q.Drv(104000,104124,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPrinter::pageOrder()
func (q *QPrinter) PageOrder() QPrinter_PageOrder {
	var __rv QPrinter_PageOrder
	q.Drv(104000,104125,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPrinter::pageRect()
func (q *QPrinter) PageRect() *QRect {
	var __rv uintptr
	q.Drv(104000,104126,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QPrinter::pageRect(QPrinter::Unit)
func (q *QPrinter) PageRectWithUnit(value QPrinter_Unit) *QRectF {
	var __rv uintptr
	q.Drv(104000,104127,unsafe.Pointer(&value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRectF{}
	_rp.SetDriver(__rv,111000,true)
	return _rp
}	
//QPrinter::pageSize()
func (q *QPrinter) PageSize() QPrinter_PageSize {
	var __rv QPrinter_PageSize
	q.Drv(104000,104128,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPrinter::paintEngine()
func (q *QPrinter) PaintEngine() *QPaintEngine {
	var __rv uintptr
	q.Drv(104000,104129,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPaintEngine{}
	_rp.SetDriver(__rv,83000,true)
	return _rp
}	
//QPrinter::paperRect()
func (q *QPrinter) PaperRect() *QRect {
	var __rv uintptr
	q.Drv(104000,104130,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QPrinter::paperRect(QPrinter::Unit)
func (q *QPrinter) PaperRectWithUnit(value QPrinter_Unit) *QRectF {
	var __rv uintptr
	q.Drv(104000,104131,unsafe.Pointer(&value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRectF{}
	_rp.SetDriver(__rv,111000,true)
	return _rp
}	
//QPrinter::paperSize()
func (q *QPrinter) PaperSize() QPrinter_PageSize {
	var __rv QPrinter_PageSize
	q.Drv(104000,104132,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPrinter::paperSize(QPrinter::Unit)
func (q *QPrinter) PaperSizeWithUnit(unit QPrinter_Unit) *QSizeF {
	var __rv uintptr
	q.Drv(104000,104133,unsafe.Pointer(&unit),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSizeF{}
	_rp.SetDriver(__rv,120000,true)
	return _rp
}	
//QPrinter::paperSource()
func (q *QPrinter) PaperSource() QPrinter_PaperSource {
	var __rv QPrinter_PaperSource
	q.Drv(104000,104134,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPrinter::printEngine()
func (q *QPrinter) PrintEngine() *QPrintEngine {
	var __rv uintptr
	q.Drv(104000,104135,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPrintEngine{}
	_rp.SetDriver(__rv,103000,true)
	return _rp
}	
//QPrinter::printProgram()
func (q *QPrinter) PrintProgram() string {
	var __rv string
	q.Drv(104000,104136,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPrinter::printRange()
func (q *QPrinter) PrintRange() QPrinter_PrintRange {
	var __rv QPrinter_PrintRange
	q.Drv(104000,104137,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPrinter::printerName()
func (q *QPrinter) PrinterName() string {
	var __rv string
	q.Drv(104000,104138,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPrinter::printerState()
func (q *QPrinter) PrinterState() QPrinter_PrinterState {
	var __rv QPrinter_PrinterState
	q.Drv(104000,104139,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPrinter::resolution()
func (q *QPrinter) Resolution() int {
	var __rv int
	q.Drv(104000,104140,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPrinter::setCollateCopies(bool)
func (q *QPrinter) SetCollateCopies(collate bool)  {
	q.Drv(104000,104141,unsafe.Pointer(&collate),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPrinter::setColorMode(QPrinter::ColorMode)
func (q *QPrinter) SetColorMode(value QPrinter_ColorMode)  {
	q.Drv(104000,104142,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPrinter::setCopyCount(int)
func (q *QPrinter) SetCopyCount(value int)  {
	q.Drv(104000,104143,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPrinter::setCreator(QString const&)
func (q *QPrinter) SetCreator(value string)  {
	q.Drv(104000,104144,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPrinter::setDocName(QString const&)
func (q *QPrinter) SetDocName(value string)  {
	q.Drv(104000,104145,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPrinter::setDoubleSidedPrinting(bool)
func (q *QPrinter) SetDoubleSidedPrinting(enable bool)  {
	q.Drv(104000,104146,unsafe.Pointer(&enable),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPrinter::setDuplex(QPrinter::DuplexMode)
func (q *QPrinter) SetDuplex(duplex QPrinter_DuplexMode)  {
	q.Drv(104000,104147,unsafe.Pointer(&duplex),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPrinter::setFontEmbeddingEnabled(bool)
func (q *QPrinter) SetFontEmbeddingEnabled(enable bool)  {
	q.Drv(104000,104148,unsafe.Pointer(&enable),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPrinter::setFromTo(int,int)
func (q *QPrinter) SetFromTo(fromPage int,toPage int)  {
	q.Drv(104000,104149,unsafe.Pointer(&fromPage),unsafe.Pointer(&toPage),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPrinter::setFullPage(bool)
func (q *QPrinter) SetFullPage(value bool)  {
	q.Drv(104000,104150,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPrinter::setNumCopies(int)
func (q *QPrinter) SetNumCopies(value int)  {
	q.Drv(104000,104151,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPrinter::setOrientation(QPrinter::Orientation)
func (q *QPrinter) SetOrientation(value QPrinter_Orientation)  {
	q.Drv(104000,104152,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPrinter::setOutputFileName(QString const&)
func (q *QPrinter) SetOutputFileName(value string)  {
	q.Drv(104000,104153,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPrinter::setOutputFormat(QPrinter::OutputFormat)
func (q *QPrinter) SetOutputFormat(format QPrinter_OutputFormat)  {
	q.Drv(104000,104154,unsafe.Pointer(&format),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPrinter::setPageMargins(double,double,double,double,QPrinter::Unit)
func (q *QPrinter) SetPageMargins(left float64,top float64,right float64,bottom float64,unit QPrinter_Unit)  {
	q.Drv(104000,104155,unsafe.Pointer(&left),unsafe.Pointer(&top),unsafe.Pointer(&right),unsafe.Pointer(&bottom),unsafe.Pointer(&unit),nil,nil,nil,nil,nil,nil,nil)
}	
//QPrinter::setPageOrder(QPrinter::PageOrder)
func (q *QPrinter) SetPageOrder(value QPrinter_PageOrder)  {
	q.Drv(104000,104156,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPrinter::setPageSize(QPrinter::PageSize)
func (q *QPrinter) SetPageSize(value QPrinter_PageSize)  {
	q.Drv(104000,104157,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPrinter::setPaperSize(QPrinter::PageSize)
func (q *QPrinter) SetPaperSize(value QPrinter_PageSize)  {
	q.Drv(104000,104158,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPrinter::setPaperSize(QSizeF const&,QPrinter::Unit)
func (q *QPrinter) SetPaperSizeFWithPapersizeUnit(paperSize *QSizeF,unit QPrinter_Unit)  {
	q.Drv(104000,104159,Native(paperSize),unsafe.Pointer(&unit),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPrinter::setPaperSource(QPrinter::PaperSource)
func (q *QPrinter) SetPaperSource(value QPrinter_PaperSource)  {
	q.Drv(104000,104160,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPrinter::setPrintProgram(QString const&)
func (q *QPrinter) SetPrintProgram(value string)  {
	q.Drv(104000,104161,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPrinter::setPrintRange(QPrinter::PrintRange)
func (q *QPrinter) SetPrintRange(_range QPrinter_PrintRange)  {
	q.Drv(104000,104162,unsafe.Pointer(&_range),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPrinter::setPrinterName(QString const&)
func (q *QPrinter) SetPrinterName(value string)  {
	q.Drv(104000,104163,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPrinter::setResolution(int)
func (q *QPrinter) SetResolution(value int)  {
	q.Drv(104000,104164,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPrinter::supportedResolutions()
func (q *QPrinter) SupportedResolutions() []int {
	var __rv []int
	q.Drv(104000,104165,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPrinter::supportsMultipleCopies()
func (q *QPrinter) SupportsMultipleCopies() bool {
	var __rv bool
	q.Drv(104000,104166,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPrinter::toPage()
func (q *QPrinter) ToPage() int {
	var __rv int
	q.Drv(104000,104167,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
