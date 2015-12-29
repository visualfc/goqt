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

#ifndef QSYNTAXHIGHLIGHTERHOOK_H
#define QSYNTAXHIGHLIGHTERHOOK_H

#include <QSyntaxHighlighter>

class QTextEdit;
class QSyntaxHighlighterHook : public QSyntaxHighlighter
{
    Q_OBJECT
public:
    QSyntaxHighlighterHook(QObject *parent);
    QSyntaxHighlighterHook(QTextDocument *parent);
    QSyntaxHighlighterHook(QTextEdit *parent);

    virtual void highlightBlock(const QString &text);
signals:
    void hook_highlightBlock(const QString &text);
};

#endif // QSYNTAXHIGHLIGHTERHOOK_H
