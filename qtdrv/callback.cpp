/********************************************************************
** Copyright 2013-2016 visualfc <visualfc@gmail.com>. All rights reserved.
**
** This library is free software; you can redistribute it and/or
** modify it under the terms of the GNU Lesser General Public
** License as published by the Free Software Foundation; either
** version 2.1 of the License, or (at your option) any later version.
**
** This library is distributed in the hope that it will be useful,
** but WITHOUT ANY WARRANTY; without even the implied warranty of
** MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
** Lesser General Public License for more details.
*********************************************************************/

#include "callback.h"
#include "uiapplication.h"
#include "cdrv.h"
#include "cdrv_event.h"

typedef void (*UTF8_TO_STRING)(void *info,const char *data, int size);
typedef void (*ARRAY_TO_SLICE)(void *slice, const void *array, int size);
typedef void (*APPEND_UTF8_TO_SLICE)(void *string,const char *str, int size);
typedef void (*APPEND_BYTES_TO_SLICE)(void *slice,const void *bytes, int size);
typedef void (*APPEND_OBJECT_TO_SLICE)(void *slice, void *obj, int id, int gc);
typedef void*(*STRING_SLICE_INDEX)(void *slice, int index);
typedef void*(*OBJECT_SLICE_INDEX)(void *slice, int index);
typedef void (*DRV_SIGNAL_CALL)(void *p, void *face, int id, const void *p1, const void *p2, const void *p3, const void *p4);
typedef void (*DRV_REMOVE_SIGNAL_CALL)(void *);
typedef void (*INSERT_STRING2OBJECT_MAP)(void *p, const char *data, int size,int id, int gc, void *obj);
typedef void (*COPY_STRING2OBJECT_MAP)(void *p, void *cp, CALLBACK_INSERT_MAP cb);
typedef void (*INSERT_INT2OBJECT_MAP)(void *p, int key,int id, int gc, void *obj);
typedef void (*COPY_INT2OBJECT_MAP)(void *p, void *cp, CALLBACK_INSERT_MAP cb);
typedef void (*INSERT_OBJECT2OBJECT_MAP)(void *p, int id1, int gc1,void *obj1, int id2, int gc2, void *obj2);
typedef void (*APP_EVENT_INIT)();
typedef int  (*DRV_EVENT_FILTER)(const void *filter, void *face, void *obj, unsigned int evid, void *event);
typedef void (*DRV_REMOVE_EVENT_FILTER)(const void *filter);
typedef void (*APP_ASYNC_TASK)();
typedef void (*APPEND_UINT32_TO_SLICE)(void *p, unsigned int v);
typedef void (*APPEND_DOUBLE_TO_SLICE)(void *p, double v);

static ARRAY_TO_SLICE pfn_array_to_slice;
static UTF8_TO_STRING pfn_utf8_to_string;
static APPEND_UTF8_TO_SLICE pfn_append_utf8_to_slice;
static APPEND_BYTES_TO_SLICE pfn_append_bytes_to_slice;
static APPEND_OBJECT_TO_SLICE pfn_append_object_to_slice;
static STRING_SLICE_INDEX pfn_string_slice_index;
static OBJECT_SLICE_INDEX pfn_object_slice_index;
static DRV_SIGNAL_CALL pfn_drv_signal_call;
static DRV_REMOVE_SIGNAL_CALL pfn_drv_remove_signal_call;
static INSERT_STRING2OBJECT_MAP pfn_insert_string2object_map;
static COPY_STRING2OBJECT_MAP pfn_copy_string2object_map;
static INSERT_INT2OBJECT_MAP pfn_insert_int2object_map;
static COPY_INT2OBJECT_MAP pfn_copy_int2object_map;
static INSERT_OBJECT2OBJECT_MAP pfn_insert_object2object_map;
static APP_EVENT_INIT pfn_app_event_init;
static DRV_EVENT_FILTER pfn_drv_event_filter;
static DRV_REMOVE_EVENT_FILTER pfn_drv_remove_event_filter;
static APP_ASYNC_TASK pfn_app_async_task;
static APPEND_UINT32_TO_SLICE pfn_append_uint_to_slice;
static APPEND_DOUBLE_TO_SLICE pfn_append_double_to_slice;

void utf8_to_string(void *info,const char *data, int size)
{
    pfn_utf8_to_string(info,data,size);
}

void array_to_slice(void *slice, const void *array, int size)
{
    pfn_array_to_slice(slice,array,size);
}

void append_utf8_to_slice(void *string,const char *str, int size)
{
    pfn_append_utf8_to_slice(string,str,size);
}

void append_bytes_to_slice(void *slice,const void *bytes, int size)
{
    pfn_append_bytes_to_slice(slice,bytes,size);
}

void  append_object_to_slice(void *slice, void *obj, int id, int gc)
{
    pfn_append_object_to_slice(slice,obj,id,gc);
}

void* string_slice_index(void *slice, int index)
{
    return pfn_string_slice_index(slice,index);
}

void* object_slice_index(void *slice, int index)
{
    return pfn_object_slice_index(slice,index);
}

void drv_signal_call(void *p, void *face, int id, const void *p1, const void *p2, const void *p3, const void *p4)
{
    pfn_drv_signal_call(p,face,id,p1,p2,p3,p4);
}

void drv_remove_signal_call(void *p)
{
    pfn_drv_remove_signal_call(p);
}

