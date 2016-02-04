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

#ifndef CDRV_H
#define CDRV_H

#include "qtdrv_global.h"
#include "uiapplication.h"
#include "callback.h"
#include "uidrv.h"
#include <QtGui>
#include <QUiLoader>
#include "qsyntaxhighlighterhook.h"

#if QT_VERSION >= 0x050000
#define QTDRV_QT5 1
#endif

#ifdef QTDRV_QT5
    #include <QtWidgets>
    #include <QtPrintSupport>
#else
    #define QT_DEPRECATED_SINCE(major, minor) 0
    #include "./qurl/qurlquery.h"
#endif

typedef	unsigned int uint32;

struct string_head {
    const char *data;
    int size;
};

typedef struct {
    const char **data;
    int size;
} string_array_head;

struct string_head_cache {
    const char *data;
    int         len;
};


struct slice_head {
    char *data;
    uint32 len;
    uint32 cap;
};

template <class T>
struct slice_head_t {
    T *data;
    uint32 len;
    uint32 cap;
};

struct Iface {
    void *tab;
    void *data;
};

Q_DECLARE_METATYPE(Iface)
Q_DECLARE_TYPEINFO(Iface, Q_MOVABLE_TYPE);

inline void drvSetString(void *p, const QString &v)
{
    if (p == 0) {
        return;
    }
    /*
    const QByteArray & ar = v.toUtf8();
    utf8_to_string(p,ar.constData(),ar.length());
    */
    utf8_cache = v.toUtf8();
    string_head_cache *sh = (string_head_cache*)(p);
    sh->data = utf8_cache.constData();
    sh->len = utf8_cache.length();
}

inline void drvSetByteArray(void *p, QByteArray v)
{
    if (p == 0) {
        return;
    }
    array_to_slice(p,v.constData(),v.size());
}


inline QByteArray drvGetByteArray(void *p)
{
    if (p == 0) {
        return QByteArray();
    }
    slice_head *sh = (slice_head*)p;
    return QByteArray((const char*)sh->data,sh->len);
}

inline void drvSetBitArray(void *p, QBitArray v)
{
    if (p == 0) {
        return;
    }
#if QT_VERSION >= 0x050000
    QVector<bool> ar;
    for (int i = 0; i < v.size(); i++) {
        ar.push_back(v.at(i));
    }
    array_to_slice(p,ar.data(),ar.size());
#else
    array_to_slice(p,v.data_ptr()->data,v.size());
#endif
}

inline QString drvGetStringHead(void *p)
{
    if (p == 0) {
        return QString();
    }
    string_head *sh = (string_head*)p;
    return QString::fromUtf8(sh->data,sh->size);
}

inline QStringList drvGetStringArray(void *p)
{
    if (p == 0) {
        return QStringList();
    }
    QStringList list;
    char **sh = (char**)p;
    while (*sh) {
        list.append(QString::fromUtf8(*sh));
        sh++;
    }
    return list;
}


inline QString drvGetStringRef(void *p)
{
    if (p == 0) {
        return QString();
    }
    string_head *h = (string_head*)p;

    return QString::fromUtf8(h->data,h->size);
}

inline void drvSetStringRef(void *p, const QString &v)
{
    if (p == 0) {
        return;
    }
    const QByteArray & ar = v.toUtf8();
    utf8_to_string(p,ar.constData(),ar.length());
}

inline void drvSetStringArray(void *p, const QStringList &list)
{
    if (p == 0) {
        return;
    }
    foreach (QString v, list) {
        const QByteArray & ar = v.toUtf8();
        append_utf8_to_slice(p,ar.constData(),ar.size());
    }
}

inline void drvSetBytesArray(void *p, const QList<QByteArray> &list)
{
    if (p == 0) {
        return;
    }
    foreach (QByteArray v, list) {
        append_bytes_to_slice(p,v.constData(),v.size());
    }
}

inline void drvSetBytesArray(void *p, char *list[])
{
    if (p == 0) {
        return;
    }
    char *head = list[0];
    while (head++) {
        append_bytes_to_slice(p,head,strlen(head));
    }
}


