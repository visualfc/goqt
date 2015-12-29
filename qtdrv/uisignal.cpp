#include "uisignal.h"
#include "uiapplication.h"
#include "uiutil.h"
#include <QDebug>

UISignal::UISignal(QObject *sender, int type) :
    QObject(sender), m_type(type), m_enable(true)
{
    this->setObjectName("UISignal");
}

UISignal::~UISignal()
{
    theApp.fnSignalRelease(this,m_context);
}

void UISignal::setEnable(bool b)
{
    if (m_enable == b) {
        return;
    }
    m_enable = b;
    if (m_enable) {
        QObject::connect(this->parent(),m_signal,this,m_method);
    } else {
        QObject::disconnect(this->parent(),m_signal,this,m_method);
    }
}

void UISignal::call()
{
    theApp.fnSignalCall(this->parent(),this,m_type,0,0,m_context);
}

void UISignal::call(bool v)
{
    theApp.fnSignalCall(this->parent(),this,m_type,&v,0,m_context);
}

void UISignal::call(int v)
{
    theApp.fnSignalCall(this->parent(),this,m_type,&v,0,m_context);
}

void UISignal::call(const QString& v)
{
    //String s = toString(v);
    //theApp.fnSignalCall(this->parent(),this,m_type,&s,0,m_context);
}

void UISignal::call(QAction* v)
{
    theApp.fnSignalCall(this->parent(),this,m_type,v,0,m_context);
}

UISignal *UISignal::createUISignal(QObject *sender, int type, const char *signal,void *context)
{
    UISignal *uis = new UISignal(sender);
    switch (type) {
    case ST_None:
        uis->m_method = SLOT(call());
        break;
    case ST_Int:
        uis->m_method = SLOT(call(int));
        break;
    case ST_Bool:
        uis->m_method = SLOT(call(bool));
    case ST_String:
        uis->m_method = SLOT(call(QString));
        break;
    case ST_Action:
        uis->m_method = SLOT(call(QAction*));
        break;
    default:
        goto end;
        break;
    }
    if (QObject::connect(sender,signal,uis,uis->m_method)) {
        uis->m_type = type;
        uis->m_signal = signal;
        uis->m_context = context;
        return uis;
    }
end:
    delete uis;
    return 0;
}
