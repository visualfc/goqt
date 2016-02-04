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

#include "uidrv.h"
#include "cdrv.h"
#include <QThread>
#include <QApplication>

UIDrv uidrv;
QByteArray utf8_cache;

UIDrv::UIDrv(QObject *parent) :
    QObject(parent), drv_index(0)
{
    m_thread = this->thread();
    connect(this,SIGNAL(asyncDrv(void*,int,int,void*,void*,void*,void*,void*,void*,void*,void*,void*,void*,void*,void*,int*)),this,SLOT(slotDrv(void*,int,int,void*,void*,void*,void*,void*,void*,void*,void*,void*,void*,void*,void*,int*)),Qt::BlockingQueuedConnection);
    connect(this,SIGNAL(asyncTask()),this,SLOT(slotAsyncTask()));
    connect(this,SIGNAL(syncTask()),this,SLOT(slotAsyncTask()),Qt::BlockingQueuedConnection);
}

void UIDrv::call_async_task()
{
    if (m_thread == QThread::currentThread()) {
        slotAsyncTask();
    } else {
        emit asyncTask();
    }
}

void UIDrv::call_sync_task()
{
    if (m_thread == QThread::currentThread()) {
        slotAsyncTask();
    } else {
        emit syncTask();
    }
}

int UIDrv::drv(void *_p, int _typeid, int funcid, void *p1, void *p2, void *p3, void *p4, void *p5, void *p6, void *p7, void *p8, void *p9, void *p10, void *p11, void *p12)
{
    if (drv_index >= 1000) {
    //    return -3;
    }
    drv_index++;
    int r = 0;    
    if (QThread::currentThread() == m_thread ) {
        slotDrv(_p,_typeid,funcid,p1,p2,p3,p4,p5,p6,p7,p8,p9,p10,p11,p12,&r);
    } else {
        emit asyncDrv(_p,_typeid,funcid,p1,p2,p3,p4,p5,p6,p7,p8,p9,p10,p11,p12,&r);
    }
    qDebug() << drv_index;
    return r;
}

void UIDrv::slotDrv(void *_p, int _typeid, int funcid, void *p1, void *p2, void *p3, void *p4, void *p5, void *p6, void *p7, void *p8, void *p9, void *p10, void *p11, void *p12, int *r)
{
    drv_index--;
    *r = qtdrv(_p,_typeid,funcid,p1,p2,p3,p4,p5,p6,p7,p8,p9,p10,p11,p12);
}

void UIDrv::slotAsyncTask()
{
    app_async_task();
}
