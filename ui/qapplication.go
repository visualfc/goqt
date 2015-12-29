// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QApplication_Type - QApplication::Type
type QApplication_Type uint32
const (
	QApplication_Tty QApplication_Type = 0
	QApplication_GuiClient QApplication_Type = 1
	QApplication_GuiServer QApplication_Type = 2
)
//enum QApplication_ColorSpec - QApplication::ColorSpec
type QApplication_ColorSpec uint32
const (
	QApplication_NormalColor QApplication_ColorSpec = 0
	QApplication_CustomColor QApplication_ColorSpec = 1
	QApplication_ManyColor QApplication_ColorSpec = 2
)
//struct QApplication : QApplication
type QApplication struct {
	QObject
}
func (q *QApplication) OnLastWindowClosed(fn func()) uintptr {
	var __rv uintptr
	q.Drv(6000,6102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QApplication) OnFocusChanged(fn func(*QWidget,*QWidget)) uintptr {
	var __rv uintptr
	q.Drv(6000,6103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QApplication) OnSaveStateRequest(fn func(*QSessionManager)) uintptr {
	var __rv uintptr
	q.Drv(6000,6104,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QApplication) OnCommitDataRequest(fn func(*QSessionManager)) uintptr {
	var __rv uintptr
	q.Drv(6000,6105,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QApplication) OnFontDatabaseChanged(fn func()) uintptr {
	var __rv uintptr
	q.Drv(6000,6106,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QApplication::QApplication(int,char**)	
func NewApplication(args []string) *QApplication {
	_argc := len(args)
	_argv := drvStringArrayToC(args)
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),6000,6107,unsafe.Pointer(&_argc),unsafe.Pointer(&_argv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QApplication{}
	_p.SetDriver(__rv,6000,false)
	return _p
} 	
//QApplication::aboutQt()	
func QApplicationAboutQt()  {
	DirectQtDrv(nil,6000,6108,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QApplication::aboutQt()
func (q *QApplication) AboutQt()  {
	q.Drv(6000,6108,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QApplication::activeModalWidget()	
func QApplicationActiveModalWidget() *QWidget {
	var __rv uintptr
	DirectQtDrv(nil,6000,6109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QApplication::activeModalWidget()
func (q *QApplication) ActiveModalWidget() *QWidget {
	var __rv uintptr
	q.Drv(6000,6109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QApplication::activePopupWidget()	
func QApplicationActivePopupWidget() *QWidget {
	var __rv uintptr
	DirectQtDrv(nil,6000,6110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QApplication::activePopupWidget()
func (q *QApplication) ActivePopupWidget() *QWidget {
	var __rv uintptr
	q.Drv(6000,6110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QApplication::activeWindow()	
func QApplicationActiveWindow() *QWidget {
	var __rv uintptr
	DirectQtDrv(nil,6000,6111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QApplication::activeWindow()
func (q *QApplication) ActiveWindow() *QWidget {
	var __rv uintptr
	q.Drv(6000,6111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QCoreApplication::addLibraryPath(QString const&)	
func QApplicationAddLibraryPath(value string)  {
	DirectQtDrv(nil,6000,6112,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCoreApplication::addLibraryPath(QString const&)
func (q *QApplication) AddLibraryPath(value string)  {
	q.Drv(6000,6112,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QApplication::alert(QWidget*)	
func QApplicationAlert(widget QWidgetInterface)  {
	DirectQtDrv(nil,6000,6113,Native(widget),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QApplication::alert(QWidget*)
func (q *QApplication) Alert(widget QWidgetInterface)  {
	q.Drv(6000,6113,Native(widget),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QApplication::alert(QWidget*,int)	
func QApplicationAlertWithWidgetDuration(widget QWidgetInterface,duration int)  {
	DirectQtDrv(nil,6000,6114,Native(widget),unsafe.Pointer(&duration),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QApplication::alert(QWidget*,int)
func (q *QApplication) AlertWithWidgetDuration(widget QWidgetInterface,duration int)  {
	q.Drv(6000,6114,Native(widget),unsafe.Pointer(&duration),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QApplication::allWidgets()	
func QApplicationAllWidgets() []*QWidget {
	var __rv []*QWidget
	DirectQtDrv(nil,6000,6115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QApplication::allWidgets()
func (q *QApplication) AllWidgets() []*QWidget {
	var __rv []*QWidget
	q.Drv(6000,6115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QCoreApplication::applicationDirPath()	
func QApplicationApplicationDirPath() string {
	var __rv string
	DirectQtDrv(nil,6000,6116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QCoreApplication::applicationDirPath()
func (q *QApplication) ApplicationDirPath() string {
	var __rv string
	q.Drv(6000,6116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QCoreApplication::applicationFilePath()	
func QApplicationApplicationFilePath() string {
	var __rv string
	DirectQtDrv(nil,6000,6117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QCoreApplication::applicationFilePath()
func (q *QApplication) ApplicationFilePath() string {
	var __rv string
	q.Drv(6000,6117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QCoreApplication::applicationName()	
func QApplicationApplicationName() string {
	var __rv string
	DirectQtDrv(nil,6000,6118,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QCoreApplication::applicationName()
func (q *QApplication) ApplicationName() string {
	var __rv string
	q.Drv(6000,6118,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QCoreApplication::applicationPid()	
func QApplicationApplicationPid() int64 {
	var __rv int64
	DirectQtDrv(nil,6000,6119,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QCoreApplication::applicationPid()
func (q *QApplication) ApplicationPid() int64 {
	var __rv int64
	q.Drv(6000,6119,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QCoreApplication::applicationVersion()	
func QApplicationApplicationVersion() string {
	var __rv string
	DirectQtDrv(nil,6000,6120,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QCoreApplication::applicationVersion()
func (q *QApplication) ApplicationVersion() string {
	var __rv string
	q.Drv(6000,6120,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QCoreApplication::arguments()	
func QApplicationArguments() []string {
	var __rv []string
	DirectQtDrv(nil,6000,6121,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QCoreApplication::arguments()
func (q *QApplication) Arguments() []string {
	var __rv []string
	q.Drv(6000,6121,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QApplication::autoSipEnabled()
func (q *QApplication) AutoSipEnabled() bool {
	var __rv bool
	q.Drv(6000,6122,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QApplication::beep()	
func QApplicationBeep()  {
	DirectQtDrv(nil,6000,6123,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QApplication::beep()
func (q *QApplication) Beep()  {
	q.Drv(6000,6123,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QApplication::changeOverrideCursor(QCursor const&)	
func QApplicationChangeOverrideCursor(value *QCursor)  {
	DirectQtDrv(nil,6000,6124,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QApplication::changeOverrideCursor(QCursor const&)
func (q *QApplication) ChangeOverrideCursor(value *QCursor)  {
	q.Drv(6000,6124,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QApplication::clipboard()	
func QApplicationClipboard() *QClipboard {
	var __rv uintptr
	DirectQtDrv(nil,6000,6125,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QClipboard{}
	_rp.SetDriver(__rv,214000,false)
	return _rp
}	
//QApplication::clipboard()
func (q *QApplication) Clipboard() *QClipboard {
	var __rv uintptr
	q.Drv(6000,6125,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QClipboard{}
	_rp.SetDriver(__rv,214000,false)
	return _rp
}	
//QApplication::closeAllWindows()	
func QApplicationCloseAllWindows()  {
	DirectQtDrv(nil,6000,6126,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QApplication::closeAllWindows()
func (q *QApplication) CloseAllWindows()  {
	q.Drv(6000,6126,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCoreApplication::closingDown()	
func QApplicationClosingDown() bool {
	var __rv bool
	DirectQtDrv(nil,6000,6127,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QCoreApplication::closingDown()
func (q *QApplication) ClosingDown() bool {
	var __rv bool
	q.Drv(6000,6127,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QApplication::colorSpec()	
func QApplicationColorSpec() int {
	var __rv int
	DirectQtDrv(nil,6000,6128,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QApplication::colorSpec()
func (q *QApplication) ColorSpec() int {
	var __rv int
	q.Drv(6000,6128,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QApplication::cursorFlashTime()	
func QApplicationCursorFlashTime() int {
	var __rv int
	DirectQtDrv(nil,6000,6129,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QApplication::cursorFlashTime()
func (q *QApplication) CursorFlashTime() int {
	var __rv int
	q.Drv(6000,6129,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QApplication::desktop()	
func QApplicationDesktop() *QDesktopWidget {
	var __rv uintptr
	DirectQtDrv(nil,6000,6130,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QDesktopWidget{}
	_rp.SetDriver(__rv,224000,false)
	return _rp
}	
//QApplication::desktop()
func (q *QApplication) Desktop() *QDesktopWidget {
	var __rv uintptr
	q.Drv(6000,6130,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QDesktopWidget{}
	_rp.SetDriver(__rv,224000,false)
	return _rp
}	
//QApplication::desktopSettingsAware()	
func QApplicationDesktopSettingsAware() bool {
	var __rv bool
	DirectQtDrv(nil,6000,6131,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QApplication::desktopSettingsAware()
func (q *QApplication) DesktopSettingsAware() bool {
	var __rv bool
	q.Drv(6000,6131,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QApplication::doubleClickInterval()	
func QApplicationDoubleClickInterval() int {
	var __rv int
	DirectQtDrv(nil,6000,6132,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QApplication::doubleClickInterval()
func (q *QApplication) DoubleClickInterval() int {
	var __rv int
	q.Drv(6000,6132,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QApplication::event(QEvent*)
func (q *QApplication) Event(value *QEvent) bool {
	var __rv bool
	q.Drv(6000,6133,Native(value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QApplication::exec()	
func QApplicationExec() int {
	var __rv int
	DirectQtDrv(nil,6000,6134,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QApplication::exec()
func (q *QApplication) Exec() int {
	var __rv int
	q.Drv(6000,6134,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QCoreApplication::exit()	
func QApplicationExit()  {
	DirectQtDrv(nil,6000,6135,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCoreApplication::exit()
func (q *QApplication) Exit()  {
	q.Drv(6000,6135,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCoreApplication::exit(int)	
func QApplicationExitWithRetcode(retcode int)  {
	DirectQtDrv(nil,6000,6136,unsafe.Pointer(&retcode),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCoreApplication::exit(int)
func (q *QApplication) ExitWithRetcode(retcode int)  {
	q.Drv(6000,6136,unsafe.Pointer(&retcode),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCoreApplication::flush()	
func QApplicationFlush()  {
	DirectQtDrv(nil,6000,6137,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCoreApplication::flush()
func (q *QApplication) Flush()  {
	q.Drv(6000,6137,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QApplication::focusWidget()	
func QApplicationFocusWidget() *QWidget {
	var __rv uintptr
	DirectQtDrv(nil,6000,6138,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QApplication::focusWidget()
func (q *QApplication) FocusWidget() *QWidget {
	var __rv uintptr
	q.Drv(6000,6138,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QApplication::font()	
func QApplicationFont() *QFont {
	var __rv uintptr
	DirectQtDrv(nil,6000,6139,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QFont{}
	_rp.SetDriver(__rv,37000,true)
	return _rp
}	
//QApplication::font()
func (q *QApplication) Font() *QFont {
	var __rv uintptr
	q.Drv(6000,6139,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QFont{}
	_rp.SetDriver(__rv,37000,true)
	return _rp
}	
//QApplication::font(QWidget const*)	
func QApplicationFontWithWidget(value QWidgetInterface) *QFont {
	var __rv uintptr
	DirectQtDrv(nil,6000,6140,Native(value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QFont{}
	_rp.SetDriver(__rv,37000,true)
	return _rp
}	
//QApplication::font(QWidget const*)
func (q *QApplication) FontWithWidget(value QWidgetInterface) *QFont {
	var __rv uintptr
	q.Drv(6000,6140,Native(value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QFont{}
	_rp.SetDriver(__rv,37000,true)
	return _rp
}	
//QApplication::font(char const*)	
func QApplicationFontWithClassname(className string) *QFont {
	var __rv uintptr
	DirectQtDrv(nil,6000,6141,unsafe.Pointer(&className),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QFont{}
	_rp.SetDriver(__rv,37000,true)
	return _rp
}	
//QApplication::font(char const*)
func (q *QApplication) FontWithClassname(className string) *QFont {
	var __rv uintptr
	q.Drv(6000,6141,unsafe.Pointer(&className),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QFont{}
	_rp.SetDriver(__rv,37000,true)
	return _rp
}	
//QApplication::fontMetrics()	
func QApplicationFontMetrics() *QFontMetrics {
	var __rv uintptr
	DirectQtDrv(nil,6000,6142,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QFontMetrics{}
	_rp.SetDriver(__rv,40000,true)
	return _rp
}	
//QApplication::fontMetrics()
func (q *QApplication) FontMetrics() *QFontMetrics {
	var __rv uintptr
	q.Drv(6000,6142,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QFontMetrics{}
	_rp.SetDriver(__rv,40000,true)
	return _rp
}	
//QApplication::globalStrut()	
func QApplicationGlobalStrut() *QSize {
	var __rv uintptr
	DirectQtDrv(nil,6000,6143,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QApplication::globalStrut()
func (q *QApplication) GlobalStrut() *QSize {
	var __rv uintptr
	q.Drv(6000,6143,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QCoreApplication::hasPendingEvents()	
func QApplicationHasPendingEvents() bool {
	var __rv bool
	DirectQtDrv(nil,6000,6144,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QCoreApplication::hasPendingEvents()
func (q *QApplication) HasPendingEvents() bool {
	var __rv bool
	q.Drv(6000,6144,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QCoreApplication::installTranslator(QTranslator*)	
func QApplicationInstallTranslator(messageFile *QTranslator)  {
	DirectQtDrv(nil,6000,6145,Native(messageFile),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCoreApplication::installTranslator(QTranslator*)
func (q *QApplication) InstallTranslator(messageFile *QTranslator)  {
	q.Drv(6000,6145,Native(messageFile),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCoreApplication::instance()	
func QApplicationInstance() *QApplication {
	var __rv uintptr
	DirectQtDrv(nil,6000,6146,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QApplication{}
	_rp.SetDriver(__rv,6000,false)
	return _rp
}	
//QCoreApplication::instance()
func (q *QApplication) Instance() *QApplication {
	var __rv uintptr
	q.Drv(6000,6146,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QApplication{}
	_rp.SetDriver(__rv,6000,false)
	return _rp
}	
//QApplication::isEffectEnabled(Qt::UIEffect)	
func QApplicationIsEffectEnabled(value Qt_UIEffect) bool {
	var __rv bool
	DirectQtDrv(nil,6000,6147,unsafe.Pointer(&value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QApplication::isEffectEnabled(Qt::UIEffect)
func (q *QApplication) IsEffectEnabled(value Qt_UIEffect) bool {
	var __rv bool
	q.Drv(6000,6147,unsafe.Pointer(&value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QApplication::isLeftToRight()	
func QApplicationIsLeftToRight() bool {
	var __rv bool
	DirectQtDrv(nil,6000,6148,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QApplication::isLeftToRight()
func (q *QApplication) IsLeftToRight() bool {
	var __rv bool
	q.Drv(6000,6148,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QApplication::isRightToLeft()	
func QApplicationIsRightToLeft() bool {
	var __rv bool
	DirectQtDrv(nil,6000,6149,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QApplication::isRightToLeft()
func (q *QApplication) IsRightToLeft() bool {
	var __rv bool
	q.Drv(6000,6149,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QApplication::isSessionRestored()
func (q *QApplication) IsSessionRestored() bool {
	var __rv bool
	q.Drv(6000,6150,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QApplication::keyboardInputInterval()	
func QApplicationKeyboardInputInterval() int {
	var __rv int
	DirectQtDrv(nil,6000,6151,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QApplication::keyboardInputInterval()
func (q *QApplication) KeyboardInputInterval() int {
	var __rv int
	q.Drv(6000,6151,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QApplication::keyboardModifiers()	
func QApplicationKeyboardModifiers() Qt_KeyboardModifier {
	var __rv Qt_KeyboardModifier
	DirectQtDrv(nil,6000,6152,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QApplication::keyboardModifiers()
func (q *QApplication) KeyboardModifiers() Qt_KeyboardModifier {
	var __rv Qt_KeyboardModifier
	q.Drv(6000,6152,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QApplication::layoutDirection()	
func QApplicationLayoutDirection() Qt_LayoutDirection {
	var __rv Qt_LayoutDirection
	DirectQtDrv(nil,6000,6153,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QApplication::layoutDirection()
func (q *QApplication) LayoutDirection() Qt_LayoutDirection {
	var __rv Qt_LayoutDirection
	q.Drv(6000,6153,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QCoreApplication::libraryPaths()	
func QApplicationLibraryPaths() []string {
	var __rv []string
	DirectQtDrv(nil,6000,6154,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QCoreApplication::libraryPaths()
func (q *QApplication) LibraryPaths() []string {
	var __rv []string
	q.Drv(6000,6154,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QApplication::mouseButtons()	
func QApplicationMouseButtons() Qt_MouseButton {
	var __rv Qt_MouseButton
	DirectQtDrv(nil,6000,6155,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QApplication::mouseButtons()
func (q *QApplication) MouseButtons() Qt_MouseButton {
	var __rv Qt_MouseButton
	q.Drv(6000,6155,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QApplication::notify(QObject*,QEvent*)
func (q *QApplication) Notify(value2 QObjectInterface,value3 *QEvent) bool {
	var __rv bool
	q.Drv(6000,6156,Native(value2),Native(value3),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QCoreApplication::organizationDomain()	
func QApplicationOrganizationDomain() string {
	var __rv string
	DirectQtDrv(nil,6000,6157,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QCoreApplication::organizationDomain()
func (q *QApplication) OrganizationDomain() string {
	var __rv string
	q.Drv(6000,6157,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QCoreApplication::organizationName()	
func QApplicationOrganizationName() string {
	var __rv string
	DirectQtDrv(nil,6000,6158,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QCoreApplication::organizationName()
func (q *QApplication) OrganizationName() string {
	var __rv string
	q.Drv(6000,6158,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QApplication::overrideCursor()	
func QApplicationOverrideCursor() *QCursor {
	var __rv uintptr
	DirectQtDrv(nil,6000,6159,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QCursor{}
	_rp.SetDriver(__rv,18000,true)
	return _rp
}	
//QApplication::overrideCursor()
func (q *QApplication) OverrideCursor() *QCursor {
	var __rv uintptr
	q.Drv(6000,6159,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QCursor{}
	_rp.SetDriver(__rv,18000,true)
	return _rp
}	
//QApplication::palette()	
func QApplicationPalette() *QPalette {
	var __rv uintptr
	DirectQtDrv(nil,6000,6160,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPalette{}
	_rp.SetDriver(__rv,91000,true)
	return _rp
}	
//QApplication::palette()
func (q *QApplication) Palette() *QPalette {
	var __rv uintptr
	q.Drv(6000,6160,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPalette{}
	_rp.SetDriver(__rv,91000,true)
	return _rp
}	
//QApplication::palette(QWidget const*)	
func QApplicationPaletteWithWidget(value QWidgetInterface) *QPalette {
	var __rv uintptr
	DirectQtDrv(nil,6000,6161,Native(value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPalette{}
	_rp.SetDriver(__rv,91000,true)
	return _rp
}	
//QApplication::palette(QWidget const*)
func (q *QApplication) PaletteWithWidget(value QWidgetInterface) *QPalette {
	var __rv uintptr
	q.Drv(6000,6161,Native(value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPalette{}
	_rp.SetDriver(__rv,91000,true)
	return _rp
}	
//QApplication::palette(char const*)	
func QApplicationPaletteWithClassname(className string) *QPalette {
	var __rv uintptr
	DirectQtDrv(nil,6000,6162,unsafe.Pointer(&className),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPalette{}
	_rp.SetDriver(__rv,91000,true)
	return _rp
}	
//QApplication::palette(char const*)
func (q *QApplication) PaletteWithClassname(className string) *QPalette {
	var __rv uintptr
	q.Drv(6000,6162,unsafe.Pointer(&className),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPalette{}
	_rp.SetDriver(__rv,91000,true)
	return _rp
}	
//QCoreApplication::postEvent(QObject*,QEvent*)	
func QApplicationPostEvent(receiver QObjectInterface,event *QEvent)  {
	DirectQtDrv(nil,6000,6163,Native(receiver),Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCoreApplication::postEvent(QObject*,QEvent*)
func (q *QApplication) PostEvent(receiver QObjectInterface,event *QEvent)  {
	q.Drv(6000,6163,Native(receiver),Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCoreApplication::quit()	
func QApplicationQuit()  {
	DirectQtDrv(nil,6000,6164,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCoreApplication::quit()
func (q *QApplication) Quit()  {
	q.Drv(6000,6164,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QApplication::quitOnLastWindowClosed()	
func QApplicationQuitOnLastWindowClosed() bool {
	var __rv bool
	DirectQtDrv(nil,6000,6165,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QApplication::quitOnLastWindowClosed()
func (q *QApplication) QuitOnLastWindowClosed() bool {
	var __rv bool
	q.Drv(6000,6165,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QCoreApplication::removeLibraryPath(QString const&)	
func QApplicationRemoveLibraryPath(value string)  {
	DirectQtDrv(nil,6000,6166,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCoreApplication::removeLibraryPath(QString const&)
func (q *QApplication) RemoveLibraryPath(value string)  {
	q.Drv(6000,6166,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCoreApplication::removePostedEvents(QObject*)	
func QApplicationRemovePostedEvents(receiver QObjectInterface)  {
	DirectQtDrv(nil,6000,6167,Native(receiver),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCoreApplication::removePostedEvents(QObject*)
func (q *QApplication) RemovePostedEvents(receiver QObjectInterface)  {
	q.Drv(6000,6167,Native(receiver),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCoreApplication::removeTranslator(QTranslator*)	
func QApplicationRemoveTranslator(messageFile *QTranslator)  {
	DirectQtDrv(nil,6000,6168,Native(messageFile),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCoreApplication::removeTranslator(QTranslator*)
func (q *QApplication) RemoveTranslator(messageFile *QTranslator)  {
	q.Drv(6000,6168,Native(messageFile),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QApplication::restoreOverrideCursor()	
func QApplicationRestoreOverrideCursor()  {
	DirectQtDrv(nil,6000,6169,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QApplication::restoreOverrideCursor()
func (q *QApplication) RestoreOverrideCursor()  {
	q.Drv(6000,6169,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCoreApplication::sendEvent(QObject*,QEvent*)	
func QApplicationSendEvent(receiver QObjectInterface,event *QEvent) bool {
	var __rv bool
	DirectQtDrv(nil,6000,6170,Native(receiver),Native(event),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QCoreApplication::sendEvent(QObject*,QEvent*)
func (q *QApplication) SendEvent(receiver QObjectInterface,event *QEvent) bool {
	var __rv bool
	q.Drv(6000,6170,Native(receiver),Native(event),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QCoreApplication::sendPostedEvents()	
func QApplicationSendPostedEvents()  {
	DirectQtDrv(nil,6000,6171,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCoreApplication::sendPostedEvents()
func (q *QApplication) SendPostedEvents()  {
	q.Drv(6000,6171,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QApplication::sessionId()
func (q *QApplication) SessionId() string {
	var __rv string
	q.Drv(6000,6172,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QApplication::sessionKey()
func (q *QApplication) SessionKey() string {
	var __rv string
	q.Drv(6000,6173,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QApplication::setActiveWindow(QWidget*)	
func QApplicationSetActiveWindow(act QWidgetInterface)  {
	DirectQtDrv(nil,6000,6174,Native(act),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QApplication::setActiveWindow(QWidget*)
func (q *QApplication) SetActiveWindow(act QWidgetInterface)  {
	q.Drv(6000,6174,Native(act),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCoreApplication::setApplicationName(QString const&)	
func QApplicationSetApplicationName(application string)  {
	DirectQtDrv(nil,6000,6175,unsafe.Pointer(&application),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCoreApplication::setApplicationName(QString const&)
func (q *QApplication) SetApplicationName(application string)  {
	q.Drv(6000,6175,unsafe.Pointer(&application),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCoreApplication::setApplicationVersion(QString const&)	
func QApplicationSetApplicationVersion(version string)  {
	DirectQtDrv(nil,6000,6176,unsafe.Pointer(&version),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCoreApplication::setApplicationVersion(QString const&)
func (q *QApplication) SetApplicationVersion(version string)  {
	q.Drv(6000,6176,unsafe.Pointer(&version),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCoreApplication::setAttribute(Qt::ApplicationAttribute)	
func QApplicationSetAttribute(attribute Qt_ApplicationAttribute)  {
	DirectQtDrv(nil,6000,6177,unsafe.Pointer(&attribute),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCoreApplication::setAttribute(Qt::ApplicationAttribute)
func (q *QApplication) SetAttribute(attribute Qt_ApplicationAttribute)  {
	q.Drv(6000,6177,unsafe.Pointer(&attribute),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCoreApplication::setAttribute(Qt::ApplicationAttribute,bool)	
func QApplicationSetAttributeWithAttributeOn(attribute Qt_ApplicationAttribute,on bool)  {
	DirectQtDrv(nil,6000,6178,unsafe.Pointer(&attribute),unsafe.Pointer(&on),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCoreApplication::setAttribute(Qt::ApplicationAttribute,bool)
func (q *QApplication) SetAttributeWithAttributeOn(attribute Qt_ApplicationAttribute,on bool)  {
	q.Drv(6000,6178,unsafe.Pointer(&attribute),unsafe.Pointer(&on),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QApplication::setAutoSipEnabled(bool const)
func (q *QApplication) SetAutoSipEnabled(enabled bool)  {
	q.Drv(6000,6179,unsafe.Pointer(&enabled),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QApplication::setColorSpec(int)	
func QApplicationSetColorSpec(value int)  {
	DirectQtDrv(nil,6000,6180,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QApplication::setColorSpec(int)
func (q *QApplication) SetColorSpec(value int)  {
	q.Drv(6000,6180,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QApplication::setCursorFlashTime(int)	
func QApplicationSetCursorFlashTime(value int)  {
	DirectQtDrv(nil,6000,6181,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QApplication::setCursorFlashTime(int)
func (q *QApplication) SetCursorFlashTime(value int)  {
	q.Drv(6000,6181,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QApplication::setDesktopSettingsAware(bool)	
func QApplicationSetDesktopSettingsAware(value bool)  {
	DirectQtDrv(nil,6000,6182,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QApplication::setDesktopSettingsAware(bool)
func (q *QApplication) SetDesktopSettingsAware(value bool)  {
	q.Drv(6000,6182,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QApplication::setDoubleClickInterval(int)	
func QApplicationSetDoubleClickInterval(value int)  {
	DirectQtDrv(nil,6000,6183,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QApplication::setDoubleClickInterval(int)
func (q *QApplication) SetDoubleClickInterval(value int)  {
	q.Drv(6000,6183,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QApplication::setEffectEnabled(Qt::UIEffect)	
func QApplicationSetEffectEnabled(value Qt_UIEffect)  {
	DirectQtDrv(nil,6000,6184,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QApplication::setEffectEnabled(Qt::UIEffect)
func (q *QApplication) SetEffectEnabled(value Qt_UIEffect)  {
	q.Drv(6000,6184,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QApplication::setEffectEnabled(Qt::UIEffect,bool)	
func QApplicationSetEffectEnabledWithUieffectEnable(value2 Qt_UIEffect,enable bool)  {
	DirectQtDrv(nil,6000,6185,unsafe.Pointer(&value2),unsafe.Pointer(&enable),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QApplication::setEffectEnabled(Qt::UIEffect,bool)
func (q *QApplication) SetEffectEnabledWithUieffectEnable(value2 Qt_UIEffect,enable bool)  {
	q.Drv(6000,6185,unsafe.Pointer(&value2),unsafe.Pointer(&enable),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QApplication::setFont(QFont const&)	
func QApplicationSetFont(value *QFont)  {
	DirectQtDrv(nil,6000,6186,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QApplication::setFont(QFont const&)
func (q *QApplication) SetFont(value *QFont)  {
	q.Drv(6000,6186,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QApplication::setFont(QFont const&,char const*)	
func QApplicationSetFontWithFontClassname(value2 *QFont,className string)  {
	DirectQtDrv(nil,6000,6187,Native(value2),unsafe.Pointer(&className),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QApplication::setFont(QFont const&,char const*)
func (q *QApplication) SetFontWithFontClassname(value2 *QFont,className string)  {
	q.Drv(6000,6187,Native(value2),unsafe.Pointer(&className),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QApplication::setGlobalStrut(QSize const&)	
func QApplicationSetGlobalStrut(value *QSize)  {
	DirectQtDrv(nil,6000,6188,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QApplication::setGlobalStrut(QSize const&)
func (q *QApplication) SetGlobalStrut(value *QSize)  {
	q.Drv(6000,6188,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QApplication::setKeyboardInputInterval(int)	
func QApplicationSetKeyboardInputInterval(value int)  {
	DirectQtDrv(nil,6000,6189,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QApplication::setKeyboardInputInterval(int)
func (q *QApplication) SetKeyboardInputInterval(value int)  {
	q.Drv(6000,6189,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QApplication::setLayoutDirection(Qt::LayoutDirection)	
func QApplicationSetLayoutDirection(direction Qt_LayoutDirection)  {
	DirectQtDrv(nil,6000,6190,unsafe.Pointer(&direction),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QApplication::setLayoutDirection(Qt::LayoutDirection)
func (q *QApplication) SetLayoutDirection(direction Qt_LayoutDirection)  {
	q.Drv(6000,6190,unsafe.Pointer(&direction),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCoreApplication::setLibraryPaths(QStringList const&)	
func QApplicationSetLibraryPaths(value []string)  {
	DirectQtDrv(nil,6000,6191,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCoreApplication::setLibraryPaths(QStringList const&)
func (q *QApplication) SetLibraryPaths(value []string)  {
	q.Drv(6000,6191,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCoreApplication::setOrganizationDomain(QString const&)	
func QApplicationSetOrganizationDomain(orgDomain string)  {
	DirectQtDrv(nil,6000,6192,unsafe.Pointer(&orgDomain),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCoreApplication::setOrganizationDomain(QString const&)
func (q *QApplication) SetOrganizationDomain(orgDomain string)  {
	q.Drv(6000,6192,unsafe.Pointer(&orgDomain),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCoreApplication::setOrganizationName(QString const&)	
func QApplicationSetOrganizationName(orgName string)  {
	DirectQtDrv(nil,6000,6193,unsafe.Pointer(&orgName),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCoreApplication::setOrganizationName(QString const&)
func (q *QApplication) SetOrganizationName(orgName string)  {
	q.Drv(6000,6193,unsafe.Pointer(&orgName),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QApplication::setOverrideCursor(QCursor const&)	
func QApplicationSetOverrideCursor(value *QCursor)  {
	DirectQtDrv(nil,6000,6194,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QApplication::setOverrideCursor(QCursor const&)
func (q *QApplication) SetOverrideCursor(value *QCursor)  {
	q.Drv(6000,6194,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QApplication::setPalette(QPalette const&)	
func QApplicationSetPalette(value *QPalette)  {
	DirectQtDrv(nil,6000,6195,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QApplication::setPalette(QPalette const&)
func (q *QApplication) SetPalette(value *QPalette)  {
	q.Drv(6000,6195,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QApplication::setPalette(QPalette const&,char const*)	
func QApplicationSetPaletteWithPaletteClassname(value2 *QPalette,className string)  {
	DirectQtDrv(nil,6000,6196,Native(value2),unsafe.Pointer(&className),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QApplication::setPalette(QPalette const&,char const*)
func (q *QApplication) SetPaletteWithPaletteClassname(value2 *QPalette,className string)  {
	q.Drv(6000,6196,Native(value2),unsafe.Pointer(&className),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QApplication::setQuitOnLastWindowClosed(bool)	
func QApplicationSetQuitOnLastWindowClosed(quit bool)  {
	DirectQtDrv(nil,6000,6197,unsafe.Pointer(&quit),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QApplication::setQuitOnLastWindowClosed(bool)
func (q *QApplication) SetQuitOnLastWindowClosed(quit bool)  {
	q.Drv(6000,6197,unsafe.Pointer(&quit),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QApplication::setStartDragDistance(int)	
func QApplicationSetStartDragDistance(l int)  {
	DirectQtDrv(nil,6000,6198,unsafe.Pointer(&l),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QApplication::setStartDragDistance(int)
func (q *QApplication) SetStartDragDistance(l int)  {
	q.Drv(6000,6198,unsafe.Pointer(&l),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QApplication::setStartDragTime(int)	
func QApplicationSetStartDragTime(ms int)  {
	DirectQtDrv(nil,6000,6199,unsafe.Pointer(&ms),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QApplication::setStartDragTime(int)
func (q *QApplication) SetStartDragTime(ms int)  {
	q.Drv(6000,6199,unsafe.Pointer(&ms),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QApplication::setStyle(QString const&)	
func QApplicationSetStyle(value string) *QStyle {
	var __rv uintptr
	DirectQtDrv(nil,6000,6200,unsafe.Pointer(&value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QStyle{}
	_rp.SetDriver(__rv,357000,false)
	return _rp
}	
//QApplication::setStyle(QString const&)
func (q *QApplication) SetStyle(value string) *QStyle {
	var __rv uintptr
	q.Drv(6000,6200,unsafe.Pointer(&value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QStyle{}
	_rp.SetDriver(__rv,357000,false)
	return _rp
}	
//QApplication::setStyle(QStyle*)	
func QApplicationSetStyleWithStyle(value *QStyle)  {
	DirectQtDrv(nil,6000,6201,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QApplication::setStyle(QStyle*)
func (q *QApplication) SetStyleWithStyle(value *QStyle)  {
	q.Drv(6000,6201,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QApplication::setStyleSheet(QString const&)
func (q *QApplication) SetStyleSheet(sheet string)  {
	q.Drv(6000,6202,unsafe.Pointer(&sheet),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QApplication::setWheelScrollLines(int)	
func QApplicationSetWheelScrollLines(value int)  {
	DirectQtDrv(nil,6000,6203,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QApplication::setWheelScrollLines(int)
func (q *QApplication) SetWheelScrollLines(value int)  {
	q.Drv(6000,6203,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QApplication::setWindowIcon(QIcon const&)	
func QApplicationSetWindowIcon(icon *QIcon)  {
	DirectQtDrv(nil,6000,6204,Native(icon),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QApplication::setWindowIcon(QIcon const&)
func (q *QApplication) SetWindowIcon(icon *QIcon)  {
	q.Drv(6000,6204,Native(icon),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QApplication::startDragDistance()	
func QApplicationStartDragDistance() int {
	var __rv int
	DirectQtDrv(nil,6000,6205,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QApplication::startDragDistance()
func (q *QApplication) StartDragDistance() int {
	var __rv int
	q.Drv(6000,6205,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QApplication::startDragTime()	
func QApplicationStartDragTime() int {
	var __rv int
	DirectQtDrv(nil,6000,6206,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QApplication::startDragTime()
func (q *QApplication) StartDragTime() int {
	var __rv int
	q.Drv(6000,6206,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QCoreApplication::startingUp()	
func QApplicationStartingUp() bool {
	var __rv bool
	DirectQtDrv(nil,6000,6207,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QCoreApplication::startingUp()
func (q *QApplication) StartingUp() bool {
	var __rv bool
	q.Drv(6000,6207,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QApplication::style()	
func QApplicationStyle() *QStyle {
	var __rv uintptr
	DirectQtDrv(nil,6000,6208,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QStyle{}
	_rp.SetDriver(__rv,357000,false)
	return _rp
}	
//QApplication::style()
func (q *QApplication) Style() *QStyle {
	var __rv uintptr
	q.Drv(6000,6208,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QStyle{}
	_rp.SetDriver(__rv,357000,false)
	return _rp
}	
//QApplication::styleSheet()
func (q *QApplication) StyleSheet() string {
	var __rv string
	q.Drv(6000,6209,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QCoreApplication::testAttribute(Qt::ApplicationAttribute)	
func QApplicationTestAttribute(attribute Qt_ApplicationAttribute) bool {
	var __rv bool
	DirectQtDrv(nil,6000,6210,unsafe.Pointer(&attribute),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QCoreApplication::testAttribute(Qt::ApplicationAttribute)
func (q *QApplication) TestAttribute(attribute Qt_ApplicationAttribute) bool {
	var __rv bool
	q.Drv(6000,6210,unsafe.Pointer(&attribute),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QApplication::topLevelAt(QPoint const&)	
func QApplicationTopLevelAt(p *QPoint) *QWidget {
	var __rv uintptr
	DirectQtDrv(nil,6000,6211,Native(p),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QApplication::topLevelAt(QPoint const&)
func (q *QApplication) TopLevelAt(p *QPoint) *QWidget {
	var __rv uintptr
	q.Drv(6000,6211,Native(p),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QApplication::topLevelAt(int,int)	
func QApplicationTopLevelAtWithXY(x int,y int) *QWidget {
	var __rv uintptr
	DirectQtDrv(nil,6000,6212,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QApplication::topLevelAt(int,int)
func (q *QApplication) TopLevelAtWithXY(x int,y int) *QWidget {
	var __rv uintptr
	q.Drv(6000,6212,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QApplication::topLevelWidgets()	
func QApplicationTopLevelWidgets() []*QWidget {
	var __rv []*QWidget
	DirectQtDrv(nil,6000,6213,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QApplication::topLevelWidgets()
func (q *QApplication) TopLevelWidgets() []*QWidget {
	var __rv []*QWidget
	q.Drv(6000,6213,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QCoreApplication::translate(char const*,char const*,char const*)	
func QApplicationTranslate(context string,key string,disambiguation string) string {
	var __rv string
	DirectQtDrv(nil,6000,6214,unsafe.Pointer(&context),unsafe.Pointer(&key),unsafe.Pointer(&disambiguation),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QCoreApplication::translate(char const*,char const*,char const*)
func (q *QApplication) Translate(context string,key string,disambiguation string) string {
	var __rv string
	q.Drv(6000,6214,unsafe.Pointer(&context),unsafe.Pointer(&key),unsafe.Pointer(&disambiguation),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QCoreApplication::translaten(char const*,char const*,char const*,int)	
func QApplicationTranslaten(context string,key string,disambiguation string,n int) string {
	var __rv string
	DirectQtDrv(nil,6000,6215,unsafe.Pointer(&context),unsafe.Pointer(&key),unsafe.Pointer(&disambiguation),unsafe.Pointer(&n),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QCoreApplication::translaten(char const*,char const*,char const*,int)
func (q *QApplication) Translaten(context string,key string,disambiguation string,n int) string {
	var __rv string
	q.Drv(6000,6215,unsafe.Pointer(&context),unsafe.Pointer(&key),unsafe.Pointer(&disambiguation),unsafe.Pointer(&n),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QApplication::wheelScrollLines()	
func QApplicationWheelScrollLines() int {
	var __rv int
	DirectQtDrv(nil,6000,6216,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QApplication::wheelScrollLines()
func (q *QApplication) WheelScrollLines() int {
	var __rv int
	q.Drv(6000,6216,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QApplication::widgetAt(QPoint const&)	
func QApplicationWidgetAt(p *QPoint) *QWidget {
	var __rv uintptr
	DirectQtDrv(nil,6000,6217,Native(p),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QApplication::widgetAt(QPoint const&)
func (q *QApplication) WidgetAt(p *QPoint) *QWidget {
	var __rv uintptr
	q.Drv(6000,6217,Native(p),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QApplication::widgetAt(int,int)	
func QApplicationWidgetAtWithXY(x int,y int) *QWidget {
	var __rv uintptr
	DirectQtDrv(nil,6000,6218,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QApplication::widgetAt(int,int)
func (q *QApplication) WidgetAtWithXY(x int,y int) *QWidget {
	var __rv uintptr
	q.Drv(6000,6218,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QApplication::windowIcon()	
func QApplicationWindowIcon() *QIcon {
	var __rv uintptr
	DirectQtDrv(nil,6000,6219,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QIcon{}
	_rp.SetDriver(__rv,51000,true)
	return _rp
}	
//QApplication::windowIcon()
func (q *QApplication) WindowIcon() *QIcon {
	var __rv uintptr
	q.Drv(6000,6219,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QIcon{}
	_rp.SetDriver(__rv,51000,true)
	return _rp
}	
