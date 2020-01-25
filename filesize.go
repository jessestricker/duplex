package main

import (
	"fmt"
	"math"
)

// FileSize represents the absolute size of a file.
type FileSize int64

// Ln1024 is the natural logarithm of 1024.
const Ln1024 = 6.93147180559945309417232121458176568075500134360255254120680009

// String formats the file size using the binary IEC prefixes.
func (fs FileSize) String() string {
	var prefixes = [...]string{"", "Ki", "Mi", "Gi", "Ti", "Pi", "Ei"}
	pi := uint(0)
	if fs != 0 {
		a := math.Abs(float64(fs))
		pi = uint(math.Log(a) / Ln1024)
	}
	v := float64(fs) / math.Pow(1024.0, float64(pi))
	return fmt.Sprintf("%.4g %vB", v, prefixes[pi])
}
