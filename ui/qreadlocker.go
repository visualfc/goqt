// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui


//struct QReadLocker : QReadLocker
type QReadLocker struct {
	BaseDrv
}
//QReadLocker::relock()
func (q *QReadLocker) Relock()  {
	q.Drv(109000,109102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QReadLocker::unlock()
func (q *QReadLocker) Unlock()  {
	q.Drv(109000,109103,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
