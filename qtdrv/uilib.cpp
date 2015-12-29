#include "uilib.h"
#include "uiapplication.h"
#include "uisignal.h"
#include <QThread>
#include <QDebug>

void InitList(List *list, int size)
{
    list->count = size;
    list->data = new const void*[size];
}

void FreeList(List *list)
{
    delete[] list->data;
    list->count = 0;
    list->data = 0;
}


UIOBJECT Instance()
{
    return (UIOBJECT)&theApp;
}

int RunLoop()
{
    return theApp.exec();
}

void Quit()
{
    theApp.quit();
}

UIHANDLE CreateHandle(int id, void *param)
{
    return theApp.createHandle(id,param);
}

void  DestroyHandle(int id, UIHANDLE uih)
{
}

int  HandleFunction(int id, UIHANDLE uih, int fn, void *p1, void *p2, void *p3, void *p4, void *p5, void *p6)
{
    return 0;
}

UIOBJECT CreateObject(int id, UIOBJECT parent)
{
    return theApp.createObject(id,(QObject*)parent);
}

void ReleaseObject(int id, UIOBJECT obj)
{
    theApp.destroyObject(id, (QObject*)obj);
}

int  ObjectFunction(int id, void *obj, int fn, void *p1, void *p2, void *p3, void *p4, void *p5, void *p6)
{
    return 0;
}


int SetSignalCallback(SignalCallProc fc, SignalReleaseProc fr)
{
    if (fc != 0)
        theApp.fnSignalCall = fc;
    if (fr != 0)
        theApp.fnSignalRelease = fr;
    return 0;
}

UISIGNAL CreateSignal(UIOBJECT uio, int st, const char *signal, void *context)
{
    return theApp.createSignal((QObject*)uio,st,signal,context);
}

void EnableSignal(UISIGNAL uis, bool b)
{
    ((UISignal*)uis)->setEnable(b);
}


UIEVENT _CreateEvent(UIOBJECT object, int type, void *context)
{
    return theApp.createEvent((QObject*)object,type,context);
}


void SetEventCallback(EventCallProc fc, EventReleaseProc fr)
{
    if (fc != 0) {
        theApp.fnEventCall = fc;
    }
    if (fr != 0) {
        theApp.fnEventRelease = fr;
    }
}
