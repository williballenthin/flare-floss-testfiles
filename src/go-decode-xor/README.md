# test-go-decode-xor

A hand-written decoder function: global ciphertext xor a repeating 16-byte key, into a new heap slice returned as a string.

Built with Go 1.27.4 (`CGO_ENABLED=0`, `-trimpath -ldflags="-s -w -buildid="`), for linux/amd64, windows/386 and
windows/amd64: `make all`.
