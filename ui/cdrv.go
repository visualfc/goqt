// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

/*
#include <stdlib.h>
#include <string.h>
#include "cdrv_def.h"

extern int qtdrv(void *p, int typeid,int funcid, void *p1,void *p2,void *p3,void *p4,void *p5, void *p6,void *p7,void *p8, void *p9, void *p10, void *p11, void *p12);
static void init()
{
	extern void utf8_to_string(void *string, void *str, int size);
	extern void array_to_slice(void *slice, void *array, int size);
	extern void append_utf8_to_slice(void *slice,void *str, int size);
	extern void append_bytes_to_slice(void *slice,void *bytes, int size);
	extern void append_object_to_slice(void *slice, void *obj, int id, int gc);
	extern void* string_slice_index(void *slice, int index);
	extern void* object_slice_index(void *slice, int index);
	extern void drv_signal_call(void *p,void *face,int id, void *p1,void *p2, void *p3, void *p4);
	extern void drv_remove_signal_call(void *p);
	extern void insert_string2object_map(void *p, char *data, int size,int id, int gc, void *obj);
	extern void copy_string2object_map(void *p, void *cp, void *cb);
	extern void insert_int2object_map(void *p, int key,int id, int gc, void *obj);
	extern void copy_int2object_map(void *p, void *cp, void *cb);
	extern void insert_object2object_map(void *p, int id1, int gc1,void *obj1, int id2, int gc2, void *obj2);
	extern void app_event_init();
	extern int drv_event_filter(void *filter, void *face, void *obj, unsigned int evid, void *event);
	extern void drv_remove_event_filter(void *filter);
	extern void app_async_task();
	qtdrv(&utf8_to_string,1,1,0,0,0,0,0,0,0,0,0,0,0,0);
	qtdrv(&array_to_slice,1,2,0,0,0,0,0,0,0,0,0,0,0,0);
	qtdrv(&append_utf8_to_slice,1,3,0,0,0,0,0,0,0,0,0,0,0,0);
	qtdrv(&append_bytes_to_slice,1,4,0,0,0,0,0,0,0,0,0,0,0,0);
	qtdrv(&append_object_to_slice,1,5,0,0,0,0,0,0,0,0,0,0,0,0);
	qtdrv(&string_slice_index,1,6,0,0,0,0,0,0,0,0,0,0,0,0);
	qtdrv(&object_slice_index,1,7,0,0,0,0,0,0,0,0,0,0,0,0);
	qtdrv(&drv_signal_call,1,8,0,0,0,0,0,0,0,0,0,0,0,0);
	qtdrv(&drv_remove_signal_call,1,9,0,0,0,0,0,0,0,0,0,0,0,0);
	qtdrv(&insert_string2object_map,1,10,0,0,0,0,0,0,0,0,0,0,0,0);
	qtdrv(&copy_string2object_map,1,11,0,0,0,0,0,0,0,0,0,0,0,0);
	qtdrv(&insert_int2object_map,1,12,0,0,0,0,0,0,0,0,0,0,0,0);
	qtdrv(&copy_int2object_map,1,13,0,0,0,0,0,0,0,0,0,0,0,0);
	qtdrv(&insert_object2object_map,1,14,0,0,0,0,0,0,0,0,0,0,0,0);
	qtdrv(&app_event_init,1,15,0,0,0,0,0,0,0,0,0,0,0,0);
	qtdrv(&drv_event_filter,1,16,0,0,0,0,0,0,0,0,0,0,0,0);
	qtdrv(&drv_remove_event_filter,1,17,0,0,0,0,0,0,0,0,0,0,0,0);
	qtdrv(&app_async_task,1,18,0,0,0,0,0,0,0,0,0,0,0,0);
}
#cgo linux LDFLAGS: -L../bin -L../../bin -L. -lqtdrv.ui
#cgo darwin LDFLAGS: -L../bin -L../../bin -L. -lqtdrv.ui
*/
import "C"
import (
	"errors"
	"log"
	"reflect"
	"runtime"
	"sync"
	"unsafe"
)

type Error int

var (
	ErrInvalid                                  = errors.New("invalid argument")
	ErrDrvcall                                  = errors.New("c function call fails")
	fnErrHandler func(error, int32, int32, int) = DefaultErrorHandler
)

var (
	qtdrv_init_error error
)

func InitError() error {
	return qtdrv_init_error
}

func IsLoaded() bool {
	return qtdrv_init_error == nil
}

func (err Error) Error() string {
	switch err {
	case -1:
		return "undefined id"
	case -2:
		return "invalid argument"
	case -3:
		return "too many calls"
	}
	return "unknown error"
}

