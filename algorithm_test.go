package main

import (
	"hash"
	"math/rand"
	"testing"
)

func TestFormatDigest(t *testing.T) {
	randHash := func(algo Algorithm) hash.Hash {
		hash := algo.New()
		dataLength := rand.Intn(1024)
		data := make([]byte, dataLength)
		rand.Read(data)
		_, err := hash.Write(data)
		if err != nil {
			panic(err)
		}
		return hash
	}

	tests := []struct {
		hash hash.Hash
		want int
	}{
		{randHash(CRC32), 8},    //  32 bits ->  4 bytes ->   8 hex-digits
		{randHash(CRC64), 16},   //  64 bits ->  8 bytes ->  16 hex-digits
		{randHash(Adler32), 8},  //  32 bits ->  4 bytes ->   8 hex-digits
		{randHash(FNV1a), 16},   //  64 bits ->  8 bytes ->  16 hex-digits
		{randHash(MD5), 32},     // 128 bits -> 16 bytes ->  32 hex-digits
		{randHash(SHA1), 40},    // 160 bits -> 20 bytes ->  40 hex-digits
		{randHash(SHA256), 64},  // 256 bits -> 32 bytes ->  64 hex-digits
		{randHash(SHA512), 128}, // 512 bits -> 64 bytes -> 128 hex-digits
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := len(FormatDigest(tt.hash)); got != tt.want {
				t.Errorf("FormatDigest() = %v, want %v", got, tt.want)
			}
		})
	}
}
