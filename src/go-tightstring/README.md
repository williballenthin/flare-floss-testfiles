# test-go-tightstring

A stack array xor-decoded in place by a tight loop, then copied into a heap string.

Built with Go 1.27.4 (`CGO_ENABLED=0`, `-trimpath -ldflags="-s -w -buildid="`), for linux/amd64, windows/386 and
windows/amd64: `make all`.

It's also built with Go 1.17.13 (2021), the `-go1.17` binaries, because older compilers emit this loop differently:
Go 1.17 keeps the bounds check of `key[j%4]`, so the loop spans several basic blocks, and on amd64 it writes the
ciphertext to the stack with immediate `MOVQ`s, where Go 1.27 copies it from read-only data.
