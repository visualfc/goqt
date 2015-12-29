CONFIG   += console
CONFIG   -= app_bundle

TEMPLATE = app
TARGET = goqt_rcc

QT += core xml
QT -= gui

DESTDIR = ../../bin
DEFINES += QT_RCC
INCLUDEPATH += .
DEPENDPATH += .

include(rcc.pri)
#HEADERS += ../../corelib/kernel/qcorecmdlineargs_p.h
SOURCES += main.cpp
#include(../bootstrap/bootstrap.pri)

#target.path=$$[QT_INSTALL_BINS]
#INSTALLS += target
#include(../../qt_targets.pri)
