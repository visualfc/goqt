// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"log"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

func mainLoop() int32 {
	for f := range mainfunc {
		return f()
	}
	return 0
}

var mainfunc = make(chan func() int32)

func runInOsMainThread(f func() int32) {
	done := make(chan bool, 1)
	mainfunc <- func() int32 {
		res := f()
		done <- true
		return res
	}
	<-done
}

func Version() string {
	return "0.1.1"
}

func Async(fn func()) {
	tasks.Append(fn)
	_DirectQtDrv(nil, 1, 200, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
}

func Run(fn func()) int32 {
	go runInOsMainThread(func() int32 {
		if qtdrv_init_error != nil {
			log.Println(qtdrv_init_error)
			return -2
		}
		app := NewApplication(nil)
		Async(func() {
			fn()
		})
		return app.Exec()
	})
	return mainLoop()
}

func RunEx(args []string, fn func()) int32 {
	go runInOsMainThread(func() int32 {
		if qtdrv_init_error != nil {
			log.Println(qtdrv_init_error)
			return -2
		}
		app := NewApplication(args)
		Async(func() {
			fn()
		})
		return app.Exec()
	})
	return mainLoop()
}

func Application() *QApplication {
	return QApplicationInstance()
}
