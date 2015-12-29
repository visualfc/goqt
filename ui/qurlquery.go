// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QUrlQuery : QUrlQuery
type QUrlQuery struct {
	QObject
}
//QUrlQuery::QUrlQuery()	
func NewUrlQuery() *QUrlQuery {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),391000,391102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QUrlQuery{}
	_p.SetDriver(__rv,391000,false)
	return _p
} 
//QUrlQuery::QUrlQuery(QString const&)	
func NewUrlQueryWithQuerystring(queryString string) *QUrlQuery {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),391000,391103,unsafe.Pointer(&queryString),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QUrlQuery{}
	_p.SetDriver(__rv,391000,false)
	return _p
} 
//QUrlQuery::QUrlQuery(QUrl const&)	
func NewUrlQueryWithUrl(url *QUrl) *QUrlQuery {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),391000,391104,Native(url),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QUrlQuery{}
	_p.SetDriver(__rv,391000,false)
	return _p
} 
//QUrlQuery::QUrlQuery(QUrlQuery const&)	
func NewUrlQueryCopy(other *QUrlQuery) *QUrlQuery {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),391000,391105,Native(other),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QUrlQuery{}
	_p.SetDriver(__rv,391000,false)
	return _p
} 
//QUrlQuery::addQueryItem(QString const&,QString const&)
func (q *QUrlQuery) AddQueryItem(key string,value string)  {
	q.Drv(391000,391106,unsafe.Pointer(&key),unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QUrlQuery::allQueryItemValues(QString const&)
func (q *QUrlQuery) AllQueryItemValues(key string) []string {
	var __rv []string
	q.Drv(391000,391107,unsafe.Pointer(&key),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QUrlQuery::clear()
func (q *QUrlQuery) Clear()  {
	q.Drv(391000,391108,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QUrlQuery::defaultQueryPairDelimiter()	
func QUrlQueryDefaultQueryPairDelimiter() rune {
	var __rv rune
	DirectQtDrv(nil,391000,391109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QUrlQuery::defaultQueryPairDelimiter()
func (q *QUrlQuery) DefaultQueryPairDelimiter() rune {
	var __rv rune
	q.Drv(391000,391109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QUrlQuery::defaultQueryValueDelimiter()	
func QUrlQueryDefaultQueryValueDelimiter() rune {
	var __rv rune
	DirectQtDrv(nil,391000,391110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QUrlQuery::defaultQueryValueDelimiter()
func (q *QUrlQuery) DefaultQueryValueDelimiter() rune {
	var __rv rune
	q.Drv(391000,391110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QUrlQuery::hasQueryItem(QString const&)
func (q *QUrlQuery) HasQueryItem(key string) bool {
	var __rv bool
	q.Drv(391000,391111,unsafe.Pointer(&key),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QUrlQuery::isDetached()
func (q *QUrlQuery) IsDetached() bool {
	var __rv bool
	q.Drv(391000,391112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QUrlQuery::isEmpty()
func (q *QUrlQuery) IsEmpty() bool {
	var __rv bool
	q.Drv(391000,391113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QUrlQuery::query()
func (q *QUrlQuery) Query() string {
	var __rv string
	q.Drv(391000,391114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QUrlQuery::queryItemValue(QString const&)
func (q *QUrlQuery) QueryItemValue(key string) string {
	var __rv string
	q.Drv(391000,391115,unsafe.Pointer(&key),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QUrlQuery::queryPairDelimiter()
func (q *QUrlQuery) QueryPairDelimiter() rune {
	var __rv rune
	q.Drv(391000,391116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QUrlQuery::queryValueDelimiter()
func (q *QUrlQuery) QueryValueDelimiter() rune {
	var __rv rune
	q.Drv(391000,391117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QUrlQuery::removeAllQueryItems(QString const&)
func (q *QUrlQuery) RemoveAllQueryItems(key string)  {
	q.Drv(391000,391118,unsafe.Pointer(&key),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QUrlQuery::removeQueryItem(QString const&)
func (q *QUrlQuery) RemoveQueryItem(key string)  {
	q.Drv(391000,391119,unsafe.Pointer(&key),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QUrlQuery::setQuery(QString const&)
func (q *QUrlQuery) SetQuery(queryString string)  {
	q.Drv(391000,391120,unsafe.Pointer(&queryString),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QUrlQuery::setQueryDelimiters(QChar,QChar)
func (q *QUrlQuery) SetQueryDelimiters(valueDelimiter rune,pairDelimiter rune)  {
	q.Drv(391000,391121,unsafe.Pointer(&valueDelimiter),unsafe.Pointer(&pairDelimiter),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QUrlQuery::swap(QUrlQuery&)
func (q *QUrlQuery) Swap(other *QUrlQuery)  {
	q.Drv(391000,391122,Native(other),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QUrlQuery::toString()
func (q *QUrlQuery) ToString() string {
	var __rv string
	q.Drv(391000,391123,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
