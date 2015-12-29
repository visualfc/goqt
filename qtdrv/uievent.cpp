#include "uievent.h"
#include "uiapplication.h"
#include "uiutil.h"
#include <QEvent>
#include <QResizeEvent>
#include <QMoveEvent>
#include <QTimerEvent>
#include <QMouseEvent>
#include <QHoverEvent>
#include <QPointF>
#include <QKeyEvent>
#include <QPaintEvent>
#include <QFocusEvent>
#include <QDebug>

UIEvent::UIEvent(QObject *parent, int type, void *context) :
    QObject(parent), m_type(type), m_context(context)
{
}

UIEvent::~UIEvent()
{
    theApp.fnEventRelease(this,m_context);
}

UIEvent *UIEvent::createUIEvent(QObject *sender, int st, void *context)
{
    UIEvent *event = new UIEvent(sender,st,context);
    sender->installEventFilter(event);
    return event;
}

bool UIEvent::eventFilter(QObject *obj, QEvent *event)
{
    if (m_type == event->type()) {
        int accept = event->isAccepted() ? 1 : 0;
        switch(m_type) {
        case QEvent::Create:
        case QEvent::Close:
        case QEvent::Show:
        case QEvent::Hide:
        case QEvent::Enter:
        case QEvent::Leave: {
            Event ev = {accept,0};
            theApp.fnEventCall(obj,this,m_type,&ev,m_context);
            event->setAccepted(ev.Accept != 0);
            return ev.Filter;
        }
        case QEvent::FocusIn:
        case QEvent::FocusOut: {
            QFocusEvent *e = (QFocusEvent*)event;
            FocusEvent ev = {accept,0,e->reason()};
            theApp.fnEventCall(obj,this,m_type,&ev,m_context);
            event->setAccepted(ev.Accept != 0);
            return ev.Filter;
        }
        case QEvent::Timer: {
            QTimerEvent *e = (QTimerEvent*)event;
            TimerEvent ev = {accept,0,e->timerId()};
            theApp.fnEventCall(obj,this,m_type,&ev,m_context);
            event->setAccepted(ev.Accept != 0);
            return ev.Filter;
        }
        case QEvent::HoverEnter:
        case QEvent::HoverLeave:
        case QEvent::HoverMove: {
            QHoverEvent *e = (QHoverEvent*)event;
            const QPoint &pt = e->pos();
            const QPoint &opt = e->oldPos();
            HoverEvent ev = {accept,0,pt.x(),pt.y(),opt.x(),opt.y()};
            theApp.fnEventCall(obj,this,m_type,&ev,m_context);
            event->setAccepted(ev.Accept != 0);
            return ev.Filter;
        }
        case QEvent::KeyPress:
        case QEvent::KeyRelease: {
            QKeyEvent *e = (QKeyEvent*)event;
            //KeyEvent ev = {accept,0,e->modifiers(),e->count(),e->isAutoRepeat()?1:0,e->key(),e->nativeModifiers(),e->nativeScanCode(),e->nativeVirtualKey(),toString(e->text())};
            //theApp.fnEventCall(obj,this,m_type,&ev,m_context);
            //event->setAccepted(ev.Accept != 0);
            //return ev.Filter;
        }
        case QEvent::MouseButtonPress:
        case QEvent::MouseButtonRelease:
        case QEvent::MouseButtonDblClick:
        case QEvent::MouseMove: {
            QMouseEvent *e = (QMouseEvent*)event;
            const QPoint &gpt = e->globalPos();
            const QPoint &pt = e->pos();
            MouseEvent ev = {accept,0,e->modifiers(),e->button(),e->buttons(),gpt.x(),gpt.y(),pt.x(),pt.y()};
            theApp.fnEventCall(obj,this,m_type,&ev,m_context);
            event->setAccepted(ev.Accept != 0);
            return ev.Filter;
        }
        case QEvent::Move: {
            QMoveEvent *e = (QMoveEvent*)event;
            const QPoint &pt = e->pos();
            const QPoint &opt = e->oldPos();
            MoveEvent ev = {accept,0,pt.x(),pt.y(),opt.x(),opt.y()};
            theApp.fnEventCall(obj,this,m_type,&ev,m_context);
            event->setAccepted(ev.Accept != 0);
            return ev.Filter;
        }
        case QEvent::Resize: {
            QResizeEvent *e = (QResizeEvent*)event;
            const QSize &sz = e->size();
            const QSize &osz = e->oldSize();
            ResizeEvent ev = {accept,0,sz.width(),sz.height(),osz.width(),osz.height()};
            theApp.fnEventCall(obj,this,m_type,&ev,m_context);
            event->setAccepted(ev.Accept != 0);
            return ev.Filter;
        }
        case QEvent::Paint: {
            QPaintEvent *e = (QPaintEvent*)event;
            const QRect &rc = e->rect();
            PaintEvent ev = {accept,rc.x(),rc.y(),rc.width(),rc.height()};
            theApp.fnEventCall(obj,this,m_type,&ev,m_context);
            event->setAccepted(ev.Accept != 0);
            return ev.Filter;
        }
        default:
            Event ev = {event->isAccepted(),0};
            theApp.fnEventCall(obj,this,m_type,&ev,m_context);
            event->setAccepted(ev.Accept != 0);
            return ev.Filter;
        }
    }
    return QObject::eventFilter(obj, event);
}