func cdrv_init() {
	C.init()
}

func DefaultErrorHandler(err error, typeid, funcid int32, call int) {
	pc1, _, _, ok1 := runtime.Caller(call)
	fn1 := runtime.FuncForPC(pc1)
	pc2, file, line, ok2 := runtime.Caller(call + 1)
	fn2 := runtime.FuncForPC(pc2)
	if ok1 && ok2 {
		log.Printf("err:%s, type:%d, func:%d; %s, %s , %s, %d \n", err, typeid, funcid, fn1.Name(), fn2.Name(), file, line)
	} else {
		log.Printf("err:%s, type:%d, func:%d\n", err, typeid, funcid)
	}
}

func PanicErrorHandler(err error, typeid, funcid int32, call int) {
	log.Panicf("err:%s, type:%d, func:%d\n", err, typeid, funcid)
}

func SetErrorHandler(fn func(error, int32, int32, int)) {
	if fn != nil {
		fnErrHandler = fn
	}
}

func DirectQtDrv(p unsafe.Pointer, typeid, funcid int32, p1, p2, p3, p4, p5, p6, p7, p8, p9, p10, p11, p12 unsafe.Pointer) error {
	r := C.qtdrv(p, C.int(typeid), C.int(funcid), p1, p2, p3, p4, p5, p6, p7, p8, p9, p10, p11, p12)
	if r != 0 {
		fnErrHandler(Error(r), typeid, funcid, 2)
		return Error(r)
	}
	return nil
}

func _DirectQtDrv(p unsafe.Pointer, typeid, funcid int32, p1, p2, p3, p4, p5, p6, p7, p8, p9, p10, p11, p12 unsafe.Pointer) error {
	r := C.qtdrv(p, C.int(typeid), C.int(funcid), p1, p2, p3, p4, p5, p6, p7, p8, p9, p10, p11, p12)
	if r != 0 {
		fnErrHandler(Error(r), typeid, funcid, 4)
		return Error(r)
	}
	return nil
}

type Driver interface {
	Drv(typeid, funcid int32, p1, p2, p3, p4, p5, p6, p7, p8, p9, p10, p11, p12 unsafe.Pointer) error
	Delete() error
	Native() unsafe.Pointer
	SetDriverFrom(Driver) error
	IsValidDriver() bool
	QtDrv() _qt_drv
}

func Native(drv Driver) unsafe.Pointer {
	if drv == nil {
		return nil
	}
	return drv.Native()
}

//QPaintDevice
type pd_head struct {
	native   unsafe.Pointer
	isWidget int32
}

func new_pd_head(i QPaintDeviceInterface) *pd_head {
	_, isWidget := i.(QWidgetInterface)
	if isWidget {
		return &pd_head{Native(i), 1}
	}
	return &pd_head{Native(i), 0}
}

func Equal(a, b Driver) bool {
	return a.Native() == b.Native()
}

type BaseDrv struct {
	drv _qt_drv
}

func (b *BaseDrv) QtDrv() _qt_drv {
	return b.drv
}

func isNilDriver(v Driver) bool {
	return !reflect.ValueOf(v).IsNil()
}

func IsValidDriver(v Driver) bool {
	return isNilDriver(v) && v.IsValidDriver()
}

func (b *BaseDrv) SetDriverFrom(v Driver) error {
	if !IsValidDriver(v) {
		return ErrInvalid
	}
	b.drv = v.QtDrv()
	return nil
}

func (b *BaseDrv) IsValidDriver() bool {
	if b == nil {
		return false
	}
	return b.drv.IsValidDriver()
}

func (b *BaseDrv) Drv(typeid, funcid int32, p1, p2, p3, p4, p5, p6, p7, p8, p9, p10, p11, p12 unsafe.Pointer) error {
	if b == nil {
		return ErrInvalid
	}
	return b.drv.Drv(typeid, funcid, p1, p2, p3, p4, p5, p6, p7, p8, p9, p10, p11, p12)
}

func (b *BaseDrv) Native() unsafe.Pointer {
	if b == nil {
		return nil
	}
	return unsafe.Pointer(b.drv.dd)
}

func (b *BaseDrv) SetDriver(dd uintptr, id int32, gc bool) {
	b.drv.dd = dd
	b.drv.id = id
	b.drv.gc = gc
	b.drv.SetAutoGC(gc)
}

func (b *BaseDrv) SetAutoGC(gc bool) {
	b.drv.SetAutoGC(gc)
}

func (b *BaseDrv) IsAutoGC() bool {
	return b.drv.IsAutoGC()
}

