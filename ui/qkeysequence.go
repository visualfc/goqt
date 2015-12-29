// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QKeySequence_SequenceFormat - QKeySequence::SequenceFormat
type QKeySequence_SequenceFormat uint32
const (
	QKeySequence_NativeText QKeySequence_SequenceFormat = 0
	QKeySequence_PortableText QKeySequence_SequenceFormat = 1
)
//enum QKeySequence_StandardKey - QKeySequence::StandardKey
type QKeySequence_StandardKey uint32
const (
	QKeySequence_UnknownKey QKeySequence_StandardKey = 0
	QKeySequence_HelpContents QKeySequence_StandardKey = 1
	QKeySequence_WhatsThis QKeySequence_StandardKey = 2
	QKeySequence_Open QKeySequence_StandardKey = 3
	QKeySequence_Close QKeySequence_StandardKey = 4
	QKeySequence_Save QKeySequence_StandardKey = 5
	QKeySequence_New QKeySequence_StandardKey = 6
	QKeySequence_Delete QKeySequence_StandardKey = 7
	QKeySequence_Cut QKeySequence_StandardKey = 8
	QKeySequence_Copy QKeySequence_StandardKey = 9
	QKeySequence_Paste QKeySequence_StandardKey = 10
	QKeySequence_Undo QKeySequence_StandardKey = 11
	QKeySequence_Redo QKeySequence_StandardKey = 12
	QKeySequence_Back QKeySequence_StandardKey = 13
	QKeySequence_Forward QKeySequence_StandardKey = 14
	QKeySequence_Refresh QKeySequence_StandardKey = 15
	QKeySequence_ZoomIn QKeySequence_StandardKey = 16
	QKeySequence_ZoomOut QKeySequence_StandardKey = 17
	QKeySequence_Print QKeySequence_StandardKey = 18
	QKeySequence_AddTab QKeySequence_StandardKey = 19
	QKeySequence_NextChild QKeySequence_StandardKey = 20
	QKeySequence_PreviousChild QKeySequence_StandardKey = 21
	QKeySequence_Find QKeySequence_StandardKey = 22
	QKeySequence_FindNext QKeySequence_StandardKey = 23
	QKeySequence_FindPrevious QKeySequence_StandardKey = 24
	QKeySequence_Replace QKeySequence_StandardKey = 25
	QKeySequence_SelectAll QKeySequence_StandardKey = 26
	QKeySequence_Bold QKeySequence_StandardKey = 27
	QKeySequence_Italic QKeySequence_StandardKey = 28
	QKeySequence_Underline QKeySequence_StandardKey = 29
	QKeySequence_MoveToNextChar QKeySequence_StandardKey = 30
	QKeySequence_MoveToPreviousChar QKeySequence_StandardKey = 31
	QKeySequence_MoveToNextWord QKeySequence_StandardKey = 32
	QKeySequence_MoveToPreviousWord QKeySequence_StandardKey = 33
	QKeySequence_MoveToNextLine QKeySequence_StandardKey = 34
	QKeySequence_MoveToPreviousLine QKeySequence_StandardKey = 35
	QKeySequence_MoveToNextPage QKeySequence_StandardKey = 36
	QKeySequence_MoveToPreviousPage QKeySequence_StandardKey = 37
	QKeySequence_MoveToStartOfLine QKeySequence_StandardKey = 38
	QKeySequence_MoveToEndOfLine QKeySequence_StandardKey = 39
	QKeySequence_MoveToStartOfBlock QKeySequence_StandardKey = 40
	QKeySequence_MoveToEndOfBlock QKeySequence_StandardKey = 41
	QKeySequence_MoveToStartOfDocument QKeySequence_StandardKey = 42
	QKeySequence_MoveToEndOfDocument QKeySequence_StandardKey = 43
	QKeySequence_SelectNextChar QKeySequence_StandardKey = 44
	QKeySequence_SelectPreviousChar QKeySequence_StandardKey = 45
	QKeySequence_SelectNextWord QKeySequence_StandardKey = 46
	QKeySequence_SelectPreviousWord QKeySequence_StandardKey = 47
	QKeySequence_SelectNextLine QKeySequence_StandardKey = 48
	QKeySequence_SelectPreviousLine QKeySequence_StandardKey = 49
	QKeySequence_SelectNextPage QKeySequence_StandardKey = 50
	QKeySequence_SelectPreviousPage QKeySequence_StandardKey = 51
	QKeySequence_SelectStartOfLine QKeySequence_StandardKey = 52
	QKeySequence_SelectEndOfLine QKeySequence_StandardKey = 53
	QKeySequence_SelectStartOfBlock QKeySequence_StandardKey = 54
	QKeySequence_SelectEndOfBlock QKeySequence_StandardKey = 55
	QKeySequence_SelectStartOfDocument QKeySequence_StandardKey = 56
	QKeySequence_SelectEndOfDocument QKeySequence_StandardKey = 57
	QKeySequence_DeleteStartOfWord QKeySequence_StandardKey = 58
	QKeySequence_DeleteEndOfWord QKeySequence_StandardKey = 59
	QKeySequence_DeleteEndOfLine QKeySequence_StandardKey = 60
	QKeySequence_InsertParagraphSeparator QKeySequence_StandardKey = 61
	QKeySequence_InsertLineSeparator QKeySequence_StandardKey = 62
	QKeySequence_SaveAs QKeySequence_StandardKey = 63
	QKeySequence_Preferences QKeySequence_StandardKey = 64
	QKeySequence_Quit QKeySequence_StandardKey = 65
)
//enum QKeySequence_SequenceMatch - QKeySequence::SequenceMatch
type QKeySequence_SequenceMatch uint32
const (
	QKeySequence_NoMatch QKeySequence_SequenceMatch = 0
	QKeySequence_PartialMatch QKeySequence_SequenceMatch = 1
	QKeySequence_ExactMatch QKeySequence_SequenceMatch = 2
)
//struct QKeySequence : QKeySequence
type QKeySequence struct {
	BaseDrv
}
//QKeySequence::QKeySequence()	
func NewKeySequence() *QKeySequence {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),65000,65102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QKeySequence{}
	_p.SetDriver(__rv,65000,true)
	return _p
} 
//QKeySequence::QKeySequence(QKeySequence const&)	
func NewKeySequenceCopy(ks *QKeySequence) *QKeySequence {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),65000,65103,Native(ks),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QKeySequence{}
	_p.SetDriver(__rv,65000,true)
	return _p
} 
//QKeySequence::QKeySequence(QKeySequence::StandardKey)	
func NewKeySequenceWithStandardkey(key QKeySequence_StandardKey) *QKeySequence {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),65000,65104,unsafe.Pointer(&key),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QKeySequence{}
	_p.SetDriver(__rv,65000,true)
	return _p
} 
//QKeySequence::QKeySequence(QString const&)	
func NewKeySequenceWithKey(key string) *QKeySequence {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),65000,65105,unsafe.Pointer(&key),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QKeySequence{}
	_p.SetDriver(__rv,65000,true)
	return _p
} 
//QKeySequence::QKeySequence(QString const&,QKeySequence::SequenceFormat)	
func NewKeySequenceWithKeyFormat(key string,format QKeySequence_SequenceFormat) *QKeySequence {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),65000,65106,unsafe.Pointer(&key),unsafe.Pointer(&format),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QKeySequence{}
	_p.SetDriver(__rv,65000,true)
	return _p
} 
//QKeySequence::QKeySequence(int,int,int,int)	
func NewKeySequenceWithK1K2K3K4(k1 int,k2 int,k3 int,k4 int) *QKeySequence {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),65000,65107,unsafe.Pointer(&k1),unsafe.Pointer(&k2),unsafe.Pointer(&k3),unsafe.Pointer(&k4),nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QKeySequence{}
	_p.SetDriver(__rv,65000,true)
	return _p
} 
//QKeySequence::count()
func (q *QKeySequence) Count() uint {
	var __rv uint
	q.Drv(65000,65108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QKeySequence::fromString(QString const&)	
func QKeySequenceFromString(str string) *QKeySequence {
	var __rv uintptr
	DirectQtDrv(nil,65000,65109,unsafe.Pointer(&str),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QKeySequence{}
	_rp.SetDriver(__rv,65000,true)
	return _rp
}	
//QKeySequence::fromString(QString const&)
func (q *QKeySequence) FromString(str string) *QKeySequence {
	var __rv uintptr
	q.Drv(65000,65109,unsafe.Pointer(&str),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QKeySequence{}
	_rp.SetDriver(__rv,65000,true)
	return _rp
}	
//QKeySequence::fromString(QString const&,QKeySequence::SequenceFormat)	
func QKeySequenceFromStringWithTextFormat(str string,format QKeySequence_SequenceFormat) *QKeySequence {
	var __rv uintptr
	DirectQtDrv(nil,65000,65110,unsafe.Pointer(&str),unsafe.Pointer(&format),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QKeySequence{}
	_rp.SetDriver(__rv,65000,true)
	return _rp
}	
//QKeySequence::fromString(QString const&,QKeySequence::SequenceFormat)
func (q *QKeySequence) FromStringWithTextFormat(str string,format QKeySequence_SequenceFormat) *QKeySequence {
	var __rv uintptr
	q.Drv(65000,65110,unsafe.Pointer(&str),unsafe.Pointer(&format),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QKeySequence{}
	_rp.SetDriver(__rv,65000,true)
	return _rp
}	
//QKeySequence::isDetached()
func (q *QKeySequence) IsDetached() bool {
	var __rv bool
	q.Drv(65000,65111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QKeySequence::isEmpty()
func (q *QKeySequence) IsEmpty() bool {
	var __rv bool
	q.Drv(65000,65112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QKeySequence::keyBindings(QKeySequence::StandardKey)	
func QKeySequenceKeyBindings(key QKeySequence_StandardKey) []*QKeySequence {
	var __rv []*QKeySequence
	DirectQtDrv(nil,65000,65113,unsafe.Pointer(&key),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QKeySequence::keyBindings(QKeySequence::StandardKey)
func (q *QKeySequence) KeyBindings(key QKeySequence_StandardKey) []*QKeySequence {
	var __rv []*QKeySequence
	q.Drv(65000,65113,unsafe.Pointer(&key),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QKeySequence::matches(QKeySequence const&)
func (q *QKeySequence) Matches(seq *QKeySequence) QKeySequence_SequenceMatch {
	var __rv QKeySequence_SequenceMatch
	q.Drv(65000,65114,Native(seq),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QKeySequence::mnemonic(QString const&)	
func QKeySequenceMnemonic(text string) *QKeySequence {
	var __rv uintptr
	DirectQtDrv(nil,65000,65115,unsafe.Pointer(&text),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QKeySequence{}
	_rp.SetDriver(__rv,65000,true)
	return _rp
}	
//QKeySequence::mnemonic(QString const&)
func (q *QKeySequence) Mnemonic(text string) *QKeySequence {
	var __rv uintptr
	q.Drv(65000,65115,unsafe.Pointer(&text),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QKeySequence{}
	_rp.SetDriver(__rv,65000,true)
	return _rp
}	
//QKeySequence::toString()
func (q *QKeySequence) ToString() string {
	var __rv string
	q.Drv(65000,65116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QKeySequence::toString(QKeySequence::SequenceFormat)
func (q *QKeySequence) ToStringWithFormat(format QKeySequence_SequenceFormat) string {
	var __rv string
	q.Drv(65000,65117,unsafe.Pointer(&format),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
