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

#ifndef CDRV_EVENT_H
#define CDRV_EVENT_H

#include "cdrv.h"
#include "cdrv_event.h"
#include <QDebug>

class UIEventFilter : public QObject
{
public:
    ~UIEventFilter()
    {
        drv_remove_event_filter(this);
    }
    bool eventFilter(QObject *obj, QEvent *event)
    {
        return drv_event_filter(this,0,obj,event->type(),event);
    }
};

inline void* drvNewFilter(void *obj)
{
    QObject *object = (QObject*)obj;
    UIEventFilter *filter = new UIEventFilter;
    filter->moveToThread(object->thread());
    filter->setParent(object);
    object->installEventFilter(filter);
    return filter;
}


#endif // DRV_EVENT_H
