package main

import (
	"encoding/base64"
	"os"
)

func emit(s string) {
	os.Stdout.WriteString(s)
	os.Stdout.WriteString("\n")
}

// a custom (shuffled) alphabet, so the encoded strings don't decode as standard base64.
var enc = base64.NewEncoding("tuZ7GscqBiXzKL6VfIhYynePANJoS8E4COUd5F9w+3Wg/DmabMr0Qpk1Tx2jHvlR")

var enc0 = "NkHDAksmAPix6d3Ooqu+AY+2A0BmNPOOoPu/NhxmNPf267fQKrvOSc5a8dGaA9nOAkvm"
var enc1 = "NkHDAksmAPix6d3US9sko0+2f03SyqiaN1iOoyIO8csSnPu5APIFSFMdokx9JeSmJwLaoC=="
var enc2 = "NkHDAksmAPix6d3dJcsrocFF6d3Lo133ocMOz0ymKZt+nkFmNcv1Sru6nZtMKZTb6ruPJeTkL7/CE7AQXf=="
var enc3 = "NkHDAksmAPix6d35NeMQAY+2yQvcns8uy5nSYeFdS9v0okNQPs83o9Ia81LSf1nrS9nm8sNFSwL3okxSywnm"
var enc4 = "NkHDAksmAPix6d3FAkOa6d30NPL0JevmzeDFEhQQNdFO"

//go:noinline
func decode(s string) string {
	b, err := enc.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func main() {
	emit(decode(enc0))
	emit(decode(enc1))
	emit(decode(enc2))
	emit(decode(enc3))
	emit(decode(enc4))
}
