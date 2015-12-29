// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QSettings_Format - QSettings::Format
type QSettings_Format uint32
const (
	QSettings_NativeFormat QSettings_Format = 0
	QSettings_IniFormat QSettings_Format = 1
	QSettings_InvalidFormat QSettings_Format = 16
	QSettings_CustomFormat1 QSettings_Format = 16+1
	QSettings_CustomFormat2 QSettings_Format = 16+1+1
	QSettings_CustomFormat3 QSettings_Format = 16+1+1+1
	QSettings_CustomFormat4 QSettings_Format = 16+1+1+1+1
	QSettings_CustomFormat5 QSettings_Format = 16+1+1+1+1+1
	QSettings_CustomFormat6 QSettings_Format = 16+1+1+1+1+1+1
	QSettings_CustomFormat7 QSettings_Format = 16+1+1+1+1+1+1+1
	QSettings_CustomFormat8 QSettings_Format = 16+1+1+1+1+1+1+1+1
	QSettings_CustomFormat9 QSettings_Format = 16+1+1+1+1+1+1+1+1+1
	QSettings_CustomFormat10 QSettings_Format = 16+1+1+1+1+1+1+1+1+1+1
	QSettings_CustomFormat11 QSettings_Format = 16+1+1+1+1+1+1+1+1+1+1+1
	QSettings_CustomFormat12 QSettings_Format = 16+1+1+1+1+1+1+1+1+1+1+1+1
	QSettings_CustomFormat13 QSettings_Format = 16+1+1+1+1+1+1+1+1+1+1+1+1+1
	QSettings_CustomFormat14 QSettings_Format = 16+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QSettings_CustomFormat15 QSettings_Format = 16+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QSettings_CustomFormat16 QSettings_Format = 16+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
)
//enum QSettings_Status - QSettings::Status
type QSettings_Status uint32
const (
	QSettings_NoError QSettings_Status = 0
	QSettings_AccessError QSettings_Status = 0
	QSettings_FormatError QSettings_Status = 1
)
//enum QSettings_Scope - QSettings::Scope
type QSettings_Scope uint32
const (
	QSettings_UserScope QSettings_Scope = 0
	QSettings_SystemScope QSettings_Scope = 1
)
//struct QSettings : QSettings
type QSettings struct {
	QObject
}
//QSettings::QSettings()	
func NewSettings() *QSettings {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),339000,339102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QSettings{}
	_p.SetDriver(__rv,339000,false)
	return _p
} 
//QSettings::QSettings(QObject*)	
func NewSettingsWithParent(parent QObjectInterface) *QSettings {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),339000,339103,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QSettings{}
	_p.SetDriver(__rv,339000,false)
	return _p
} 
//QSettings::QSettings(QString const&,QSettings::Format,QObject*)	
func NewSettingsWithFilenameFormatParent(fileName string,format QSettings_Format,parent QObjectInterface) *QSettings {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),339000,339104,unsafe.Pointer(&fileName),unsafe.Pointer(&format),Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QSettings{}
	_p.SetDriver(__rv,339000,false)
	return _p
} 
//QSettings::QSettings(QString const&,QString const&,QObject*)	
func NewSettingsWithOrganizationApplicationParent(organization string,application string,parent QObjectInterface) *QSettings {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),339000,339105,unsafe.Pointer(&organization),unsafe.Pointer(&application),Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QSettings{}
	_p.SetDriver(__rv,339000,false)
	return _p
} 
//QSettings::QSettings(QSettings::Scope,QString const&,QString const&,QObject*)	
func NewSettingsWithScopeOrganizationApplicationParent(scope QSettings_Scope,organization string,application string,parent QObjectInterface) *QSettings {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),339000,339106,unsafe.Pointer(&scope),unsafe.Pointer(&organization),unsafe.Pointer(&application),Native(parent),nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QSettings{}
	_p.SetDriver(__rv,339000,false)
	return _p
} 
//QSettings::QSettings(QSettings::Format,QSettings::Scope,QString const&,QString const&,QObject*)	
func NewSettingsWithFormatScopeOrganizationApplicationParent(format QSettings_Format,scope QSettings_Scope,organization string,application string,parent QObjectInterface) *QSettings {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),339000,339107,unsafe.Pointer(&format),unsafe.Pointer(&scope),unsafe.Pointer(&organization),unsafe.Pointer(&application),Native(parent),nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QSettings{}
	_p.SetDriver(__rv,339000,false)
	return _p
} 
//QSettings::allKeys()
func (q *QSettings) AllKeys() []string {
	var __rv []string
	q.Drv(339000,339108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSettings::applicationName()
func (q *QSettings) ApplicationName() string {
	var __rv string
	q.Drv(339000,339109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSettings::beginGroup(QString const&)
func (q *QSettings) BeginGroup(prefix string)  {
	q.Drv(339000,339110,unsafe.Pointer(&prefix),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSettings::beginReadArray(QString const&)
func (q *QSettings) BeginReadArray(prefix string) int {
	var __rv int
	q.Drv(339000,339111,unsafe.Pointer(&prefix),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSettings::beginWriteArray(QString const&)
func (q *QSettings) BeginWriteArray(prefix string)  {
	q.Drv(339000,339112,unsafe.Pointer(&prefix),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSettings::beginWriteArray(QString const&,int)
func (q *QSettings) BeginWriteArrayWithPrefixSize(prefix string,size int)  {
	q.Drv(339000,339113,unsafe.Pointer(&prefix),unsafe.Pointer(&size),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSettings::childGroups()
func (q *QSettings) ChildGroups() []string {
	var __rv []string
	q.Drv(339000,339114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSettings::childKeys()
func (q *QSettings) ChildKeys() []string {
	var __rv []string
	q.Drv(339000,339115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSettings::clear()
func (q *QSettings) Clear()  {
	q.Drv(339000,339116,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSettings::contains(QString const&)
func (q *QSettings) Contains(key string) bool {
	var __rv bool
	q.Drv(339000,339117,unsafe.Pointer(&key),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSettings::defaultFormat()	
func QSettingsDefaultFormat() QSettings_Format {
	var __rv QSettings_Format
	DirectQtDrv(nil,339000,339118,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSettings::defaultFormat()
func (q *QSettings) DefaultFormat() QSettings_Format {
	var __rv QSettings_Format
	q.Drv(339000,339118,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSettings::endArray()
func (q *QSettings) EndArray()  {
	q.Drv(339000,339119,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSettings::endGroup()
func (q *QSettings) EndGroup()  {
	q.Drv(339000,339120,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSettings::event(QEvent*)
func (q *QSettings) Event(event *QEvent) bool {
	var __rv bool
	q.Drv(339000,339121,Native(event),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSettings::fallbacksEnabled()
func (q *QSettings) FallbacksEnabled() bool {
	var __rv bool
	q.Drv(339000,339122,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSettings::fileName()
func (q *QSettings) FileName() string {
	var __rv string
	q.Drv(339000,339123,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSettings::format()
func (q *QSettings) Format() QSettings_Format {
	var __rv QSettings_Format
	q.Drv(339000,339124,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSettings::group()
func (q *QSettings) Group() string {
	var __rv string
	q.Drv(339000,339125,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSettings::iniCodec()
func (q *QSettings) IniCodec() *QTextCodec {
	var __rv uintptr
	q.Drv(339000,339126,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextCodec{}
	_rp.SetDriver(__rv,143000,true)
	return _rp
}	
//QSettings::isWritable()
func (q *QSettings) IsWritable() bool {
	var __rv bool
	q.Drv(339000,339127,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSettings::organizationName()
func (q *QSettings) OrganizationName() string {
	var __rv string
	q.Drv(339000,339128,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSettings::remove(QString const&)
func (q *QSettings) Remove(key string)  {
	q.Drv(339000,339129,unsafe.Pointer(&key),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSettings::scope()
func (q *QSettings) Scope() QSettings_Scope {
	var __rv QSettings_Scope
	q.Drv(339000,339130,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSettings::setArrayIndex(int)
func (q *QSettings) SetArrayIndex(i int)  {
	q.Drv(339000,339131,unsafe.Pointer(&i),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSettings::setDefaultFormat(QSettings::Format)	
func QSettingsSetDefaultFormat(format QSettings_Format)  {
	DirectQtDrv(nil,339000,339132,unsafe.Pointer(&format),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSettings::setDefaultFormat(QSettings::Format)
func (q *QSettings) SetDefaultFormat(format QSettings_Format)  {
	q.Drv(339000,339132,unsafe.Pointer(&format),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSettings::setFallbacksEnabled(bool)
func (q *QSettings) SetFallbacksEnabled(b bool)  {
	q.Drv(339000,339133,unsafe.Pointer(&b),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSettings::setIniCodec(QTextCodec*)
func (q *QSettings) SetIniCodec(codec *QTextCodec)  {
	q.Drv(339000,339134,Native(codec),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSettings::setIniCodec(char const*)
func (q *QSettings) SetIniCodecWithCodecname(codecName string)  {
	q.Drv(339000,339135,unsafe.Pointer(&codecName),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSettings::setPath(QSettings::Format,QSettings::Scope,QString const&)	
func QSettingsSetPath(format QSettings_Format,scope QSettings_Scope,path string)  {
	DirectQtDrv(nil,339000,339136,unsafe.Pointer(&format),unsafe.Pointer(&scope),unsafe.Pointer(&path),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSettings::setPath(QSettings::Format,QSettings::Scope,QString const&)
func (q *QSettings) SetPath(format QSettings_Format,scope QSettings_Scope,path string)  {
	q.Drv(339000,339136,unsafe.Pointer(&format),unsafe.Pointer(&scope),unsafe.Pointer(&path),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSettings::setSystemIniPath(QString const&)	
func QSettingsSetSystemIniPath(dir string)  {
	DirectQtDrv(nil,339000,339137,unsafe.Pointer(&dir),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSettings::setSystemIniPath(QString const&)
func (q *QSettings) SetSystemIniPath(dir string)  {
	q.Drv(339000,339137,unsafe.Pointer(&dir),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSettings::setUserIniPath(QString const&)	
func QSettingsSetUserIniPath(dir string)  {
	DirectQtDrv(nil,339000,339138,unsafe.Pointer(&dir),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSettings::setUserIniPath(QString const&)
func (q *QSettings) SetUserIniPath(dir string)  {
	q.Drv(339000,339138,unsafe.Pointer(&dir),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSettings::setValue(QString const&,QVariant const&)
func (q *QSettings) SetValue(key string,value *QVariant)  {
	q.Drv(339000,339139,unsafe.Pointer(&key),Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSettings::status()
func (q *QSettings) Status() QSettings_Status {
	var __rv QSettings_Status
	q.Drv(339000,339140,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSettings::sync()
func (q *QSettings) Sync()  {
	q.Drv(339000,339141,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSettings::value(QString const&)
func (q *QSettings) Value(key string) *QVariant {
	var __rv uintptr
	q.Drv(339000,339142,unsafe.Pointer(&key),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVariant{}
	_rp.SetDriver(__rv,182000,true)
	return _rp
}	
//QSettings::value(QString const&,QVariant const&)
func (q *QSettings) ValueWithKeyDefaultvalue(key string,defaultValue *QVariant) *QVariant {
	var __rv uintptr
	q.Drv(339000,339143,unsafe.Pointer(&key),Native(defaultValue),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVariant{}
	_rp.SetDriver(__rv,182000,true)
	return _rp
}	
