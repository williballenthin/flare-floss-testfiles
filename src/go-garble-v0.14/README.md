# test-go-garble-v0.14

String literals obfuscated by garble's -literals, before its v0.15 redesign: each literal becomes a closure, called in place, that decodes it into a heap byte slice.

Built with [garble](https://github.com/burrowers/garble) v0.14.2 and Go 1.23.12:
`garble -literals -seed=HAAAAAAAAAA build -trimpath -ldflags="-s -w -buildid="`, for linux/amd64, windows/386 and
windows/amd64: `make all`.

garble picks an obfuscator at random for each string literal; with this seed, the ten strings use all five of its
obfuscators (simple, swap, split, shuffle and seed). garble also obfuscates the program's identifiers.

This is the last release (April 2025) of the design that garble used since its first releases, and that
Mandiant's [GoStringUngarbler](https://github.com/mandiant/GoStringUngarbler) targets. See `go-garble` for the current design.
