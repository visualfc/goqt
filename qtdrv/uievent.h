#ifndef UIEVENT_H
#define UIEVENT_H

#include <QObject>

class UIEvent : public QObject
{
    Q_OBJECT
public:
    explicit UIEvent(QObject *parent = 0, int type = 0,void *context = 0);
    virtual ~UIEvent();
public:
    static UIEvent *createUIEvent(QObject *sender, int st, void *context);
protected:
    virtual bool eventFilter(QObject *, QEvent *);
    int m_type;
    void *m_context;
};

#endif // UIEVENT_H
