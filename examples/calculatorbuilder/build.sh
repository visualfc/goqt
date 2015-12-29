#!/bin/sh
../../bin/goqt_rcc -go main -o calculatorbuilder_qrc.go calculatorbuilder.qrc
go build -ldflags "-r ." -o ../../bin/calculatorbuilder