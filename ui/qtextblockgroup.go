// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QTextBlockGroup : QTextBlockGroup
type QTextBlockGroup struct {
	QTextObject
}
//QTextBlockGroup::blockFormatChanged(QTextBlock const&)
func (q *QTextBlockGroup) BlockFormatChanged(block *QTextBlock)  {
	q.Drv(370000,370102,Native(block),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextBlockGroup::blockInserted(QTextBlock const&)
func (q *QTextBlockGroup) BlockInserted(block *QTextBlock)  {
	q.Drv(370000,370103,Native(block),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextBlockGroup::blockList()
func (q *QTextBlockGroup) BlockList() []*QTextBlock {
	var __rv []*QTextBlock
	q.Drv(370000,370104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextBlockGroup::blockRemoved(QTextBlock const&)
func (q *QTextBlockGroup) BlockRemoved(block *QTextBlock)  {
	q.Drv(370000,370105,Native(block),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
