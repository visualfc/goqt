// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui


//struct QAbstractUndoItem : QAbstractUndoItem
type QAbstractUndoItem struct {
	BaseDrv
}
//QAbstractUndoItem::redo()
func (q *QAbstractUndoItem) Redo()  {
	q.Drv(4000,4102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractUndoItem::undo()
func (q *QAbstractUndoItem) Undo()  {
	q.Drv(4000,4103,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
