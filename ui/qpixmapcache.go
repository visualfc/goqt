// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QPixmapCache : QPixmapCache
type QPixmapCache struct {
	BaseDrv
}
//QPixmapCache::QPixmapCache()	
func NewPixmapCache() *QPixmapCache {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),97000,97102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPixmapCache{}
	_p.SetDriver(__rv,97000,true)
	return _p
} 
//QPixmapCache::cacheLimit()	
func QPixmapCacheCacheLimit() int {
	var __rv int
	DirectQtDrv(nil,97000,97103,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPixmapCache::cacheLimit()
func (q *QPixmapCache) CacheLimit() int {
	var __rv int
	q.Drv(97000,97103,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPixmapCache::clear()	
func QPixmapCacheClear()  {
	DirectQtDrv(nil,97000,97104,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPixmapCache::clear()
func (q *QPixmapCache) Clear()  {
	q.Drv(97000,97104,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPixmapCache::find(QString const&)	
func QPixmapCacheFind(key string) *QPixmap {
	var __rv uintptr
	DirectQtDrv(nil,97000,97105,unsafe.Pointer(&key),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPixmap{}
	_rp.SetDriver(__rv,96000,true)
	return _rp
}	
//QPixmapCache::find(QString const&)
func (q *QPixmapCache) Find(key string) *QPixmap {
	var __rv uintptr
	q.Drv(97000,97105,unsafe.Pointer(&key),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPixmap{}
	_rp.SetDriver(__rv,96000,true)
	return _rp
}	
//QPixmapCache::find(QPixmapCache::Key const&,QPixmap*)	
func QPixmapCacheFindWithPixmapCacheKeyPixmap(key *QPixmapCacheKey,pixmap *QPixmap) bool {
	var __rv bool
	DirectQtDrv(nil,97000,97106,Native(key),Native(pixmap),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPixmapCache::find(QPixmapCache::Key const&,QPixmap*)
func (q *QPixmapCache) FindWithPixmapCacheKeyPixmap(key *QPixmapCacheKey,pixmap *QPixmap) bool {
	var __rv bool
	q.Drv(97000,97106,Native(key),Native(pixmap),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPixmapCache::find(QString const&,QPixmap&)	
func QPixmapCacheFindWithKeyPixmap(key string,pixmap *QPixmap) bool {
	var __rv bool
	DirectQtDrv(nil,97000,97107,unsafe.Pointer(&key),Native(pixmap),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPixmapCache::find(QString const&,QPixmap&)
func (q *QPixmapCache) FindWithKeyPixmap(key string,pixmap *QPixmap) bool {
	var __rv bool
	q.Drv(97000,97107,unsafe.Pointer(&key),Native(pixmap),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPixmapCache::insert(QPixmap const&)	
func QPixmapCacheInsert(pixmap *QPixmap) *QPixmapCacheKey {
	var __rv uintptr
	DirectQtDrv(nil,97000,97108,Native(pixmap),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPixmapCacheKey{}
	_rp.SetDriver(__rv,98000,true)
	return _rp
}	
//QPixmapCache::insert(QPixmap const&)
func (q *QPixmapCache) Insert(pixmap *QPixmap) *QPixmapCacheKey {
	var __rv uintptr
	q.Drv(97000,97108,Native(pixmap),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPixmapCacheKey{}
	_rp.SetDriver(__rv,98000,true)
	return _rp
}	
//QPixmapCache::insert(QString const&,QPixmap const&)	
func QPixmapCacheInsertWithKeyPixmap(key string,pixmap *QPixmap) bool {
	var __rv bool
	DirectQtDrv(nil,97000,97109,unsafe.Pointer(&key),Native(pixmap),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPixmapCache::insert(QString const&,QPixmap const&)
func (q *QPixmapCache) InsertWithKeyPixmap(key string,pixmap *QPixmap) bool {
	var __rv bool
	q.Drv(97000,97109,unsafe.Pointer(&key),Native(pixmap),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPixmapCache::remove(QPixmapCache::Key const&)	
func QPixmapCacheRemove(key *QPixmapCacheKey)  {
	DirectQtDrv(nil,97000,97110,Native(key),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPixmapCache::remove(QPixmapCache::Key const&)
func (q *QPixmapCache) Remove(key *QPixmapCacheKey)  {
	q.Drv(97000,97110,Native(key),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPixmapCache::remove(QString const&)	
func QPixmapCacheRemoveWithKey(key string)  {
	DirectQtDrv(nil,97000,97111,unsafe.Pointer(&key),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPixmapCache::remove(QString const&)
func (q *QPixmapCache) RemoveWithKey(key string)  {
	q.Drv(97000,97111,unsafe.Pointer(&key),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPixmapCache::replace(QPixmapCache::Key const&,QPixmap const&)	
func QPixmapCacheReplace(key *QPixmapCacheKey,pixmap *QPixmap) bool {
	var __rv bool
	DirectQtDrv(nil,97000,97112,Native(key),Native(pixmap),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPixmapCache::replace(QPixmapCache::Key const&,QPixmap const&)
func (q *QPixmapCache) Replace(key *QPixmapCacheKey,pixmap *QPixmap) bool {
	var __rv bool
	q.Drv(97000,97112,Native(key),Native(pixmap),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPixmapCache::setCacheLimit(int)	
func QPixmapCacheSetCacheLimit(value int)  {
	DirectQtDrv(nil,97000,97113,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPixmapCache::setCacheLimit(int)
func (q *QPixmapCache) SetCacheLimit(value int)  {
	q.Drv(97000,97113,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
