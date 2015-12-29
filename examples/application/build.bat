..\..\bin\goqt_rcc -go main -o application_qrc.go application.qrc
go build -ldflags "-H windowsgui" -o ..\..\bin\application.exe