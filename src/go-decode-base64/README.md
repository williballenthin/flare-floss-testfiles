# test-go-decode-base64

Strings encoded with base64 using a custom alphabet (encoding/base64), decoded at runtime.

Built with Go 1.27.4 (`CGO_ENABLED=0`, `-trimpath -ldflags="-s -w -buildid="`), for linux/amd64, windows/386 and
windows/amd64: `make all`.
