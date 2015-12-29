// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QPrinterInfo : QPrinterInfo
type QPrinterInfo struct {
	BaseDrv
}
//QPrinterInfo::QPrinterInfo()	
func NewPrinterInfo() *QPrinterInfo {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),105000,105102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPrinterInfo{}
	_p.SetDriver(__rv,105000,true)
	return _p
} 
//QPrinterInfo::QPrinterInfo(QPrinter const&)	
func NewPrinterInfoWithPrinter(printer *QPrinter) *QPrinterInfo {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),105000,105103,Native(printer),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPrinterInfo{}
	_p.SetDriver(__rv,105000,true)
	return _p
} 
//QPrinterInfo::QPrinterInfo(QPrinterInfo const&)	
func NewPrinterInfoCopy(src *QPrinterInfo) *QPrinterInfo {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),105000,105104,Native(src),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPrinterInfo{}
	_p.SetDriver(__rv,105000,true)
	return _p
} 
//QPrinterInfo::availablePrinters()	
func QPrinterInfoAvailablePrinters() []*QPrinterInfo {
	var __rv []*QPrinterInfo
	DirectQtDrv(nil,105000,105105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPrinterInfo::availablePrinters()
func (q *QPrinterInfo) AvailablePrinters() []*QPrinterInfo {
	var __rv []*QPrinterInfo
	q.Drv(105000,105105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPrinterInfo::defaultPrinter()	
func QPrinterInfoDefaultPrinter() *QPrinterInfo {
	var __rv uintptr
	DirectQtDrv(nil,105000,105106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPrinterInfo{}
	_rp.SetDriver(__rv,105000,true)
	return _rp
}	
//QPrinterInfo::defaultPrinter()
func (q *QPrinterInfo) DefaultPrinter() *QPrinterInfo {
	var __rv uintptr
	q.Drv(105000,105106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPrinterInfo{}
	_rp.SetDriver(__rv,105000,true)
	return _rp
}	
//QPrinterInfo::isDefault()
func (q *QPrinterInfo) IsDefault() bool {
	var __rv bool
	q.Drv(105000,105107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPrinterInfo::isNull()
func (q *QPrinterInfo) IsNull() bool {
	var __rv bool
	q.Drv(105000,105108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPrinterInfo::printerName()
func (q *QPrinterInfo) PrinterName() string {
	var __rv string
	q.Drv(105000,105109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPrinterInfo::supportedPaperSizes()
func (q *QPrinterInfo) SupportedPaperSizes() []QPrinter_PaperSize {
	var __rv []QPrinter_PaperSize
	q.Drv(105000,105110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
