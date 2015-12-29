// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QPrintEngine_PrintEnginePropertyKey - QPrintEngine::PrintEnginePropertyKey
type QPrintEngine_PrintEnginePropertyKey uint32
const (
	QPrintEngine_PPK_CollateCopies QPrintEngine_PrintEnginePropertyKey = 0
	QPrintEngine_PPK_ColorMode QPrintEngine_PrintEnginePropertyKey = 1
	QPrintEngine_PPK_Creator QPrintEngine_PrintEnginePropertyKey = 2
	QPrintEngine_PPK_DocumentName QPrintEngine_PrintEnginePropertyKey = 3
	QPrintEngine_PPK_FullPage QPrintEngine_PrintEnginePropertyKey = 4
	QPrintEngine_PPK_NumberOfCopies QPrintEngine_PrintEnginePropertyKey = 5
	QPrintEngine_PPK_Orientation QPrintEngine_PrintEnginePropertyKey = 6
	QPrintEngine_PPK_OutputFileName QPrintEngine_PrintEnginePropertyKey = 7
	QPrintEngine_PPK_PageOrder QPrintEngine_PrintEnginePropertyKey = 8
	QPrintEngine_PPK_PageRect QPrintEngine_PrintEnginePropertyKey = 9
	QPrintEngine_PPK_PageSize QPrintEngine_PrintEnginePropertyKey = 10
	QPrintEngine_PPK_PaperRect QPrintEngine_PrintEnginePropertyKey = 11
	QPrintEngine_PPK_PaperSource QPrintEngine_PrintEnginePropertyKey = 12
	QPrintEngine_PPK_PrinterName QPrintEngine_PrintEnginePropertyKey = 13
	QPrintEngine_PPK_PrinterProgram QPrintEngine_PrintEnginePropertyKey = 14
	QPrintEngine_PPK_Resolution QPrintEngine_PrintEnginePropertyKey = 15
	QPrintEngine_PPK_SelectionOption QPrintEngine_PrintEnginePropertyKey = 16
	QPrintEngine_PPK_SupportedResolutions QPrintEngine_PrintEnginePropertyKey = 17
	QPrintEngine_PPK_WindowsPageSize QPrintEngine_PrintEnginePropertyKey = 18
	QPrintEngine_PPK_FontEmbedding QPrintEngine_PrintEnginePropertyKey = 19
	QPrintEngine_PPK_SuppressSystemPrintStatus QPrintEngine_PrintEnginePropertyKey = 20
	QPrintEngine_PPK_Duplex QPrintEngine_PrintEnginePropertyKey = 21
	QPrintEngine_PPK_PaperSources QPrintEngine_PrintEnginePropertyKey = 22
	QPrintEngine_PPK_CustomPaperSize QPrintEngine_PrintEnginePropertyKey = 23
	QPrintEngine_PPK_PageMargins QPrintEngine_PrintEnginePropertyKey = 24
	QPrintEngine_PPK_CopyCount QPrintEngine_PrintEnginePropertyKey = 25
	QPrintEngine_PPK_SupportsMultipleCopies QPrintEngine_PrintEnginePropertyKey = 26
	QPrintEngine_PPK_PaperSize QPrintEngine_PrintEnginePropertyKey = QPrintEngine_PPK_PageSize
	QPrintEngine_PPK_CustomBase QPrintEngine_PrintEnginePropertyKey = 0xff00
)
//struct QPrintEngine : QPrintEngine
type QPrintEngine struct {
	BaseDrv
}
//QPrintEngine::abort()
func (q *QPrintEngine) Abort() bool {
	var __rv bool
	q.Drv(103000,103102,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPrintEngine::metric(QPaintDevice::PaintDeviceMetric)
func (q *QPrintEngine) Metric(value QPaintDevice_PaintDeviceMetric) int {
	var __rv int
	q.Drv(103000,103103,unsafe.Pointer(&value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPrintEngine::newPage()
func (q *QPrintEngine) NewPage() bool {
	var __rv bool
	q.Drv(103000,103104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPrintEngine::printerState()
func (q *QPrintEngine) PrinterState() QPrinter_PrinterState {
	var __rv QPrinter_PrinterState
	q.Drv(103000,103105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPrintEngine::property(QPrintEngine::PrintEnginePropertyKey)
func (q *QPrintEngine) Property(key QPrintEngine_PrintEnginePropertyKey) *QVariant {
	var __rv uintptr
	q.Drv(103000,103106,unsafe.Pointer(&key),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVariant{}
	_rp.SetDriver(__rv,182000,true)
	return _rp
}	
//QPrintEngine::setProperty(QPrintEngine::PrintEnginePropertyKey,QVariant const&)
func (q *QPrintEngine) SetProperty(key QPrintEngine_PrintEnginePropertyKey,value *QVariant)  {
	q.Drv(103000,103107,unsafe.Pointer(&key),Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
