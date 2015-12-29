// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QFileDialog_FileMode - QFileDialog::FileMode
type QFileDialog_FileMode uint32
const (
	QFileDialog_AnyFile QFileDialog_FileMode = 0
	QFileDialog_ExistingFile QFileDialog_FileMode = 1
	QFileDialog_Directory QFileDialog_FileMode = 2
	QFileDialog_ExistingFiles QFileDialog_FileMode = 3
	QFileDialog_DirectoryOnly QFileDialog_FileMode = 4
)
//enum QFileDialog_Option - QFileDialog::Option
type QFileDialog_Option uint32
const (
	QFileDialog_ShowDirsOnly QFileDialog_Option = 0x00000001
	QFileDialog_DontResolveSymlinks QFileDialog_Option = 0x00000002
	QFileDialog_DontConfirmOverwrite QFileDialog_Option = 0x00000004
	QFileDialog_DontUseSheet QFileDialog_Option = 0x00000008
	QFileDialog_DontUseNativeDialog QFileDialog_Option = 0x00000010
	QFileDialog_ReadOnly QFileDialog_Option = 0x00000020
	QFileDialog_HideNameFilterDetails QFileDialog_Option = 0x00000040
)
//enum QFileDialog_DialogLabel - QFileDialog::DialogLabel
type QFileDialog_DialogLabel uint32
const (
	QFileDialog_LookIn QFileDialog_DialogLabel = 0
	QFileDialog_FileName QFileDialog_DialogLabel = 1
	QFileDialog_FileType QFileDialog_DialogLabel = 2
	QFileDialog_Accept QFileDialog_DialogLabel = 3
	QFileDialog_Reject QFileDialog_DialogLabel = 4
)
//enum QFileDialog_ViewMode - QFileDialog::ViewMode
type QFileDialog_ViewMode uint32
const (
	QFileDialog_Detail QFileDialog_ViewMode = 0
	QFileDialog_List QFileDialog_ViewMode = 1
)
//enum QFileDialog_AcceptMode - QFileDialog::AcceptMode
type QFileDialog_AcceptMode uint32
const (
	QFileDialog_AcceptOpen QFileDialog_AcceptMode = 0
	QFileDialog_AcceptSave QFileDialog_AcceptMode = 1
)
//struct QFileDialog : QFileDialog
type QFileDialog struct {
	QDialog
}
func (q *QFileDialog) OnFilesSelected(fn func([]string)) uintptr {
	var __rv uintptr
	q.Drv(237000,237102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QFileDialog) OnCurrentChanged(fn func(string)) uintptr {
	var __rv uintptr
	q.Drv(237000,237103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QFileDialog) OnDirectoryEntered(fn func(string)) uintptr {
	var __rv uintptr
	q.Drv(237000,237104,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QFileDialog) OnFileSelected(fn func(string)) uintptr {
	var __rv uintptr
	q.Drv(237000,237105,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QFileDialog) OnFilterSelected(fn func(string)) uintptr {
	var __rv uintptr
	q.Drv(237000,237106,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QFileDialog::QFileDialog()	
func NewFileDialog() *QFileDialog {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),237000,237107,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QFileDialog{}
	_p.SetDriver(__rv,237000,false)
	return _p
} 
//QFileDialog::QFileDialog(QWidget*,QFlags<Qt::WindowType>)	
func NewFileDialogWithParentFlags(parent QWidgetInterface,f Qt_WindowType) *QFileDialog {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),237000,237108,Native(parent),unsafe.Pointer(&f),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QFileDialog{}
	_p.SetDriver(__rv,237000,false)
	return _p
} 
//QFileDialog::QFileDialog(QWidget*,QString const&,QString const&,QString const&)	
func NewFileDialogWithParentCaptionDirectoryFilter(parent QWidgetInterface,caption string,directory string,filter string) *QFileDialog {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),237000,237109,Native(parent),unsafe.Pointer(&caption),unsafe.Pointer(&directory),unsafe.Pointer(&filter),nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QFileDialog{}
	_p.SetDriver(__rv,237000,false)
	return _p
} 
//QFileDialog::accept()
func (q *QFileDialog) Accept()  {
	q.Drv(237000,237110,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFileDialog::acceptMode()
func (q *QFileDialog) AcceptMode() QFileDialog_AcceptMode {
	var __rv QFileDialog_AcceptMode
	q.Drv(237000,237111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileDialog::changeEvent(QEvent*)
func (q *QFileDialog) ChangeEvent(e *QEvent)  {
	q.Drv(237000,237112,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFileDialog::confirmOverwrite()
func (q *QFileDialog) ConfirmOverwrite() bool {
	var __rv bool
	q.Drv(237000,237113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileDialog::defaultSuffix()
func (q *QFileDialog) DefaultSuffix() string {
	var __rv string
	q.Drv(237000,237114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileDialog::directory()
func (q *QFileDialog) Directory() *QDir {
	var __rv uintptr
	q.Drv(237000,237115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QDir{}
	_rp.SetDriver(__rv,22000,true)
	return _rp
}	
//QFileDialog::done(int)
func (q *QFileDialog) Done(result int)  {
	q.Drv(237000,237116,unsafe.Pointer(&result),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFileDialog::fileMode()
func (q *QFileDialog) FileMode() QFileDialog_FileMode {
	var __rv QFileDialog_FileMode
	q.Drv(237000,237117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileDialog::filter()
func (q *QFileDialog) Filter() QDir_Filter {
	var __rv QDir_Filter
	q.Drv(237000,237118,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileDialog::getExistingDirectory()	
func QFileDialogGetExistingDirectory() string {
	var __rv string
	DirectQtDrv(nil,237000,237119,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileDialog::getExistingDirectory()
func (q *QFileDialog) GetExistingDirectory() string {
	var __rv string
	q.Drv(237000,237119,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileDialog::getExistingDirectory(QWidget*,QString const&,QString const&,QFlags<QFileDialog::Option>)	
func QFileDialogGetExistingDirectoryWithParentCaptionDirOptions(parent QWidgetInterface,caption string,dir string,options QFileDialog_Option) string {
	var __rv string
	DirectQtDrv(nil,237000,237120,Native(parent),unsafe.Pointer(&caption),unsafe.Pointer(&dir),unsafe.Pointer(&options),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileDialog::getExistingDirectory(QWidget*,QString const&,QString const&,QFlags<QFileDialog::Option>)
func (q *QFileDialog) GetExistingDirectoryWithParentCaptionDirOptions(parent QWidgetInterface,caption string,dir string,options QFileDialog_Option) string {
	var __rv string
	q.Drv(237000,237120,Native(parent),unsafe.Pointer(&caption),unsafe.Pointer(&dir),unsafe.Pointer(&options),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileDialog::getOpenFileName()	
func QFileDialogGetOpenFileName() string {
	var __rv string
	DirectQtDrv(nil,237000,237121,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileDialog::getOpenFileName()
func (q *QFileDialog) GetOpenFileName() string {
	var __rv string
	q.Drv(237000,237121,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileDialog::getOpenFileName(QWidget*,QString const&,QString const&,QString const&,QString*,QFlags<QFileDialog::Option>)	
func QFileDialogGetOpenFileNameWithParentCaptionDirFilterSelectedfilterOptions(parent QWidgetInterface,caption string,dir string,filter string,selectedFilter *string,options QFileDialog_Option) string {
	var __rv string
	DirectQtDrv(nil,237000,237122,Native(parent),unsafe.Pointer(&caption),unsafe.Pointer(&dir),unsafe.Pointer(&filter),unsafe.Pointer(&selectedFilter),unsafe.Pointer(&options),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil)
	return __rv
}	
//QFileDialog::getOpenFileName(QWidget*,QString const&,QString const&,QString const&,QString*,QFlags<QFileDialog::Option>)
func (q *QFileDialog) GetOpenFileNameWithParentCaptionDirFilterSelectedfilterOptions(parent QWidgetInterface,caption string,dir string,filter string,selectedFilter *string,options QFileDialog_Option) string {
	var __rv string
	q.Drv(237000,237122,Native(parent),unsafe.Pointer(&caption),unsafe.Pointer(&dir),unsafe.Pointer(&filter),unsafe.Pointer(&selectedFilter),unsafe.Pointer(&options),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil)
	return __rv
}	
//QFileDialog::getOpenFileNames()	
func QFileDialogGetOpenFileNames() []string {
	var __rv []string
	DirectQtDrv(nil,237000,237123,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileDialog::getOpenFileNames()
func (q *QFileDialog) GetOpenFileNames() []string {
	var __rv []string
	q.Drv(237000,237123,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileDialog::getOpenFileNames(QWidget*,QString const&,QString const&,QString const&,QString*,QFlags<QFileDialog::Option>)	
func QFileDialogGetOpenFileNamesWithParentCaptionDirFilterSelectedfilterOptions(parent QWidgetInterface,caption string,dir string,filter string,selectedFilter *string,options QFileDialog_Option) []string {
	var __rv []string
	DirectQtDrv(nil,237000,237124,Native(parent),unsafe.Pointer(&caption),unsafe.Pointer(&dir),unsafe.Pointer(&filter),unsafe.Pointer(&selectedFilter),unsafe.Pointer(&options),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil)
	return __rv
}	
//QFileDialog::getOpenFileNames(QWidget*,QString const&,QString const&,QString const&,QString*,QFlags<QFileDialog::Option>)
func (q *QFileDialog) GetOpenFileNamesWithParentCaptionDirFilterSelectedfilterOptions(parent QWidgetInterface,caption string,dir string,filter string,selectedFilter *string,options QFileDialog_Option) []string {
	var __rv []string
	q.Drv(237000,237124,Native(parent),unsafe.Pointer(&caption),unsafe.Pointer(&dir),unsafe.Pointer(&filter),unsafe.Pointer(&selectedFilter),unsafe.Pointer(&options),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil)
	return __rv
}	
//QFileDialog::getSaveFileName()	
func QFileDialogGetSaveFileName() string {
	var __rv string
	DirectQtDrv(nil,237000,237125,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileDialog::getSaveFileName()
func (q *QFileDialog) GetSaveFileName() string {
	var __rv string
	q.Drv(237000,237125,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileDialog::getSaveFileName(QWidget*,QString const&,QString const&,QString const&,QString*,QFlags<QFileDialog::Option>)	
func QFileDialogGetSaveFileNameWithParentCaptionDirFilterSelectedfilterOptions(parent QWidgetInterface,caption string,dir string,filter string,selectedFilter *string,options QFileDialog_Option) string {
	var __rv string
	DirectQtDrv(nil,237000,237126,Native(parent),unsafe.Pointer(&caption),unsafe.Pointer(&dir),unsafe.Pointer(&filter),unsafe.Pointer(&selectedFilter),unsafe.Pointer(&options),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil)
	return __rv
}	
//QFileDialog::getSaveFileName(QWidget*,QString const&,QString const&,QString const&,QString*,QFlags<QFileDialog::Option>)
func (q *QFileDialog) GetSaveFileNameWithParentCaptionDirFilterSelectedfilterOptions(parent QWidgetInterface,caption string,dir string,filter string,selectedFilter *string,options QFileDialog_Option) string {
	var __rv string
	q.Drv(237000,237126,Native(parent),unsafe.Pointer(&caption),unsafe.Pointer(&dir),unsafe.Pointer(&filter),unsafe.Pointer(&selectedFilter),unsafe.Pointer(&options),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil)
	return __rv
}	
//QFileDialog::history()
func (q *QFileDialog) History() []string {
	var __rv []string
	q.Drv(237000,237127,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileDialog::iconProvider()
func (q *QFileDialog) IconProvider() *QFileIconProvider {
	var __rv uintptr
	q.Drv(237000,237128,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QFileIconProvider{}
	_rp.SetDriver(__rv,33000,true)
	return _rp
}	
//QFileDialog::isNameFilterDetailsVisible()
func (q *QFileDialog) IsNameFilterDetailsVisible() bool {
	var __rv bool
	q.Drv(237000,237129,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileDialog::isReadOnly()
func (q *QFileDialog) IsReadOnly() bool {
	var __rv bool
	q.Drv(237000,237130,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileDialog::itemDelegate()
func (q *QFileDialog) ItemDelegate() *QAbstractItemDelegate {
	var __rv uintptr
	q.Drv(237000,237131,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAbstractItemDelegate{}
	_rp.SetDriver(__rv,194000,false)
	return _rp
}	
//QFileDialog::labelText(QFileDialog::DialogLabel)
func (q *QFileDialog) LabelText(label QFileDialog_DialogLabel) string {
	var __rv string
	q.Drv(237000,237132,unsafe.Pointer(&label),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileDialog::nameFilters()
func (q *QFileDialog) NameFilters() []string {
	var __rv []string
	q.Drv(237000,237133,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileDialog::open()
func (q *QFileDialog) Open()  {
	q.Drv(237000,237134,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFileDialog::open(QObject*,char const*)
func (q *QFileDialog) OpenWithObjectMember(receiver QObjectInterface,member string)  {
	q.Drv(237000,237135,Native(receiver),unsafe.Pointer(&member),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFileDialog::options()
func (q *QFileDialog) Options() QFileDialog_Option {
	var __rv QFileDialog_Option
	q.Drv(237000,237136,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileDialog::proxyModel()
func (q *QFileDialog) ProxyModel() *QAbstractProxyModel {
	var __rv uintptr
	q.Drv(237000,237137,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAbstractProxyModel{}
	_rp.SetDriver(__rv,199000,false)
	return _rp
}	
//QFileDialog::resolveSymlinks()
func (q *QFileDialog) ResolveSymlinks() bool {
	var __rv bool
	q.Drv(237000,237138,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileDialog::restoreState(QByteArray const&)
func (q *QFileDialog) RestoreState(state []byte) bool {
	var __rv bool
	q.Drv(237000,237139,unsafe.Pointer(&state),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileDialog::saveState()
func (q *QFileDialog) SaveState() []byte {
	var __rv []byte
	q.Drv(237000,237140,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileDialog::selectFile(QString const&)
func (q *QFileDialog) SelectFile(filename string)  {
	q.Drv(237000,237141,unsafe.Pointer(&filename),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFileDialog::selectNameFilter(QString const&)
func (q *QFileDialog) SelectNameFilter(filter string)  {
	q.Drv(237000,237142,unsafe.Pointer(&filter),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFileDialog::selectedFiles()
func (q *QFileDialog) SelectedFiles() []string {
	var __rv []string
	q.Drv(237000,237143,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileDialog::selectedNameFilter()
func (q *QFileDialog) SelectedNameFilter() string {
	var __rv string
	q.Drv(237000,237144,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileDialog::setAcceptMode(QFileDialog::AcceptMode)
func (q *QFileDialog) SetAcceptMode(mode QFileDialog_AcceptMode)  {
	q.Drv(237000,237145,unsafe.Pointer(&mode),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFileDialog::setConfirmOverwrite(bool)
func (q *QFileDialog) SetConfirmOverwrite(enabled bool)  {
	q.Drv(237000,237146,unsafe.Pointer(&enabled),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFileDialog::setDefaultSuffix(QString const&)
func (q *QFileDialog) SetDefaultSuffix(suffix string)  {
	q.Drv(237000,237147,unsafe.Pointer(&suffix),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFileDialog::setDirectory(QDir const&)
func (q *QFileDialog) SetDirectory(directory *QDir)  {
	q.Drv(237000,237148,Native(directory),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFileDialog::setDirectory(QString const&)
func (q *QFileDialog) SetDirectoryWithDirectory(directory string)  {
	q.Drv(237000,237149,unsafe.Pointer(&directory),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFileDialog::setFileMode(QFileDialog::FileMode)
func (q *QFileDialog) SetFileMode(mode QFileDialog_FileMode)  {
	q.Drv(237000,237150,unsafe.Pointer(&mode),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFileDialog::setFilter(QFlags<QDir::Filter>)
func (q *QFileDialog) SetFilter(filters QDir_Filter)  {
	q.Drv(237000,237151,unsafe.Pointer(&filters),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFileDialog::setHistory(QStringList const&)
func (q *QFileDialog) SetHistory(paths []string)  {
	q.Drv(237000,237152,unsafe.Pointer(&paths),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFileDialog::setIconProvider(QFileIconProvider*)
func (q *QFileDialog) SetIconProvider(provider *QFileIconProvider)  {
	q.Drv(237000,237153,Native(provider),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFileDialog::setItemDelegate(QAbstractItemDelegate*)
func (q *QFileDialog) SetItemDelegate(delegate *QAbstractItemDelegate)  {
	q.Drv(237000,237154,Native(delegate),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFileDialog::setLabelText(QFileDialog::DialogLabel,QString const&)
func (q *QFileDialog) SetLabelText(label QFileDialog_DialogLabel,text string)  {
	q.Drv(237000,237155,unsafe.Pointer(&label),unsafe.Pointer(&text),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFileDialog::setNameFilter(QString const&)
func (q *QFileDialog) SetNameFilter(filter string)  {
	q.Drv(237000,237156,unsafe.Pointer(&filter),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFileDialog::setNameFilterDetailsVisible(bool)
func (q *QFileDialog) SetNameFilterDetailsVisible(enabled bool)  {
	q.Drv(237000,237157,unsafe.Pointer(&enabled),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFileDialog::setNameFilters(QStringList const&)
func (q *QFileDialog) SetNameFilters(filters []string)  {
	q.Drv(237000,237158,unsafe.Pointer(&filters),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFileDialog::setOption(QFileDialog::Option)
func (q *QFileDialog) SetOption(option QFileDialog_Option)  {
	q.Drv(237000,237159,unsafe.Pointer(&option),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFileDialog::setOption(QFileDialog::Option,bool)
func (q *QFileDialog) SetOptionWithOptionOn(option QFileDialog_Option,on bool)  {
	q.Drv(237000,237160,unsafe.Pointer(&option),unsafe.Pointer(&on),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFileDialog::setOptions(QFlags<QFileDialog::Option>)
func (q *QFileDialog) SetOptions(options QFileDialog_Option)  {
	q.Drv(237000,237161,unsafe.Pointer(&options),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFileDialog::setProxyModel(QAbstractProxyModel*)
func (q *QFileDialog) SetProxyModel(model *QAbstractProxyModel)  {
	q.Drv(237000,237162,Native(model),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFileDialog::setReadOnly(bool)
func (q *QFileDialog) SetReadOnly(enabled bool)  {
	q.Drv(237000,237163,unsafe.Pointer(&enabled),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFileDialog::setResolveSymlinks(bool)
func (q *QFileDialog) SetResolveSymlinks(enabled bool)  {
	q.Drv(237000,237164,unsafe.Pointer(&enabled),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFileDialog::setSidebarUrls(QList<QUrl> const&)
func (q *QFileDialog) SetSidebarUrls(urls []*QUrl)  {
	q.Drv(237000,237165,unsafe.Pointer(&urls),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFileDialog::setViewMode(QFileDialog::ViewMode)
func (q *QFileDialog) SetViewMode(mode QFileDialog_ViewMode)  {
	q.Drv(237000,237166,unsafe.Pointer(&mode),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFileDialog::setVisible(bool)
func (q *QFileDialog) SetVisible(visible bool)  {
	q.Drv(237000,237167,unsafe.Pointer(&visible),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFileDialog::sidebarUrls()
func (q *QFileDialog) SidebarUrls() []*QUrl {
	var __rv []*QUrl
	q.Drv(237000,237168,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileDialog::testOption(QFileDialog::Option)
func (q *QFileDialog) TestOption(option QFileDialog_Option) bool {
	var __rv bool
	q.Drv(237000,237169,unsafe.Pointer(&option),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileDialog::viewMode()
func (q *QFileDialog) ViewMode() QFileDialog_ViewMode {
	var __rv QFileDialog_ViewMode
	q.Drv(237000,237170,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
