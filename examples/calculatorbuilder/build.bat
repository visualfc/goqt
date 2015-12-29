..\..\bin\goqt_rcc -go main -o calculatorbuilder_qrc.go calculatorbuilder.qrc
go build -ldflags "-H windowsgui" -o ..\..\bin\calculatorbuilder.exe