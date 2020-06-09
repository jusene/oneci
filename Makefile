# Go parameters
GOCMD=go


default:
	@echo "[make all] for linux_amd64 linux_arm64 windows_amd64"
	@echo "[make linux_amd64] for linux_amd64"
	@echo "[make linux_arm64] for linux_arm64"
	@echo "[make windows_amd64] for windows_amd64"
all: linux_amd64 linux_arm64 windows_amd64
clean:
	${GOCMD} clean

linux_amd64:
	@${GOCMD} env -w GOARCH=amd64
	@${GOCMD} env -w GOOS=linux
	${GOCMD} install zjhw.com/oneci/

linux_arm64:
	@${GOCMD} env -w GOARCH=arm64
	@${GOCMD} env -w GOOS=linux
	${GOCMD} install zjhw.com/oneci/

windows_amd64:
	@${GOCMD} env -w GOARCH=amd64
	@${GOCMD} env -w GOOS=windows
	${GOCMD} install zjhw.com/oneci/