func (b *BaseDrv) Delete() error {
	if b == nil {
		return ErrInvalid
	}
	return b.drv.Delete()
}

type _qt_drv struct {
	dd uintptr
	id int32
	gc bool
}

func (d *_qt_drv) Native() unsafe.Pointer {
	if d == nil {
		return nil
	}
	return unsafe.Pointer(d.dd)
}

func (d *_qt_drv) IsValidDriver() bool {
	return d != nil && d.dd != 0
}

func (d *_qt_drv) SetAutoGC(b bool) {
	d.gc = b
	if d.gc {
		runtime.SetFinalizer(d, (*_qt_drv).Delete)
	} else {
		runtime.SetFinalizer(d, nil)
	}
}

func (d *_qt_drv) IsAutoGC() bool {
	return d.gc
}

func (d *_qt_drv) Drv(typeid, funcid int32, p1, p2, p3, p4, p5, p6, p7, p8, p9, p10, p11, p12 unsafe.Pointer) error {
	if d == nil || d.dd == 0 {
		return ErrInvalid
	}
	return _DirectQtDrv(unsafe.Pointer(d.dd), typeid, funcid, p1, p2, p3, p4, p5, p6, p7, p8, p9, p10, p11, p12)
}

func (d *_qt_drv) Delete() error {
	if d == nil || d.dd == 0 {
		return ErrInvalid
	}
	err := _DirectQtDrv(unsafe.Pointer(d.dd), d.id, d.id+1, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
	d.dd = 0
	if d.gc {
		runtime.SetFinalizer(d, nil)
	}
	return err
}

func NewCPtrArrayHead(sa []uintptr) *C.ptr_array_head {
	size := len(sa)
	if size == 0 {
		return &C.ptr_array_head{}
	}
	var p *C.pvoid
	p = (*C.pvoid)(C.malloc(C.size_t(size * C.pvoid_size)))
	C.memcpy(unsafe.Pointer(p), unsafe.Pointer(&sa[0]), C.size_t(size*C.pvoid_size))
	v := &C.ptr_array_head{}
	v.data = p
	v.size = C.int(size)
	return v
}

func UnsafePtrSize() int {
	return C.pvoid_size
}

func FreeCPtrArrayHead(h *C.ptr_array_head) {
	if h.size == 0 {
		return
	}
	C.free(unsafe.Pointer(h.data))
}

func NewCIntArrayHead(sa []int32) *C.int_array_head {
	size := len(sa)
	if size == 0 {
		return &C.int_array_head{}
	}
	var p *C.int
	p = (*C.int)(C.malloc(C.size_t(size * 4)))
	C.memcpy(unsafe.Pointer(p), unsafe.Pointer(&sa[0]), C.size_t(size*4))
	v := &C.int_array_head{}
	v.data = p
	v.size = C.int(size)
	return v

}

func FreeCIntArrayHead(h *C.int_array_head) {
	if h.size == 0 {
		return
	}
	C.free(unsafe.Pointer(h.data))
}

func NewCBoolArrayHead(sa []bool) *C.bool_array_head {
	size := len(sa)
	if size == 0 {
		return &C.bool_array_head{}
	}
	var p *C.bool
	p = (*C.bool)(C.malloc(C.size_t(size)))
	C.memcpy(unsafe.Pointer(p), unsafe.Pointer(&sa[0]), C.size_t(size))
	v := &C.bool_array_head{}
	v.data = p
	v.size = C.int(size)
	return v
}

func FreeCBoolArrayHead(h *C.bool_array_head) {
	if h.size == 0 {
		return
	}
	C.free(unsafe.Pointer(h.data))
}

func NewCByteArrayHead(sa []byte) *C.byte_array_head {
	size := len(sa)
	if size == 0 {
		return &C.byte_array_head{}
	}
	var p *C.char
	p = (*C.char)(C.malloc(C.size_t(size)))
	C.memcpy(unsafe.Pointer(p), unsafe.Pointer(&sa[0]), C.size_t(size))
	v := &C.byte_array_head{}
	v.data = p
	v.size = C.int(size)
	return v
}

func FreeCByteArrayHead(h *C.byte_array_head) {
	if h.size == 0 {
		return
	}
	C.free(unsafe.Pointer(h.data))
}

func NewCStringHead(s string) *C.string_head {
	v := &C.string_head{}
	v.data = C.CString(s)
	v.size = C.int(len(s))
	return v
}

func FreeCStringHead(s *C.string_head) {
	C.free(unsafe.Pointer(s.data))
}

func NewCBytesArray(sa [][]byte) ([]*C.char, int) {
	max := len(sa)
	csa := make([]*C.char, max+1)
	for i := 0; i < max; i++ {
		csa[i] = C.CString(string(sa[i]))
	}
	csa[max] = nil
	return csa, max
}

func NewCSArray(sa []string) ([]*C.char, int) {
	max := len(sa)
	csa := make([]*C.char, max+1)
	for i := 0; i < max; i++ {
		csa[i] = C.CString(sa[i])
	}
	csa[max] = nil
	return csa, max
}

func FreeCSArray(csa []*C.char, max int) {
	for i := 0; i < max; i++ {
		if csa[i] != nil {
			C.free(unsafe.Pointer(csa[i]))
		}
	}
}

//func NewQtDriver(dd uintptr, id int32, gc bool) Driver {
//	d := &qtdrv{dd, id, gc}
//	if gc {
//		runtime.SetFinalizer(d, (*qtdrv).Delete)
//	}
//	return d
//}

//export utf8_to_string
func utf8_to_string(dst unsafe.Pointer, src unsafe.Pointer, size int32) {
	*(*string)(dst) = C.GoStringN((*C.char)(src), C.int(size))
}

//export array_to_slice
func array_to_slice(dst, src unsafe.Pointer, size int32) {
	*(*[]byte)(dst) = C.GoBytes(src, C.int(size))
}

//export append_utf8_to_slice
func append_utf8_to_slice(dst unsafe.Pointer, src unsafe.Pointer, size int32) {
	*(*[]string)(dst) = append(*(*[]string)(dst), C.GoStringN((*C.char)(src), C.int(size)))
}

//export append_bytes_to_slice
func append_bytes_to_slice(dst unsafe.Pointer, src unsafe.Pointer, size int32) {
	*(*[][]byte)(dst) = append(*(*[][]byte)(dst), C.GoBytes(src, C.int(size)))
}

//export append_object_to_slice
func append_object_to_slice(dst unsafe.Pointer, src unsafe.Pointer, id C.int, gc int32) {
	obj := &QObject{}
	obj.SetDriver(uintptr(src), int32(id), gc != 0)
	*(*[]*QObject)(dst) = append(*(*[]*QObject)(dst), obj)
}

//export string_slice_index
func string_slice_index(slice unsafe.Pointer, index C.int) unsafe.Pointer {
	return unsafe.Pointer(&(*(*[]string)(slice))[index])
}

//export object_slice_index
func object_slice_index(slice unsafe.Pointer, index C.int) unsafe.Pointer {
	return ((*(*[]*QObject)(slice))[index]).Native()
}

var (
	signalMap = make(map[uintptr]interface{})
)

//export drv_signal_call
func drv_signal_call(p unsafe.Pointer, face unsafe.Pointer, id int32, p1, p2, p3, p4 unsafe.Pointer) {
	//fn := (*(*Iface)(face)).Interface()
	if fn, ok := signalMap[uintptr(p)]; ok {
		drvSignalCall(fn, id, p1, p2, p3, p4)
	}
}

//export drv_remove_signal_call
func drv_remove_signal_call(p unsafe.Pointer) {
	delete(signalMap, uintptr(p))
}

//export insert_string2object_map
func insert_string2object_map(p unsafe.Pointer, data *C.char, size C.int, id int32, gc int32, src unsafe.Pointer) {
	obj := &QObject{}
	obj.SetDriver(uintptr(src), id, gc != 0)
	(*(*map[string]*QObject)(p))[C.GoStringN(data, size)] = obj
}

//export copy_string2object_map
func copy_string2object_map(p, cp, cb unsafe.Pointer) {
	for k, v := range *(*map[string]*QObject)(p) {
		drvInsertMapCb(cp, unsafe.Pointer(&k), v.Native(), cb)
	}
}

//export insert_int2object_map
func insert_int2object_map(p unsafe.Pointer, key int32, id int32, gc int32, src unsafe.Pointer) {
	obj := &QObject{}
	obj.SetDriver(uintptr(src), id, gc != 0)
	(*(*map[int]*QObject)(p))[int(key)] = obj
}

//export copy_int2object_map
func copy_int2object_map(p, cp, cb unsafe.Pointer) {
	for k, v := range *(*map[int]*QObject)(p) {
		drvInsertMapCb(cp, unsafe.Pointer(&k), v.Native(), cb)
	}
}

//export insert_object2object_map
func insert_object2object_map(p unsafe.Pointer, id1, gc1 int32, src1 unsafe.Pointer, id2, gc2 int32, src2 unsafe.Pointer) {
	obj1 := &QObject{}
	obj1.SetDriver(uintptr(src1), id1, gc1 != 0)
	obj2 := &QObject{}
	obj2.SetDriver(uintptr(src2), id2, gc2 != 0)
	(*(*map[*QObject]*QObject)(p))[obj1] = obj2
}

func defEventInit() {
	log.Println("def_event_init")
}

var fnEventInit func()

func OnEventInit(fn func()) {
	fnEventInit = fn
	log.Println("setup init", fnEventInit)
}

//export app_event_init
func app_event_init() {
	log.Println("app_event_init")
	if fnEventInit != nil {
		go fnEventInit()
	}
}

//string,T => QMap
func drvInsertMapCb(cp, k, v, cb unsafe.Pointer) {
	C.qtdrv(cp, 1, 104, k, v, cb, nil, nil, nil, nil, nil, nil, nil, nil, nil)
}

//QString => string
func drvGetString(p unsafe.Pointer, cp unsafe.Pointer) error {
	C.qtdrv(p, 1, 100, cp, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
	return nil
}

//QStringList => []string
func drvGetStringArray(p unsafe.Pointer, cp unsafe.Pointer) error {
	C.qtdrv(p, 1, 101, cp, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
	return nil
}

//QList<QRectF> => []*QRectF
func drvGetRectFArray(p unsafe.Pointer, cp unsafe.Pointer, id uint32) error {
	C.qtdrv(p, 1, 102, cp, unsafe.Pointer(&id), nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
	return nil
}

//QList<QModelIndex> => []*QModelIndex
func drvGetModelIndexArray(p unsafe.Pointer, cp unsafe.Pointer, id uint32) error {
	C.qtdrv(p, 1, 103, cp, unsafe.Pointer(&id), nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
	return nil
}

func drvBytesArrayToC(src [][]byte) []*byte {
	dst := make([]*byte, len(src)+1)
	for n, v := range src {
		cv := make([]byte, len(v)+1)
		copy(cv, v)
		dst[n] = &cv[0]
	}
	return dst
}

func drvStringArrayToC(src []string) []*byte {
	dst := make([]*byte, len(src)+1)
	for n, v := range src {
		cv := make([]byte, len(v)+1)
		copy(cv, []byte(v))
		dst[n] = &cv[0]
	}
	return dst
}

type QRgb uint32

type Iface struct {
	tab  uintptr
	data uintptr
}

func NewIface(i interface{}) Iface {
	return *(*Iface)(unsafe.Pointer(&i))
}

func drvNewIfaceRef(i interface{}) *Iface {
	return (*Iface)(unsafe.Pointer(&i))
}

func (face Iface) Interface() interface{} {
	var i interface{}
	*(*Iface)(unsafe.Pointer(&i)) = face
	return i
}

//export drv_event_filter
func drv_event_filter(filter unsafe.Pointer, face unsafe.Pointer, obj unsafe.Pointer, evid uint32, event unsafe.Pointer) C.int {
	//	i := ((*Iface)(face)).Interface()
	//	if drvEventCall(i, drvNewObject(uintptr(obj)), evid, uintptr(event)) {
	//		return 1
	//	}

	if i, ok := eventMap[uintptr(filter)]; ok {
		if drvEventCall(i, drvNewObject(uintptr(obj)), evid, uintptr(event)) {
			return 1
		}
	}
	return 0
}

//export drv_remove_event_filter
func drv_remove_event_filter(filter unsafe.Pointer) {
	delete(eventMap, uintptr(filter))
}

var (
	eventMap = make(map[uintptr]interface{})
)

func drvInstallEventFilter(obj QObjectInterface, filter interface{}) uintptr {
	if filter == nil {
		return 0
	}
	var __rv uintptr
	DirectQtDrv(Native(obj), 1, 105, nil, unsafe.Pointer(&__rv), nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
	if __rv != 0 {
		eventMap[__rv] = filter
	}
	return __rv
}

type TaskList struct {
	task  []func()
	mutex sync.Mutex
}

func (l *TaskList) Append(fn func()) {
	l.mutex.Lock()
	l.task = append(l.task, fn)
	l.mutex.Unlock()
}

func (l *TaskList) Run() {
	l.mutex.Lock()
	for _, fn := range l.task {
		fn()
	}
	l.task = make([]func(), 0)
	l.mutex.Unlock()
}

var (
	tasks TaskList
)

//export app_async_task
func app_async_task() {
	tasks.Run()
}

func QtVersion() string {
	var __rv string
	_DirectQtDrv(nil, 1, 106, unsafe.Pointer(&__rv), nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
	return __rv
}
