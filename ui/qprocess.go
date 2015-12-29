// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QProcess_ProcessError - QProcess::ProcessError
type QProcess_ProcessError uint32
const (
	QProcess_FailedToStart QProcess_ProcessError = 0
	QProcess_Crashed QProcess_ProcessError = 1
	QProcess_Timedout QProcess_ProcessError = 2
	QProcess_ReadError QProcess_ProcessError = 3
	QProcess_WriteError QProcess_ProcessError = 4
	QProcess_UnknownError QProcess_ProcessError = 5
)
//enum QProcess_ProcessChannelMode - QProcess::ProcessChannelMode
type QProcess_ProcessChannelMode uint32
const (
	QProcess_SeparateChannels QProcess_ProcessChannelMode = 0
	QProcess_MergedChannels QProcess_ProcessChannelMode = 1
	QProcess_ForwardedChannels QProcess_ProcessChannelMode = 2
)
//enum QProcess_ProcessChannel - QProcess::ProcessChannel
type QProcess_ProcessChannel uint32
const (
	QProcess_StandardOutput QProcess_ProcessChannel = 0
	QProcess_StandardError QProcess_ProcessChannel = 1
)
//enum QProcess_ProcessState - QProcess::ProcessState
type QProcess_ProcessState uint32
const (
	QProcess_NotRunning QProcess_ProcessState = 0
	QProcess_Starting QProcess_ProcessState = 1
	QProcess_Running QProcess_ProcessState = 2
)
//enum QProcess_ExitStatus - QProcess::ExitStatus
type QProcess_ExitStatus uint32
const (
	QProcess_NormalExit QProcess_ExitStatus = 0
	QProcess_CrashExit QProcess_ExitStatus = 1
)
//struct QProcess : QProcess
type QProcess struct {
	QIODevice
}
func (q *QProcess) OnError(fn func(QProcess_ProcessError)) uintptr {
	var __rv uintptr
	q.Drv(326000,326102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QProcess) OnReadyReadStandardError(fn func()) uintptr {
	var __rv uintptr
	q.Drv(326000,326103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QProcess) OnStateChanged(fn func(QProcess_ProcessState)) uintptr {
	var __rv uintptr
	q.Drv(326000,326104,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QProcess) OnStarted(fn func()) uintptr {
	var __rv uintptr
	q.Drv(326000,326105,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QProcess) OnReadyReadStandardOutput(fn func()) uintptr {
	var __rv uintptr
	q.Drv(326000,326106,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QProcess) OnFinished(fn func(int,QProcess_ExitStatus)) uintptr {
	var __rv uintptr
	q.Drv(326000,326107,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QProcess) OnFinishedWithExitcode(fn func(int)) uintptr {
	var __rv uintptr
	q.Drv(326000,326108,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QProcess::QProcess()	
func NewProcess() *QProcess {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),326000,326109,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QProcess{}
	_p.SetDriver(__rv,326000,false)
	return _p
} 
//QProcess::QProcess(QObject*)	
func NewProcessWithParent(parent QObjectInterface) *QProcess {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),326000,326110,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QProcess{}
	_p.SetDriver(__rv,326000,false)
	return _p
} 
//QProcess::atEnd()
func (q *QProcess) AtEnd() bool {
	var __rv bool
	q.Drv(326000,326111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QProcess::bytesAvailable()
func (q *QProcess) BytesAvailable() int64 {
	var __rv int64
	q.Drv(326000,326112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QProcess::bytesToWrite()
func (q *QProcess) BytesToWrite() int64 {
	var __rv int64
	q.Drv(326000,326113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QProcess::canReadLine()
func (q *QProcess) CanReadLine() bool {
	var __rv bool
	q.Drv(326000,326114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QProcess::close()
func (q *QProcess) Close()  {
	q.Drv(326000,326115,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QProcess::closeReadChannel(QProcess::ProcessChannel)
func (q *QProcess) CloseReadChannel(channel QProcess_ProcessChannel)  {
	q.Drv(326000,326116,unsafe.Pointer(&channel),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QProcess::closeWriteChannel()
func (q *QProcess) CloseWriteChannel()  {
	q.Drv(326000,326117,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QProcess::environment()
func (q *QProcess) Environment() []string {
	var __rv []string
	q.Drv(326000,326118,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QProcess::error()
func (q *QProcess) Error() QProcess_ProcessError {
	var __rv QProcess_ProcessError
	q.Drv(326000,326119,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QProcess::execute(QString const&)	
func QProcessExecute(program string) int {
	var __rv int
	DirectQtDrv(nil,326000,326120,unsafe.Pointer(&program),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QProcess::execute(QString const&)
func (q *QProcess) Execute(program string) int {
	var __rv int
	q.Drv(326000,326120,unsafe.Pointer(&program),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QProcess::execute(QString const&,QStringList const&)	
func QProcessExecuteWithProgramArguments(program string,arguments []string) int {
	var __rv int
	DirectQtDrv(nil,326000,326121,unsafe.Pointer(&program),unsafe.Pointer(&arguments),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QProcess::execute(QString const&,QStringList const&)
func (q *QProcess) ExecuteWithProgramArguments(program string,arguments []string) int {
	var __rv int
	q.Drv(326000,326121,unsafe.Pointer(&program),unsafe.Pointer(&arguments),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QProcess::exitCode()
func (q *QProcess) ExitCode() int {
	var __rv int
	q.Drv(326000,326122,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QProcess::exitStatus()
func (q *QProcess) ExitStatus() QProcess_ExitStatus {
	var __rv QProcess_ExitStatus
	q.Drv(326000,326123,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QProcess::isSequential()
func (q *QProcess) IsSequential() bool {
	var __rv bool
	q.Drv(326000,326124,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QProcess::kill()
func (q *QProcess) Kill()  {
	q.Drv(326000,326125,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QProcess::processChannelMode()
func (q *QProcess) ProcessChannelMode() QProcess_ProcessChannelMode {
	var __rv QProcess_ProcessChannelMode
	q.Drv(326000,326126,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QProcess::processEnvironment()
func (q *QProcess) ProcessEnvironment() *QProcessEnvironment {
	var __rv uintptr
	q.Drv(326000,326127,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QProcessEnvironment{}
	_rp.SetDriver(__rv,106000,true)
	return _rp
}	
//QProcess::readAllStandardError()
func (q *QProcess) ReadAllStandardError() []byte {
	var __rv []byte
	q.Drv(326000,326128,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QProcess::readAllStandardOutput()
func (q *QProcess) ReadAllStandardOutput() []byte {
	var __rv []byte
	q.Drv(326000,326129,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QProcess::readChannel()
func (q *QProcess) ReadChannel() QProcess_ProcessChannel {
	var __rv QProcess_ProcessChannel
	q.Drv(326000,326130,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QProcess::readChannelMode()
func (q *QProcess) ReadChannelMode() QProcess_ProcessChannelMode {
	var __rv QProcess_ProcessChannelMode
	q.Drv(326000,326131,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QProcess::readData(char*,qint64)
func (q *QProcess) ReadData(data *byte,maxlen int64) int64 {
	var __rv int64
	q.Drv(326000,326132,unsafe.Pointer(&data),unsafe.Pointer(&maxlen),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QProcess::setEnvironment(QStringList const&)
func (q *QProcess) SetEnvironment(environment []string)  {
	q.Drv(326000,326133,unsafe.Pointer(&environment),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QProcess::setProcessChannelMode(QProcess::ProcessChannelMode)
func (q *QProcess) SetProcessChannelMode(mode QProcess_ProcessChannelMode)  {
	q.Drv(326000,326134,unsafe.Pointer(&mode),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QProcess::setProcessEnvironment(QProcessEnvironment const&)
func (q *QProcess) SetProcessEnvironment(environment *QProcessEnvironment)  {
	q.Drv(326000,326135,Native(environment),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QProcess::setProcessState(QProcess::ProcessState)
func (q *QProcess) SetProcessState(state QProcess_ProcessState)  {
	q.Drv(326000,326136,unsafe.Pointer(&state),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QProcess::setReadChannel(QProcess::ProcessChannel)
func (q *QProcess) SetReadChannel(channel QProcess_ProcessChannel)  {
	q.Drv(326000,326137,unsafe.Pointer(&channel),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QProcess::setReadChannelMode(QProcess::ProcessChannelMode)
func (q *QProcess) SetReadChannelMode(mode QProcess_ProcessChannelMode)  {
	q.Drv(326000,326138,unsafe.Pointer(&mode),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QProcess::setStandardErrorFile(QString const&)
func (q *QProcess) SetStandardErrorFile(fileName string)  {
	q.Drv(326000,326139,unsafe.Pointer(&fileName),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QProcess::setStandardErrorFile(QString const&,QFlags<QIODevice::OpenModeFlag>)
func (q *QProcess) SetStandardErrorFileWithFilenameMode(fileName string,mode QIODevice_OpenModeFlag)  {
	q.Drv(326000,326140,unsafe.Pointer(&fileName),unsafe.Pointer(&mode),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QProcess::setStandardInputFile(QString const&)
func (q *QProcess) SetStandardInputFile(fileName string)  {
	q.Drv(326000,326141,unsafe.Pointer(&fileName),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QProcess::setStandardOutputFile(QString const&)
func (q *QProcess) SetStandardOutputFile(fileName string)  {
	q.Drv(326000,326142,unsafe.Pointer(&fileName),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QProcess::setStandardOutputFile(QString const&,QFlags<QIODevice::OpenModeFlag>)
func (q *QProcess) SetStandardOutputFileWithFilenameMode(fileName string,mode QIODevice_OpenModeFlag)  {
	q.Drv(326000,326143,unsafe.Pointer(&fileName),unsafe.Pointer(&mode),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QProcess::setStandardOutputProcess(QProcess*)
func (q *QProcess) SetStandardOutputProcess(destination *QProcess)  {
	q.Drv(326000,326144,Native(destination),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QProcess::setWorkingDirectory(QString const&)
func (q *QProcess) SetWorkingDirectory(dir string)  {
	q.Drv(326000,326145,unsafe.Pointer(&dir),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QProcess::setupChildProcess()
func (q *QProcess) SetupChildProcess()  {
	q.Drv(326000,326146,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QProcess::start(QString const&)
func (q *QProcess) Start(program string)  {
	q.Drv(326000,326147,unsafe.Pointer(&program),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QProcess::start(QString const&,QFlags<QIODevice::OpenModeFlag>)
func (q *QProcess) StartWithProgramMode(program string,mode QIODevice_OpenModeFlag)  {
	q.Drv(326000,326148,unsafe.Pointer(&program),unsafe.Pointer(&mode),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QProcess::start(QString const&,QStringList const&,QFlags<QIODevice::OpenModeFlag>)
func (q *QProcess) StartWithProgramArgumentsMode(program string,arguments []string,mode QIODevice_OpenModeFlag)  {
	q.Drv(326000,326149,unsafe.Pointer(&program),unsafe.Pointer(&arguments),unsafe.Pointer(&mode),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QProcess::startDetached(QString const&)	
func QProcessStartDetached(program string) bool {
	var __rv bool
	DirectQtDrv(nil,326000,326150,unsafe.Pointer(&program),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QProcess::startDetached(QString const&)
func (q *QProcess) StartDetached(program string) bool {
	var __rv bool
	q.Drv(326000,326150,unsafe.Pointer(&program),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QProcess::startDetached(QString const&,QStringList const&)	
func QProcessStartDetachedWithProgramArguments(program string,arguments []string) bool {
	var __rv bool
	DirectQtDrv(nil,326000,326151,unsafe.Pointer(&program),unsafe.Pointer(&arguments),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QProcess::startDetached(QString const&,QStringList const&)
func (q *QProcess) StartDetachedWithProgramArguments(program string,arguments []string) bool {
	var __rv bool
	q.Drv(326000,326151,unsafe.Pointer(&program),unsafe.Pointer(&arguments),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QProcess::startDetached(QString const&,QStringList const&,QString const&,qint64*)	
func QProcessStartDetachedWithProgramArgumentsWorkingdirectoryPid(program string,arguments []string,workingDirectory string,pid *int64) bool {
	var __rv bool
	DirectQtDrv(nil,326000,326152,unsafe.Pointer(&program),unsafe.Pointer(&arguments),unsafe.Pointer(&workingDirectory),unsafe.Pointer(&pid),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QProcess::startDetached(QString const&,QStringList const&,QString const&,qint64*)
func (q *QProcess) StartDetachedWithProgramArgumentsWorkingdirectoryPid(program string,arguments []string,workingDirectory string,pid *int64) bool {
	var __rv bool
	q.Drv(326000,326152,unsafe.Pointer(&program),unsafe.Pointer(&arguments),unsafe.Pointer(&workingDirectory),unsafe.Pointer(&pid),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QProcess::state()
func (q *QProcess) State() QProcess_ProcessState {
	var __rv QProcess_ProcessState
	q.Drv(326000,326153,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QProcess::systemEnvironment()	
func QProcessSystemEnvironment() []string {
	var __rv []string
	DirectQtDrv(nil,326000,326154,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QProcess::systemEnvironment()
func (q *QProcess) SystemEnvironment() []string {
	var __rv []string
	q.Drv(326000,326154,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QProcess::terminate()
func (q *QProcess) Terminate()  {
	q.Drv(326000,326155,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QProcess::waitForBytesWritten()
func (q *QProcess) WaitForBytesWritten() bool {
	var __rv bool
	q.Drv(326000,326156,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QProcess::waitForBytesWritten(int)
func (q *QProcess) WaitForBytesWrittenWithMsecs(msecs int) bool {
	var __rv bool
	q.Drv(326000,326157,unsafe.Pointer(&msecs),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QProcess::waitForFinished()
func (q *QProcess) WaitForFinished() bool {
	var __rv bool
	q.Drv(326000,326158,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QProcess::waitForFinished(int)
func (q *QProcess) WaitForFinishedWithMsecs(msecs int) bool {
	var __rv bool
	q.Drv(326000,326159,unsafe.Pointer(&msecs),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QProcess::waitForReadyRead()
func (q *QProcess) WaitForReadyRead() bool {
	var __rv bool
	q.Drv(326000,326160,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QProcess::waitForReadyRead(int)
func (q *QProcess) WaitForReadyReadWithMsecs(msecs int) bool {
	var __rv bool
	q.Drv(326000,326161,unsafe.Pointer(&msecs),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QProcess::waitForStarted()
func (q *QProcess) WaitForStarted() bool {
	var __rv bool
	q.Drv(326000,326162,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QProcess::waitForStarted(int)
func (q *QProcess) WaitForStartedWithMsecs(msecs int) bool {
	var __rv bool
	q.Drv(326000,326163,unsafe.Pointer(&msecs),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QProcess::workingDirectory()
func (q *QProcess) WorkingDirectory() string {
	var __rv string
	q.Drv(326000,326164,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QProcess::writeData(char const*,qint64)
func (q *QProcess) WriteData(data string,len int64) int64 {
	var __rv int64
	q.Drv(326000,326165,unsafe.Pointer(&data),unsafe.Pointer(&len),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
