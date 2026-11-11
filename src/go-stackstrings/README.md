# test-go-stackstrings

Strings built byte by byte in a stack array, then copied into a heap string.

Built with Go 1.27.4 (`CGO_ENABLED=0`, `-trimpath -ldflags="-s -w -buildid="`), for linux/amd64, windows/386 and
windows/amd64: `make all`.
