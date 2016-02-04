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

#ifndef UIDRV_H
#define UIDRV_H

#include <QObject>
#include <QByteArray>

class UIDrv : public QObject
{
    Q_OBJECT
public:
    explicit UIDrv(QObject *parent = 0);    
public:
    void call_async_task();
    void call_sync_task();
    int drv(void *_p, int _typeid, int funcid, void *p1, void *p2, void *p3, void *p4, void *p5, void *p6, void *p7, void *p8, void *p9, void *p10, void *p11, void *p12);
signals:
    void asyncDrv(void *_p, int _typeid, int funcid, void *p1, void *p2, void *p3, void *p4, void *p5, void *p6, void *p7, void *p8, void *p9, void *p10, void *p11, void *p12, int *r);
    void asyncTask();
    void syncTask();
public slots:
    void slotDrv(void *_p, int _typeid, int funcid, void *p1, void *p2, void *p3, void *p4, void *p5, void *p6, void *p7, void *p8, void *p9, void *p10, void *p11, void *p12, int *r);
    void slotAsyncTask();
protected:
    int drv_index;
    QThread *m_thread;
};

extern UIDrv uidrv;
extern QByteArray utf8_cache;

#endif // UIDRV_H
