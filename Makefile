build_win:
	set GOOS=windows
	go build -o convertor_win.exe

build_linux:
	set GOOS=linux
	go build -o convertor_linux

build_mac:
	set GOOS=darwin
	go build -o convertor_mac

clean-build:
	del convertor_win.exe convertor_linux convertor_mac

go-build: clean-build build_win build_mac build_linux