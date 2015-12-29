#ifndef UIUTIL_H
#define UIUTIL_H

#include "uilib.h"
#include <QString>
#include <QByteArray>
#include <QPoint>
#include <QPointF>
#include <QSize>
#include <QSizeF>
#include <QRect>
#include <QRectF>
#include <QList>
/*
inline QByteArray fromByteArray(const ByteArray &v)
{
    return QByteArray(v.data,v.size);
}

inline QByteArray fromByteArray(void *p)
{
    ByteArray *pb = (ByteArray*)p;
    return QByteArray(pb->data,pb->size);
}

inline ByteArray toByteArray(const QByteArray &v)
{
    return (ByteArray){v.constData(),v.size()};
}

inline ByteArray toByteArray(const void *p)
{
    QByteArray *v = (QByteArray*)p;
    return (ByteArray){v->constData(),v->size()};
}

inline ByteArray toByteArray(void *p)
{
    QByteArray *v = (QByteArray*)p;
    return (ByteArray){v->constData(),v->size()};
}

inline void toByteArray(void *p, const QByteArray &v)
{
    ByteArray *pb = (ByteArray*)p;
    pb->data = v.constData();
    pb->size = v.size();
}

inline QString fromString(const String &v)
{
    return QString::fromUtf16(v.data,v.size);
}

inline String toString(const QString &v)
{
    return (String){v.utf16(),v.size()};
}

inline void toString(void *p, const QString &v)
{
    String *ps = (String*)p;
    ps->data = v.utf16();
    ps->size = v.size();
}

inline QPoint fromPoint(const Point &v)
{
    return QPoint(v.X,v.Y);
}

inline Point toPoint(const QPoint &v)
{
    return (Point){v.x(),v.y()};
}

inline QPointF fromPointF(const PointF &v)
{
    return QPointF(v.X,v.Y);
}

inline PointF toPointF(const QPointF &v)
{
    return (PointF){v.x(),v.y()};
}

inline QSize fromSize(const Size &v)
{
    return QSize(v.Width,v.Height);
}

inline Size toSize(const QSize &v)
{
    return (Size){v.width(),v.height()};
}

inline QSizeF fromSizeF(const SizeF &v)
{
    return QSizeF(v.Width,v.Height);
}

inline SizeF toSizeF(const QSizeF &v)
{
    return (SizeF){v.width(),v.height()};
}

inline QRect fromRect(const Rect &v)
{
    return QRect(v.X,v.Y,v.Widht,v.Height);
}

inline Rect toRect(const QRect &v)
{
    return (Rect){v.x(),v.y(),v.width(),v.height()};
}

inline QRectF fromRectF(const RectF &v)
{
    return QRectF(v.X,v.Y,v.Width,v.Height);
}

inline RectF toRect(const QRectF &v)
{
    return (RectF){v.x(),v.y(),v.width(),v.height()};
}

template <typename T>
inline List toList(const QList<T> ar)
{
    List list;
    InitList(&list,ar.size());
    for (int i = 0; i < ar.size(); i++) {
        list.data[i] = (void*)&ar[i];
    }
    return list;
}
*/

#endif // UIUTIL_H
