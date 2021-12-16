all: build-linux build-mac

build-linux:
	wails build -p -x linux/amd64
	mv link.app link-linux-amd64.app

build-windows:
	wails build -p -x windows/amd64
	mv link.app link-windows-amd64.app

build-mac:
	wails build -p -x darwin/amd64
	mv link.app link-darwin-amd64.app

build-mac-m1:
	wails build -p
	mv link.app link-darwin-arm64

clean:
	rm -rf link-server