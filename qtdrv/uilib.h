#ifndef UILIB_H
#define UILIB_H

#include "qtui_global.h"

typedef int    INT;
typedef double REAL;
typedef unsigned char BYTE;
typedef unsigned short USHORT;

typedef void *UIOBJECT;
typedef void *UIHANDLE;
typedef void *UISIGNAL;
typedef void *UIEVENT;

enum SignalType {
    ST_None,
    ST_Bool,    //bool&
    ST_Int,     //int&
    ST_String,  //UIString&
    ST_Action   //void*
};

typedef struct _ByteArray {
    const char *data;
    int size;
} ByteArray;

typedef struct _String {
    const ushort *data;
    int size;
} String;

typedef struct _List {
    const void **data;
    int count;
} List;

typedef struct _Point {
    INT X;
    INT Y;
} Point;

typedef struct _PointF {
    REAL X;
    REAL Y;
} PointF;

typedef struct _Size {
    INT Width;
    INT Height;
} Size;

typedef struct _SizeF {
    REAL Width;
    REAL Height;
} SizeF;

typedef struct _Rect {
    INT X;
    INT Y;
    INT Widht;
    INT Height;
} Rect;

typedef struct _RectF {
    REAL X;
    REAL Y;
    REAL Width;
    REAL Height;
} RectF;

typedef struct _PathData {
    INT Count;
    Point* Points;
    BYTE* Types;
} PathData;

typedef struct _PathDataF {
    INT Count;
    PointF* Points;
    BYTE* Types;
} PathDataF;

typedef struct _Event {
    int Accept;
    int Filter;
} Event;

typedef struct _TimerEvent {
    int Accept;
    int Filter;
    int Timerid;
} TimerEvent;

typedef struct _MoveEvent {
    int Accept;
    int Filter;
    int X;
    int Y;
    int OX;
    int OY;
} MoveEvent;

typedef struct _HoverEvent {
    int Accept;
    int Filter;
    int X;
    int Y;
    int OX;
    int OY;
} HoverEvent;

typedef struct ResizeEvent {
    int Accept;
    int Filter;
    int W;
    int H;
    int OW;
    int OH;
} ResizeEvent;

typedef struct _MouseEvent {
    int Accept;
    int Filter;
    int Modify;
    int Button;
    int Buttons;
    int GX;
    int GY;
    int X;
    int Y;
} MouseEvent;

typedef struct _KeyEvent {
    int Accept;
    int Filter;
    int Modify;
    int Count;
    int Autorepeat;
    int Key;
    quint32 NativeModifiers;
    quint32 NativeScanCode;
    quint32 NativeVirtualKey;
    String text;
} KeyEvent;

typedef struct _PaintEvent {
    int Accept;
    int Filter;
    int X;
    int Y;
    int W;
    int H;
} PaintEvent;

typedef struct _FocusEvent {
    int Accept;
    int Filter;
    int Reason;
} FocusEvent;

typedef void (*SignalCallProc)(UIOBJECT obj, UISIGNAL uis, int type, void *p1, void *p2, void *context);
typedef void (*SignalReleaseProc)(UISIGNAL uis, void *context);
typedef void (*EventCallProc)(UIOBJECT obj, UIEVENT uie, int type, void *event, void *context);
typedef void (*EventReleaseProc)(UIEVENT uie, void *context);

extern "C" {
QTUISHARED_EXPORT void InitList(List *list, int size);
QTUISHARED_EXPORT void FreeList(List *list);

//Application
QTUISHARED_EXPORT UIOBJECT Instance();
QTUISHARED_EXPORT int RunLoop();
QTUISHARED_EXPORT void Quit();

//widget signal
QTUISHARED_EXPORT int SetSignalCallback(SignalCallProc fc, SignalReleaseProc fr);
QTUISHARED_EXPORT UISIGNAL CreateSignal(UIOBJECT uio, int st, const char *signal, void *context);
QTUISHARED_EXPORT void  EnableSignal(UISIGNAL signal, bool b);

//safe thread ui object
QTUISHARED_EXPORT UIOBJECT CreateObject(int id, UIOBJECT parent);
QTUISHARED_EXPORT int ObjectFunction(int id, UIOBJECT object, int fn, void *p1, void *p2, void *p3, void *p4, void *p5, void *p6);
QTUISHARED_EXPORT void ReleaseObject(int id, UIOBJECT obj);

//widget event
QTUISHARED_EXPORT void SetEventCallback(EventCallProc fc, EventReleaseProc fr);
QTUISHARED_EXPORT UIEVENT _CreateEvent(UIOBJECT object, int type, void *context);

//util tool handle
QTUISHARED_EXPORT UIHANDLE CreateHandle(int id, void *param);
QTUISHARED_EXPORT void  DestroyHandle(int id, UIHANDLE uih);
QTUISHARED_EXPORT int   HandleFunction(int id, UIHANDLE uih, int fn, void *p1, void *p2, void *p3, void *p4, void *p5, void *p6);

}

#endif // UILIB_H
