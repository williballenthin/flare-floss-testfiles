# test-go-garble

String literals obfuscated by garble v0.18's -literals: each literal is decoded by its own top-level function, reached through a tree of global structs of function values (a proxy), with part of the key passed by the caller.

Built with [garble](https://github.com/burrowers/garble) v0.18.0 and Go 1.27.4:
`garble -literals -seed=HAAAAAAAAAA build -trimpath -ldflags="-s -w -buildid="`, for linux/amd64, windows/386 and
windows/amd64: `make all`.

garble picks an obfuscator at random for each string literal; with this seed, the ten strings use all five of its
obfuscators (simple, swap, split, shuffle and seed). garble also obfuscates the program's identifiers.

garble has hidden its literal decoders behind such a proxy, and passed part of the key from the caller, since v0.15.0
(August 2025). See `go-garble-v0.14` for the earlier design.
