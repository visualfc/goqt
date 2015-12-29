#!/bin/sh
../../bin/goqt_rcc -go main -o application_qrc.go application.qrc
go build -ldflags "-r ." -o ../../bin/application