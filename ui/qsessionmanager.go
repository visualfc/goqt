// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QSessionManager_RestartHint - QSessionManager::RestartHint
type QSessionManager_RestartHint uint32
const (
	QSessionManager_RestartIfRunning QSessionManager_RestartHint = 0
	QSessionManager_RestartAnyway QSessionManager_RestartHint = 1
	QSessionManager_RestartImmediately QSessionManager_RestartHint = 2
	QSessionManager_RestartNever QSessionManager_RestartHint = 3
)
//struct QSessionManager : QSessionManager
type QSessionManager struct {
	QObject
}
//QSessionManager::allowsErrorInteraction()
func (q *QSessionManager) AllowsErrorInteraction() bool {
	var __rv bool
	q.Drv(338000,338102,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSessionManager::allowsInteraction()
func (q *QSessionManager) AllowsInteraction() bool {
	var __rv bool
	q.Drv(338000,338103,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSessionManager::cancel()
func (q *QSessionManager) Cancel()  {
	q.Drv(338000,338104,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSessionManager::discardCommand()
func (q *QSessionManager) DiscardCommand() []string {
	var __rv []string
	q.Drv(338000,338105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSessionManager::isPhase2()
func (q *QSessionManager) IsPhase2() bool {
	var __rv bool
	q.Drv(338000,338106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSessionManager::release()
func (q *QSessionManager) Release()  {
	q.Drv(338000,338107,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSessionManager::requestPhase2()
func (q *QSessionManager) RequestPhase2()  {
	q.Drv(338000,338108,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSessionManager::restartCommand()
func (q *QSessionManager) RestartCommand() []string {
	var __rv []string
	q.Drv(338000,338109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSessionManager::restartHint()
func (q *QSessionManager) RestartHint() QSessionManager_RestartHint {
	var __rv QSessionManager_RestartHint
	q.Drv(338000,338110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSessionManager::sessionId()
func (q *QSessionManager) SessionId() string {
	var __rv string
	q.Drv(338000,338111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSessionManager::sessionKey()
func (q *QSessionManager) SessionKey() string {
	var __rv string
	q.Drv(338000,338112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSessionManager::setDiscardCommand(QStringList const&)
func (q *QSessionManager) SetDiscardCommand(value []string)  {
	q.Drv(338000,338113,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSessionManager::setManagerProperty(QString const&,QString const&)
func (q *QSessionManager) SetManagerPropertyWithNameValue(name string,value string)  {
	q.Drv(338000,338114,unsafe.Pointer(&name),unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSessionManager::setManagerProperty(QString const&,QStringList const&)
func (q *QSessionManager) SetManagerPropertyWithNameValues(name string,value []string)  {
	q.Drv(338000,338115,unsafe.Pointer(&name),unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSessionManager::setRestartCommand(QStringList const&)
func (q *QSessionManager) SetRestartCommand(value []string)  {
	q.Drv(338000,338116,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSessionManager::setRestartHint(QSessionManager::RestartHint)
func (q *QSessionManager) SetRestartHint(value QSessionManager_RestartHint)  {
	q.Drv(338000,338117,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
