// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QPixmapCacheKey : QPixmapCache::Key
type QPixmapCacheKey struct {
	BaseDrv
}
//QPixmapCache::Key::Key()	
func NewPixmapCacheKey() *QPixmapCacheKey {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),98000,98102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPixmapCacheKey{}
	_p.SetDriver(__rv,98000,true)
	return _p
} 
//QPixmapCache::Key::Key(QPixmapCache::Key const&)	
func NewPixmapCacheKeyCopy(other *QPixmapCacheKey) *QPixmapCacheKey {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),98000,98103,Native(other),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPixmapCacheKey{}
	_p.SetDriver(__rv,98000,true)
	return _p
} 