void insert_string2object_map(void *p, const char *data, int size,int id, int gc, void *obj)
{
    pfn_insert_string2object_map(p,data,size,id,gc,obj);
}

void copy_string2object_map(void *p, void *cp, CALLBACK_INSERT_MAP cb)
{
    pfn_copy_string2object_map(p,cp,cb);
}

void insert_int2object_map(void *p, int key,int id, int gc, void *obj)
{
    pfn_insert_int2object_map(p,key,id,gc,obj);
}

void copy_int2object_map(void *p, void *cp, CALLBACK_INSERT_MAP cb)
{
    pfn_copy_int2object_map(p,cp,cb);
}

void insert_object2object_map(void *p, int id1, int gc1,void *obj1, int id2, int gc2, void *obj2)
{
    pfn_insert_object2object_map(p,id1,gc1,obj1,id2,gc2,obj2);
}

void app_event_init()
{
    pfn_app_event_init();
}

void app_async_task()
{
    pfn_app_async_task();
}

bool drv_event_filter(const void *filter, void *face, void *obj, uint32 evid, void *event)
{
    return pfn_drv_event_filter(filter, face, obj, evid, event) != 0;
}

void drv_remove_event_filter(const void *filter)
{
    pfn_drv_remove_event_filter(filter);
}

void append_uint32_to_slice(void *p, unsigned int v)
{
    pfn_append_uint_to_slice(p,v);
}

void append_double_to_slice(void *p, double v)
{
    pfn_append_double_to_slice(p,v);
}

int init_callback(void *p, int id, void *p1, void *p2, void *p3)
{
    switch (id) {
    case 1:
        pfn_utf8_to_string = (UTF8_TO_STRING)p;
        break;
    case 2:
        pfn_array_to_slice = (ARRAY_TO_SLICE)p;
        break;
    case 3:
        pfn_append_utf8_to_slice = (APPEND_UTF8_TO_SLICE)p;
        break;
    case 4:
        pfn_append_bytes_to_slice = (APPEND_BYTES_TO_SLICE)p;
        break;
    case 5:
        pfn_append_object_to_slice = (APPEND_OBJECT_TO_SLICE)p;
        break;
    case 6:
        pfn_string_slice_index = (STRING_SLICE_INDEX)p;
        break;
    case 7:
        pfn_object_slice_index = (OBJECT_SLICE_INDEX)p;
        break;
    case 8:
        pfn_drv_signal_call = (DRV_SIGNAL_CALL)p;
        break;
    case 9:
        pfn_drv_remove_signal_call = (DRV_REMOVE_SIGNAL_CALL)p;
        break;
    case 10:
        pfn_insert_string2object_map = (INSERT_STRING2OBJECT_MAP)p;
        break;
    case 11:
        pfn_copy_string2object_map = (COPY_STRING2OBJECT_MAP)p;
        break;
    case 12:
        pfn_insert_int2object_map = (INSERT_INT2OBJECT_MAP)p;
        break;
    case 13:
        pfn_copy_int2object_map = (COPY_INT2OBJECT_MAP)p;
        break;
    case 14:
        pfn_insert_object2object_map = (INSERT_OBJECT2OBJECT_MAP)p;
        break;
    case 15:
        pfn_app_event_init = (APP_EVENT_INIT)p;
        break;
    case 16:
        pfn_drv_event_filter = (DRV_EVENT_FILTER)p;
        break;
    case 17:
        pfn_drv_remove_event_filter = (DRV_REMOVE_EVENT_FILTER)p;
        break;
    case 18:
        pfn_app_async_task = (APP_ASYNC_TASK)p;
        break;
    case 19:
        pfn_append_uint_to_slice = (APPEND_UINT32_TO_SLICE)p;
        break;
    case 20:
        pfn_append_double_to_slice = (APPEND_DOUBLE_TO_SLICE)p;
        break;
    case 100: //QString => string
        drvSetString(p,*(QString*)p1);
        break;
    case 101: //QStringList => []string
        drvSetStringArray(p,*(QStringList*)p1);
        break;
    case 102: //QList<QRectF> => []*QRectF
        drvSetListPtr(p,*(int*)p2,*(QList<QRectF>*)p1);
        break;
    case 103: //QList<QModelIndex> => []*QModelIndex
        drvSetListPtr(p,*(int*)p2,*(QList<QModelIndex>*)p1);
        break;
    case 104: //string,T => QMap
        (CALLBACK_INSERT_MAP(p3))(p,p1,p2);
        break;
    case 105:
        *(void**)p2 = drvNewFilter(p);
        break;
    case 106:
        drvSetString(p1,qVersion());
        break;
    case 200:
        uidrv.call_async_task();
        break;
    case 201:
        uidrv.call_sync_task();
        break;
    default:
        return -1;
    }
    return 0;
}

extern Q_CORE_EXPORT bool qRegisterResourceData
    (int, const unsigned char *, const unsigned char *, const unsigned char *);

extern Q_CORE_EXPORT bool qUnregisterResourceData
    (int, const unsigned char *, const unsigned char *, const unsigned char *);

bool QResHelp::registerResourceData(int version, const unsigned char *tree, const unsigned char *name, const unsigned char *data)
{
    return qRegisterResourceData(version,tree,name,data);
}

bool QResHelp::unregisterResourceData(int version, const unsigned char *tree, const unsigned char *name, const unsigned char *data)
{
    return qUnregisterResourceData(version,tree,name,data);
}
