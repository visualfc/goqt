// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QVariant_Type - QVariant::Type
type QVariant_Type uint32
const (
	QVariant_Invalid QVariant_Type = 0
	QVariant_Bool QVariant_Type = 1
	QVariant_Int QVariant_Type = 2
	QVariant_UInt QVariant_Type = 3
	QVariant_LongLong QVariant_Type = 4
	QVariant_ULongLong QVariant_Type = 5
	QVariant_Double QVariant_Type = 6
	QVariant_Char QVariant_Type = 7
	QVariant_Map QVariant_Type = 8
	QVariant_List QVariant_Type = 9
	QVariant_String QVariant_Type = 10
	QVariant_StringList QVariant_Type = 11
	QVariant_ByteArray QVariant_Type = 12
	QVariant_BitArray QVariant_Type = 13
	QVariant_Date QVariant_Type = 14
	QVariant_Time QVariant_Type = 15
	QVariant_DateTime QVariant_Type = 16
	QVariant_Url QVariant_Type = 17
	QVariant_Locale QVariant_Type = 18
	QVariant_Rect QVariant_Type = 19
	QVariant_RectF QVariant_Type = 20
	QVariant_Size QVariant_Type = 21
	QVariant_SizeF QVariant_Type = 22
	QVariant_Line QVariant_Type = 23
	QVariant_LineF QVariant_Type = 24
	QVariant_Point QVariant_Type = 25
	QVariant_PointF QVariant_Type = 26
	QVariant_RegExp QVariant_Type = 27
	QVariant_Hash QVariant_Type = 28
	QVariant_EasingCurve QVariant_Type = 29
	QVariant_LastCoreType QVariant_Type = QVariant_EasingCurve
	QVariant_Font QVariant_Type = 64
	QVariant_Pixmap QVariant_Type = 65
	QVariant_Brush QVariant_Type = 66
	QVariant_Color QVariant_Type = 67
	QVariant_Palette QVariant_Type = 68
	QVariant_Icon QVariant_Type = 69
	QVariant_Image QVariant_Type = 70
	QVariant_Polygon QVariant_Type = 71
	QVariant_Region QVariant_Type = 72
	QVariant_Bitmap QVariant_Type = 73
	QVariant_Cursor QVariant_Type = 74
	QVariant_SizePolicy QVariant_Type = 75
	QVariant_KeySequence QVariant_Type = 76
	QVariant_Pen QVariant_Type = 77
	QVariant_TextLength QVariant_Type = 78
	QVariant_TextFormat QVariant_Type = 79
	QVariant_Matrix QVariant_Type = 80
	QVariant_Transform QVariant_Type = 81
	QVariant_Matrix4x4 QVariant_Type = 82
	QVariant_Vector2D QVariant_Type = 83
	QVariant_Vector3D QVariant_Type = 84
	QVariant_Vector4D QVariant_Type = 85
	QVariant_Quaternion QVariant_Type = 86
	QVariant_LastGuiType QVariant_Type = QVariant_Quaternion
	QVariant_UserType QVariant_Type = 127
	QVariant_LastType QVariant_Type = 0xffffffff
)
//struct QVariant : QVariant
type QVariant struct {
	BaseDrv
}
//QVariant::QVariant()	
func NewVariant() *QVariant {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),182000,182102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QVariant{}
	_p.SetDriver(__rv,182000,true)
	return _p
} 
//QVariant::QVariant(QBitArray const&)	
func NewVariantWithBoolArray(bitarray []bool) *QVariant {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),182000,182103,unsafe.Pointer(&bitarray),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QVariant{}
	_p.SetDriver(__rv,182000,true)
	return _p
} 
//QVariant::QVariant(QByteArray const&)	
func NewVariantWithByteArray(bytearray []byte) *QVariant {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),182000,182104,unsafe.Pointer(&bytearray),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QVariant{}
	_p.SetDriver(__rv,182000,true)
	return _p
} 
//QVariant::QVariant(QChar const&)	
func NewVariantWithRune(qchar rune) *QVariant {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),182000,182105,unsafe.Pointer(&qchar),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QVariant{}
	_p.SetDriver(__rv,182000,true)
	return _p
} 
//QVariant::QVariant(QDate const&)	
func NewVariantWithDate(date *QDate) *QVariant {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),182000,182106,Native(date),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QVariant{}
	_p.SetDriver(__rv,182000,true)
	return _p
} 
//QVariant::QVariant(QDateTime const&)	
func NewVariantWithDateTime(datetime *QDateTime) *QVariant {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),182000,182107,Native(datetime),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QVariant{}
	_p.SetDriver(__rv,182000,true)
	return _p
} 
//QVariant::QVariant(QEasingCurve const&)	
func NewVariantWithEasingCurve(easing *QEasingCurve) *QVariant {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),182000,182108,Native(easing),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QVariant{}
	_p.SetDriver(__rv,182000,true)
	return _p
} 
//QVariant::QVariant(QLine const&)	
func NewVariantWithLine(line *QLine) *QVariant {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),182000,182109,Native(line),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QVariant{}
	_p.SetDriver(__rv,182000,true)
	return _p
} 
//QVariant::QVariant(QLineF const&)	
func NewVariantWithLineF(line *QLineF) *QVariant {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),182000,182110,Native(line),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QVariant{}
	_p.SetDriver(__rv,182000,true)
	return _p
} 
//QVariant::QVariant(QList<QVariant> const&)	
func NewVariantWithVariantArray(list []*QVariant) *QVariant {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),182000,182111,unsafe.Pointer(&list),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QVariant{}
	_p.SetDriver(__rv,182000,true)
	return _p
} 
//QVariant::QVariant(QLocale const&)	
func NewVariantWithLocale(locale *QLocale) *QVariant {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),182000,182112,Native(locale),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QVariant{}
	_p.SetDriver(__rv,182000,true)
	return _p
} 
//QVariant::QVariant(QMap<QString,QVariant> const&)	
func NewVariantWithMap(_map map[string]*QVariant) *QVariant {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),182000,182113,unsafe.Pointer(&_map),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QVariant{}
	_p.SetDriver(__rv,182000,true)
	return _p
} 
//QVariant::QVariant(QPoint const&)	
func NewVariantWithPoint(pt *QPoint) *QVariant {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),182000,182114,Native(pt),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QVariant{}
	_p.SetDriver(__rv,182000,true)
	return _p
} 
//QVariant::QVariant(QPointF const&)	
func NewVariantWithPointF(pt *QPointF) *QVariant {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),182000,182115,Native(pt),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QVariant{}
	_p.SetDriver(__rv,182000,true)
	return _p
} 
//QVariant::QVariant(QRect const&)	
func NewVariantWithRect(rect *QRect) *QVariant {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),182000,182116,Native(rect),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QVariant{}
	_p.SetDriver(__rv,182000,true)
	return _p
} 
//QVariant::QVariant(QRectF const&)	
func NewVariantWithRectF(rect *QRectF) *QVariant {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),182000,182117,Native(rect),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QVariant{}
	_p.SetDriver(__rv,182000,true)
	return _p
} 
//QVariant::QVariant(QRegExp const&)	
func NewVariantWithRegExp(regExp *QRegExp) *QVariant {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),182000,182118,Native(regExp),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QVariant{}
	_p.SetDriver(__rv,182000,true)
	return _p
} 
//QVariant::QVariant(QSize const&)	
func NewVariantWithSize(size *QSize) *QVariant {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),182000,182119,Native(size),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QVariant{}
	_p.SetDriver(__rv,182000,true)
	return _p
} 
//QVariant::QVariant(QSizeF const&)	
func NewVariantWithSizeF(size *QSizeF) *QVariant {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),182000,182120,Native(size),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QVariant{}
	_p.SetDriver(__rv,182000,true)
	return _p
} 
//QVariant::QVariant(QString const&)	
func NewVariantWithString(string string) *QVariant {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),182000,182121,unsafe.Pointer(&string),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QVariant{}
	_p.SetDriver(__rv,182000,true)
	return _p
} 
//QVariant::QVariant(QStringList const&)	
func NewVariantWithStringArray(stringlist []string) *QVariant {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),182000,182122,unsafe.Pointer(&stringlist),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QVariant{}
	_p.SetDriver(__rv,182000,true)
	return _p
} 
//QVariant::QVariant(QTime const&)	
func NewVariantWithTime(time *QTime) *QVariant {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),182000,182123,Native(time),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QVariant{}
	_p.SetDriver(__rv,182000,true)
	return _p
} 
//QVariant::QVariant(QUrl const&)	
func NewVariantWithUrl(url *QUrl) *QVariant {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),182000,182124,Native(url),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QVariant{}
	_p.SetDriver(__rv,182000,true)
	return _p
} 
//QVariant::QVariant(QVariant const&)	
func NewVariantCopy(other *QVariant) *QVariant {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),182000,182125,Native(other),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QVariant{}
	_p.SetDriver(__rv,182000,true)
	return _p
} 
//QVariant::QVariant(QVariant::Type)	
func NewVariantWithType(_type QVariant_Type) *QVariant {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),182000,182126,unsafe.Pointer(&_type),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QVariant{}
	_p.SetDriver(__rv,182000,true)
	return _p
} 
//QVariant::QVariant(bool)	
func NewVariantWithBool(b bool) *QVariant {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),182000,182127,unsafe.Pointer(&b),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QVariant{}
	_p.SetDriver(__rv,182000,true)
	return _p
} 
//QVariant::QVariant(double)	
func NewVariantWithFloat64(d float64) *QVariant {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),182000,182128,unsafe.Pointer(&d),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QVariant{}
	_p.SetDriver(__rv,182000,true)
	return _p
} 
//QVariant::QVariant(float)	
func NewVariantWithFloat32(f float32) *QVariant {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),182000,182129,unsafe.Pointer(&f),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QVariant{}
	_p.SetDriver(__rv,182000,true)
	return _p
} 
//QVariant::QVariant(int)	
func NewVariantWithInt(i int) *QVariant {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),182000,182130,unsafe.Pointer(&i),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QVariant{}
	_p.SetDriver(__rv,182000,true)
	return _p
} 
//QVariant::QVariant(qint64)	
func NewVariantWithInt64(ll int64) *QVariant {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),182000,182131,unsafe.Pointer(&ll),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QVariant{}
	_p.SetDriver(__rv,182000,true)
	return _p
} 
//QVariant::QVariant(unsigned int)	
func NewVariantWithUint(ui uint) *QVariant {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),182000,182132,unsafe.Pointer(&ui),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QVariant{}
	_p.SetDriver(__rv,182000,true)
	return _p
} 
//QVariant::QVariant(int,void const*)	
func NewVariantWithIntUintptr(typeOrUserType int,copy uintptr) *QVariant {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),182000,182133,unsafe.Pointer(&typeOrUserType),unsafe.Pointer(&copy),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QVariant{}
	_p.SetDriver(__rv,182000,true)
	return _p
} 
//QVariant::QVariant(int,void const*,unsigned int)	
func NewVariantWithIntUintptrUint(typeOrUserType int,copy uintptr,flags uint) *QVariant {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),182000,182134,unsafe.Pointer(&typeOrUserType),unsafe.Pointer(&copy),unsafe.Pointer(&flags),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QVariant{}
	_p.SetDriver(__rv,182000,true)
	return _p
} 
//QVariant::canConvert(QVariant::Type)
func (q *QVariant) CanConvert(t QVariant_Type) bool {
	var __rv bool
	q.Drv(182000,182135,unsafe.Pointer(&t),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QVariant::clear()
func (q *QVariant) Clear()  {
	q.Drv(182000,182136,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QVariant::constData()
func (q *QVariant) ConstData() uintptr {
	var __rv uintptr
	q.Drv(182000,182137,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QVariant::convert(QVariant::Type)
func (q *QVariant) Convert(t QVariant_Type) bool {
	var __rv bool
	q.Drv(182000,182138,unsafe.Pointer(&t),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QVariant::data()
func (q *QVariant) Data() uintptr {
	var __rv uintptr
	q.Drv(182000,182139,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QVariant::detach()
func (q *QVariant) Detach()  {
	q.Drv(182000,182140,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QVariant::isDetached()
func (q *QVariant) IsDetached() bool {
	var __rv bool
	q.Drv(182000,182141,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QVariant::isNull()
func (q *QVariant) IsNull() bool {
	var __rv bool
	q.Drv(182000,182142,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QVariant::isValid()
func (q *QVariant) IsValid() bool {
	var __rv bool
	q.Drv(182000,182143,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QVariant::nameToType(char const*)	
func QVariantNameToType(name string) QVariant_Type {
	var __rv QVariant_Type
	DirectQtDrv(nil,182000,182144,unsafe.Pointer(&name),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QVariant::nameToType(char const*)
func (q *QVariant) NameToType(name string) QVariant_Type {
	var __rv QVariant_Type
	q.Drv(182000,182144,unsafe.Pointer(&name),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QVariant::toBitArray()
func (q *QVariant) ToBitArray() []bool {
	var __rv []bool
	q.Drv(182000,182145,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QVariant::toBool()
func (q *QVariant) ToBool() bool {
	var __rv bool
	q.Drv(182000,182146,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QVariant::toByteArray()
func (q *QVariant) ToByteArray() []byte {
	var __rv []byte
	q.Drv(182000,182147,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QVariant::toChar()
func (q *QVariant) ToChar() rune {
	var __rv rune
	q.Drv(182000,182148,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QVariant::toDate()
func (q *QVariant) ToDate() *QDate {
	var __rv uintptr
	q.Drv(182000,182149,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QDate{}
	_rp.SetDriver(__rv,19000,true)
	return _rp
}	
//QVariant::toDateTime()
func (q *QVariant) ToDateTime() *QDateTime {
	var __rv uintptr
	q.Drv(182000,182150,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QDateTime{}
	_rp.SetDriver(__rv,20000,true)
	return _rp
}	
//QVariant::toDouble(bool*)
func (q *QVariant) ToDouble(ok *bool) float64 {
	var __rv float64
	q.Drv(182000,182151,unsafe.Pointer(&ok),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QVariant::toEasingCurve()
func (q *QVariant) ToEasingCurve() *QEasingCurve {
	var __rv uintptr
	q.Drv(182000,182152,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QEasingCurve{}
	_rp.SetDriver(__rv,29000,true)
	return _rp
}	
//QVariant::toFloat(bool*)
func (q *QVariant) ToFloat(ok *bool) float32 {
	var __rv float32
	q.Drv(182000,182153,unsafe.Pointer(&ok),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QVariant::toInt(bool*)
func (q *QVariant) ToInt(ok *bool) int {
	var __rv int
	q.Drv(182000,182154,unsafe.Pointer(&ok),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QVariant::toLine()
func (q *QVariant) ToLine() *QLine {
	var __rv uintptr
	q.Drv(182000,182155,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QLine{}
	_rp.SetDriver(__rv,68000,true)
	return _rp
}	
//QVariant::toLineF()
func (q *QVariant) ToLineF() *QLineF {
	var __rv uintptr
	q.Drv(182000,182156,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QLineF{}
	_rp.SetDriver(__rv,69000,true)
	return _rp
}	
//QVariant::toList()
func (q *QVariant) ToList() []*QVariant {
	var __rv []*QVariant
	q.Drv(182000,182157,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QVariant::toLocale()
func (q *QVariant) ToLocale() *QLocale {
	var __rv uintptr
	q.Drv(182000,182158,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QLocale{}
	_rp.SetDriver(__rv,72000,true)
	return _rp
}	
//QVariant::toLongLong(bool*)
func (q *QVariant) ToLongLong(ok *bool) int64 {
	var __rv int64
	q.Drv(182000,182159,unsafe.Pointer(&ok),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QVariant::toMap()
func (q *QVariant) ToMap() map[string]*QVariant {
	__rv := make(map[string]*QVariant)
	q.Drv(182000,182160,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QVariant::toPoint()
func (q *QVariant) ToPoint() *QPoint {
	var __rv uintptr
	q.Drv(182000,182161,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPoint{}
	_rp.SetDriver(__rv,99000,true)
	return _rp
}	
//QVariant::toPointF()
func (q *QVariant) ToPointF() *QPointF {
	var __rv uintptr
	q.Drv(182000,182162,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QVariant::toReal(bool*)
func (q *QVariant) ToReal(ok *bool) float64 {
	var __rv float64
	q.Drv(182000,182163,unsafe.Pointer(&ok),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QVariant::toRect()
func (q *QVariant) ToRect() *QRect {
	var __rv uintptr
	q.Drv(182000,182164,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QVariant::toRectF()
func (q *QVariant) ToRectF() *QRectF {
	var __rv uintptr
	q.Drv(182000,182165,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRectF{}
	_rp.SetDriver(__rv,111000,true)
	return _rp
}	
//QVariant::toRegExp()
func (q *QVariant) ToRegExp() *QRegExp {
	var __rv uintptr
	q.Drv(182000,182166,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRegExp{}
	_rp.SetDriver(__rv,112000,true)
	return _rp
}	
//QVariant::toSize()
func (q *QVariant) ToSize() *QSize {
	var __rv uintptr
	q.Drv(182000,182167,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QVariant::toSizeF()
func (q *QVariant) ToSizeF() *QSizeF {
	var __rv uintptr
	q.Drv(182000,182168,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSizeF{}
	_rp.SetDriver(__rv,120000,true)
	return _rp
}	
//QVariant::toString()
func (q *QVariant) ToString() string {
	var __rv string
	q.Drv(182000,182169,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QVariant::toStringList()
func (q *QVariant) ToStringList() []string {
	var __rv []string
	q.Drv(182000,182170,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QVariant::toTime()
func (q *QVariant) ToTime() *QTime {
	var __rv uintptr
	q.Drv(182000,182171,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTime{}
	_rp.SetDriver(__rv,170000,true)
	return _rp
}	
//QVariant::toUInt(bool*)
func (q *QVariant) ToUInt(ok *bool) uint {
	var __rv uint
	q.Drv(182000,182172,unsafe.Pointer(&ok),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QVariant::toULongLong(bool*)
func (q *QVariant) ToULongLong(ok *bool) int64 {
	var __rv int64
	q.Drv(182000,182173,unsafe.Pointer(&ok),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QVariant::toUrl()
func (q *QVariant) ToUrl() *QUrl {
	var __rv uintptr
	q.Drv(182000,182174,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QUrl{}
	_rp.SetDriver(__rv,180000,true)
	return _rp
}	
//QVariant::type()
func (q *QVariant) Type() QVariant_Type {
	var __rv QVariant_Type
	q.Drv(182000,182175,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QVariant::typeName()
func (q *QVariant) TypeName() string {
	var __rv string
	q.Drv(182000,182176,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QVariant::typeToName(QVariant::Type)	
func QVariantTypeToName(_type QVariant_Type) string {
	var __rv string
	DirectQtDrv(nil,182000,182177,unsafe.Pointer(&_type),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QVariant::typeToName(QVariant::Type)
func (q *QVariant) TypeToName(_type QVariant_Type) string {
	var __rv string
	q.Drv(182000,182177,unsafe.Pointer(&_type),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QVariant::userType()
func (q *QVariant) UserType() int {
	var __rv int
	q.Drv(182000,182178,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
