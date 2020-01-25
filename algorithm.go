package main

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash"
	"hash/adler32"
	"hash/crc32"
	"hash/crc64"
	"hash/fnv"
	"sort"
	"strings"
)

// Algorithm identifies the mathematical reduction operation
// that is performed on the file contents.
type Algorithm uint

// Supported CRC/checksum/hash algorithms.
const (
	// CRCs
	CRC32 Algorithm = iota
	CRC64

	// checksums
	Adler32

	// general-purpose hashes
	FNV1a
	MD5

	// crypto-safe hashes
	SHA256
	SHA512
)

// Name returns a descriptive name of the algorithm,
// suitable for end-user presentation.
func (a Algorithm) Name() string {
	return algorithmInfos[a].name
}

// Identifier returns a simple name of the algorithm,
// suitable as an argument for a command line parameter.
func (a Algorithm) Identifier() string {
	return algorithmInfos[a].identifier
}

func (a Algorithm) String() string {
	return a.Name()
}

// New creates a new instance of hash.Hash for the given algorithm.
func (a Algorithm) New() hash.Hash {
	return algorithmInfos[a].creator()
}

// Set overwrites this instance to an algorithm which is parsed from a string,
// trying to match it case-insensitive against one of the possible identifiers.
func (a *Algorithm) Set(s string) error {
	choices := []string{}
	for k, v := range algorithmInfos {
		choices = append(choices, v.identifier)
		if strings.EqualFold(v.identifier, s) {
			*a = k
			return nil
		}
	}

	sort.Strings(choices)
	return fmt.Errorf("algorithm is unknown\ntry one of: %s", strings.Join(choices, ", "))
}

type algorithmInfo struct {
	identifier string
	name       string
	creator    func() hash.Hash
}

var algorithmInfos = map[Algorithm]algorithmInfo{
	CRC32:   {"crc32", "CRC-32 (Castagnoli)", newCRC32},
	CRC64:   {"crc64", "CRC-64 (ISO)", newCRC64},
	Adler32: {"adler32", "Adler-32", newAdler32},
	FNV1a:   {"fnv1a", "FNV-1a (64 bit)", newFNV1a},
	MD5:     {"md5", "MD5", md5.New},
	SHA256:  {"sha256", "SHA-256", sha256.New},
	SHA512:  {"sha512", "SHA-512", sha512.New},
}

func newCRC32() hash.Hash   { return crc32.New(crc32.MakeTable(crc32.Castagnoli)) }
func newCRC64() hash.Hash   { return crc64.New(crc64.MakeTable(crc64.ISO)) }
func newAdler32() hash.Hash { return adler32.New() }
func newFNV1a() hash.Hash   { return fnv.New64a() }
