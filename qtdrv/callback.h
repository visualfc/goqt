/********************************************************************
** Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
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

#ifndef CALLBACK_H
#define CALLBACK_H

typedef void (*CALLBACK_INSERT_MAP)(void *p, void *p1, void *p2);

void utf8_to_string(void *info,const char *data, int size);
void array_to_slice(void *slice, const void *array, int size);
void append_utf8_to_slice(void *string,const char *str, int size);
void append_bytes_to_slice(void *slice,const void *bytes, int size);
void  append_object_to_slice(void *slice, void *obj, int id, int gc);
void* string_slice_index(void *slice, int index);
void* object_slice_index(void *slice, int index);
void drv_signal_call(void *p, void *face,int id, const void *p1, const void *p2, const void *p3, const void *p4);
void drv_remove_signal_call(void *p);
void insert_string2object_map(void *p, const char *data, int size,int id, int gc, void *obj);
void copy_string2object_map(void *p, void *cp, CALLBACK_INSERT_MAP cb);
void insert_int2object_map(void *p, int key,int id, int gc, void *obj);
void copy_int2object_map(void *p, void *cp, CALLBACK_INSERT_MAP cb);
void insert_object2object_map(void *p, int id1, int gc1,void *obj1, int id2, int gc2, void *obj2);
void app_event_init();
bool drv_event_filter(const void *filter, void *face, void *obj, unsigned int evid, void *event);
void drv_remove_event_filter(const void *filter);
void app_async_task();
void append_uint32_to_slice(void *p, unsigned int v);
void append_double_to_slice(void *p, double v);

int init_callback(void *p, int id, void *p1, void *p2, void *p3);

class QResHelp
{
public:
    static bool registerResourceData(int version, const unsigned char *tree, const unsigned char *name, const unsigned char *data);
    static bool unregisterResourceData(int version, const unsigned char *tree, const unsigned char *name, const unsigned char *data);
};

#endif // CALLBACK_H
