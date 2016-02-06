// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#ifndef _CDRV_DEF_H_
#define _CDRV_DEF_H_

#include <stdbool.h>

typedef struct {
    const char *data;
    int size;
} string_head;

typedef struct {
    const char *data;
    int size;
} byte_array_head;

typedef struct {
	const bool *data;
	int	size;
} bool_array_head;

typedef struct {
    const int *data;
	int	size;
} int_array_head;

typedef struct {
	const unsigned int *data;
	int	size;
} uint_array_head;

typedef void* pvoid;

typedef struct {
	const pvoid *data;
	int	 size;
} ptr_array_head;

typedef struct {
	double *data;
	int		size;
} double_array_head;

#define pvoid_size sizeof(pvoid)

#endif //_CDRV_DEF_H_
