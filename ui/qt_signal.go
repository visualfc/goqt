// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

func drvSignalCall(fn interface{},id int32, v1,v2,v3,v4 unsafe.Pointer) error{
	switch id {
	case 1:
	fn.(func())()
	case 2:
	_v1 := &QAbstractAnimation{}
	_v1.SetDriver(uintptr(v1),192000,false)
	fn.(func(*QAbstractAnimation))(_v1)
	case 3:
	fn.(func(QAbstractAnimation_Direction))(*(*QAbstractAnimation_Direction)(v1))
	case 4:
	fn.(func(QAbstractAnimation_State,QAbstractAnimation_State))(*(*QAbstractAnimation_State)(v1),*(*QAbstractAnimation_State)(v2))
	case 5:
	_v1 := &QAbstractButton{}
	_v1.SetDriver(uintptr(v1),193000,false)
	fn.(func(*QAbstractButton))(_v1)
	case 6:
	_v1 := &QAction{}
	_v1.SetDriver(uintptr(v1),207000,false)
	fn.(func(*QAction))(_v1)
	case 7:
	_v1 := &QBrush{}
	_v1.SetDriver(uintptr(v1),9000,false)
	fn.(func(*QBrush))(_v1)
	case 8:
	fn.(func(QClipboard_Mode))(*(*QClipboard_Mode)(v1))
	case 9:
	_v1 := &QColor{}
	_v1.SetDriver(uintptr(v1),13000,false)
	fn.(func(*QColor))(_v1)
	case 10:
	_v1 := &QDate{}
	_v1.SetDriver(uintptr(v1),19000,false)
	fn.(func(*QDate))(_v1)
	case 11:
	_v1 := &QDateTime{}
	_v1.SetDriver(uintptr(v1),20000,false)
	fn.(func(*QDateTime))(_v1)
	case 12:
	fn.(func(QDockWidget_DockWidgetFeature))(*(*QDockWidget_DockWidgetFeature)(v1))
	case 13:
	fn.(func(QGraphicsBlurEffect_BlurHint))(*(*QGraphicsBlurEffect_BlurHint)(v1))
	case 14:
	fn.(func(Qt_DockWidgetArea))(*(*Qt_DockWidgetArea)(v1))
	case 15:
	fn.(func(Qt_ToolBarArea))(*(*Qt_ToolBarArea)(v1))
	case 16:
	fn.(func(Qt_WindowState,Qt_WindowState))(*(*Qt_WindowState)(v1),*(*Qt_WindowState)(v2))
	case 17:
	_v1 := &QFont{}
	_v1.SetDriver(uintptr(v1),37000,false)
	fn.(func(*QFont))(_v1)
	case 18:
	fn.(func(QImageReader_ImageReaderError))(*(*QImageReader_ImageReaderError)(v1))
	case 19:
	_v1 := &QItemSelection{}
	_v1.SetDriver(uintptr(v1),62000,false)
	_v2 := &QItemSelection{}
	_v2.SetDriver(uintptr(v2),62000,false)
	fn.(func(*QItemSelection,*QItemSelection))(_v1,_v2)
	case 20:
	var _v1 []*QModelIndex
		drvGetModelIndexArray(unsafe.Pointer(&_v1),v1,79000)
	fn.(func([]*QModelIndex))(_v1)
	case 21:
	var _v1 []*QRectF
		drvGetRectFArray(unsafe.Pointer(&_v1),v1,111000)
	fn.(func([]*QRectF))(_v1)
	case 22:
	_v1 := &QListWidgetItem{}
	_v1.SetDriver(uintptr(v1),71000,false)
	fn.(func(*QListWidgetItem))(_v1)
	case 23:
	_v1 := &QListWidgetItem{}
	_v1.SetDriver(uintptr(v1),71000,false)
	_v2 := &QListWidgetItem{}
	_v2.SetDriver(uintptr(v2),71000,false)
	fn.(func(*QListWidgetItem,*QListWidgetItem))(_v1,_v2)
	case 24:
	_v1 := &QMdiSubWindow{}
	_v1.SetDriver(uintptr(v1),307000,false)
	fn.(func(*QMdiSubWindow))(_v1)
	case 25:
	_v1 := &QModelIndex{}
	_v1.SetDriver(uintptr(v1),79000,false)
	fn.(func(*QModelIndex))(_v1)
	case 26:
	_v1 := &QModelIndex{}
	_v1.SetDriver(uintptr(v1),79000,false)
	_v2 := &QModelIndex{}
	_v2.SetDriver(uintptr(v2),79000,false)
	fn.(func(*QModelIndex,*QModelIndex))(_v1,_v2)
	case 27:
	fn.(func(QMovie_MovieState))(*(*QMovie_MovieState)(v1))
	case 28:
	_v1 := &QObject{}
	_v1.SetDriver(uintptr(v1),314000,false)
	fn.(func(*QObject))(_v1)
	case 29:
	_v1 := &QPoint{}
	_v1.SetDriver(uintptr(v1),99000,false)
	fn.(func(*QPoint))(_v1)
	case 30:
	_v1 := &QPointF{}
	_v1.SetDriver(uintptr(v1),100000,false)
	fn.(func(*QPointF))(_v1)
	case 31:
	_v1 := &QPrinter{}
	_v1.SetDriver(uintptr(v1),104000,false)
	fn.(func(*QPrinter))(_v1)
	case 32:
	fn.(func(QProcess_ProcessError))(*(*QProcess_ProcessError)(v1))
	case 33:
	fn.(func(QProcess_ProcessState))(*(*QProcess_ProcessState)(v1))
	case 34:
	_v1 := &QRect{}
	_v1.SetDriver(uintptr(v1),110000,false)
	fn.(func(*QRect))(_v1)
	case 35:
	_v1 := &QRect{}
	_v1.SetDriver(uintptr(v1),110000,false)
	fn.(func(*QRect,int32))(_v1,*(*int32)(v2))
	case 36:
	_v1 := &QRectF{}
	_v1.SetDriver(uintptr(v1),111000,false)
	fn.(func(*QRectF))(_v1)
	case 37:
	_v1 := &QSessionManager{}
	_v1.SetDriver(uintptr(v1),338000,false)
	fn.(func(*QSessionManager))(_v1)
	case 38:
	_v1 := &QSize{}
	_v1.SetDriver(uintptr(v1),119000,false)
	fn.(func(*QSize))(_v1)
	case 39:
	_v1 := &QSizeF{}
	_v1.SetDriver(uintptr(v1),120000,false)
	fn.(func(*QSizeF))(_v1)
	case 40:
	_v1 := &QStandardItem{}
	_v1.SetDriver(uintptr(v1),123000,false)
	fn.(func(*QStandardItem))(_v1)
	case 41:
	var _v1 string
		drvGetString(unsafe.Pointer(&_v1),v1)
	fn.(func(string))(_v1)
	case 42:
	var _v1 string
		drvGetString(unsafe.Pointer(&_v1),v1)
var _v2 string
		drvGetString(unsafe.Pointer(&_v2),v2)
var _v3 string
		drvGetString(unsafe.Pointer(&_v3),v3)
	fn.(func(string,string,string))(_v1,_v2,_v3)
	case 43:
	var _v1 []string
		drvGetStringArray(unsafe.Pointer(&_v1),v1)
	fn.(func([]string))(_v1)
	case 44:
	fn.(func(QSystemTrayIcon_ActivationReason))(*(*QSystemTrayIcon_ActivationReason)(v1))
	case 45:
	_v1 := &QTableWidgetItem{}
	_v1.SetDriver(uintptr(v1),134000,false)
	fn.(func(*QTableWidgetItem))(_v1)
	case 46:
	_v1 := &QTableWidgetItem{}
	_v1.SetDriver(uintptr(v1),134000,false)
	_v2 := &QTableWidgetItem{}
	_v2.SetDriver(uintptr(v2),134000,false)
	fn.(func(*QTableWidgetItem,*QTableWidgetItem))(_v1,_v2)
	case 47:
	_v1 := &QTextBlock{}
	_v1.SetDriver(uintptr(v1),137000,false)
	fn.(func(*QTextBlock))(_v1)
	case 48:
	_v1 := &QTextCharFormat{}
	_v1.SetDriver(uintptr(v1),142000,false)
	fn.(func(*QTextCharFormat))(_v1)
	case 49:
	_v1 := &QTextCursor{}
	_v1.SetDriver(uintptr(v1),145000,false)
	fn.(func(*QTextCursor))(_v1)
	case 50:
	_v1 := &QTime{}
	_v1.SetDriver(uintptr(v1),170000,false)
	fn.(func(*QTime))(_v1)
	case 51:
	fn.(func(QTimeLine_State))(*(*QTimeLine_State)(v1))
	case 52:
	_v1 := &QTreeWidgetItem{}
	_v1.SetDriver(uintptr(v1),177000,false)
	fn.(func(*QTreeWidgetItem))(_v1)
	case 53:
	_v1 := &QTreeWidgetItem{}
	_v1.SetDriver(uintptr(v1),177000,false)
	_v2 := &QTreeWidgetItem{}
	_v2.SetDriver(uintptr(v2),177000,false)
	fn.(func(*QTreeWidgetItem,*QTreeWidgetItem))(_v1,_v2)
	case 54:
	_v1 := &QTreeWidgetItem{}
	_v1.SetDriver(uintptr(v1),177000,false)
	fn.(func(*QTreeWidgetItem,int32))(_v1,*(*int32)(v2))
	case 55:
	_v1 := &QUndoStack{}
	_v1.SetDriver(uintptr(v1),389000,false)
	fn.(func(*QUndoStack))(_v1)
	case 56:
	_v1 := &QUrl{}
	_v1.SetDriver(uintptr(v1),180000,false)
	fn.(func(*QUrl))(_v1)
	case 57:
	_v1 := &QVariant{}
	_v1.SetDriver(uintptr(v1),182000,false)
	fn.(func(*QVariant))(_v1)
	case 58:
	_v1 := &QWidget{}
	_v1.SetDriver(uintptr(v1),395000,false)
	fn.(func(*QWidget))(_v1)
	case 59:
	_v1 := &QWidget{}
	_v1.SetDriver(uintptr(v1),395000,false)
	fn.(func(*QWidget,QAbstractItemDelegate_EndEditHint))(_v1,*(*QAbstractItemDelegate_EndEditHint)(v2))
	case 60:
	_v1 := &QWidget{}
	_v1.SetDriver(uintptr(v1),395000,false)
	_v2 := &QWidget{}
	_v2.SetDriver(uintptr(v2),395000,false)
	fn.(func(*QWidget,*QWidget))(_v1,_v2)
	case 61:
	fn.(func(Qt_DockWidgetArea))(*(*Qt_DockWidgetArea)(v1))
	case 62:
	fn.(func(Qt_DropAction))(*(*Qt_DropAction)(v1))
	case 63:
	fn.(func(Qt_Orientation))(*(*Qt_Orientation)(v1))
	case 64:
	fn.(func(Qt_Orientation,int32,int32))(*(*Qt_Orientation)(v1),*(*int32)(v2),*(*int32)(v3))
	case 65:
	fn.(func(Qt_ToolButtonStyle))(*(*Qt_ToolButtonStyle)(v1))
	case 66:
	fn.(func(bool))(*(*bool)(v1))
	case 67:
	fn.(func(float64))(*(*float64)(v1))
	case 68:
	fn.(func(int32))(*(*int32)(v1))
	case 69:
	fn.(func(int32,QHeaderView_ResizeMode))(*(*int32)(v1),*(*QHeaderView_ResizeMode)(v2))
	case 70:
	fn.(func(int32,QProcess_ExitStatus))(*(*int32)(v1),*(*QProcess_ExitStatus)(v2))
	case 71:
	fn.(func(int32,Qt_SortOrder))(*(*int32)(v1),*(*Qt_SortOrder)(v2))
	case 72:
	fn.(func(int32,int32))(*(*int32)(v1),*(*int32)(v2))
	case 73:
	fn.(func(int32,int32,int32))(*(*int32)(v1),*(*int32)(v2),*(*int32)(v3))
	case 74:
	fn.(func(int32,int32,int32,int32))(*(*int32)(v1),*(*int32)(v2),*(*int32)(v3),*(*int32)(v4))
	case 75:
	fn.(func(int64))(*(*int64)(v1))
	default:
		return ErrInvalid
	}
	return nil
}	
