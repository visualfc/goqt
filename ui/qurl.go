// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QUrl_ParsingMode - QUrl::ParsingMode
type QUrl_ParsingMode uint32
const (
	QUrl_TolerantMode QUrl_ParsingMode = 0
	QUrl_StrictMode QUrl_ParsingMode = 1
	QUrl_DecodedMode QUrl_ParsingMode = 2
)
//enum QUrl_FormattingOption - QUrl::FormattingOption
type QUrl_FormattingOption uint32
const (
	QUrl_None QUrl_FormattingOption = 0x0
	QUrl_RemoveScheme QUrl_FormattingOption = 0x1
	QUrl_RemovePassword QUrl_FormattingOption = 0x2
	QUrl_RemoveUserInfo QUrl_FormattingOption = QUrl_RemovePassword|0x4
	QUrl_RemovePort QUrl_FormattingOption = 0x8
	QUrl_RemoveAuthority QUrl_FormattingOption = QUrl_RemoveUserInfo|QUrl_RemovePort|0x10
	QUrl_RemovePath QUrl_FormattingOption = 0x20
	QUrl_RemoveQuery QUrl_FormattingOption = 0x40
	QUrl_RemoveFragment QUrl_FormattingOption = 0x80
	QUrl_StripTrailingSlash QUrl_FormattingOption = 0x10000
)
//enum QUrl_ComponentFormattingOption - QUrl::ComponentFormattingOption
type QUrl_ComponentFormattingOption uint32
const (
	QUrl_PrettyDecoded QUrl_ComponentFormattingOption = 0x000000
	QUrl_EncodeSpaces QUrl_ComponentFormattingOption = 0x100000
	QUrl_EncodeUnicode QUrl_ComponentFormattingOption = 0x200000
	QUrl_EncodeDelimiters QUrl_ComponentFormattingOption = 0x400000 | 0x800000
	QUrl_EncodeReserved QUrl_ComponentFormattingOption = 0x1000000
	QUrl_DecodeReserved QUrl_ComponentFormattingOption = 0x2000000
	QUrl_FullyEncoded QUrl_ComponentFormattingOption = QUrl_EncodeSpaces | QUrl_EncodeUnicode | QUrl_EncodeDelimiters | QUrl_EncodeReserved
	QUrl_FullyDecoded QUrl_ComponentFormattingOption = QUrl_FullyEncoded | QUrl_DecodeReserved | 0x4000000
)
//struct QUrl : QUrl
type QUrl struct {
	BaseDrv
}
//QUrl::QUrl()	
func NewUrl() *QUrl {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),180000,180102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QUrl{}
	_p.SetDriver(__rv,180000,true)
	return _p
} 
//QUrl::QUrl(QString const&)	
func NewUrlWithUrl(url string) *QUrl {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),180000,180103,unsafe.Pointer(&url),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QUrl{}
	_p.SetDriver(__rv,180000,true)
	return _p
} 
//QUrl::QUrl(QUrl const&)	
func NewUrlCopy(copy *QUrl) *QUrl {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),180000,180104,Native(copy),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QUrl{}
	_p.SetDriver(__rv,180000,true)
	return _p
} 
//QUrl::QUrl(QString const&,QUrl::ParsingMode)	
func NewUrlWithUrlMode(url string,mode QUrl_ParsingMode) *QUrl {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),180000,180105,unsafe.Pointer(&url),unsafe.Pointer(&mode),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QUrl{}
	_p.SetDriver(__rv,180000,true)
	return _p
} 
//QUrl::authority()
func (q *QUrl) Authority() string {
	var __rv string
	q.Drv(180000,180106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QUrl::clear()
func (q *QUrl) Clear()  {
	q.Drv(180000,180107,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QUrl::detach()
func (q *QUrl) Detach()  {
	q.Drv(180000,180108,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QUrl::errorString()
func (q *QUrl) ErrorString() string {
	var __rv string
	q.Drv(180000,180109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QUrl::fragment()
func (q *QUrl) Fragment() string {
	var __rv string
	q.Drv(180000,180110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QUrl::fromAce(QByteArray const&)	
func QUrlFromAce(value []byte) string {
	var __rv string
	DirectQtDrv(nil,180000,180111,unsafe.Pointer(&value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QUrl::fromAce(QByteArray const&)
func (q *QUrl) FromAce(value []byte) string {
	var __rv string
	q.Drv(180000,180111,unsafe.Pointer(&value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QUrl::fromEncoded(QByteArray const&)	
func QUrlFromEncoded(url []byte) *QUrl {
	var __rv uintptr
	DirectQtDrv(nil,180000,180112,unsafe.Pointer(&url),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QUrl{}
	_rp.SetDriver(__rv,180000,true)
	return _rp
}	
//QUrl::fromEncoded(QByteArray const&)
func (q *QUrl) FromEncoded(url []byte) *QUrl {
	var __rv uintptr
	q.Drv(180000,180112,unsafe.Pointer(&url),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QUrl{}
	_rp.SetDriver(__rv,180000,true)
	return _rp
}	
//QUrl::fromEncoded(QByteArray const&,QUrl::ParsingMode)	
func QUrlFromEncodedWithUrlMode(url []byte,mode QUrl_ParsingMode) *QUrl {
	var __rv uintptr
	DirectQtDrv(nil,180000,180113,unsafe.Pointer(&url),unsafe.Pointer(&mode),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QUrl{}
	_rp.SetDriver(__rv,180000,true)
	return _rp
}	
//QUrl::fromEncoded(QByteArray const&,QUrl::ParsingMode)
func (q *QUrl) FromEncodedWithUrlMode(url []byte,mode QUrl_ParsingMode) *QUrl {
	var __rv uintptr
	q.Drv(180000,180113,unsafe.Pointer(&url),unsafe.Pointer(&mode),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QUrl{}
	_rp.SetDriver(__rv,180000,true)
	return _rp
}	
//QUrl::fromLocalFile(QString const&)	
func QUrlFromLocalFile(localfile string) *QUrl {
	var __rv uintptr
	DirectQtDrv(nil,180000,180114,unsafe.Pointer(&localfile),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QUrl{}
	_rp.SetDriver(__rv,180000,true)
	return _rp
}	
//QUrl::fromLocalFile(QString const&)
func (q *QUrl) FromLocalFile(localfile string) *QUrl {
	var __rv uintptr
	q.Drv(180000,180114,unsafe.Pointer(&localfile),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QUrl{}
	_rp.SetDriver(__rv,180000,true)
	return _rp
}	
//QUrl::fromPercentEncoding(QByteArray const&)	
func QUrlFromPercentEncoding(value []byte) string {
	var __rv string
	DirectQtDrv(nil,180000,180115,unsafe.Pointer(&value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QUrl::fromPercentEncoding(QByteArray const&)
func (q *QUrl) FromPercentEncoding(value []byte) string {
	var __rv string
	q.Drv(180000,180115,unsafe.Pointer(&value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QUrl::fromUserInput(QString const&)	
func QUrlFromUserInput(userInput string) *QUrl {
	var __rv uintptr
	DirectQtDrv(nil,180000,180116,unsafe.Pointer(&userInput),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QUrl{}
	_rp.SetDriver(__rv,180000,true)
	return _rp
}	
//QUrl::fromUserInput(QString const&)
func (q *QUrl) FromUserInput(userInput string) *QUrl {
	var __rv uintptr
	q.Drv(180000,180116,unsafe.Pointer(&userInput),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QUrl{}
	_rp.SetDriver(__rv,180000,true)
	return _rp
}	
//QUrl::hasFragment()
func (q *QUrl) HasFragment() bool {
	var __rv bool
	q.Drv(180000,180117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QUrl::hasQuery()
func (q *QUrl) HasQuery() bool {
	var __rv bool
	q.Drv(180000,180118,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QUrl::host()
func (q *QUrl) Host() string {
	var __rv string
	q.Drv(180000,180119,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QUrl::idnWhitelist()	
func QUrlIdnWhitelist() []string {
	var __rv []string
	DirectQtDrv(nil,180000,180120,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QUrl::idnWhitelist()
func (q *QUrl) IdnWhitelist() []string {
	var __rv []string
	q.Drv(180000,180120,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QUrl::isDetached()
func (q *QUrl) IsDetached() bool {
	var __rv bool
	q.Drv(180000,180121,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QUrl::isEmpty()
func (q *QUrl) IsEmpty() bool {
	var __rv bool
	q.Drv(180000,180122,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QUrl::isParentOf(QUrl const&)
func (q *QUrl) IsParentOf(url *QUrl) bool {
	var __rv bool
	q.Drv(180000,180123,Native(url),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QUrl::isRelative()
func (q *QUrl) IsRelative() bool {
	var __rv bool
	q.Drv(180000,180124,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QUrl::isValid()
func (q *QUrl) IsValid() bool {
	var __rv bool
	q.Drv(180000,180125,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QUrl::password()
func (q *QUrl) Password() string {
	var __rv string
	q.Drv(180000,180126,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QUrl::path()
func (q *QUrl) Path() string {
	var __rv string
	q.Drv(180000,180127,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QUrl::port()
func (q *QUrl) Port() int {
	var __rv int
	q.Drv(180000,180128,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QUrl::port(int)
func (q *QUrl) PortWithDefaultport(defaultPort int) int {
	var __rv int
	q.Drv(180000,180129,unsafe.Pointer(&defaultPort),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QUrl::query()
func (q *QUrl) Query() string {
	var __rv string
	q.Drv(180000,180130,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QUrl::query(QUrl::ComponentFormattingOption)
func (q *QUrl) QueryWithOptions(options QUrl_ComponentFormattingOption) string {
	var __rv string
	q.Drv(180000,180131,unsafe.Pointer(&options),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QUrl::resolved(QUrl const&)
func (q *QUrl) Resolved(relative *QUrl) *QUrl {
	var __rv uintptr
	q.Drv(180000,180132,Native(relative),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QUrl{}
	_rp.SetDriver(__rv,180000,true)
	return _rp
}	
//QUrl::scheme()
func (q *QUrl) Scheme() string {
	var __rv string
	q.Drv(180000,180133,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QUrl::setAuthority(QString const&)
func (q *QUrl) SetAuthority(authority string)  {
	q.Drv(180000,180134,unsafe.Pointer(&authority),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QUrl::setFragment(QString const&)
func (q *QUrl) SetFragment(fragment string)  {
	q.Drv(180000,180135,unsafe.Pointer(&fragment),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QUrl::setHost(QString const&)
func (q *QUrl) SetHost(host string)  {
	q.Drv(180000,180136,unsafe.Pointer(&host),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QUrl::setIdnWhitelist(QStringList const&)	
func QUrlSetIdnWhitelist(value []string)  {
	DirectQtDrv(nil,180000,180137,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QUrl::setIdnWhitelist(QStringList const&)
func (q *QUrl) SetIdnWhitelist(value []string)  {
	q.Drv(180000,180137,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QUrl::setPassword(QString const&)
func (q *QUrl) SetPassword(password string)  {
	q.Drv(180000,180138,unsafe.Pointer(&password),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QUrl::setPath(QString const&)
func (q *QUrl) SetPath(path string)  {
	q.Drv(180000,180139,unsafe.Pointer(&path),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QUrl::setPort(int)
func (q *QUrl) SetPort(port int)  {
	q.Drv(180000,180140,unsafe.Pointer(&port),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QUrl::setQuery(QString const&)
func (q *QUrl) SetQuery(query string)  {
	q.Drv(180000,180141,unsafe.Pointer(&query),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QUrl::setQuery(QUrlQuery const&)
func (q *QUrl) SetQueryWithQuery(query *QUrlQuery)  {
	q.Drv(180000,180142,Native(query),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QUrl::setQuery(QString const&,QUrl::ParsingMode)
func (q *QUrl) SetQueryWithQueryMode(query string,mode QUrl_ParsingMode)  {
	q.Drv(180000,180143,unsafe.Pointer(&query),unsafe.Pointer(&mode),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QUrl::setScheme(QString const&)
func (q *QUrl) SetScheme(scheme string)  {
	q.Drv(180000,180144,unsafe.Pointer(&scheme),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QUrl::setUrl(QString const&)
func (q *QUrl) SetUrl(url string)  {
	q.Drv(180000,180145,unsafe.Pointer(&url),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QUrl::setUrl(QString const&,QUrl::ParsingMode)
func (q *QUrl) SetUrlWithUrlMode(url string,mode QUrl_ParsingMode)  {
	q.Drv(180000,180146,unsafe.Pointer(&url),unsafe.Pointer(&mode),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QUrl::setUserInfo(QString const&)
func (q *QUrl) SetUserInfo(userInfo string)  {
	q.Drv(180000,180147,unsafe.Pointer(&userInfo),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QUrl::setUserName(QString const&)
func (q *QUrl) SetUserName(userName string)  {
	q.Drv(180000,180148,unsafe.Pointer(&userName),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QUrl::toAce(QString const&)	
func QUrlToAce(value string) []byte {
	var __rv []byte
	DirectQtDrv(nil,180000,180149,unsafe.Pointer(&value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QUrl::toAce(QString const&)
func (q *QUrl) ToAce(value string) []byte {
	var __rv []byte
	q.Drv(180000,180149,unsafe.Pointer(&value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QUrl::toEncoded()
func (q *QUrl) ToEncoded() []byte {
	var __rv []byte
	q.Drv(180000,180150,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QUrl::toEncoded(QFlags<QUrl::FormattingOption>)
func (q *QUrl) ToEncodedWithOptions(options QUrl_FormattingOption) []byte {
	var __rv []byte
	q.Drv(180000,180151,unsafe.Pointer(&options),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QUrl::toLocalFile()
func (q *QUrl) ToLocalFile() string {
	var __rv string
	q.Drv(180000,180152,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QUrl::toPercentEncoding(QString const&)	
func QUrlToPercentEncoding(value string) []byte {
	var __rv []byte
	DirectQtDrv(nil,180000,180153,unsafe.Pointer(&value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QUrl::toPercentEncoding(QString const&)
func (q *QUrl) ToPercentEncoding(value string) []byte {
	var __rv []byte
	q.Drv(180000,180153,unsafe.Pointer(&value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QUrl::toPercentEncoding(QString const&,QByteArray const&,QByteArray const&)	
func QUrlToPercentEncodingWithStringExcludeInclude(value2 string,exclude []byte,include []byte) []byte {
	var __rv []byte
	DirectQtDrv(nil,180000,180154,unsafe.Pointer(&value2),unsafe.Pointer(&exclude),unsafe.Pointer(&include),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QUrl::toPercentEncoding(QString const&,QByteArray const&,QByteArray const&)
func (q *QUrl) ToPercentEncodingWithStringExcludeInclude(value2 string,exclude []byte,include []byte) []byte {
	var __rv []byte
	q.Drv(180000,180154,unsafe.Pointer(&value2),unsafe.Pointer(&exclude),unsafe.Pointer(&include),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QUrl::toString()
func (q *QUrl) ToString() string {
	var __rv string
	q.Drv(180000,180155,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QUrl::toString(QFlags<QUrl::FormattingOption>)
func (q *QUrl) ToStringWithOptions(options QUrl_FormattingOption) string {
	var __rv string
	q.Drv(180000,180156,unsafe.Pointer(&options),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QUrl::userInfo()
func (q *QUrl) UserInfo() string {
	var __rv string
	q.Drv(180000,180157,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QUrl::userName()
func (q *QUrl) UserName() string {
	var __rv string
	q.Drv(180000,180158,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
