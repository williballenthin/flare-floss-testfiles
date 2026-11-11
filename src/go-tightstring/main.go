package main

import "os"

func emit(s string) {
	os.Stdout.WriteString(s)
	os.Stdout.WriteString("\n")
}

// the key, in a variable so the compiler can't decode the strings at build time.
var key = [4]byte{105, 165, 81, 23}

//go:noinline
func tight0() string {
	buf := [51]byte{14, 202, 124, 116, 8, 203, 48, 101, 16, 159, 107, 118, 5, 213, 57, 118, 83, 159, 50, 37, 71, 192, 41, 118, 4, 213, 61, 114, 71, 203, 52, 99, 83, 157, 101, 35, 90, 138, 48, 103, 0, 138, 39, 38, 70, 199, 52, 118, 10, 202, 63}
	for j := range buf {
		buf[j] ^= key[j%4]
	}
	return string(buf[:])
}

//go:noinline
func tight1() string {
	buf := [52]byte{14, 202, 124, 116, 8, 203, 48, 101, 16, 159, 107, 117, 27, 196, 39, 120, 83, 159, 18, 45, 53, 245, 35, 120, 14, 215, 48, 122, 45, 196, 37, 118, 53, 240, 33, 115, 8, 209, 52, 101, 53, 198, 62, 121, 15, 204, 54, 57, 3, 214, 62, 121}
	for j := range buf {
		buf[j] ^= key[j%4]
	}
	return string(buf[:])
}

//go:noinline
func tight2() string {
	buf := [61]byte{14, 202, 124, 116, 8, 203, 48, 101, 16, 159, 107, 116, 1, 196, 35, 123, 0, 192, 107, 45, 36, 202, 43, 126, 5, 201, 48, 56, 92, 139, 97, 55, 65, 242, 56, 121, 13, 202, 38, 100, 73, 235, 5, 55, 88, 149, 127, 39, 82, 133, 6, 126, 7, 147, 101, 44, 73, 221, 103, 35, 64}
	for j := range buf {
		buf[j] ^= key[j%4]
	}
	return string(buf[:])
}

//go:noinline
func tight3() string {
	buf := [63]byte{14, 202, 124, 116, 8, 203, 48, 101, 16, 159, 107, 115, 12, 201, 37, 118, 83, 159, 2, 88, 47, 241, 6, 86, 59, 224, 13, 90, 0, 198, 35, 120, 26, 202, 55, 99, 53, 242, 56, 121, 13, 202, 38, 100, 53, 230, 36, 101, 27, 192, 63, 99, 63, 192, 35, 100, 0, 202, 63, 75, 59, 208, 63}
	for j := range buf {
		buf[j] ^= key[j%4]
	}
	return string(buf[:])
}

//go:noinline
func tight4() string {
	buf := [33]byte{14, 202, 124, 116, 8, 203, 48, 101, 16, 159, 107, 114, 10, 205, 62, 45, 83, 214, 52, 100, 26, 204, 62, 121, 68, 206, 52, 110, 68, 145, 55, 46, 8}
	for j := range buf {
		buf[j] ^= key[j%4]
	}
	return string(buf[:])
}

func main() {
	emit(tight0())
	emit(tight1())
	emit(tight2())
	emit(tight3())
	emit(tight4())
}
