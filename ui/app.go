// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

func Version() string {
	return "0.1.1"
}

func Async(fn func()) {
	tasks.Append(fn)
	_DirectQtDrv(nil, 1, 200, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
}

func Run(fn func()) int {
	app := NewQApplication(nil)
	Async(func() {
		fn()
	})
	return app.Exec()
}

func RunEx(args []string, fn func()) int {
	app := NewQApplication(args)
	Async(func() {
		fn()
	})
	return app.Exec()
}

func Application() *QApplication {
	return QApplicationInstance()
}
