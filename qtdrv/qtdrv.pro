#-------------------------------------------------
#
# Project created by QtCreator
#
#-------------------------------------------------
greaterThan(QT_MAJOR_VERSION, 4) {
    QT += widgets uitools printsupport
} else {
    QT += core gui
    CONFIG += uitools
}

TARGET = qtdrv.ui
TEMPLATE = lib

BUILD_SOURCE_TREE = $$PWD
BUILD_LIB_PATH = $$BUILD_SOURCE_TREE/../bin

DESTDIR = $$BUILD_LIB_PATH

DEFINES += QTDRV_LIBRARY

SOURCES += uiapplication.cpp \
    cdrv.cpp \
    callback.cpp \
    cdrv_event.cpp \
    uidrv.cpp \
    cdrv_model.cpp \
    qsyntaxhighlighterhook.cpp


HEADERS += qtdrv_global.h \
    uiapplication.h \
    cdrv.h  \
    cdrv_signal.h \
    callback.h \
    cdrv_event.h \
    uidrv.h \
    cdrv_model.h \
    qsyntaxhighlighterhook.h


greaterThan(QT_MAJOR_VERSION, 4) {
} else {
    HEADERS += qurl/url_help.h \
               qurl/qurlquery.h

    SOURCES += qurl/qurlquery.cpp \
               qurl/qurlrecode.cpp

}

unix:!symbian {
    maemo5 {
        target.path = /opt/usr/lib
    } else {
        target.path = /usr/lib
    }
    INSTALLS += target
}
