# test-go-decode-rc4

Strings decrypted with RC4 from the standard library (crypto/rc4) into heap slices.

Built with Go 1.27.4 (`CGO_ENABLED=0`, `-trimpath -ldflags="-s -w -buildid="`), for linux/amd64, windows/386 and
windows/amd64: `make all`.
