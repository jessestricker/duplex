package main

import (
	"math"
	"testing"
)

func TestFileSize_String(t *testing.T) {
	tests := []struct {
		name string
		fs   FileSize
		want string
	}{
		// edge cases
		{"", 0, "0 B"},
		{"", 1, "1 B"},
		{"", 512, "512 B"},
		{"", 1023, "1023 B"},

		// positive kibi
		{"", 1024, "1 KiB"},
		{"", 1025, "1.001 KiB"},
		{"", 1024 * 512, "512 KiB"},
		{"", 1024 * 1023.5, "1024 KiB"},

		// negative kibi
		{"", -1024, "-1 KiB"},
		{"", -1025, "-1.001 KiB"},
		{"", -1024 * 512, "-512 KiB"},
		{"", -1024 * 1023.5, "-1024 KiB"},

		// all positive units
		{"", 1024 * 1024, "1 MiB"},
		{"", 1024 * 1024 * 1024, "1 GiB"},
		{"", 1024 * 1024 * 1024 * 1024, "1 TiB"},
		{"", 1024 * 1024 * 1024 * 1024 * 1024, "1 PiB"},
		{"", 1024 * 1024 * 1024 * 1024 * 1024 * 1024, "1 EiB"},
		{"", math.MaxInt64, "8 EiB"},

		// all negative units
		{"", -1024 * 1024, "-1 MiB"},
		{"", -1024 * 1024 * 1024, "-1 GiB"},
		{"", -1024 * 1024 * 1024 * 1024, "-1 TiB"},
		{"", -1024 * 1024 * 1024 * 1024 * 1024, "-1 PiB"},
		{"", -1024 * 1024 * 1024 * 1024 * 1024 * 1024, "-1 EiB"},
		{"", math.MinInt64, "-8 EiB"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fs.String(); got != tt.want {
				t.Errorf("FileSize.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
