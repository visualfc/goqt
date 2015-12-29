// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QCryptographicHash_Algorithm - QCryptographicHash::Algorithm
type QCryptographicHash_Algorithm uint32
const (
	QCryptographicHash_Md4 QCryptographicHash_Algorithm = 0
	QCryptographicHash_Md5 QCryptographicHash_Algorithm = 1
	QCryptographicHash_Sha1 QCryptographicHash_Algorithm = 2
)
//struct QCryptographicHash : QCryptographicHash
type QCryptographicHash struct {
	BaseDrv
}
//QCryptographicHash::QCryptographicHash(QCryptographicHash::Algorithm)	
func NewCryptographicHash(method QCryptographicHash_Algorithm) *QCryptographicHash {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),17000,17102,unsafe.Pointer(&method),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QCryptographicHash{}
	_p.SetDriver(__rv,17000,true)
	return _p
} 
//QCryptographicHash::addData(QByteArray const&)
func (q *QCryptographicHash) AddData(data []byte)  {
	q.Drv(17000,17103,unsafe.Pointer(&data),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCryptographicHash::hash(QByteArray const&,QCryptographicHash::Algorithm)	
func QCryptographicHashHash(data []byte,method QCryptographicHash_Algorithm) []byte {
	var __rv []byte
	DirectQtDrv(nil,17000,17104,unsafe.Pointer(&data),unsafe.Pointer(&method),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QCryptographicHash::hash(QByteArray const&,QCryptographicHash::Algorithm)
func (q *QCryptographicHash) Hash(data []byte,method QCryptographicHash_Algorithm) []byte {
	var __rv []byte
	q.Drv(17000,17104,unsafe.Pointer(&data),unsafe.Pointer(&method),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QCryptographicHash::reset()
func (q *QCryptographicHash) Reset()  {
	q.Drv(17000,17105,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCryptographicHash::result()
func (q *QCryptographicHash) Result() []byte {
	var __rv []byte
	q.Drv(17000,17106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
