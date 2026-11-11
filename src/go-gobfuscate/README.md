# test-go-gobfuscate

String literals encrypted by gobfuscate: each literal becomes a closure, called in place, that xors a ciphertext with a random mask of the same length into a new heap slice.

[gobfuscate](https://github.com/unixpickle/gobfuscate) (at 13796335, December 2021, its last commit) rewrites a copy
of the program's GOPATH: it turns string constants into variables, replaces each string literal with a decoding
closure, and renames packages and symbols. `build.sh` runs it in GOPATH mode (`GO111MODULE=off`) with a fixed
`-padding`, then builds the result with Go 1.27.4 (`-trimpath -ldflags="-s -w -buildid="`), for linux/amd64,
windows/386 and windows/amd64: `make all`.

gobfuscate draws its masks from `math/rand` without seeding it, so they're the same on every run only when gobfuscate
itself is built with Go 1.19 or older (here, 1.17.13).
