#ifndef DRV_EVENT_H
#define DRV_EVENT_H

#include "cdrv.h"
#include "drv_event.h"
#include <QDebug>

class UIEventFilter : public QObject
{
public:
    bool eventFilter(QObject *, QEvent *event)
    {
        return drv_event_filter(this,event->type(),event);
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