inline void drvSetObjectArray(void *p, int id, const QList<QObject*> &list)
{
    if (p == 0) {
        return;
    }
    foreach (QObject *v, list) {
        append_object_to_slice(p,v,id,false);
    }
}

inline void drvSetRgbArray(void *p, QVector<QRgb> ar)
{
    if (p == 0) {
        return;
    }
    array_to_slice(p,&ar[0],ar.size());
}

inline void drvSetDoubleArray(void *p, QVector<double> ar)
{
    if (p == 0) {
        return;
    }
    array_to_slice(p,&ar[0],ar.size());
}

inline QVector<QRgb> drvGetRgbArray(void *p)
{
    if (p == 0) {
        return QVector<QRgb>();
    }
    QVector<QRgb> ar;
    slice_head_t<QRgb> *sh = (slice_head_t<QRgb>*)p;
    for (uint32 i = 0; i < sh->len; i++) {
        ar.append(sh->data[i]);
    }
    return ar;
}

inline QVector<double> drvGetDoubleArray(void *p)
{
    if (p == 0) {
        return QVector<double>();
    }
    QVector<double> ar;
    slice_head_t<double> *sh = (slice_head_t<double>*)p;
    for (uint32 i = 0; i < sh->len; i++) {
        ar.append(sh->data[i]);
    }
    return ar;
}

inline QBitArray drvGetBitArray(void *p)
{
    if (p == 0) {
        return QBitArray();
    }
    slice_head_t<bool> *sh = (slice_head_t<bool>*)p;
    QBitArray ar(sh->len);
    for (uint32 i = 0; i < sh->len; i++) {
        ar[i] = sh->data[i];
    }
    return ar;
}

inline char ** drvGetBytesArray(void *p)
{
    if (p == 0) {
        return 0;
    }
    slice_head_t<char*> *sh = (slice_head_t<char*>*)(p);
    return &sh->data[0];
}

template <class T>
inline QVector<T> drvGetVector(void *p)
{
    if (p == 0) {
        return QVector<T>();
    }
    QVector<T> ar;
    slice_head_t<T> *sh = (slice_head_t<T>*)p;
    for (uint32 i = 0; i < sh->len; i++) {
        ar.append(sh->data[i]);
    }
    return ar;
}

template <class T>
inline void drvSetVector(void *p, QVector<T> ar)
{
    if (p == 0) {
        return;
    }
    array_to_slice(p,&ar[0],ar.size());
}

template <class T>
inline void drvSetList(void *p, QList<T> ar)
{
    if (p == 0) {
        return;
    }
    array_to_slice(p,&ar.toVector()[0],ar.size());
}

template <class T>
inline QList<T> drvGetList(void *p)
{
    if (p == 0) {
        return QList<T>();
    }
    QList<T> ar;
    slice_head_t<T> *sh = (slice_head_t<T>*)p;
    for (uint32 i = 0; i < sh->len; i++) {
        ar.append(sh->data[i]);
    }
    return ar;
}

template <class T>
inline QVector<T> drvGetVectorPtr(void *p)
{
    if (p == 0) {
        return QVector<T>();
    }
    QVector<T> ar;
    slice_head *sh = (slice_head*)p;
    for (uint32 i = 0; i < sh->len; i++) {
        ar.append(*(T*)object_slice_index(p,i));
    }
    return ar;
}

template <class T>
inline QList<T> drvGetListPtr(void *p)
{
    if (p == 0) {
        return QList<T>();
    }
    QList<T> ar;
    slice_head *sh = (slice_head*)p;
    for (uint32 i = 0; i < sh->len; i++) {
        ar.append(*(T*)object_slice_index(p,i));
    }
    return ar;
}

template <class T>
inline QList<T> drvGetListObj(void *p)
{
    if (p == 0) {
        return QList<T>();
    }
    QList<T> ar;
    slice_head *sh = (slice_head*)p;
    for (uint32 i = 0; i < sh->len; i++) {
        ar.append((T)object_slice_index(p,i));
    }
    return ar;
}

template <class T>
inline void drvSetVectorObj(void *p, int id, QVector<T> ar)
{
    if (p == 0) {
        return;
    }
    foreach (T v, ar) {
       append_object_to_slice(p,v,id,false);
    }
}

