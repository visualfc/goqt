/********************************************************************
** Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
**
** This library is free software; you can redistribute it and/or
** modify it under the terms of the GNU Lesser General Public
** License as published by the Free Software Foundation; either
** version 2.1 of the License, or (at your option) any later version.
**
** This library is distributed in the hope that it will be useful,
** but WITHOUT ANY WARRANTY; without even the implied warranty of
** MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
** Lesser General Public License for more details.
*********************************************************************/

#ifndef CDRV_SIGNAL_H
#define CDRV_SIGNAL_H	
#include "cdrv.h"	

//go func()
class UISignal1 : public QObject
{
	Q_OBJECT
public:
    explicit UISignal1(QObject *sender, Iface *_face) : face(*_face)
	{
		QObject::connect(sender,SIGNAL(destroyed()),this,SLOT(deleteLater()));
	}
	~UISignal1()
	{
		drv_remove_signal_call(this);
	}
public slots:
	void call() {
		drv_signal_call(this,&face,1,0,0,0,0);
	}
public:
	Iface face;	
};
//go func(*QAbstractAnimation)
class UISignal2 : public QObject
{
	Q_OBJECT
public:
    explicit UISignal2(QObject *sender, Iface *_face) : face(*_face)
	{
		QObject::connect(sender,SIGNAL(destroyed()),this,SLOT(deleteLater()));
	}
	~UISignal2()
	{
		drv_remove_signal_call(this);
	}
public slots:
	void call(QAbstractAnimation* v1) {
		drv_signal_call(this,&face,2,v1,0,0,0);
	}
public:
	Iface face;	
};
//go func(QAbstractAnimation_Direction)
class UISignal3 : public QObject
{
	Q_OBJECT
public:
    explicit UISignal3(QObject *sender, Iface *_face) : face(*_face)
	{
		QObject::connect(sender,SIGNAL(destroyed()),this,SLOT(deleteLater()));
	}
	~UISignal3()
	{
		drv_remove_signal_call(this);
	}
public slots:
	void call(QAbstractAnimation::Direction v1) {
		drv_signal_call(this,&face,3,&v1,0,0,0);
	}
public:
	Iface face;	
};
//go func(QAbstractAnimation_State,QAbstractAnimation_State)
class UISignal4 : public QObject
{
	Q_OBJECT
public:
    explicit UISignal4(QObject *sender, Iface *_face) : face(*_face)
	{
		QObject::connect(sender,SIGNAL(destroyed()),this,SLOT(deleteLater()));
	}
	~UISignal4()
	{
		drv_remove_signal_call(this);
	}
public slots:
	void call(QAbstractAnimation::State v1,QAbstractAnimation::State v2) {
		drv_signal_call(this,&face,4,&v1,&v2,0,0);
	}
public:
	Iface face;	
};
//go func(*QAbstractButton)
class UISignal5 : public QObject
{
	Q_OBJECT
public:
    explicit UISignal5(QObject *sender, Iface *_face) : face(*_face)
	{
		QObject::connect(sender,SIGNAL(destroyed()),this,SLOT(deleteLater()));
	}
	~UISignal5()
	{
		drv_remove_signal_call(this);
	}
public slots:
	void call(QAbstractButton* v1) {
		drv_signal_call(this,&face,5,v1,0,0,0);
	}
public:
	Iface face;	
};
//go func(*QAction)
class UISignal6 : public QObject
{
	Q_OBJECT
public:
    explicit UISignal6(QObject *sender, Iface *_face) : face(*_face)
	{
		QObject::connect(sender,SIGNAL(destroyed()),this,SLOT(deleteLater()));
	}
	~UISignal6()
	{
		drv_remove_signal_call(this);
	}
public slots:
	void call(QAction* v1) {
		drv_signal_call(this,&face,6,v1,0,0,0);
	}
public:
	Iface face;	
};
//go func(*QBrush)
class UISignal7 : public QObject
{
	Q_OBJECT
public:
    explicit UISignal7(QObject *sender, Iface *_face) : face(*_face)
	{
		QObject::connect(sender,SIGNAL(destroyed()),this,SLOT(deleteLater()));
	}
	~UISignal7()
	{
		drv_remove_signal_call(this);
	}
public slots:
	void call(QBrush const& v1) {
		drv_signal_call(this,&face,7,new QBrush(v1),0,0,0);
	}
public:
	Iface face;	
};
//go func(QClipboard_Mode)
class UISignal8 : public QObject
{
	Q_OBJECT
public:
    explicit UISignal8(QObject *sender, Iface *_face) : face(*_face)
	{
		QObject::connect(sender,SIGNAL(destroyed()),this,SLOT(deleteLater()));
	}
	~UISignal8()
	{
		drv_remove_signal_call(this);
	}
public slots:
	void call(QClipboard::Mode v1) {
		drv_signal_call(this,&face,8,&v1,0,0,0);
	}
public:
	Iface face;	
};
//go func(*QColor)
class UISignal9 : public QObject
{
	Q_OBJECT
public:
    explicit UISignal9(QObject *sender, Iface *_face) : face(*_face)
	{
		QObject::connect(sender,SIGNAL(destroyed()),this,SLOT(deleteLater()));
	}
	~UISignal9()
	{
		drv_remove_signal_call(this);
	}
public slots:
	void call(QColor const& v1) {
		drv_signal_call(this,&face,9,new QColor(v1),0,0,0);
	}
public:
	Iface face;	
};
//go func(*QDate)
class UISignal10 : public QObject
{
	Q_OBJECT
public:
    explicit UISignal10(QObject *sender, Iface *_face) : face(*_face)
	{
		QObject::connect(sender,SIGNAL(destroyed()),this,SLOT(deleteLater()));
	}
	~UISignal10()
	{
		drv_remove_signal_call(this);
	}
public slots:
	void call(QDate const& v1) {
		drv_signal_call(this,&face,10,new QDate(v1),0,0,0);
	}
public:
	Iface face;	
};
//go func(*QDateTime)
class UISignal11 : public QObject
{
	Q_OBJECT
public:
    explicit UISignal11(QObject *sender, Iface *_face) : face(*_face)
	{
		QObject::connect(sender,SIGNAL(destroyed()),this,SLOT(deleteLater()));
	}
	~UISignal11()
	{
		drv_remove_signal_call(this);
	}
public slots:
	void call(QDateTime const& v1) {
		drv_signal_call(this,&face,11,new QDateTime(v1),0,0,0);
	}
public:
	Iface face;	
};
//go func(QDockWidget_DockWidgetFeature)
class UISignal12 : public QObject
{
	Q_OBJECT
public:
    explicit UISignal12(QObject *sender, Iface *_face) : face(*_face)
	{
		QObject::connect(sender,SIGNAL(destroyed()),this,SLOT(deleteLater()));
	}
	~UISignal12()
	{
		drv_remove_signal_call(this);
	}
public slots:
	void call(QFlags<QDockWidget::DockWidgetFeature> v1) {
		drv_signal_call(this,&face,12,&v1,0,0,0);
	}
public:
	Iface face;	
};
//go func(QGraphicsBlurEffect_BlurHint)
class UISignal13 : public QObject
{
	Q_OBJECT
public:
    explicit UISignal13(QObject *sender, Iface *_face) : face(*_face)
	{
		QObject::connect(sender,SIGNAL(destroyed()),this,SLOT(deleteLater()));
	}
	~UISignal13()
	{
		drv_remove_signal_call(this);
	}
public slots:
	void call(QFlags<QGraphicsBlurEffect::BlurHint> v1) {
		drv_signal_call(this,&face,13,&v1,0,0,0);
	}
public:
	Iface face;	
};
//go func(Qt_DockWidgetArea)
class UISignal14 : public QObject
{
	Q_OBJECT
public:
    explicit UISignal14(QObject *sender, Iface *_face) : face(*_face)
	{
		QObject::connect(sender,SIGNAL(destroyed()),this,SLOT(deleteLater()));
	}
	~UISignal14()
	{
		drv_remove_signal_call(this);
	}
public slots:
	void call(QFlags<Qt::DockWidgetArea> v1) {
		drv_signal_call(this,&face,14,&v1,0,0,0);
	}
public:
	Iface face;	
};
//go func(Qt_ToolBarArea)
class UISignal15 : public QObject
{
	Q_OBJECT
public:
    explicit UISignal15(QObject *sender, Iface *_face) : face(*_face)
	{
		QObject::connect(sender,SIGNAL(destroyed()),this,SLOT(deleteLater()));
	}
	~UISignal15()
	{
		drv_remove_signal_call(this);
	}
public slots:
	void call(QFlags<Qt::ToolBarArea> v1) {
		drv_signal_call(this,&face,15,&v1,0,0,0);
	}
public:
	Iface face;	
};
//go func(Qt_WindowState,Qt_WindowState)
class UISignal16 : public QObject
{
	Q_OBJECT
public:
    explicit UISignal16(QObject *sender, Iface *_face) : face(*_face)
	{
		QObject::connect(sender,SIGNAL(destroyed()),this,SLOT(deleteLater()));
	}
	~UISignal16()
	{
		drv_remove_signal_call(this);
	}
public slots:
	void call(QFlags<Qt::WindowState> v1,QFlags<Qt::WindowState> v2) {
		drv_signal_call(this,&face,16,&v1,&v2,0,0);
	}
public:
	Iface face;	
};
//go func(*QFont)
class UISignal17 : public QObject
{
	Q_OBJECT
public:
    explicit UISignal17(QObject *sender, Iface *_face) : face(*_face)
	{
		QObject::connect(sender,SIGNAL(destroyed()),this,SLOT(deleteLater()));
	}
	~UISignal17()
	{
		drv_remove_signal_call(this);
	}
public slots:
	void call(QFont const& v1) {
		drv_signal_call(this,&face,17,new QFont(v1),0,0,0);
	}
public:
	Iface face;	
};
//go func(QImageReader_ImageReaderError)
class UISignal18 : public QObject
{
	Q_OBJECT
public:
    explicit UISignal18(QObject *sender, Iface *_face) : face(*_face)
	{
		QObject::connect(sender,SIGNAL(destroyed()),this,SLOT(deleteLater()));
	}
	~UISignal18()
	{
		drv_remove_signal_call(this);
	}
public slots:
	void call(QImageReader::ImageReaderError v1) {
		drv_signal_call(this,&face,18,&v1,0,0,0);
	}
public:
	Iface face;	
};
//go func(*QItemSelection,*QItemSelection)
class UISignal19 : public QObject
{
	Q_OBJECT
public:
    explicit UISignal19(QObject *sender, Iface *_face) : face(*_face)
	{
		QObject::connect(sender,SIGNAL(destroyed()),this,SLOT(deleteLater()));
	}
	~UISignal19()
	{
		drv_remove_signal_call(this);
	}
public slots:
	void call(QItemSelection const& v1,QItemSelection const& v2) {
		drv_signal_call(this,&face,19,new QItemSelection(v1),new QItemSelection(v2),0,0);
	}
public:
	Iface face;	
};
//go func([]*QModelIndex)
class UISignal20 : public QObject
{
	Q_OBJECT
public:
    explicit UISignal20(QObject *sender, Iface *_face) : face(*_face)
	{
		QObject::connect(sender,SIGNAL(destroyed()),this,SLOT(deleteLater()));
	}
	~UISignal20()
	{
		drv_remove_signal_call(this);
	}
public slots:
	void call(QList<QModelIndex> const& v1) {
		drv_signal_call(this,&face,20,new QList<QModelIndex>(v1),0,0,0);
	}
public:
	Iface face;	
};
//go func([]*QRectF)
class UISignal21 : public QObject
{
	Q_OBJECT
public:
    explicit UISignal21(QObject *sender, Iface *_face) : face(*_face)
	{
		QObject::connect(sender,SIGNAL(destroyed()),this,SLOT(deleteLater()));
	}
	~UISignal21()
	{
		drv_remove_signal_call(this);
	}
public slots:
	void call(QList<QRectF> const& v1) {
		drv_signal_call(this,&face,21,new QList<QRectF>(v1),0,0,0);
	}
public:
	Iface face;	
};
//go func(*QListWidgetItem)
class UISignal22 : public QObject
{
	Q_OBJECT
public:
    explicit UISignal22(QObject *sender, Iface *_face) : face(*_face)
	{
		QObject::connect(sender,SIGNAL(destroyed()),this,SLOT(deleteLater()));
	}
	~UISignal22()
	{
		drv_remove_signal_call(this);
	}
public slots:
	void call(QListWidgetItem* v1) {
		drv_signal_call(this,&face,22,v1,0,0,0);
	}
public:
	Iface face;	
};
//go func(*QListWidgetItem,*QListWidgetItem)
class UISignal23 : public QObject
{
	Q_OBJECT
public:
    explicit UISignal23(QObject *sender, Iface *_face) : face(*_face)
	{
		QObject::connect(sender,SIGNAL(destroyed()),this,SLOT(deleteLater()));
	}
	~UISignal23()
	{
		drv_remove_signal_call(this);
	}
public slots:
	void call(QListWidgetItem* v1,QListWidgetItem* v2) {
		drv_signal_call(this,&face,23,v1,v2,0,0);
	}
public:
	Iface face;	
};
//go func(*QMdiSubWindow)
class UISignal24 : public QObject
{
	Q_OBJECT
public:
    explicit UISignal24(QObject *sender, Iface *_face) : face(*_face)
	{
		QObject::connect(sender,SIGNAL(destroyed()),this,SLOT(deleteLater()));
	}
	~UISignal24()
	{
		drv_remove_signal_call(this);
	}
public slots:
	void call(QMdiSubWindow* v1) {
		drv_signal_call(this,&face,24,v1,0,0,0);
	}
public:
	Iface face;	
};
//go func(*QModelIndex)
class UISignal25 : public QObject
{
	Q_OBJECT
public:
    explicit UISignal25(QObject *sender, Iface *_face) : face(*_face)
	{
		QObject::connect(sender,SIGNAL(destroyed()),this,SLOT(deleteLater()));
	}
	~UISignal25()
	{
		drv_remove_signal_call(this);
	}
public slots:
	void call(QModelIndex const& v1) {
		drv_signal_call(this,&face,25,new QModelIndex(v1),0,0,0);
	}
public:
	Iface face;	
};
//go func(*QModelIndex,*QModelIndex)
class UISignal26 : public QObject
{
	Q_OBJECT
public:
    explicit UISignal26(QObject *sender, Iface *_face) : face(*_face)
	{
		QObject::connect(sender,SIGNAL(destroyed()),this,SLOT(deleteLater()));
	}
	~UISignal26()
	{
		drv_remove_signal_call(this);
	}
public slots:
	void call(QModelIndex const& v1,QModelIndex const& v2) {
		drv_signal_call(this,&face,26,new QModelIndex(v1),new QModelIndex(v2),0,0);
	}
public:
	Iface face;	
};
//go func(QMovie_MovieState)
class UISignal27 : public QObject
{
	Q_OBJECT
public:
    explicit UISignal27(QObject *sender, Iface *_face) : face(*_face)
	{
		QObject::connect(sender,SIGNAL(destroyed()),this,SLOT(deleteLater()));
	}
	~UISignal27()
	{
		drv_remove_signal_call(this);
	}
public slots:
	void call(QMovie::MovieState v1) {
		drv_signal_call(this,&face,27,&v1,0,0,0);
	}
public:
	Iface face;	
};
//go func(*QObject)
class UISignal28 : public QObject
{
	Q_OBJECT
public:
    explicit UISignal28(QObject *sender, Iface *_face) : face(*_face)
	{
		QObject::connect(sender,SIGNAL(destroyed()),this,SLOT(deleteLater()));
	}
	~UISignal28()
	{
		drv_remove_signal_call(this);
	}
public slots:
	void call(QObject* v1) {
		drv_signal_call(this,&face,28,v1,0,0,0);
	}
public:
	Iface face;	
};
//go func(*QPoint)
class UISignal29 : public QObject
{
	Q_OBJECT
public:
    explicit UISignal29(QObject *sender, Iface *_face) : face(*_face)
	{
		QObject::connect(sender,SIGNAL(destroyed()),this,SLOT(deleteLater()));
	}
	~UISignal29()
	{
		drv_remove_signal_call(this);
	}
public slots:
	void call(QPoint const& v1) {
		drv_signal_call(this,&face,29,new QPoint(v1),0,0,0);
	}
public:
	Iface face;	
};
//go func(*QPointF)
class UISignal30 : public QObject
{
	Q_OBJECT
public:
    explicit UISignal30(QObject *sender, Iface *_face) : face(*_face)
	{
		QObject::connect(sender,SIGNAL(destroyed()),this,SLOT(deleteLater()));
	}
	~UISignal30()
	{
		drv_remove_signal_call(this);
	}
public slots:
	void call(QPointF const& v1) {
		drv_signal_call(this,&face,30,new QPointF(v1),0,0,0);
	}
public:
	Iface face;	
};
//go func(*QPrinter)
class UISignal31 : public QObject
{
	Q_OBJECT
public:
    explicit UISignal31(QObject *sender, Iface *_face) : face(*_face)
	{
		QObject::connect(sender,SIGNAL(destroyed()),this,SLOT(deleteLater()));
	}
	~UISignal31()
	{
		drv_remove_signal_call(this);
	}
public slots:
	void call(QPrinter* v1) {
		drv_signal_call(this,&face,31,v1,0,0,0);
	}
public:
	Iface face;	
};
//go func(QProcess_ProcessError)
class UISignal32 : public QObject
{
	Q_OBJECT
public:
    explicit UISignal32(QObject *sender, Iface *_face) : face(*_face)
	{
		QObject::connect(sender,SIGNAL(destroyed()),this,SLOT(deleteLater()));
	}
	~UISignal32()
	{
		drv_remove_signal_call(this);
	}
public slots:
	void call(QProcess::ProcessError v1) {
		drv_signal_call(this,&face,32,&v1,0,0,0);
	}
public:
	Iface face;	
};
//go func(QProcess_ProcessState)
class UISignal33 : public QObject
{
	Q_OBJECT
public:
    explicit UISignal33(QObject *sender, Iface *_face) : face(*_face)
	{
		QObject::connect(sender,SIGNAL(destroyed()),this,SLOT(deleteLater()));
	}
	~UISignal33()
	{
		drv_remove_signal_call(this);
	}
public slots:
	void call(QProcess::ProcessState v1) {
		drv_signal_call(this,&face,33,&v1,0,0,0);
	}
public:
	Iface face;	
};
//go func(*QRect)
class UISignal34 : public QObject
{
	Q_OBJECT
public:
    explicit UISignal34(QObject *sender, Iface *_face) : face(*_face)
	{
		QObject::connect(sender,SIGNAL(destroyed()),this,SLOT(deleteLater()));
	}
	~UISignal34()
	{
		drv_remove_signal_call(this);
	}
public slots:
	void call(QRect const& v1) {
		drv_signal_call(this,&face,34,new QRect(v1),0,0,0);
	}
public:
	Iface face;	
};
//go func(*QRect,int)
class UISignal35 : public QObject
{
	Q_OBJECT
public:
    explicit UISignal35(QObject *sender, Iface *_face) : face(*_face)
	{
		QObject::connect(sender,SIGNAL(destroyed()),this,SLOT(deleteLater()));
	}
	~UISignal35()
	{
		drv_remove_signal_call(this);
	}
public slots:
	void call(QRect const& v1,int v2) {
		drv_signal_call(this,&face,35,new QRect(v1),&v2,0,0);
	}
public:
	Iface face;	
};
//go func(*QRectF)
class UISignal36 : public QObject
{
	Q_OBJECT
public:
    explicit UISignal36(QObject *sender, Iface *_face) : face(*_face)
	{
		QObject::connect(sender,SIGNAL(destroyed()),this,SLOT(deleteLater()));
	}
	~UISignal36()
	{
		drv_remove_signal_call(this);
	}
public slots:
	void call(QRectF const& v1) {
		drv_signal_call(this,&face,36,new QRectF(v1),0,0,0);
	}
public:
	Iface face;	
};
//go func(*QSessionManager)
class UISignal37 : public QObject
{
	Q_OBJECT
public:
    explicit UISignal37(QObject *sender, Iface *_face) : face(*_face)
	{
		QObject::connect(sender,SIGNAL(destroyed()),this,SLOT(deleteLater()));
	}
	~UISignal37()
	{
		drv_remove_signal_call(this);
	}
public slots:
	void call(QSessionManager& v1) {
		drv_signal_call(this,&face,37,&v1,0,0,0);
	}
public:
	Iface face;	
};
//go func(*QSize)
class UISignal38 : public QObject
{
	Q_OBJECT
public:
    explicit UISignal38(QObject *sender, Iface *_face) : face(*_face)
	{
		QObject::connect(sender,SIGNAL(destroyed()),this,SLOT(deleteLater()));
	}
	~UISignal38()
	{
		drv_remove_signal_call(this);
	}
public slots:
	void call(QSize const& v1) {
		drv_signal_call(this,&face,38,new QSize(v1),0,0,0);
	}
public:
	Iface face;	
};
//go func(*QSizeF)
class UISignal39 : public QObject
{
	Q_OBJECT
public:
    explicit UISignal39(QObject *sender, Iface *_face) : face(*_face)
	{
		QObject::connect(sender,SIGNAL(destroyed()),this,SLOT(deleteLater()));
	}
	~UISignal39()
	{
		drv_remove_signal_call(this);
	}
public slots:
	void call(QSizeF const& v1) {
		drv_signal_call(this,&face,39,new QSizeF(v1),0,0,0);
	}
public:
	Iface face;	
};
//go func(*QStandardItem)
class UISignal40 : public QObject
{
	Q_OBJECT
public:
    explicit UISignal40(QObject *sender, Iface *_face) : face(*_face)
	{
		QObject::connect(sender,SIGNAL(destroyed()),this,SLOT(deleteLater()));
	}
	~UISignal40()
	{
		drv_remove_signal_call(this);
	}
public slots:
	void call(QStandardItem* v1) {
		drv_signal_call(this,&face,40,v1,0,0,0);
	}
public:
	Iface face;	
};
//go func(string)
class UISignal41 : public QObject
{
	Q_OBJECT
public:
    explicit UISignal41(QObject *sender, Iface *_face) : face(*_face)
	{
		QObject::connect(sender,SIGNAL(destroyed()),this,SLOT(deleteLater()));
	}
	~UISignal41()
	{
		drv_remove_signal_call(this);
	}
public slots:
	void call(QString const& v1) {
		drv_signal_call(this,&face,41,new QString(v1),0,0,0);
	}
public:
	Iface face;	
};
//go func(string,string,string)
class UISignal42 : public QObject
{
	Q_OBJECT
public:
    explicit UISignal42(QObject *sender, Iface *_face) : face(*_face)
	{
		QObject::connect(sender,SIGNAL(destroyed()),this,SLOT(deleteLater()));
	}
	~UISignal42()
	{
		drv_remove_signal_call(this);
	}
public slots:
	void call(QString const& v1,QString const& v2,QString const& v3) {
		drv_signal_call(this,&face,42,new QString(v1),new QString(v2),new QString(v3),0);
	}
public:
	Iface face;	
};
//go func([]string)
class UISignal43 : public QObject
{
	Q_OBJECT
public:
    explicit UISignal43(QObject *sender, Iface *_face) : face(*_face)
	{
		QObject::connect(sender,SIGNAL(destroyed()),this,SLOT(deleteLater()));
	}
	~UISignal43()
	{
		drv_remove_signal_call(this);
	}
public slots:
	void call(QStringList const& v1) {
		drv_signal_call(this,&face,43,new QStringList(v1),0,0,0);
	}
public:
	Iface face;	
};
//go func(QSystemTrayIcon_ActivationReason)
class UISignal44 : public QObject
{
	Q_OBJECT
public:
    explicit UISignal44(QObject *sender, Iface *_face) : face(*_face)
	{
		QObject::connect(sender,SIGNAL(destroyed()),this,SLOT(deleteLater()));
	}
	~UISignal44()
	{
		drv_remove_signal_call(this);
	}
public slots:
	void call(QSystemTrayIcon::ActivationReason v1) {
		drv_signal_call(this,&face,44,&v1,0,0,0);
	}
public:
	Iface face;	
};
//go func(*QTableWidgetItem)
class UISignal45 : public QObject
{
	Q_OBJECT
public:
    explicit UISignal45(QObject *sender, Iface *_face) : face(*_face)
	{
		QObject::connect(sender,SIGNAL(destroyed()),this,SLOT(deleteLater()));
	}
	~UISignal45()
	{
		drv_remove_signal_call(this);
	}
public slots:
	void call(QTableWidgetItem* v1) {
		drv_signal_call(this,&face,45,v1,0,0,0);
	}
public:
	Iface face;	
};
//go func(*QTableWidgetItem,*QTableWidgetItem)
class UISignal46 : public QObject
{
	Q_OBJECT
public:
    explicit UISignal46(QObject *sender, Iface *_face) : face(*_face)
	{
		QObject::connect(sender,SIGNAL(destroyed()),this,SLOT(deleteLater()));
	}
	~UISignal46()
	{
		drv_remove_signal_call(this);
	}
public slots:
	void call(QTableWidgetItem* v1,QTableWidgetItem* v2) {
		drv_signal_call(this,&face,46,v1,v2,0,0);
	}
public:
	Iface face;	
};
//go func(*QTextBlock)
class UISignal47 : public QObject
{
	Q_OBJECT
public:
    explicit UISignal47(QObject *sender, Iface *_face) : face(*_face)
	{
		QObject::connect(sender,SIGNAL(destroyed()),this,SLOT(deleteLater()));
	}
	~UISignal47()
	{
		drv_remove_signal_call(this);
	}
public slots:
	void call(QTextBlock const& v1) {
		drv_signal_call(this,&face,47,new QTextBlock(v1),0,0,0);
	}
public:
	Iface face;	
};
//go func(*QTextCharFormat)
class UISignal48 : public QObject
{
	Q_OBJECT
public:
    explicit UISignal48(QObject *sender, Iface *_face) : face(*_face)
	{
		QObject::connect(sender,SIGNAL(destroyed()),this,SLOT(deleteLater()));
	}
	~UISignal48()
	{
		drv_remove_signal_call(this);
	}
public slots:
	void call(QTextCharFormat const& v1) {
		drv_signal_call(this,&face,48,new QTextCharFormat(v1),0,0,0);
	}
public:
	Iface face;	
};
//go func(*QTextCursor)
class UISignal49 : public QObject
{
	Q_OBJECT
public:
    explicit UISignal49(QObject *sender, Iface *_face) : face(*_face)
	{
		QObject::connect(sender,SIGNAL(destroyed()),this,SLOT(deleteLater()));
	}
	~UISignal49()
	{
		drv_remove_signal_call(this);
	}
public slots:
	void call(QTextCursor const& v1) {
		drv_signal_call(this,&face,49,new QTextCursor(v1),0,0,0);
	}
public:
	Iface face;	
};
//go func(*QTime)
class UISignal50 : public QObject
{
	Q_OBJECT
public:
    explicit UISignal50(QObject *sender, Iface *_face) : face(*_face)
	{
		QObject::connect(sender,SIGNAL(destroyed()),this,SLOT(deleteLater()));
	}
	~UISignal50()
	{
		drv_remove_signal_call(this);
	}
public slots:
	void call(QTime const& v1) {
		drv_signal_call(this,&face,50,new QTime(v1),0,0,0);
	}
public:
	Iface face;	
};
//go func(QTimeLine_State)
class UISignal51 : public QObject
{
	Q_OBJECT
public:
    explicit UISignal51(QObject *sender, Iface *_face) : face(*_face)
	{
		QObject::connect(sender,SIGNAL(destroyed()),this,SLOT(deleteLater()));
	}
	~UISignal51()
	{
		drv_remove_signal_call(this);
	}
public slots:
	void call(QTimeLine::State v1) {
		drv_signal_call(this,&face,51,&v1,0,0,0);
	}
public:
	Iface face;	
};
//go func(*QTreeWidgetItem)
class UISignal52 : public QObject
{
	Q_OBJECT
public:
    explicit UISignal52(QObject *sender, Iface *_face) : face(*_face)
	{
		QObject::connect(sender,SIGNAL(destroyed()),this,SLOT(deleteLater()));
	}
	~UISignal52()
	{
		drv_remove_signal_call(this);
	}
public slots:
	void call(QTreeWidgetItem* v1) {
		drv_signal_call(this,&face,52,v1,0,0,0);
	}
public:
	Iface face;	
};
//go func(*QTreeWidgetItem,*QTreeWidgetItem)
class UISignal53 : public QObject
{
	Q_OBJECT
public:
    explicit UISignal53(QObject *sender, Iface *_face) : face(*_face)
	{
		QObject::connect(sender,SIGNAL(destroyed()),this,SLOT(deleteLater()));
	}
	~UISignal53()
	{
		drv_remove_signal_call(this);
	}
public slots:
	void call(QTreeWidgetItem* v1,QTreeWidgetItem* v2) {
		drv_signal_call(this,&face,53,v1,v2,0,0);
	}
public:
	Iface face;	
};
//go func(*QTreeWidgetItem,int)
class UISignal54 : public QObject
{
	Q_OBJECT
public:
    explicit UISignal54(QObject *sender, Iface *_face) : face(*_face)
	{
		QObject::connect(sender,SIGNAL(destroyed()),this,SLOT(deleteLater()));
	}
	~UISignal54()
	{
		drv_remove_signal_call(this);
	}
public slots:
	void call(QTreeWidgetItem* v1,int v2) {
		drv_signal_call(this,&face,54,v1,&v2,0,0);
	}
public:
	Iface face;	
};
//go func(*QUndoStack)
class UISignal55 : public QObject
{
	Q_OBJECT
public:
    explicit UISignal55(QObject *sender, Iface *_face) : face(*_face)
	{
		QObject::connect(sender,SIGNAL(destroyed()),this,SLOT(deleteLater()));
	}
	~UISignal55()
	{
		drv_remove_signal_call(this);
	}
public slots:
	void call(QUndoStack* v1) {
		drv_signal_call(this,&face,55,v1,0,0,0);
	}
public:
	Iface face;	
};
//go func(*QUrl)
class UISignal56 : public QObject
{
	Q_OBJECT
public:
    explicit UISignal56(QObject *sender, Iface *_face) : face(*_face)
	{
		QObject::connect(sender,SIGNAL(destroyed()),this,SLOT(deleteLater()));
	}
	~UISignal56()
	{
		drv_remove_signal_call(this);
	}
public slots:
	void call(QUrl const& v1) {
		drv_signal_call(this,&face,56,new QUrl(v1),0,0,0);
	}
public:
	Iface face;	
};
//go func(*QVariant)
class UISignal57 : public QObject
{
	Q_OBJECT
public:
    explicit UISignal57(QObject *sender, Iface *_face) : face(*_face)
	{
		QObject::connect(sender,SIGNAL(destroyed()),this,SLOT(deleteLater()));
	}
	~UISignal57()
	{
		drv_remove_signal_call(this);
	}
public slots:
	void call(QVariant const& v1) {
		drv_signal_call(this,&face,57,new QVariant(v1),0,0,0);
	}
public:
	Iface face;	
};
//go func(*QWidget)
class UISignal58 : public QObject
{
	Q_OBJECT
public:
    explicit UISignal58(QObject *sender, Iface *_face) : face(*_face)
	{
		QObject::connect(sender,SIGNAL(destroyed()),this,SLOT(deleteLater()));
	}
	~UISignal58()
	{
		drv_remove_signal_call(this);
	}
public slots:
	void call(QWidget* v1) {
		drv_signal_call(this,&face,58,v1,0,0,0);
	}
public:
	Iface face;	
};
//go func(*QWidget,QAbstractItemDelegate_EndEditHint)
class UISignal59 : public QObject
{
	Q_OBJECT
public:
    explicit UISignal59(QObject *sender, Iface *_face) : face(*_face)
	{
		QObject::connect(sender,SIGNAL(destroyed()),this,SLOT(deleteLater()));
	}
	~UISignal59()
	{
		drv_remove_signal_call(this);
	}
public slots:
	void call(QWidget* v1,QAbstractItemDelegate::EndEditHint v2) {
		drv_signal_call(this,&face,59,v1,&v2,0,0);
	}
public:
	Iface face;	
};
//go func(*QWidget,*QWidget)
class UISignal60 : public QObject
{
	Q_OBJECT
public:
    explicit UISignal60(QObject *sender, Iface *_face) : face(*_face)
	{
		QObject::connect(sender,SIGNAL(destroyed()),this,SLOT(deleteLater()));
	}
	~UISignal60()
	{
		drv_remove_signal_call(this);
	}
public slots:
	void call(QWidget* v1,QWidget* v2) {
		drv_signal_call(this,&face,60,v1,v2,0,0);
	}
public:
	Iface face;	
};
//go func(Qt_DockWidgetArea)
class UISignal61 : public QObject
{
	Q_OBJECT
public:
    explicit UISignal61(QObject *sender, Iface *_face) : face(*_face)
	{
		QObject::connect(sender,SIGNAL(destroyed()),this,SLOT(deleteLater()));
	}
	~UISignal61()
	{
		drv_remove_signal_call(this);
	}
public slots:
	void call(Qt::DockWidgetArea v1) {
		drv_signal_call(this,&face,61,&v1,0,0,0);
	}
public:
	Iface face;	
};
//go func(Qt_DropAction)
class UISignal62 : public QObject
{
	Q_OBJECT
public:
    explicit UISignal62(QObject *sender, Iface *_face) : face(*_face)
	{
		QObject::connect(sender,SIGNAL(destroyed()),this,SLOT(deleteLater()));
	}
	~UISignal62()
	{
		drv_remove_signal_call(this);
	}
public slots:
	void call(Qt::DropAction v1) {
		drv_signal_call(this,&face,62,&v1,0,0,0);
	}
public:
	Iface face;	
};
//go func(Qt_Orientation)
class UISignal63 : public QObject
{
	Q_OBJECT
public:
    explicit UISignal63(QObject *sender, Iface *_face) : face(*_face)
	{
		QObject::connect(sender,SIGNAL(destroyed()),this,SLOT(deleteLater()));
	}
	~UISignal63()
	{
		drv_remove_signal_call(this);
	}
public slots:
	void call(Qt::Orientation v1) {
		drv_signal_call(this,&face,63,&v1,0,0,0);
	}
public:
	Iface face;	
};
//go func(Qt_Orientation,int,int)
class UISignal64 : public QObject
{
	Q_OBJECT
public:
    explicit UISignal64(QObject *sender, Iface *_face) : face(*_face)
	{
		QObject::connect(sender,SIGNAL(destroyed()),this,SLOT(deleteLater()));
	}
	~UISignal64()
	{
		drv_remove_signal_call(this);
	}
public slots:
	void call(Qt::Orientation v1,int v2,int v3) {
		drv_signal_call(this,&face,64,&v1,&v2,&v3,0);
	}
public:
	Iface face;	
};
//go func(Qt_ToolButtonStyle)
class UISignal65 : public QObject
{
	Q_OBJECT
public:
    explicit UISignal65(QObject *sender, Iface *_face) : face(*_face)
	{
		QObject::connect(sender,SIGNAL(destroyed()),this,SLOT(deleteLater()));
	}
	~UISignal65()
	{
		drv_remove_signal_call(this);
	}
public slots:
	void call(Qt::ToolButtonStyle v1) {
		drv_signal_call(this,&face,65,&v1,0,0,0);
	}
public:
	Iface face;	
};
//go func(bool)
class UISignal66 : public QObject
{
	Q_OBJECT
public:
    explicit UISignal66(QObject *sender, Iface *_face) : face(*_face)
	{
		QObject::connect(sender,SIGNAL(destroyed()),this,SLOT(deleteLater()));
	}
	~UISignal66()
	{
		drv_remove_signal_call(this);
	}
public slots:
	void call(bool v1) {
		drv_signal_call(this,&face,66,&v1,0,0,0);
	}
public:
	Iface face;	
};
//go func(float64)
class UISignal67 : public QObject
{
	Q_OBJECT
public:
    explicit UISignal67(QObject *sender, Iface *_face) : face(*_face)
	{
		QObject::connect(sender,SIGNAL(destroyed()),this,SLOT(deleteLater()));
	}
	~UISignal67()
	{
		drv_remove_signal_call(this);
	}
public slots:
	void call(double v1) {
		drv_signal_call(this,&face,67,&v1,0,0,0);
	}
public:
	Iface face;	
};
//go func(int)
class UISignal68 : public QObject
{
	Q_OBJECT
public:
    explicit UISignal68(QObject *sender, Iface *_face) : face(*_face)
	{
		QObject::connect(sender,SIGNAL(destroyed()),this,SLOT(deleteLater()));
	}
	~UISignal68()
	{
		drv_remove_signal_call(this);
	}
public slots:
	void call(int v1) {
		drv_signal_call(this,&face,68,&v1,0,0,0);
	}
public:
	Iface face;	
};
//go func(int,QHeaderView_ResizeMode)
class UISignal69 : public QObject
{
	Q_OBJECT
public:
    explicit UISignal69(QObject *sender, Iface *_face) : face(*_face)
	{
		QObject::connect(sender,SIGNAL(destroyed()),this,SLOT(deleteLater()));
	}
	~UISignal69()
	{
		drv_remove_signal_call(this);
	}
public slots:
	void call(int v1,QHeaderView::ResizeMode v2) {
		drv_signal_call(this,&face,69,&v1,&v2,0,0);
	}
public:
	Iface face;	
};
//go func(int,QProcess_ExitStatus)
class UISignal70 : public QObject
{
	Q_OBJECT
public:
    explicit UISignal70(QObject *sender, Iface *_face) : face(*_face)
	{
		QObject::connect(sender,SIGNAL(destroyed()),this,SLOT(deleteLater()));
	}
	~UISignal70()
	{
		drv_remove_signal_call(this);
	}
public slots:
	void call(int v1,QProcess::ExitStatus v2) {
		drv_signal_call(this,&face,70,&v1,&v2,0,0);
	}
public:
	Iface face;	
};
//go func(int,Qt_SortOrder)
class UISignal71 : public QObject
{
	Q_OBJECT
public:
    explicit UISignal71(QObject *sender, Iface *_face) : face(*_face)
	{
		QObject::connect(sender,SIGNAL(destroyed()),this,SLOT(deleteLater()));
	}
	~UISignal71()
	{
		drv_remove_signal_call(this);
	}
public slots:
	void call(int v1,Qt::SortOrder v2) {
		drv_signal_call(this,&face,71,&v1,&v2,0,0);
	}
public:
	Iface face;	
};
//go func(int,int)
class UISignal72 : public QObject
{
	Q_OBJECT
public:
    explicit UISignal72(QObject *sender, Iface *_face) : face(*_face)
	{
		QObject::connect(sender,SIGNAL(destroyed()),this,SLOT(deleteLater()));
	}
	~UISignal72()
	{
		drv_remove_signal_call(this);
	}
public slots:
	void call(int v1,int v2) {
		drv_signal_call(this,&face,72,&v1,&v2,0,0);
	}
public:
	Iface face;	
};
//go func(int,int,int)
class UISignal73 : public QObject
{
	Q_OBJECT
public:
    explicit UISignal73(QObject *sender, Iface *_face) : face(*_face)
	{
		QObject::connect(sender,SIGNAL(destroyed()),this,SLOT(deleteLater()));
	}
	~UISignal73()
	{
		drv_remove_signal_call(this);
	}
public slots:
	void call(int v1,int v2,int v3) {
		drv_signal_call(this,&face,73,&v1,&v2,&v3,0);
	}
public:
	Iface face;	
};
//go func(int,int,int,int)
class UISignal74 : public QObject
{
	Q_OBJECT
public:
    explicit UISignal74(QObject *sender, Iface *_face) : face(*_face)
	{
		QObject::connect(sender,SIGNAL(destroyed()),this,SLOT(deleteLater()));
	}
	~UISignal74()
	{
		drv_remove_signal_call(this);
	}
public slots:
	void call(int v1,int v2,int v3,int v4) {
		drv_signal_call(this,&face,74,&v1,&v2,&v3,&v4);
	}
public:
	Iface face;	
};
//go func(int64)
class UISignal75 : public QObject
{
	Q_OBJECT
public:
    explicit UISignal75(QObject *sender, Iface *_face) : face(*_face)
	{
		QObject::connect(sender,SIGNAL(destroyed()),this,SLOT(deleteLater()));
	}
	~UISignal75()
	{
		drv_remove_signal_call(this);
	}
public slots:
	void call(qint64 v1) {
		drv_signal_call(this,&face,75,&v1,0,0,0);
	}
public:
	Iface face;	
};
#endif //CDRV_SIGNAL_H	
