# Install GoQt

## Platforms Support

### System

* Windows x86 (32-bit or 64-bit) 
* Linux x86 (32-bit or 64-bit)
* MacOS X10.6 

### Golang

* Go1.4.2
* Go1.5.2

### Qt Version

* Qt4.8.5

* Qt5.5.1

		My test Qt4.8.5 on Windows Linux and MacOS X
		My test Qt5.5.1 only on Windows.


## Install GoQt

1. Download github.com/visualfc/goqt
2. Use QtCreator or make build goqt/qtdrv and goqt/tools/rcc, need install QtSDK
3. Install goqt/ui

### Windows

	QtSDK x86_32 build for GOARCH=386 and CGO_ENABLED=1
	QtSDK x86_64 build for GOARCH=amd64 and CGO_ENABLED=1

#### 1.get goqt
    > go get github.com/visualfc/goqt
#### 2.build qtdrv, need install QtSDK
    > cd goqt\qtdrv
    > qmake "CONFIG+=release"
    > mingw32-make
#### 3.build rcc	
	> cd goqt\tools\rcc
	> qmake "CONFIG+=release"	
	> mingw32-make
#### 4.build ui, need CGO_ENABLED=1 and install gcc
	> set GOARCH=386
	> set CGO_ENABLED=1
    > cd goqt\ui
    > go install -v
#### 5.build examples
	> set GOARCH=386
	> set CGO_ENABLED=1
    > cd goqt\examples\minimal
    > build.bat
    > ..\..\bin\minimal.exe
	
### Linux

#### 1.get goqt
    > go get github.com/visualfc/goqt
#### 2.build qtdrv, need install QtSDK
    > cd goqt/qtdrv
    > qmake "CONFIG+=release"
    > make
#### 3.build rcc	
	> cd goqt/tools/rcc
	> qmake "CONFIG+=release"	
	> make
#### 4.build ui, need CGO_ENABLED=1 and install gcc
    > cd goqt/ui
    > go install -v
#### 5.build examples
    > cd goqt/examples/minimal
    > ./build.sh
    > ../../bin/minimal
	
### MacOS X

#### 1.get goqt
    > go get github.com/visualfc/goqt
#### 2.build qtdrv, need install QtSDK
    > cd goqt/qtdrv
    > qmake -spec macx-g++ "CONFIG+=release"
    > make
#### 3.build rcc	
	> cd goqt/tools/rcc
	> qmake -spec macx-g++ "CONFIG+=release"	
	> make
#### 4.build ui, need CGO_ENABLED=1 and install gcc
    > cd goqt/ui
    > go install -v
#### 5.build examples
    > cd goqt/examples/minimal
    > ./build.sh
    > ../../bin/minimal	