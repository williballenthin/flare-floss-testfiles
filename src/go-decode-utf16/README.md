# test-go-decode-utf16

A xor decoder that produces NUL-terminated UTF-16 buffers, as passed to wide Windows APIs.

Built with Go 1.27.4 (`CGO_ENABLED=0`, `-trimpath -ldflags="-s -w -buildid="`), for linux/amd64, windows/386 and
windows/amd64: `make all`.
