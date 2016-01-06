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

#include "uiapplication.h"
#include "cdrv.h"
#include <QVariant>
#include <QThread>
#include <QTextCodec>
#include <QDebug>
#include <QMutex>

UIApplication::UIApplication(int argc, char **argv) :
    QApplication(argc,argv)
{
}

static bool event_init = false;

bool UIApplication::event(QEvent *e)
{
    if (!event_init && e->type() == QEvent::User+1) {
        event_init = true;
        app_event_init();
    }

    if (e->type() == QEvent::User+2) {
        app_async_task();
    }

    return QApplication::event(e);
}

void UIApplication::sendAsyncTask()
{
    this->postEvent(this,new QEvent(QEvent::Type(QEvent::User+2)));
}