template <class T>
inline void drvSetVectorPtr(void *p, int id, QVector<T> ar)
{
    if (p == 0) {
        return;
    }
    foreach (T v, ar) {
       append_object_to_slice(p,new T(v),id,true);
    }
}

template <class T>
inline void drvSetListPtr(void *p, int id, QList<T> ar)
{
    if (p == 0) {
        return;
    }
    foreach (T v, ar) {
       append_object_to_slice(p,new T(v),id,true);
    }
}

template <class T>
inline void drvSetListObj(void *p, int id, QList<T> ar)
{
    if (p == 0) {
        return;
    }
    foreach (T v, ar) {
       append_object_to_slice(p,v,id,false);
    }
}


//template <class T>
//inline T drvNative(void *p)
//{
//    if (p == 0) {
//        return 0;
//    }
//    return *(T*)p;
//}

//template <class T>
//inline T* drvNativePtr(void *p)
//{
//    if (p == 0) {
//        return 0;
//    }
//    return (T*)p;
//}

//template <>
//inline QPaintDevice* drvNativePtr(void *p)
//{
//    qDebug() << "painter";
//    return (QWidget*)p;
//}

struct pd_head {
    void *native;
    int isWidget;
};

inline QPaintDevice* drvGetPaintDevice(void *p)
{
    if (p == 0) {
        return 0;
    }
    pd_head *ph = (pd_head*)p;
    if (ph->isWidget) {
        return (QWidget*)ph->native;
    }
    return (QPaintDevice*)ph->native;
}

extern QList<string_head_cache*> sh_casche;
inline const char *drvGet_const_char(void *p) {
    if (p == 0) {
        return "";
    }
    string_head *sh = (string_head*)p;

    if (sh->size == 0) {
        return "";
    }
    return sh->data;
}

template <typename T>
void drvSetStringToObjectMap(void *p, int id, bool gc, QMap<QString,T> map)
{
    QMapIterator<QString, T> i(map);
    while (i.hasNext()) {
        i.next();
        const QByteArray & ar = i.key().toUtf8();
        insert_string2object_map(p,ar.constData(),ar.size(),id,gc?1:0,new T(i.value()));
    }
}

template <typename T>
void drvSetIntToObjectMap(void *p, int id, bool gc, QMap<int,T> map)
{
    QMapIterator<int, T> i(map);
    while (i.hasNext()) {
        i.next();
        insert_int2object_map(p,i.key(),id,gc?1:0,new T(i.value()));
    }
}

template <typename T1, typename T2>
void drvSetObjectToObjectMap(void *p, int id1, bool gc1, int id2, bool gc2, QMap<T1,T2> map)
{
    QMapIterator<T1,T2> i(map);
    while (i.hasNext()) {
        i.next();
        insert_object2object_map(p,id1,gc1?1:0,new T1(i.key()),id2,gc2?1:0,new T2(i.value()));
    }
}
template <typename T>
inline void drvInsertString2Object(void *p, void *p1, void *p2)
{
    string_head *sh = (string_head*)p1;
    (*((QMap<QString,T>*)p)).insert(QString::fromUtf8(sh->data,sh->size),*(T*)p2);
}

template <typename T>
QMap<QString,T> drvGetStringToObjectMap(void *p)
{
    if (p == 0) {
        return QMap<QString,T>();
    }
    QMap<QString,T> m;
    copy_string2object_map(p,&m,&drvInsertString2Object<T>);
    return m;
}

template <typename T>
inline void drvInsertInt2Object(void *p, void *p1, void *p2)
{
    (*((QMap<int,T>*)p)).insert(*(int*)p1,*(T*)p2);
}


template <typename T>
QMap<int,T> drvGetIntToObjectMap(void *p)
{
    if (p == 0) {
        return QMap<int,T>();
    }
    QMap<int,T> m;
    copy_int2object_map(p,&m,&drvInsertInt2Object<T>);
    return m;
}

extern "C"
QTUISHARED_EXPORT int qtdrv(void *_p, int _typeid, int funcid, void *p1, void *p2, void *p3, void *p4, void *p5, void *p6, void *p7, void *p8, void *p9, void *p10, void *p11, void *p12);


#endif // CDRV_H
