#ifndef UISIGNAL_H
#define UISIGNAL_H

#include <QObject>
#include <QAction>

class UISignal : public QObject
{
    Q_OBJECT
public:
    explicit UISignal(QObject *sender = 0, int type = 0);
    virtual ~UISignal();
    void setEnable(bool b);
public:
    static UISignal *createUISignal(QObject *sender, int type, const char *signal, void *context);
public slots:
    void call();
    void call(bool);
    void call(int);
    void call(const QString&);
    void call(QAction*);    
protected:
    int     m_type;
    bool    m_enable;
    const char *m_signal;
    const char *m_method;
    void    *m_context;
};

#endif // UISIGNAL_H
