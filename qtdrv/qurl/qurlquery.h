/****************************************************************************
**
** Copyright (C) 2012 Intel Corporation.
** Contact: http://www.qt-project.org/legal
**
** This file is part of the QtCore module of the Qt Toolkit.
**
** $QT_BEGIN_LICENSE:LGPL$
** Commercial License Usage
** Licensees holding valid commercial Qt licenses may use this file in
** accordance with the commercial license agreement provided with the
** Software or, alternatively, in accordance with the terms contained in
** a written agreement between you and Digia.  For licensing terms and
** conditions see http://qt.digia.com/licensing.  For further information
** use the contact form at http://qt.digia.com/contact-us.
**
** GNU Lesser General Public License Usage
** Alternatively, this file may be used under the terms of the GNU Lesser
** General Public License version 2.1 as published by the Free Software
** Foundation and appearing in the file LICENSE.LGPL included in the
** packaging of this file.  Please review the following information to
** ensure the GNU Lesser General Public License version 2.1 requirements
** will be met: http://www.gnu.org/licenses/old-licenses/lgpl-2.1.html.
**
** In addition, as a special exception, Digia gives you certain additional
** rights.  These rights are described in the Digia Qt LGPL Exception
** version 1.1, included in the file LGPL_EXCEPTION.txt in this package.
**
** GNU General Public License Usage
** Alternatively, this file may be used under the terms of the GNU
** General Public License version 3.0 as published by the Free Software
** Foundation and appearing in the file LICENSE.GPL included in the
** packaging of this file.  Please review the following information to
** ensure the GNU General Public License version 3.0 requirements will be
** met: http://www.gnu.org/copyleft/gpl.html.
**
**
** $QT_END_LICENSE$
**
****************************************************************************/

#ifndef QURLQUERY_H
#define QURLQUERY_H

#include <QtCore/qpair.h>
#include <QtCore/qshareddata.h>
#include <QtCore/qurl.h>

#include "url_help.h"

QT_BEGIN_HEADER

QT_BEGIN_NAMESPACE

QString fromEncodedComponent_helper(const QByteArray &ba);
inline QByteArray toLatin1_helper(const QString &string)
{
    if (string.isEmpty())
        return string.isNull() ? QByteArray() : QByteArray("");
    return string.toLatin1();
}

class QUrlQueryPrivate;
class QUrlQuery
{
public:
    QUrlQuery();
    explicit QUrlQuery(const QUrl &url);
    explicit QUrlQuery(const QString &queryString);
    QUrlQuery(const QUrlQuery &other);
    QUrlQuery &operator=(const QUrlQuery &other);
#ifdef Q_COMPILER_RVALUE_REFS
    QUrlQuery &operator=(QUrlQuery &&other)
    { qSwap(d, other.d); return *this; }
#endif
    ~QUrlQuery();

    bool operator==(const QUrlQuery &other) const;
    bool operator!=(const QUrlQuery &other) const
    { return !(*this == other); }

    void swap(QUrlQuery &other) { qSwap(d, other.d); }

    bool isEmpty() const;
    bool isDetached() const;
    void clear();

    QString query(QUrl2::ComponentFormattingOptions encoding = QUrl2::PrettyDecoded) const;
    void setQuery(const QString &queryString);
    QString toString(QUrl2::ComponentFormattingOptions encoding = QUrl2::PrettyDecoded) const
    { return query(encoding); }

    void setQueryDelimiters(QChar valueDelimiter, QChar pairDelimiter);
    QChar queryValueDelimiter() const;
    QChar queryPairDelimiter() const;

    void setQueryItems(const QList<QPair<QString, QString> > &query);
    QList<QPair<QString, QString> > queryItems(QUrl2::ComponentFormattingOptions encoding = QUrl2::PrettyDecoded) const;

    bool hasQueryItem(const QString &key) const;
    void addQueryItem(const QString &key, const QString &value);
    void removeQueryItem(const QString &key);
    QString queryItemValue(const QString &key, QUrl2::ComponentFormattingOptions encoding = QUrl2::PrettyDecoded) const;
    QStringList allQueryItemValues(const QString &key, QUrl2::ComponentFormattingOptions encoding = QUrl2::PrettyDecoded) const;
    void removeAllQueryItems(const QString &key);

    static QChar defaultQueryValueDelimiter()
    { return QChar(ushort('=')); }
    static QChar defaultQueryPairDelimiter()
    { return QChar(ushort('&')); }

private:
    friend class QUrl;
    QSharedDataPointer<QUrlQueryPrivate> d;
public:
    typedef QSharedDataPointer<QUrlQueryPrivate> DataPtr;
    inline DataPtr &data_ptr() { return d; }
};

Q_DECLARE_SHARED(QUrlQuery)

QT_END_NAMESPACE

QT_END_HEADER

#endif // QURLQUERY_H
