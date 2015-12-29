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

#ifndef UIAPPLICATION_H
#define UIAPPLICATION_H

#include <QApplication>
#include <QEvent>

class UIApplication : public QApplication
{
    Q_OBJECT
public:
    explicit UIApplication(int argc = 0, char **argv = 0);
    virtual bool event(QEvent *);
public:
    void sendAsyncTask();
};

#endif // UIAPPLICATION_H
